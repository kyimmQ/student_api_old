package repository

import (
	"database/sql"
	"kyimmQ/student_api/internal/models"
)

type DatabaseRepo interface {
	Connect() *sql.DB
	// database interaction for creating new student
	CreateStudent(stdId int, stdName string, acaYear int) (bool, error)
	// database interaction for searching student
	SearchStudentByID(stdID int) (*models.Student, error)
	SearchStudentByName(stdName string) (*[]models.Student, error)
}
