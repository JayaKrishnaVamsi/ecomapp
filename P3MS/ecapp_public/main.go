package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type Cart struct {
	OrderId uint `gorm:"primary_key" json:"oid"`
	CId     uint `json:"cid"`
	PId     uint `gorm:"primary_key;auto_increment:false" json:"pid"`
	PQty    uint `json:"pqty"`
}
type Payment struct {
	OId  uint `gorm:"primary_key;auto_increment:false" json:"oid"`
	CId  uint `json:"cid"`
	Bill uint `json:"bill"`
}

var db *sql.DB

func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "./public/signup.html")
		return
	}
	usr := r.FormValue("username")
	pwd := r.FormValue("password")
	fmt.Print(usr)
	fmt.Print(pwd)
	/*if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}*/
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pwd), 8)

	// Next, insert the username, along with the hashed password into the database
	if _, err := db.Query("insert into users(username,password) values (?, ?)", usr, string(hashedPassword)); err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		fmt.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/product", http.StatusFound)
}
func Signin(w http.ResponseWriter, r *http.Request) {
	//http.ServeFile(w, r, "./public/home.html")
	if r.Method != "POST" {
		http.ServeFile(w, r, "./public/home.html")
		return
	}
	usr := r.FormValue("username")
	pwd := r.FormValue("password")

	fmt.Print(usr)
	fmt.Print(pwd)
	var result string
	db.QueryRow("select password from users where username=?", usr).Scan(&result)
	/*if err != nil {
		// If there is an issue with the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}*/
	if err := bcrypt.CompareHashAndPassword([]byte(result), []byte(pwd)); err != nil {
		// If the two passwords don't match, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		http.Post("/account", "application/json")
		http.Redirect(w, r, "/product", http.StatusFound)
	}

}
func Products(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/product.html")
}
func Carticon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/cart.html")
}
func ProdBut(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/product.html")
}

func ContBut(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/contact.html")
}
func Purchase(w http.ResponseWriter, r *http.Request) {
	postBody, _ := json.Marshal(Payment{
		OId:  1,
		CId:  1,
		Bill: 0,
	})
	responseBody := bytes.NewBuffer(postBody)
	url := "http://web_pay:8003/payment/pay"
	response, _ := http.Post(url, "application/json", responseBody)
	fmt.Println(response)
}
func Account(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "./public/account.html")
		return
	}
}
func Buy(w http.ResponseWriter, r *http.Request) {
	q1 := r.FormValue("xbox")
	q2 := r.FormValue("switch")
	q3 := r.FormValue("ps5")
	q4 := r.FormValue("3ds")
	q5 := r.FormValue("sam")
	q6 := r.FormValue("ant")
	fmt.Print(q1, q2, q3, q4, q5, q6)

	q := []string{q1, q2, q3, q4, q5, q6}
	for i := 1; i <= len(q); i++ {
		if q[i-1] != "" {
			qt, _ := strconv.ParseUint(q[i-1], 10, 32)
			postBody, _ := json.Marshal(Cart{
				OrderId: 1,
				CId:     1,
				PId:     uint(i),
				PQty:    uint(qt),
			})
			responseBody := bytes.NewBuffer(postBody)
			http.Post("http://web_cart:8002/mycart/additem/", "application/json", responseBody)
		}
	}
	http.Redirect(w, r, "/viewcart", http.StatusFound)
}
func Viewcart(w http.ResponseWriter, r *http.Request) {

	po := 1
	fmt.Print("here")
	resp, _ := http.Get("http://web_cart:8002/mycart/cart/" + strconv.FormatUint(uint64(po), 10))
	body, _ := ioutil.ReadAll(resp.Body)
	var responseObject []Cart
	json.Unmarshal(body, &responseObject)
	fmt.Println(responseObject)
	/*if r.Method != "POST" {
		http.ServeFile(w, r, "./public/cart.html")
		return
	}*/
	t, _ := template.ParseFiles("./public/cart.html")
	t.Execute(w, responseObject)
}
func initDB() {
	var err error
	db, err = sql.Open("mysql", "root:vamsi4mysql@tcp(db_detls)/ecom")
	if err != nil {
		panic(err)
	}
}
func main() {
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/signup", Signup)
	http.HandleFunc("/product", Products)
	http.HandleFunc("/account.html", Account)
	http.HandleFunc("/buy", Buy)
	http.HandleFunc("/viewcart", Viewcart)
	http.HandleFunc("/purchase", Purchase)
	http.HandleFunc("/cart.html", Carticon)
	http.HandleFunc("/contact.html", ContBut)
	http.HandleFunc("/product.html", ProdBut)
	// initialize our database connection
	initDB()
	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8004", nil))
}
