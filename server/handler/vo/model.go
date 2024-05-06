package vo

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
	Name     string          `json:"name"`
	Path     string          `json:"path"`
	Title    string          `json:"title"`
	IsMenu   bool            `json:"is_menu"`
	Children []*RouterRecord `json:"children"`
}

type RouterTreeResponseData struct {
	Routers []*RouterRecord `json:"routers"`
}

type RouterTreeResponse struct {
	BaseResponse
	Data *RouterTreeResponseData `json:"data"`
}
