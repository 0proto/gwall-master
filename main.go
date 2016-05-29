package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	mux := make(map[string]func(http.ResponseWriter, *http.Request))

	server := http.Server{
		Addr: ":29000",
		Handler: &masterHandler{
			mux: mux,
		},
	}

	mux["/"] = index

	err := server.ListenAndServe()
	fmt.Println(err.Error())
}

type masterHandler struct {
	mux map[string]func(http.ResponseWriter, *http.Request)
}

func (m *masterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := m.mux[r.URL.String()]; ok {
		h(w, r)
		return
	}

	io.WriteString(w, "My server: "+r.URL.String())
}
