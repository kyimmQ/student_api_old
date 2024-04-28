package main

import (
	"errors"
	"fmt"
	"kyimmQ/student_api/internal/models"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (app *application) CreateStudent(w http.ResponseWriter, r *http.Request) {

	var student models.Student
	err := app.readJSON(w, r, &student)
	if err != nil {
		fmt.Println("cant read json")
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	result, err := app.DB.CreateStudent(student.STD_ID, student.STD_NAME, student.ACADEMIC_YEAR)

	if !result {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, &student)
}

func (app *application) GetStudentById(w http.ResponseWriter, r *http.Request) {

	sid, err := strconv.Atoi(chi.URLParam(r, "sid"))
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	student, err := app.DB.SearchStudentByID(sid)
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}
	_ = app.writeJSON(w, http.StatusOK, student)
}

func (app *application) GetStudentByName(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()

	if !queryValues.Has("name") {
		app.errorJSON(w, errors.New("receive no query string, try again"), http.StatusBadRequest)
		return
	}
	name := queryValues.Get("name")
	students, err := app.DB.SearchStudentByName(name)
	if err != nil {
		app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}
	_ = app.writeJSON(w, http.StatusOK, students)
}
