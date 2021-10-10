package RestServer

import (
	"net/http"
	"time"
)

func Server(handler http.Handler) *http.Server{
	return &http.Server{
		Handler: handler,
		Addr: ":8080",
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
