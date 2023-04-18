package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"todolist.go/db"
	"todolist.go/service"
)

const port = 8000

func main() {
	// initialize DB connection
	dsn := db.DefaultDSN(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
	if err := db.Connect(dsn); err != nil {
		log.Fatal(err)
	}

	// initialize Gin engine
	engine := gin.Default()
	engine.LoadHTMLGlob("views/*.html")

	// routing
	engine.Static("/assets", "./assets")
	engine.GET("/", service.Login)
	engine.GET("/first", service.First)
	engine.POST("/certify", service.Certify)
	engine.POST("/new_account", service.NewAccount)
	engine.GET("/delete_user", service.DeleteUser)
	engine.GET("/edit_user", service.EditUser)
	engine.POST("/edit_username", service.EditUsername)
	engine.POST("/edit_password", service.EditPassword)
	engine.GET("/list", service.TaskList)
	engine.POST("/task_delete", service.TaskDelete)
	engine.POST("/task_toFalse", service.TaskToFalse)
	engine.POST("/task_toTrue", service.TaskToTrue)
	engine.POST("/task_newTitle", service.ChangeTitle)
	engine.GET("/task_add", service.Add)
	engine.GET("/true_list", service.TrueList)
	engine.GET("/false_list", service.FalseList)
	engine.POST("/register_add", service.TaskAdd)
	engine.GET("/logout", service.Logout)
	engine.GET("/task/:id", service.ShowTask) // ":id" is a parameter
	

	// start server
	engine.Run(fmt.Sprintf(":%d", port))
}

