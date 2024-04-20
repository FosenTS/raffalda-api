package entity

type RefreshToken struct {
	ID                 uint
	Token              string
	Login              string
	ExpirationTimeUnix int64
	CreateTimeUnix     int64
}
