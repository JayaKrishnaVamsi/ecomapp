//Models/UserModel.go
package Models

type Payment struct {
	OId  uint `gorm:"primary_key;auto_increment:false" json:"oid"`
	CId  uint `json:"cid"`
	Bill uint `json:"bill"`
}

func (b *Payment) TableName() string {
	return "payment"
}
