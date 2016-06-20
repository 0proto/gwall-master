package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/0prototype/gwall-master/entities"
)

func newhost(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Page for adding new host")
}

func hosts(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Available hosts:")
}

func alerts(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Current Alerts:")
}

func (a *Application) index(w http.ResponseWriter, r *http.Request) {
	indexPage := struct {
		Title    string
		Body     string
		UserName string
		Home     bool
	}{
		"Home",
		"Index page body",
		"Admin",
		true,
	}

	a.templates.ExecuteTemplate(w, "header", indexPage)
	a.templates.ExecuteTemplate(w, "navbar", indexPage)
	a.templates.ExecuteTemplate(w, "indexPage", indexPage)
	return
}

func (a *Application) hosts(w http.ResponseWriter, r *http.Request) {
	allHosts := entities.LoadAllHosts("", a.db)
	hostsPage := struct {
		Title string
		Hosts []entities.Host
	}{
		"Hosts",
		allHosts,
	}

	a.templates.ExecuteTemplate(w, "header", hostsPage)
	a.templates.ExecuteTemplate(w, "navbar", hostsPage)
	a.templates.ExecuteTemplate(w, "hostsPage", hostsPage)
	return
}

func (a *Application) hostedit(w http.ResponseWriter, r *http.Request) {
	// Actual updating request
	if r.Method == http.MethodPost {
		r.ParseForm()
		fmt.Println(r.Form["hostName"])
		return
	}

	// Show update form
	if r.Method == http.MethodGet {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			log.Fatal(err)
		}

		host := entities.Host{ID: id}
		host.Load(a.db)

		hostEditPage := struct {
			Title string
			Host  entities.Host
		}{
			"Edit Host",
			host,
		}

		a.templates.ExecuteTemplate(w, "header", hostEditPage)
		a.templates.ExecuteTemplate(w, "navbar", hostEditPage)
		a.templates.ExecuteTemplate(w, "hostEditPage", hostEditPage)
		return
	}
}

func (a *Application) users(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/users.html")
	t.Execute(w, nil)
}

func (a *Application) userlist(w http.ResponseWriter, r *http.Request) {
	users := entities.LoadAllUsers("", a.db)
	result, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, string(result))
	return
}
