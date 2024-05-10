package main

import (
	"log"
	"net/http"
	"os"
	todotrpt "social-todo-list/transport"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
		log.Println("DB connection: ", db)
	}

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", todotrpt.HandleCreateItem(db))         // create item
			items.GET("", todotrpt.HandleListItem(db))            // list items
			items.GET("/:id", todotrpt.HandleFindAnItem(db))      // get an item by ID
			items.PUT("/:id", todotrpt.HandleUpdateAnItem(db))    // edit an item by ID
			items.DELETE("/:id", todotrpt.HandleDeleteAnItem(db)) // delete an item by ID
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
