package main

import (
	"FirstProject/datafile"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type GuestBook struct {
	SignatureCount int
	Signature      []string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)

}

func createHandler(writer http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("signature")

	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
	check(err)
	_, err = fmt.Fprintln(file, signature)
	check(err)
	err = file.Close()
	check(err)
	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}

func newHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("new.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signature, err := datafile.GetStrings("signatures.txt")
	check(err)
	html, err := template.ParseFiles("view.html")
	check(err)

	guestBook := GuestBook{
		SignatureCount: len(signature),
		Signature:      signature,
	}

	err = html.Execute(writer, guestBook)
	check(err)

}
