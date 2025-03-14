package webapi

import (
	"fmt"
	"net/http"
)

func Start(path string) {
	if path != "" {
		router := http.NewServeMux()

		server := http.Server{
			Addr:    path,
			Handler: router,
		}

		setupEndpoints(router)

		fmt.Println("wevapi server starting at:", path)

		server.ListenAndServe() // returns error if applicable
	}
}
