package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type ValgfagsStruct struct {
	Names []string
}

var Valgfags ValgfagsStruct

func main() {
	http.HandleFunc("/", home)
	http.Handle("/Valgfags/", http.StripPrefix("/Valgfags/", http.FileServer(http.Dir("./Valgfags/"))))
	http.Handle("/JS/", http.StripPrefix("/JS/", http.FileServer(http.Dir("./JS/"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	fmt.Println("Go to http://127.0.0.1:8080")

	alle_valgfags, _ := os.ReadDir("./Valgfags")
	for _, k := range alle_valgfags {
		http.HandleFunc("/"+k.Name(), valgfag)
		Valgfags.Names = append(Valgfags.Names, k.Name())
	}

	http.ListenAndServe("127.0.0.1:8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles(filepath.Join("./templates/home.html"))
	tmpl.Execute(w, Valgfags)
}

func valgfag(w http.ResponseWriter, r *http.Request) {
	name := strings.Trim(r.URL.String(), "/")

	var data struct {
		Name        string
		Information string
	}

	data.Name = name

	file, err := os.ReadFile(filepath.Join("Valgfags", name, "Information.txt"))
	if err != nil {
		data.Information = "Kun ikke finde information"
	} else {
		data.Information = string(file)
	}

	tmpl, _ := template.ParseFiles(filepath.Join("./templates/valgfag.html"))
	tmpl.Execute(w, data)
}
