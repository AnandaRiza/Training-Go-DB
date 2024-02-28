package main

import (
	"bcas/bookstore-go/internals/routes"
	"bcas/bookstore-go/pkg"
	"log"
)

//Dependency Injection (DI)

func main() {
	//fungsi utama
	// Initialize DB
	db, err := pkg.InitMySql()
	if err != nil {
		log.Fatal(err)
		return
	}
	// Initialize Router
	router := routes.InitRouter(db)

	// Initialize Server
	server := pkg.InitServer(router)
	// jalankan server
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
