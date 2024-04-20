package scheme

type SoldPoint struct {
	ID      uint `gorm:"primaryKey;<-:create"`
	Region  string
	Name    string
	Address string
}
