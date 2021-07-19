//Models/UserModel.go
package Models

type Customer struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func (b *Customer) TableName() string {
	return "customer"
}
