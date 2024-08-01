package meta

import (
	"fmt"
	"gorm.io/gorm/schema"
	"strings"
)

type QueryStruct struct {
	Schema *schema.Schema

	QueryPkgName string // 包名
	Name         string // Model对象名称
	StructName   string // 小写开头的Model名称
	WhereFns     []*Method
	OrderFns     []*Method
}

// ParseQueryStruct 解析Query对象
func ParseQueryStruct(QueryPkgName string, s *schema.Schema) (*QueryStruct, error) {
	meta := &QueryStruct{
		Schema:       s,
		QueryPkgName: QueryPkgName,
		Name:         s.Name,
		StructName:   strings.ToLower(s.Name[:1]) + s.Name[1:],
		WhereFns:     make([]*Method, 0, 10),
		OrderFns:     make([]*Method, 0, 10),
	}

	for _, field := range s.Fields {
		meta.WhereFns = append(meta.WhereFns, parseField2WhereFn(field)...)
		meta.OrderFns = append(meta.OrderFns, parseField2OrderFn(field)...)
	}

	return meta, nil
}

type Method struct {
	Field *schema.Field

	MethodName string
	Params     MethodParams
	Query      string
}

// parseField2WhereFn 字段解析为where
func parseField2WhereFn(field *schema.Field) []*Method {
	whereFns := make([]*Method, 0, 5)
	// TODO
	if field.Name != "Username" {
		return whereFns
	}

	fn := &Method{
		Field:      field,
		MethodName: "WhereUsernameEq",
		Params: []*Param{
			{Name: "username", Type: "string"},
		},
		Query: "username = ?",
	}
	whereFns = append(whereFns, fn)

	return whereFns
}

// parseField2OrderFn 字段解析为Order
func parseField2OrderFn(field *schema.Field) []*Method {
	orderFns := make([]*Method, 0, 5)
	// TODO
	if field.Name != "Sort" {
		return orderFns
	}
	fn := &Method{
		Field:      field,
		MethodName: "OrderBySort",
		Params:     []*Param{},
		Query:      fmt.Sprintf("%s ASC", field.DBName),
	}
	orderFns = append(orderFns, fn)
	return orderFns
}

type MethodParams []*Param
type Param struct { // (user model.User)
	PkgPath   string // package's path: internal/model
	Package   string // package's name: model
	Name      string // param's name: user
	Type      string // param's type: User
	IsArray   bool   // is array or not
	IsPointer bool   // is pointer or not
}

func (params MethodParams) Body() string {
	paramBody := make([]string, 0, len(params))

	for _, param := range params {
		paramName := param.Name
		if param.Package != "" {
			paramName = param.Package + "." + paramName
		}
		if param.IsPointer {
			paramName = "*" + paramName
		}
		if param.IsArray {
			paramName = "[]" + paramName
		}

		paramBody = append(paramBody, paramName+" "+param.Type)
	}

	return strings.Join(paramBody, ",")
}

func (params MethodParams) Args() string {
	paramNames := make([]string, len(params))
	for i, param := range params {
		paramNames[i] = param.Name
	}
	return strings.Join(paramNames, ",")
}
