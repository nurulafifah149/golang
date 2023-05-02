package user

type UserCreateResponse struct {
	Username string `json:"Username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}
