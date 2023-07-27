package main

import (
	db "hexagonal-architecture/data"
	router "hexagonal-architecture/router"
)

func main() {
	// Setup connection to database
	db.Connect()

	router.Router().Run(":" + "4000")
}
