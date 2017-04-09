package httpapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"io"

	"github.com/danlin/mutago/parser"
)

func getTracks(w http.ResponseWriter, r *http.Request) {
	tracks, err := service.DumpTracks()
	if err != nil {
		http.Error(w, fmt.Sprint(err), 500)
		return
	}
	data, err := json.Marshal(tracks)
	if err != nil {
		http.Error(w, fmt.Sprint(err), 500)
		return
	}
	w.Write(data)
}

type postTracksJSON struct {
	path string
}

func postTracks(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	data := make([]byte, r.ContentLength)
	_, err := r.Body.Read(data)
	if err != io.EOF && err != nil {
		return
	}
	var j postTracksJSON
	err = json.Unmarshal(data, j)

	path := j.path
	parser.Parse(path, service)
	elapsed := time.Since(start)
	w.Write([]byte(fmt.Sprintf("Execution time: %s", elapsed)))
}
