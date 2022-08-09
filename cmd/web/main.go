package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aiym182/news/internal/config"
	"github.com/aiym182/news/internal/handlers"
	"github.com/aiym182/news/internal/render"
	"github.com/joho/godotenv"
)

var webPort string
var errorLog *log.Logger
var infoLog *log.Logger
var app = &config.AppConfig{}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Addr:    webPort,
		Handler: Routes(app),
	}

	log.Printf("Application running on port %s", webPort)
	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}

func run() error {

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("error loading env")
		return err
	}
	webPort = os.Getenv("WEBPORT")

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Println("error creating Template Cache")
		return err
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(app)
	handlers.NewHandlers(repo)
	render.NewTemplates(app)

	return nil
}
