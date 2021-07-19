//Models/UserModel.go
package Models

type Cart struct {
	OrderId uint `gorm:"primary_key" json:"oid"`
	CId     uint `json:"cid"`
	PId     uint `gorm:"primary_key;auto_increment:false" json:"pid"`
	PQty    uint `json:"pqty"`
}

func (b *Cart) TableName() string {
	return "cart"
}
