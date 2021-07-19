//Models/User.go
package Models

import (
	"ecapp_payment/infrastructure"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db infrastructure.Database

//GetAllUsers Fetch all user data
func GetPayment(payment *[]Payment) (err error) {
	db = infrastructure.NewDatabase()
	if err = db.DB.Find(payment).Error; err != nil {
		return err
	}
	return nil
}

//CreateUser ... Insert New data
func CreatePayment(payment *Payment) (err error) {
	db = infrastructure.NewDatabase()
	if err = db.DB.Create(payment).Error; err != nil {
		return err
	}
	//db.DB.Exec("UPDATE product SET quantity = ? WHERE id = ?", gorm.Expr("quantity - ?", cart.PQty), cart.PId)
	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetPaymentByID(payment *Payment, id string) (err error) {
	if err = db.DB.Where("o_id = ?", id).First(payment).Error; err != nil {
		return err
	}
	return nil
}

//UpdateUser ... Update user
func UpdatePayment(payment *Payment, id string) (err error) {
	fmt.Println(payment)
	db.DB.Save(payment)
	return nil
}

//DeleteUser ... Delete user
/*func DeleteProduct(product *Product, id string) (err error) {
	db.DB.Where("id = ?", id).Delete(product)
	return nil
}
*/
