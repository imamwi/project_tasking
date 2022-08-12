package controllers

import (
	"html/template"
	"log"
	"net/http"
	"project_tasking/models"

	"github.com/julienschmidt/httprouter"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type NoteControllers struct{}

func (controller *NoteControllers) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	files := []string{
		"./views/base.html",
		"./views/index.html",
	}

	htmlTemplate, err := template.ParseFiles(files...)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	var notes []models.Note
	db.Find(&notes)

	datas := map[string]interface{}{
		"Notes": notes,
	}

	err = htmlTemplate.ExecuteTemplate(w, "base", datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())
	}
}

func (controller *NoteControllers) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	if r.Method == "POST" {
		// fmt.Println(r.FormValue("content"))
		note := models.Note{
			Assignee: r.FormValue("assigne"),
			Content:  r.FormValue("content"),
			Date:     r.FormValue("deadline"),
		}

		result := db.Create(&note)
		if result.Error != nil {
			log.Println(result.Error)
		}

		http.Redirect(w, r, "/", http.StatusFound)

	} else {
		files := []string{
			"./views/base.html",
			"./views/create.html",
		}

		htmlTemplate, err := template.ParseFiles(files...)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		htmlTemplate.ExecuteTemplate(w, "base", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err.Error())
		}
	}

}
