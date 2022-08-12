package main

import (
	"fmt"
	"log"
	"net/http"
	"project_tasking/controllers"
	"project_tasking/models"

	"github.com/julienschmidt/httprouter"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	//setup db
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&models.Note{})
	if err != nil {
		panic("failed to connect models table")
	}

	// controller
	noteController := &controllers.NoteControllers{}

	// routes
	router := httprouter.New()

	router.GET("/", noteController.Index)
	router.GET("/create", noteController.Create)
	router.POST("/store", noteController.Store)
	router.GET("/edit/:id", noteController.Edit)
	router.POST("/update/:id", noteController.Update)
	router.POST("/done/:id", noteController.Done)
	router.POST("/delete/:id", noteController.Delete)

	fmt.Println("Aplikasi berjalan di http://localhost:9000")
	log.Fatal(http.ListenAndServe(":9000", router))

}
