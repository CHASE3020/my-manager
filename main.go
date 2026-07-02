package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Create a blueprint for an Activity
type Activity struct {
	Name      string
	StartTime string
	EndTime   string
}

// Create a list to store your schedule in memory
var schedule []Activity

func main() {
	// 1. Serve the CSS file so both pages look good
	http.HandleFunc("/index.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.css")
	})

	// 2. The Landing Page Route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// This ensures that ONLY the exact root URL shows the landing page
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "index.html")
			return
		}
		// If they type a random URL, show a 404 error
		http.NotFound(w, r)
	})

	// 3. Handle the Registration
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// Later, you will grab the username and password here to save them to a database.
			// For now, we simply redirect the user to the dashboard!
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		} else {
			// Blocks anyone trying to just type /register into their URL bar
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// 4. The Dashboard Route
	http.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("dashboard.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, schedule)
	})

	// 5. Handle Adding a New Activity
	http.HandleFunc("/add-activity", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			newActivity := Activity{
				Name:      r.FormValue("task-name"),
				StartTime: r.FormValue("start-time"),
				EndTime:   r.FormValue("end-time"),
			}

			schedule = append(schedule, newActivity)
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		}
	})

	// Start the Server
	port := ":8080"
	fmt.Printf("Server is starting... Go to http://localhost%s/\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}