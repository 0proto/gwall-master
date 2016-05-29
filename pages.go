package main

import (
    "io"
    "net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func newhost(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Page for adding new host")
}

func hosts(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Available hosts:")
}

func alerts(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Current Alerts:")
}
