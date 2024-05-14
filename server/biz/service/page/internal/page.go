package internal

import "github.com/moumou/server/biz/model"

type PageService struct {
}

func NewPageService() *PageService {
	return &PageService{}
}

func (svc *PageService) GetPage(pageID uint) (*model.Page, error) {
	page := &model.Page{
		BaseModel: model.BaseModel{
			ID: pageID,
		},
		Title: "测试页面",
		Schema: &model.SysPageSchema{
			PageID: pageID,
			Attributes: []*model.SysPageSchemaAttribute{
				{
					FieldAttribute: &model.SysPageSchemaFieldAttribute{
						Key:   "name",
						Label: "名称",
					},
				},
				{
					FieldAttribute: &model.SysPageSchemaFieldAttribute{
						Key:   "age",
						Label: "年龄",
					},
				},
				{
					FieldAttribute: &model.SysPageSchemaFieldAttribute{
						Key:   "address",
						Label: "地址",
					},
				},
			},
		},
	}
	return page, nil
}
