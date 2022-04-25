package lyrics

import (
	"fmt"
	"github.com/goccy/go-yaml"
	"github.com/k4fer74/promptee/pkg/log"
	"io/ioutil"
	"os"
	"path"
	"sync"
)

type Store struct {
	m     sync.Mutex
	Songs []*Song
}

func NewStore(dirPath string) (Registry, error) {
	f, err := os.Open(dirPath)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	s := &Store{
		Songs: []*Song{},
	}
	files, err := f.Readdir(-1)
	chErr := make(chan error, len(files))
	var wg sync.WaitGroup
	for _, file := range files {
		wg.Add(1)
		go s.parseSpecFile(path.Join(dirPath, file.Name()), &wg, chErr)
	}
	wg.Wait()
	close(chErr)
	for range chErr {
		log.Logger.Error(<-chErr)
	}
	return s, nil
}

func (s *Store) parseSpecFile(fileName string, wg *sync.WaitGroup, errCh chan error) {
	defer wg.Done()
	f, err := os.Open(fileName)
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		errCh <- fmt.Errorf("cannot read file: %v", err)
		return
	}
	var song Song
	if err = yaml.Unmarshal(b, &song); err != nil {
		errCh <- fmt.Errorf("cannot unmarshal file: %v", err)
		return
	}
	s.m.Lock()
	song.ID = len(s.Songs) + 1
	s.Songs = append(s.Songs, &song)
	s.m.Unlock()
}

func (s *Store) GetSongs() ([]*Song, error) {
	return s.Songs, nil
}

func (s *Store) GetSong(id int) (*Song, error) {
	for _, s := range s.Songs {
		if s.ID == id {
			return s, nil
		}
	}
	return nil, ErrSongNotFound
}
