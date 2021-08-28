package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

var templateHtml = template.Must(template.ParseGlob("templates/*"))

// Home - execute Template
func Home(writer http.ResponseWriter, request *http.Request) {
	var customers []Customer
	customers = GetCustomers()
	log.Println(customers)
	templateHtml.ExecuteTemplate(writer, "Home", customers)
}

// Create - execute Template
func Create(writer http.ResponseWriter, request *http.Request) {
	templateHtml.ExecuteTemplate(writer, "Create", nil)
}

// Insert - execute template
func Insert(writer http.ResponseWriter, request *http.Request) {
	var customer Customer
	customer.CustomerName = request.FormValue("customername")
	customer.SSN = request.FormValue("ssn")
	InsertCustomer(customer)
	var customers []Customer
	customers = GetCustomers()
	templateHtml.ExecuteTemplate(writer, "Home", customers)
}

// Alter - execute template
func Alter(writer http.ResponseWriter, request *http.Request) {
	var customer Customer
	var customerId int
	var customerIdStr string
	customerIdStr = request.FormValue("id")
	fmt.Sscanf(customerIdStr, "%d", &customerId)
	customer.CustomerId = customerId
	customer.CustomerName = request.FormValue("customername")
	customer.SSN = request.FormValue("ssn")
	UpdateCustomer(customer)
	var customers []Customer
	customers = GetCustomers()
	templateHtml.ExecuteTemplate(writer, "Home", customers)
}

// Update - execute template
func Update(writer http.ResponseWriter, request *http.Request) {
	var customerId int
	var customerIdStr string
	customerIdStr = request.FormValue("id")
	fmt.Sscanf(customerIdStr, "%d", &customerId)
	var customer Customer
	customer = GetCustomerById(customerId)
	templateHtml.ExecuteTemplate(writer, "Update", customer)
}

// Delete - execute Template
func Delete(writer http.ResponseWriter, request *http.Request) {
	var customerId int
	var customerIdStr string
	customerIdStr = request.FormValue("id")
	_, _ = fmt.Sscanf(customerIdStr, "%d", &customerId)
	var customer Customer
	customer = GetCustomerById(customerId)
	DeleteCustomer(customer)
	var customers []Customer
	customers = GetCustomers()
	templateHtml.ExecuteTemplate(writer, "Home", customers)
}

// View - execute Template
func View(writer http.ResponseWriter, request *http.Request) {
	var customerId int
	var customerIdStr string
	customerIdStr = request.FormValue("id")
	_, _ = fmt.Sscanf(customerIdStr, "%d", &customerId)
	var customer Customer
	customer = GetCustomerById(customerId)
	fmt.Println(customer)
	var customers []Customer
	customers = []Customer{customer}
	//customers.append(customers, customer)
	templateHtml.ExecuteTemplate(writer, "View", customers)
}

//// Home method renders the main.html
//func Home(writer http.ResponseWriter, reader *http.Request) {
//	var html_template *template.Template
//	html_template = template.Must(template.ParseFiles("main.html"))
//	_ = html_template.Execute(writer, nil)
//}

//// main method
//func main() {
//	log.Println("Server started on: http://localhost:8000")
//	http.HandleFunc("/", Home)
//	_ = http.ListenAndServe(":8000", nil)
//}

// main method
func main() {
	log.Println("Server started on: http://localhost:8000")
	http.HandleFunc("/", Home)
	http.HandleFunc("/alter", Alter)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/view", View)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8000", nil)
}
