package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Customer struct
type Customer struct {
	CustomerId   int
	CustomerName string
	SSN          string
}

// GetConnection method which returns sql.DB
func GetConnection() (database *sql.DB) {
	databaseDriver := "mysql"
	databaseUser := "root"
	databasePass := "123456"
	databaseName := "crm"
	database, err := sql.Open(databaseDriver, databaseUser+":"+databasePass+"@tcp(121.5.167.147:13306)/"+databaseName)
	if err != nil {
		panic(err.Error())
	}
	return database
}

// GetCustomerById with parameter customerId returns Customer
func GetCustomerById(customerId int) Customer {
	var database *sql.DB
	database = GetConnection()
	var err error
	var rows *sql.Rows

	rows, err = database.Query("SELECT * FROM Customer WHERE CustomerId=?", customerId)
	if err != nil {
		panic(err.Error())
	}

	var customer Customer
	customer = Customer{}
	for rows.Next() {
		var customerId int
		var customerName string
		var SSN string
		err = rows.Scan(&customerId, &customerName, &SSN)
		if err != nil {
			panic(err.Error())
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = SSN
	}

	defer database.Close()
	return customer
}

// GetCustomers method returns Customer array
func GetCustomers() []Customer {
	var database *sql.DB
	database = GetConnection()
	var err error
	var rows *sql.Rows
	rows, err = database.Query("SELECT * FROM Customer ORDER BY Customerid DESC")
	if err != nil {
		panic(err.Error())
	}
	var customer Customer
	customer = Customer{}
	var customers []Customer
	customers = []Customer{}
	for rows.Next() {
		var customerId int
		var customerName string
		var ssn string
		err = rows.Scan(&customerId, &customerName, &ssn)
		if err != nil {
			panic(err.Error())
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = ssn
		customers = append(customers, customer)
	}
	defer database.Close()
	return customers
}

// InsertCustomer method with parameter customer
func InsertCustomer(customer Customer) {
	var database *sql.DB
	database = GetConnection()
	var err error
	var insert *sql.Stmt
	insert, err = database.Prepare("INSERT INTO Customer(CustomerName, SSN) VALUES(?, ?)")
	if err != nil {
		panic(err.Error())
	}
	_, _ = insert.Exec(customer.CustomerName, customer.SSN)
	defer database.Close()
}

// Update Customer method with parameter customer
func UpdateCustomer(customer Customer) {
	var database *sql.DB
	database = GetConnection()
	var err error
	var update *sql.Stmt
	update, err = database.Prepare("UPDATE Customer SET CustomerName=?, SSN=? WHERE CustomerId=?")
	if err != nil {
		panic(err.Error())
	}
	_, _ = update.Exec(customer.CustomerName, customer.SSN, customer.CustomerId)
	defer database.Close()
}

// Delete Customer method with parameter customer
func DeleteCustomer(customer Customer) {
	var database *sql.DB
	database = GetConnection()
	var err error
	var del *sql.Stmt
	del, err = database.Prepare("DELETE FROM Customer WHERE Customerid=?")
	if err != nil {
		panic(err.Error())
	}
	_, _ = del.Exec(customer.CustomerId)
	defer database.Close()
}
