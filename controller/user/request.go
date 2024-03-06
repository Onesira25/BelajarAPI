package user

type LoginRequest struct {
	HP       string `json:"hp" form:"hp" validate:"required,max=13,min=10,number"`
	Password string `json:"password" form:"password" validate:"required,alphanum"`
}

type RegisterRequest struct {
	HP       string `json:"hp" form:"hp" validate:"required,max=13,min=10,number"`
	Name     string `json:"name" form:"name" validate:"required,min=3"`
	Password string `json:"password" form:"password" validate:"required,alphanum"`
}

type UpdateRequest struct {
	Name     string `json:"name" form:"name" validate:"required,min=3"`
	Password string `json:"password" form:"password" validate:"required,alphanum"`
}
