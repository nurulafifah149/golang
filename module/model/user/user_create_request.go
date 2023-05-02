package user

type UserCreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
