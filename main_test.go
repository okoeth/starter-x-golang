package main

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	initDB("test.db", true)
	migrateModel(db)
	r := gin.Default()
	addGroup(r, "/api/v1")
	TodoURL = "http://localhost:8081/api/v1"
	Log.Printf("Creating test server on port 8081")
	go r.Run(":8081")
	Log.Printf("Running tests")
	i := m.Run()
	os.Exit(i)
}
