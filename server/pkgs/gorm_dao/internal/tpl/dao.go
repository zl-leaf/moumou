package tpl

const (
	DaoStructTpl = `
package {{.QueryPkgName}}

type Dao struct {
	db        *gorm.DB

	{{range .QueryStructList -}}
		{{.Name}}Dao *{{.StructName}}Dao
	{{end}}
}

func NewDao(db *gorm.DB) *Dao {
	return &Dao{
		db:        db,
		{{range .QueryStructList -}}
			{{.Name}}Dao: new{{.Name}}Dao(db),
		{{end}}
	}
}

func (d *Dao) Transaction(fc func(tx *Dao) error, opts ...*sql.TxOptions) (err error) {
	return d.db.Transaction(func(tx *gorm.DB) error {
		newDao := NewDao(tx)
		return fc(newDao)
	}, opts...)
}
`
)
