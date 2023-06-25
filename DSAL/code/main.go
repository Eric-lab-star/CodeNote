package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// Customer struct  
type Customer struct {
	CustomerId   int
	CustomerName string
	SSN          string
}

// get connection method which returns sql.DB
func GetConnection() (database *sql.DB) {
	databaseDriver := "mysql"
	databaseUser := "root"
	databseName := "crm"
	database, error := sql.Open(databaseDriver, databaseUser+":@/"+databseName)
	if error != nil {
		panic(error.Error())
	}
	return database
}

//Get Customer method returns Customer Array

func GetCustomers() []Customer {
	var database *sql.DB
	database = GetConnection()
	rows, error := database.Query("Select * from Customer order by CustomerId desc")
	if error != nil {
		panic(error.Error())
	}
	var customer Customer
	customer = Customer{}
	customers := []Customer{}
	for rows.Next() {
		var customerId int
		var customerName string
		var ssn string
		error = rows.Scan(&customerId, &customerName, &ssn)
		if error != nil {
			panic(error.Error())
		}
		customer.CustomerId = customerId
	}
	defer database.Close()
	return customers
}

// InsertCustomer method with parameter customer
func InsertCustomer(customer Customer) {
	var database *sql.DB
	database = GetConnection()

	var error error
	var insert *sql.Stmt
	insert, error = database.Prepare("INSERT INTO CUSTOMER(CustomerName,SSN) VALUES(?,?)")
	if error != nil {
		panic(error.Error())
	}
	insert.Exec(customer.CustomerName, customer.SSN)

	defer database.Close()

}

func main() {
	customers := GetCustomers()
	fmt.Println("Customers", customers)
}
