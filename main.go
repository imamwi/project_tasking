package main

import (
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

	noteController := &controllers.NoteControllers{}

	router := httprouter.New()

	router.GET("/", noteController.Index)

	log.Fatal(http.ListenAndServe(":9000", router))

}
