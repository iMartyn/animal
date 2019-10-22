package animal

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	
	"os"
	"fmt"
	"net/http"
	"html/template"
	"time"
	"path/filepath"
)

func HealthHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-type","text/plain")
	fmt.Fprint(response, "I'm okay jack!")
}

func NotFoundHandler(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/404.html"))
	tmpl.Execute(response, nil)
}

func CSSHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-type","text/css")
	tmpl := template.Must(template.ParseFiles("html/cover.css"))
	tmpl.Execute(response, nil)
}

func RootHandler(response http.ResponseWriter, request *http.Request) {
	type TemplateData struct {
		Animal     AnimalData
	}
	animalID:= FindAnimalID(AnimalName)
	if animalID == -1 {
		http.Error(response, "Animal "+AnimalName+" not found", http.StatusInternalServerError)
		return
	}
	tmpl := template.Must(template.ParseFiles(filepath.FromSlash("html/index.html")))
	animal := Animals[animalID]
	data := TemplateData{
		Animal: animal, 
	}
	tmpl.Execute(response, data)
}

func HandleHTTP() {
	r := mux.NewRouter()
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
	r.HandleFunc("/", RootHandler)
	r.HandleFunc("/healthz", HealthHandler)
	r.HandleFunc("/cover.css", CSSHandler)
	http.Handle("/", r)
	srv := &http.Server {
		Handler: loggedRouter,
		Addr: "0.0.0.0:5353",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Listening on 0.0.0.0:5353")
	srv.ListenAndServe()
}

