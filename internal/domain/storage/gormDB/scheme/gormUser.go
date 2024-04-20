package scheme

type User struct {
	ID         uint   `gorm:"primaryKey;<-:create"`
	Login      string `gorm:"index;unique"`
	Password   string
	Permission uint
}
