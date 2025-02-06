package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
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


	alle_valgfags,_ := os.ReadDir("./Valgfags")
	for _,k := range alle_valgfags {
		file_path := "/Valgfags/" + k.Name()
		fmt.Println(file_path)
		http.Handle(file_path, http.StripPrefix(file_path, http.FileServer(http.Dir(file_path))))
		Valgfags.Names = append(Valgfags.Names, k.Name())

	}

	http.ListenAndServe("127.0.0.1:8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)

	tmpl,_ := template.ParseFiles(filepath.Join("index.html"))
	tmpl.Execute(w, Valgfags)
}

