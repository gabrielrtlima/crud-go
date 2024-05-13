package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" binding:"required,min=3,max=100"`
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%*"`
	Age      int8   `json:"age" binding:"required,min=18,max=120"`
}

type UserUpdateRequest struct {
	Name string `json:"name" binding:"omitempty,min=3,max=100"`
	Age  int8   `json:"age" binding:"omitempty,min=18,max=120"`
}
