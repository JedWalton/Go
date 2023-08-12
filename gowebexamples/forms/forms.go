package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type ContactDetails struct {
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

func main() {
	tmpl := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		log.Println("Received request:", r.Method)

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		var details ContactDetails
		err := json.NewDecoder(r.Body).Decode(&details)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Logging each value
		log.Println("Message:", details.Message)
		log.Println("Subject:", details.Subject)
		log.Println("Email:", details.Email)

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":8080", nil)
}
