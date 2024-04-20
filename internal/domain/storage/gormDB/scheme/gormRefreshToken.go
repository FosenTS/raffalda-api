package scheme

type RefreshToken struct {
	ID                 uint   `gorm:"primaryKey;<-:create"`
	Token              string `gorm:"index;unique"`
	Login              string `gorm:"index;unique"`
	ExpirationTimeUnix int64
	CreateTimeUnix     int64
}
