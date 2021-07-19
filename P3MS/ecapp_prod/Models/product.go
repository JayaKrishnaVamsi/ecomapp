//Models/User.go
package Models

import (
	"ecapp_prod/infrastructure"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db infrastructure.Database

//GetAllUsers Fetch all user data
func GetAllProducts(product *[]Product) (err error) {
	db = infrastructure.NewDatabase()
	if err = db.DB.Find(product).Error; err != nil {
		return err
	}
	return nil
}

//CreateUser ... Insert New data
func CreateProduct(product *Product) (err error) {
	db = infrastructure.NewDatabase()
	if err = db.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetProductByID(product *Product, id string) (err error) {
	db = infrastructure.NewDatabase()
	if err = db.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}
func UpdateProduct(product *Product, id string) (err error) {
	db = infrastructure.NewDatabase()
	fmt.Println(product)
	db.DB.Save(product)
	return nil
}

//DeleteUser ... Delete user
func DeleteProduct(product *Product, id string) (err error) {
	db = infrastructure.NewDatabase()
	db.DB.Where("id = ?", id).Delete(product)
	return nil
}
