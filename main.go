package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"social-todo-list/component/tokenProvider/jwt"
	"social-todo-list/middleware"
	todotrpt "social-todo-list/module/item/transport"
	"social-todo-list/module/user/storage"
	ginuser "social-todo-list/module/user/transport/gin"
)

func main() {
	// Checking that an environment variable is present or not.
	mysqlConnStr, ok := os.LookupEnv("MYSQL_CONNECTION")
	systemSecret := os.Getenv("SECRET")

	if !ok {
		log.Fatalln("Missing MySQL connection string.")
	}

	dsn := mysqlConnStr
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}

	log.Println("DB connection: ", db.Debug())

	authStore := storage.NewSQLStore(db)
	tokenProvider := jwt.NewTokenJWTProvider("jwt", systemSecret)
	middlewareAuth := middleware.RequiredAuth(authStore, tokenProvider)

	router := gin.Default()
	// Response error json when panic
	router.Use(middleware.Recover())

	v1 := router.Group("/v1")
	{
		v1.POST("/register", ginuser.Register(db))
		v1.POST("/login", ginuser.Login(db, tokenProvider))
		v1.GET("/profile", middlewareAuth, ginuser.Profile())
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
