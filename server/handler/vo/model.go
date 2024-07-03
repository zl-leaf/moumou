package vo

import "github.com/moumou/server/biz/model"

var errorResponse = &BaseResponse{
	Code:    -1,
	Message: "error",
}

type BaseRequest struct {
	Token string `header:"x-token"`
}

type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type BaseRequestPage struct {
	CurrentPage int `json:"current_page"`
	PageSize    int `json:"page_size"`
	// TODO order by
}

type BaseResponsePageData struct {
	Total int64 `json:"total"`
}

type PingRequest struct {
}

type PingResponse struct {
	BaseResponse
}

type HelloRequest struct {
}

type HelloResponseData struct {
	Key string `json:"key"`
	IV  string `json:"iv"`
}

type HelloResponse struct {
	BaseResponse
	Data *HelloResponseData `json:"data"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LogoutResponse struct {
	BaseResponse
}

type LogoutRequest struct {
	BaseRequest
}

type UserData struct {
	UserID string `json:"user_id"`
}

type LoginResponseData struct {
	User  *UserData `json:"user"`
	Token string    `json:"token"`
}

type LoginResponse struct {
	BaseResponse
	Data *LoginResponseData `json:"data"`
}

type SelfRequest struct {
	BaseRequest
}

type SelfResponse struct {
	BaseResponse
	Data *UserData `json:"data"`
}

type RouterTreeRequest struct {
	BaseRequest
}

type RouterRecord struct {
	Name      string          `json:"name"`
	Path      string          `json:"path"`
	Title     string          `json:"title"`
	IsMenu    bool            `json:"is_menu"`
	Component string          `json:"component"`
	Children  []*RouterRecord `json:"children"`
}

type RouterTreeResponseData struct {
	Routers []*RouterRecord `json:"routers"`
}

type RouterTreeResponse struct {
	BaseResponse
	Data *RouterTreeResponseData `json:"data"`
}

type GetPageRequest struct {
	BaseRequest
	PageID uint `json:"page_id"`
}

type GetPageResponseData struct {
	Page *model.Page `json:"page"`
}

type GetPageResponse struct {
	BaseResponse
	Data *GetPageResponseData `json:"data"`
}

type GetPageDataListRequest struct {
	BaseRequest
	PageID      uint `json:"page_id"`
	CurrentPage int  `json:"current_page"`
	PageSize    int  `json:"page_size"`
}

type GetPageDataListResponseData struct {
	Total int64            `json:"total"`
	List  []map[string]any `json:"list"`
}

type GetPageDataListResponse struct {
	BaseResponse
	Data *GetPageDataListResponseData `json:"data"`
}
