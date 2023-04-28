package data

type BusinessKnowledge struct {
	Business_id  int                  `gorm:"primaryKey" json:"Business_id"`
	Business_pid int                  `json:"Business_pid"`
	Name         string               `json:"Name"`
	Subject_id   int                  `json:"Subject_id"`
	Base_id      int                  `json:"Base_id"`
	Child        []*BusinessKnowledge `gorm:"-" json:"Child"`
}

func (BusinessKnowledge) TableName() string {
	return "business_knowledge"
}

// func (b *BusinessKnowledge) String() string {
// 	return fmt.Sprintf("{Business_id:%v,Business_pid:%v,Name:%v,Subject_id:%v,Base_id:%v,Child:%v}", b.Business_id, b.Business_pid, b.Name, b.Subject_id, b.Base_id, b.Child)
// }
