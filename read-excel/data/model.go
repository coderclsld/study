package data

type BusinessKnowledge struct {
	Business_id  int `gorm:"primaryKey"`
	Business_pid int
	Name         string
	Subject_id   int
	Base_id      int
}

func (BusinessKnowledge) TableName() string {
	return "business_knowledge"
}
