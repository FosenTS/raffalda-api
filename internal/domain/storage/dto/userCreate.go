package dto

type UserCreate struct {
	Login      string `json:"login" binding:"true"`
	Password   string `json:"password" binding:"true"`
	Permission int    `json:"permission" binding:"true"`
}

func NewUserCreate(login string, password string, permission int) *UserCreate {
	return &UserCreate{
		Login:      login,
		Password:   password,
		Permission: permission,
	}
}
