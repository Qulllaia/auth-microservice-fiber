package main

import (
	"log"
	"main/config"
	"main/database"
	"main/router"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"main/gRPC"
)


func main() {

	gRPC.InitGRPC();

	var connectionString string = config.CreateConnectionString(); 
	db, err := sqlx.Connect("postgres", connectionString);
	if err != nil{
		log.Fatal("Failed Connection", err)
	}
	defer db.Close();
	_database := database.InitDatabse(db);

	app := fiber.New()
	api := app.Group("/api")
	router.SetupAuthRoutes(api, _database)
	router.SetupUserRoutes(api,_database)
	app.Listen(":3000")
}