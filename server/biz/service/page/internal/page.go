package internal

import "github.com/moumou/server/biz/model"

type PageService struct {
}

func NewPageService() *PageService {
	return &PageService{}
}

func (svc *PageService) GetPage(pageID int64) (*model.Page, error) {
	page := &model.Page{
		BaseModel: model.BaseModel{
			ID: pageID,
		},
		Title: "测试页面",
		Schema: &model.PageSchema{
			PageID: pageID,
			Attributes: []*model.PageSchemaAttribute{
				{
					FieldAttribute: &model.PageSchemaFieldAttribute{
						Key:   "name",
						Label: "名称",
					},
				},
				{
					FieldAttribute: &model.PageSchemaFieldAttribute{
						Key:   "age",
						Label: "年龄",
					},
				},
				{
					FieldAttribute: &model.PageSchemaFieldAttribute{
						Key:   "address",
						Label: "地址",
					},
				},
			},
		},
	}
	return page, nil
}
