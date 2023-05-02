package reponse

const (
	InvalidParam       = "invalid param request"
	InvalidBody        = "invalid body request"
	InvalidPayload     = "invalid payload request"
	InvalidQuery       = "invalid query request"
	InternalServer     = "internal server error"
	SomethingWentWrong = "something went wrong"
	Unauthorized       = "unauthorized request"
	Forbidden          = "client does not have access rights to do this action"
)

type WebResponseFailed struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

type WebResponseSuccess struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
