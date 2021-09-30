package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
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

// func redirect handles non secure requests and
// redirects to https
func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://samtrouy.com", http.StatusMovedPermanently)
}

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
	// name, email, message := r.Form.Get("name"), r.Form.Get("email"), r.Form.Get("messagebody")
	// err = sendEmail(name, email, message)
	if err != nil {
		CurrentStatus = Failed
	} else {
		CurrentStatus = Sent
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func viewResume(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("site/assets/Sam-Trouy-Resume.pdf")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}
	defer f.Close()

	w.Header().Set("Content-type", "application/pdf")

	if _, err := io.Copy(w, f); err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	}
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
	fmt.Println("Listening on port 443...")
	fmt.Println("Listening on port 80...")

	go http.ListenAndServe(":80", http.HandlerFunc(redirect))

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("site"))
	mux.Handle("/css/", fs)
	mux.Handle("/js/", fs)
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/send/", sendHandler)
	mux.HandleFunc("/Sam-Trouy-Resume.pdf", viewResume)

	log.Fatal(http.ListenAndServeTLS(":443", "/home/ubuntu/pem/samtrouy.com.pem", "/home/ubuntu/pem/private.key.pem", mux))

}
