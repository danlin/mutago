package backend

import (
	"log"

	"github.com/HouzuoGuo/tiedot/db"
	"github.com/danlin/mutago/parser"
)

// Backend service
type Service struct {
	database *db.DB
}

const trackFeed string = "Tracks"

func Open(path string) (*Service, error) {
	database, err := db.OpenDB(path)

	if err != nil {
		return nil, err
	}

	if database.Use(trackFeed) == nil {
		if err := database.Create(trackFeed); err != nil {
			return nil, err
		}
	} else {
		if err := database.Scrub(trackFeed); err != nil {
			return nil, err
		}
	}

	return &Service{database: database}, nil
}

func (s *Service) InsertTrack(track parser.Track) (int, error) {
	tracks := s.database.Use(trackFeed)
	docID, err := tracks.Insert(map[string]interface{}{
		"path": track.Path, "hash": track.Hash})
	if err != nil {
		return 0, err
	}
	return docID, nil
}

func (s *Service) DumpTracks() {
	tracks := s.database.Use(trackFeed)
	tracks.ForEachDoc(func(id int, doc []byte) (moveOn bool) {
		track, err := tracks.Read(id)
		if err != nil {
			log.Print(err)
			return true
		}
		log.Print(track)

		return true
	})
}

func (s *Service) Close() error {
	if err := s.database.Close(); err != nil {
		return err
	}
	return nil
}
