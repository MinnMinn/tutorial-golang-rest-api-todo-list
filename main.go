package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"social-todo-list/middleware"
	todotrpt2 "social-todo-list/module/item/transport"
)

func main() {
	// Checking that an environment variable is present or not.
	mysqlConnStr, ok := os.LookupEnv("MYSQL_CONNECTION")

	if !ok {
		log.Fatalln("Missing MySQL connection string.")
	}

	dsn := mysqlConnStr
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}

	log.Println("DB connection: ", db.Debug())

	router := gin.Default()
	// Response error json when panic
	router.Use(middleware.Recover())

	v1 := router.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", todotrpt2.HandleCreateItem(db))         // create item
			items.GET("", todotrpt2.HandleListItem(db))            // list items
			items.GET("/:id", todotrpt2.HandleFindAnItem(db))      // get an item by ID
			items.PUT("/:id", todotrpt2.HandleUpdateAnItem(db))    // edit an item by ID
			items.DELETE("/:id", todotrpt2.HandleDeleteAnItem(db)) // delete an item by ID
		}
	}

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": http.StatusOK,
		})
	})

	if err := router.Run(":3000"); err != nil {
		log.Fatalln(err)
	}

	router.Run()
}
