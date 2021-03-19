package main

import (
	"html/template"
	"log"
	"net/http"
)

type Status int

const (
	Visit Status = iota
	Sent
	Failed
)

var CurrentStatus = Visit

var MessageSent = "<p id='messageSent'>Message sent successfully</p>"

var MessageFailed = "<p id='messageFailed'>Message failed to send. Please try again momentarily or copy this email address: trouys16@gmail.com</p>"

func homeHandler(w http.ResponseWriter, r *http.Request) {
	switch CurrentStatus {
	case Visit:
		renderTemplate(w, Visit)
	case Sent:
		renderTemplate(w, Sent)
		CurrentStatus = Visit
	case Failed:
		renderTemplate(w, Failed)
		CurrentStatus = Visit
	}
}

func sendHandler(w http.ResponseWriter, r *http.Request) {
	err := sendEmail("body", "me")
	if err != nil {
		CurrentStatus = Failed
	} else {
		CurrentStatus = Sent
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

var t = template.Must(template.ParseFiles("./site/index.html"))

func renderTemplate(w http.ResponseWriter, status Status) {
	switch status {
	case Visit:
		err := t.ExecuteTemplate(w, "index.html", "")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case Sent:
		err := t.ExecuteTemplate(w, "index.html", template.HTML(MessageSent))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case Failed:
		err := t.ExecuteTemplate(w, "index.html", template.HTML(MessageFailed))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func main() {
	fs := http.FileServer(http.Dir("site"))
	http.Handle("/css/", fs)
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/send/", sendHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
