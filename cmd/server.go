package main

import (
	"fmt"
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

var MessageSent = "<p id='messageStatus'>Message sent successfully</p>"
var MessageFailed = "<p id='messageStatus'>Message failed to send. Please try again momentarily or click <a href='mailto:trouys16@gmail.com'>here</a>"
var scroll = "<script>location.href = '#message'</script>"

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
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	name, email, message := r.Form.Get("name"), r.Form.Get("email"), r.Form.Get("message")
	err = sendEmail(name, email, message)
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
		err := t.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case Sent:
		data := struct {
			Message template.HTML
			Scroll  template.HTML
		}{
			Message: template.HTML(MessageSent),
			Scroll:  template.HTML(scroll),
		}
		err := t.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case Failed:
		data := struct {
			Message template.HTML
			Scroll  template.HTML
		}{
			Message: template.HTML(MessageFailed),
			Scroll:  template.HTML(scroll),
		}
		err := t.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func main() {
	fmt.Println("Listening on port 80...")
	fs := http.FileServer(http.Dir("site"))
	http.Handle("/css/", fs)
	http.Handle("/js/", fs)
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/send/", sendHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
