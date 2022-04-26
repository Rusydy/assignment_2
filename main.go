package main

import (
	"assignment_2/connection"
	"assignment_2/routers"
)

func main() {
	connection.StartDB()

	routers.StartRoute().Run(":8080")
}
