package webapi

import (
	"io"
	"net/http"
)

var chCommand chan string

func SetCommandChannel(ch chan string) {
	chCommand = ch
}

var chPost chan string

func SetPostChannel(ch chan string) {
	chPost = ch
}

func setupEndpoints(router *http.ServeMux) error {

	router.HandleFunc("GET /echo/{message}", echo)
	router.HandleFunc("GET /cmd/{command}", command)
	router.HandleFunc("POST /employee", postEmployee)
	return nil
}

func echo(w http.ResponseWriter, r *http.Request) {
	message := r.PathValue("message")
	w.Write([]byte("Echo: " + message))
}

func command(w http.ResponseWriter, r *http.Request) {
	cmd := r.PathValue("command")

	if chCommand != nil {
		chCommand <- cmd
	}

	w.Write([]byte("Cmd: " + cmd))
}

func postEmployee(w http.ResponseWriter, r *http.Request) {

	bytedata, err := io.ReadAll(r.Body)

	if err == nil {
		chPost <- string(bytedata)
	}
}
