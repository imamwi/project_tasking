package controllers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type NoteControllers struct{}

func (controller *NoteControllers) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := gorm.Open(sqlite.Open("data.db", &gorm.Config{}))

	if err != nil {
		panic("failed to connect database")
	}

	files := []string{
		".views/base.html",
		"./views/index.html",
	}

	htmlTemplate, err := template.ParseFiles(files...)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	datas := map[string]interface{}{}

	err = htmlTemplate.ExecuteTemplate(w, "base", datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())
	}
}
