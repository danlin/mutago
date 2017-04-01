package parser

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/danlin/mutago/backend"
)

var (
	Service *backend.Service
)

func parse(path string) (backend.Track, error) {
	var track backend.Track
	track.Path = path
	file, err := os.Open(path)
	if err != nil {
		track.Error = err
		return track, err
	}
	defer file.Close()

	track.Tags, err = tags(file)
	if err != nil {
		track.Error = err
		return track, err
	}

	track.Hash, err = hash(file)
	if err != nil {
		track.Error = err
		return track, err
	}

	return track, nil
}

func walker(path string, fileInfo os.FileInfo, err error) error {
	if err != nil {
		panic(err)
	} else {
		if fileInfo.IsDir() {
			return nil
		}

		if Service.Exists(path) {
			return nil
		}

		track, err := parse(path)
		if err != nil {
			panic(err)
		}
		Service.InsertTrack(track)
		fmt.Printf("%s: %s %s\r\n", track.Path, track.Hash, track.Tags)
	}

	return nil
}

// Parse walks to a directory and read all Music data
func Parse(path string, service *backend.Service) error {
	Service = service
	return filepath.Walk(path, walker)
}
