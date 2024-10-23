package main

import (
	"fmt"
	"os"
)

func main() {
	// Set Environment Variables
	os.Setenv("SERVICE_NAME", "200Lab Service")

	// Get value from an Environment Variable
	srvName := os.Getenv("SERVICE_NAME")
	fmt.Printf("Service name: %s\n", srvName)

	// Unset an Environment Variable
	os.Unsetenv("SERVICE_NAME")
	fmt.Printf("After unset, Service name: '%s'\n", os.Getenv("SERVICE_NAME"))

	// Checking that an environment variable is present or not.
	mysqlConnStr, ok := os.LookupEnv("MYSQL_CONNECTION")
	if !ok {
		fmt.Println("MySQL connection string is not present")
	} else {
		fmt.Printf("MySQL connection string: %s\n", mysqlConnStr)
	}
}
