package internal

type PageDataService struct {
}

func NewPageDataService() *PageDataService {
	return &PageDataService{}
}

func (svc *PageDataService) GetList(pageID uint, currentPage, pageSize int) ([]map[string]any, int64, error) {
	dataList := make([]map[string]any, 0, pageSize)
	total := int64(200)

	pageTotal := pageSize
	if int64(currentPage*pageSize) > total {
		pageTotal = int(total) - (currentPage-1)*pageSize
	}
	if pageTotal < 0 {
		pageTotal = 0
	}
	for i := 0; i < pageTotal; i++ {
		dataList = append(dataList, map[string]any{
			"name":    "叶志良",
			"age":     (currentPage-1)*pageSize + i,
			"address": "广州",
		})
	}

	return dataList, total, nil
}
