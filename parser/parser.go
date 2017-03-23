package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Track holds parsed file informations
type Track struct {
	hash string
	path string
	tags map[string]string
	err  error
}

func parse(path string) (Track, error) {
	var track Track
	track.path = path
	file, err := os.Open(path)
	if err != nil {
		track.err = err
		return track, err
	}
	defer file.Close()

	track.tags, err = tags(file)
	if err != nil {
		track.err = err
		return track, err
	}

	track.hash, err = hash(file)
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
		fmt.Printf("%s: %s %s", track.path, track.hash, track.tags)
	}

	return nil
}

// Parse walks to a directiory and read all Musik data
func Parse(path string) error {
	start := time.Now()
	err := filepath.Walk(path, walker)
	elapsed := time.Since(start)

	fmt.Printf("Execution time: %s", elapsed)
	return err
}
