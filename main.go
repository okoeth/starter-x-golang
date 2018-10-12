package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

var db *gorm.DB

// Log is the main logger
var Log = log.New(os.Stdout, "STARTER-X: ", log.Lshortfile|log.LstdFlags)

func initDB(name string, drop bool) {
	Log.Printf("Initialising database as sqlite3")
	// TODO: Drop db if drop==true
	db, _ = gorm.Open("sqlite3", name)
	if db == nil {
		panic("Could not connect to database")
	}
	//defer db.Close()
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Hello World"})
}

func main() {
	initDB("todo.db", false)
	migrateModel(db)
	r := gin.Default()
	r.GET("", hello)
	addGroup(r, "/api/v1")
	r.Run(":8080")
}
