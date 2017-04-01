package httpapi

import (
	"fmt"
	"net/http"

	"github.com/danlin/mutago/backend"
)

var (
	service *backend.Service
)

func tracks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getTracks(w, r)
		break
	case "POST":
		postTracks(w, r)
		break
	}
}

func Start(path string, bind string, port int) {
	b, err := backend.Open(path)
	if err != nil {
		panic(err)
	}
	defer b.Close()
	service = b

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "Invalid API endpoint", 404)
	})
	http.HandleFunc("/tracks", tracks)

	fmt.Printf("Server listen to %s:%d", bind, port)

	http.ListenAndServe(fmt.Sprintf("%s:%d", bind, port), nil)
}
