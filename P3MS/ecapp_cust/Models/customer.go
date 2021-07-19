//Models/Customer.go
package Models

import (
	"ecapp_cust/infrastructure"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db infrastructure.Database

//GetAllCustomers Fetch all user data
func GetAllCustomers(customer *[]Customer) (err error) {
	db = infrastructure.NewDatabase()
	if err = db.DB.Find(customer).Error; err != nil {
		return err
	}
	return nil
}

//CreateCustomer ... Insert New data
func CreateCustomer(customer *Customer) (err error) {
	db = infrastructure.NewDatabase()
	if err = db.DB.Create(customer).Error; err != nil {
		return err
	}
	return nil
}

//GetCustomerByID ... Fetch only one user by Id
func GetCustomerByID(customer *Customer, id string) (err error) {
	db = infrastructure.NewDatabase()
	if err = db.DB.Where("id = ?", id).First(customer).Error; err != nil {
		return err
	}
	return nil
}

//UpdateCustomer ... Update user
func UpdateCustomer(customer *Customer, id string) (err error) {
	db = infrastructure.NewDatabase()
	fmt.Println(customer)
	db.DB.Save(customer)
	return nil
}

//DeleteCustomer ... Delete user
func DeleteCustomer(customer *Customer, id string) (err error) {
	db = infrastructure.NewDatabase()
	db.DB.Where("id = ?", id).Delete(customer)
	return nil
}
