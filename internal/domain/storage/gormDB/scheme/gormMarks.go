package scheme

type Mark struct {
	Id        uint `gorm:"primaryKey;<-:create"`
	Type      string
	ObjectId  uint
	Latitude  float64
	Longitude float64
}
