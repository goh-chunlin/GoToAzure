package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("static/img"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("static/fonts"))))
	http.HandleFunc("/", home)
	http.HandleFunc("/api", api)
	http.ListenAndServe(getPort(), nil)
}

func getPort() string {
	p := os.Getenv("HTTP_PLATFORM_PORT")
	if p != "" {
		return ":" + p
	}
	return ":80"
}

func render(response http.ResponseWriter, templateFileName string) {

	templateFileName = fmt.Sprintf("static/templates/%s", templateFileName)
	template, err := template.ParseFiles(templateFileName)

	if err != nil {
		log.Print("template parsing error: ", err)
	}

	err = template.Execute(response, nil)

	if err != nil {
		log.Print("template executing error: ", err)
	}
}

func home(response http.ResponseWriter, request *http.Request) {
	render(response, "index.html")
}

func api(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	apiStatus := modelAPIStatus{
		Status:  true,
		Message: "The API method is called successfully.",
	}
	output, _ := json.MarshalIndent(&apiStatus, "", "\t")

	response.WriteHeader(200)
	response.Write(output)
}
