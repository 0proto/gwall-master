package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Application struct {
	db        *sql.DB
	cfg       *Config
	templates *template.Template
}

func main() {

	confPtr := flag.String("conf", "./config.json", "Config file path")

	flag.Parse()

	config := LoadConfigFromFile(*confPtr)

	db, err := sql.Open("sqlite3", config.Dbfile)
	if err != nil {
		log.Fatal(err)
	}

	app := Application{
		db:        db,
		cfg:       config,
		templates: template.Must(template.ParseGlob("templates/*")),
	}

	http.HandleFunc("/", app.index)
	http.HandleFunc("/hosts", app.hosts)
	http.HandleFunc("/hosts/edit", app.hostedit)
	http.HandleFunc("/user/", app.users)
	http.HandleFunc("/user/list/", app.userlist)

	addr := app.cfg.BindAddress + ":" + strconv.Itoa(app.cfg.Port)

	// Serve static files from static subdirectory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.ListenAndServe(addr, nil)
}
