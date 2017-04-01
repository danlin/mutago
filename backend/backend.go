package backend

import (
	"github.com/HouzuoGuo/tiedot/db"
)

// Track holds parsed file information's
type Track struct {
	Hash  string
	Path  string
	Tags  map[string]string
	Error error
}

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
		tracks := database.Use(trackFeed)
		if err := tracks.Index([]string{"path"}); err != nil {
			return nil, err
		}
	} else {
		if err := database.Scrub(trackFeed); err != nil {
			return nil, err
		}
	}

	return &Service{database: database}, nil
}

//func (s *Service) GetTrack(path string) (Track, error) {
//	tracks := s.database.Use(trackFeed)
//	query := map[string]interface{}{
//		"eq":    path,
//		"in":    []interface{}{"hash"},
//		"limit": 1,
//	}
//	// Evaluate the query
//	queryResult := make(map[int]struct{})
//	if err := db.EvalQuery(query, tracks, &queryResult); nil != err {
//		panic(err)
//	}
//	// Fetch the results
//	for id := range queryResult {
//		readBack, err := tracks.Read(id)
//		if nil != err {
//			panic(err)
//		}
//		fmt.Printf("Query returned document %v\n", readBack)
//	}
//}

func (s *Service) Exists(path string) bool {
	tracks := s.database.Use(trackFeed)
	query := map[string]interface{}{
		"eq":    path,
		"in":    []interface{}{"path"},
		"limit": 1,
	}
	// Evaluate the query
	queryResult := make(map[int]struct{})
	if err := db.EvalQuery(query, tracks, &queryResult); nil != err {
		panic(err)
	}
	return len(queryResult) > 0
}

func (s *Service) InsertTrack(track Track) (int, error) {
	tracks := s.database.Use(trackFeed)
	docID, err := tracks.Insert(map[string]interface{}{
		"path": track.Path, "hash": track.Hash, "tags": track.Tags})
	if err != nil {
		return 0, err
	}
	return docID, nil
}

func (s *Service) DumpTracks() (map[string]Track, error) {
	var result = make(map[string]Track)

	tracks := s.database.Use(trackFeed)
	tracks.ForEachDoc(func(id int, doc []byte) (moveOn bool) {
		track, err := tracks.Read(id)
		if err != nil {
			return false
		}

		path := track["path"].(string)
		rt := Track{Path: path}
		result[path] = rt

		return true
	})

	return result, nil
}

func (s *Service) Close() error {
	if err := s.database.Close(); err != nil {
		return err
	}
	return nil
}
