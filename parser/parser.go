package parser

import (
	"fmt"
	"os"
	"path/filepath"
)

// Track holds parsed file information's
type Track struct {
	Hash string
	Path string
	Tags map[string]string
	err  error
}

func parse(path string) (Track, error) {
	var track Track
	track.Path = path
	file, err := os.Open(path)
	if err != nil {
		track.err = err
		return track, err
	}
	defer file.Close()

	track.Tags, err = tags(file)
	if err != nil {
		track.err = err
		return track, err
	}

	track.Hash, err = hash(file)
	if err != nil {
		track.err = err
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

		track, err := parse(path)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s: %s %s", track.Path, track.Hash, track.Tags)
	}

	return nil
}

// Parse walks to a directory and read all Music data
func Parse(path string) error {
	return filepath.Walk(path, walker)
}
