package controllers

import (
	"html/template"
	"log"
	"net/http"
	"project_tasking/config"
	"project_tasking/models"

	"github.com/julienschmidt/httprouter"
)

type NoteControllers struct{}

// index
func (controller *NoteControllers) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db, err := config.ConnectionDb()

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

// create
func (controller *NoteControllers) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	_, err := config.ConnectionDb()
	if err != nil {
		panic(err.Error())
	}

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

// store
func (controller *NoteControllers) Store(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	db, err := config.ConnectionDb()
	if err != nil {
		panic(err.Error())
	}

	note := models.Note{
		Assignee: r.FormValue("pegawai"),
		Content:  r.FormValue("content"),
		Date:     r.FormValue("deadline"),
	}

	result := db.Create(&note)
	if result.Error != nil {
		log.Println(result.Error)
	}

	http.Redirect(w, r, "/", http.StatusFound)

}

// edit
func (controller *NoteControllers) Edit(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	db, err := config.ConnectionDb()
	if err != nil {
		panic(err.Error())
	}

	files := []string{
		"./views/base.html",
		"./views/edit.html",
	}
	htmlTemplate, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	var note models.Note
	db.Where("ID = ?", params.ByName("id")).Find(&note)

	data := map[string]interface{}{
		"Note": note,
		"ID":   params.ByName("id"),
	}

	err = htmlTemplate.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())
	}
}

// update
func (controller *NoteControllers) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	db, err := config.ConnectionDb()

	if err != nil {
		panic(err.Error())
	}

	noteID := params.ByName("id")

	var note models.Note
	db.Where("ID = ?", noteID).First(&note)

	note.Assignee = r.FormValue("pegawai")
	note.Date = r.FormValue("deadline")
	note.Content = r.FormValue("content")

	db.Save(&note)

	http.Redirect(w, r, "/", http.StatusFound)
}

// done
func (controller *NoteControllers) Done(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	db, err := config.ConnectionDb()

	if err != nil {
		panic(err.Error())
	}

	var note models.Note
	db.Find(&note, params.ByName("id"))

	note.IsDone = !note.IsDone

	db.Save(&note)

	http.Redirect(w, r, "/", http.StatusFound)
}

// delete
func (controller *NoteControllers) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	db, err := config.ConnectionDb()
	if err != nil {
		panic(err.Error())
	}

	var note models.Note
	db.Delete(&note, params.ByName("id"))

	http.Redirect(w, r, "/", http.StatusFound)

}
