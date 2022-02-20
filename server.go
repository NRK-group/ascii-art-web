package main

import (
	"asciiweb/banner"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var T = template.Must(template.ParseGlob("static/*.html")) // calling the html file

type Output struct { // struct to output the input, banner, and the art
	Input  string
	Banner string
	Art    string
}

func AsciiArt(input, ban string) (string, error) { // this function accept two argument, the "input" and  the type of the "banner"
	if !(ban == "standard" || ban == "shadow" || ban == "thinkertoy") {
		return "", errors.New("invalid") // return if theirs an error
	}
	return banner.PrintAsciiArt(input, "banner/"+ban+".txt"), nil // return if theirs no error in the banner and input
}

func Process(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/ascii-art" { // only accept the / and /ascii-art
		http.Error(w, "404 Status not found", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		input := ""
		banner := ""
		var out Output
		out.Input = input
		out.Banner = banner
		output, _ := AsciiArt(input, banner) // converts the input to an asciiart
		out.Art = output
		if err := T.Execute(w, out); err != nil { // Execute the AsciiArt to prevent printing {{.}}
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}
	case "POST":
		input := r.FormValue("input-name") // getting the input in <textarea>
		banner := r.FormValue("banner")    // getting the banner value in the radio button
		if input == "" || banner == "" {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		output := ""
		if strings.Contains(input, "\r\n") { // cathing the "return or enter" value
			output = strings.Replace(input, "\r\n", "\\n", -1) // and replace it with newline
		} else {
			output = input
		}
		var out Output
		output, err := AsciiArt(output, banner) //call the AsciiArt() to convert the input into Ascii Art
		if err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}
		out.Input = input
		out.Banner = banner
		out.Art = output
		if err := T.Execute(w, out); err != nil { // execute the struct to the html
			http.Error(w, "Internal Error", http.StatusInternalServerError) // error if their is a prolblem in the server
			return
		}
	default:
		http.Error(w, "Bad request", http.StatusBadRequest) // if the request method is not GET or POST
	}
}
func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs)) // handling the CSS
	http.HandleFunc("/", Process)                             // handle the GET request
	http.HandleFunc("/ascii-art", Process)                    // handle the POST request
	fmt.Printf("Starting server at port 5500\n")
	log.Fatal(http.ListenAndServe(":5500", nil)) // run the server
}
