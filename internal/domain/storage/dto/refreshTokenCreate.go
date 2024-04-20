package dto

type RefreshTokenCreate struct {
	Token              string
	Login              string
	ExpirationTimeUnix int64
	CreateTimeUnix     int64
}

func NewRefreshTokenCreate(token string, login string, expirationTimeUnix int64, createTimeUnix int64) *RefreshTokenCreate {
	return &RefreshTokenCreate{Token: token, Login: login, ExpirationTimeUnix: expirationTimeUnix, CreateTimeUnix: createTimeUnix}
}
