package main

import (
	"flag"
	"fmt"
	"kyimmQ/student_api/internal/repository"
	"kyimmQ/student_api/internal/repository/dbrepo"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	DBInfo string
	DB     repository.DatabaseRepo
}

func main() {

	var port int
	var (
		env        string
		DBUsername string
		DBPassword string
		DBHost     string
		DBPort     string
	)
	flag.StringVar(&env, "env", "dev", "specify environment")
	flag.StringVar(&DBUsername, "DBUsername", "", "specify username for database connection")
	flag.StringVar(&DBPassword, "DBPassword", "", "specify password for database connection")
	flag.StringVar(&DBHost, "DBHost", "", "specify host for database connection")
	flag.StringVar(&DBPort, "DBPort", "", "specify port for database connection")
	// Parse command-line flags
	flag.Parse()
	if flag.NFlag() != 5 {
		fmt.Println("Usage: go run ./cmd/api -env= -DBUsername= -DBPassword= -DBHost= -DBPort=")
		os.Exit(1)
	}

	if env == "dev" {
		port = 8080
	} else {
		port = 443
	}
	// main app
	var app application
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/student_api", DBUsername, DBPassword, DBHost, DBPort)
	app.DBInfo = connectionString
	// create a database object which can be used
	// to connect with database.
	conn, err := app.connectToDB()
	if err != nil {
		panic(err)
	}
	app.DB = &dbrepo.MySQLDBRepo{DB: conn}
	defer app.DB.Connect().Close()

	log.Println("Starting application on port", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
