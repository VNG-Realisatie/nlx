package directoryservice

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

var (
	store = NewStore("./directory-store.json")
)

// In the absence of a database, just store services in-memory and save to JSON file.
// This file is poorly documented as it is just a temporary store for PoC.

type StoredService struct {
	OrganizationName string
	ServiceName      string
	InwayAddresses   map[string]bool
}

func NewStoredService(orgName, serviceName string) *StoredService {
	return &StoredService{
		OrganizationName: orgName,
		ServiceName:      serviceName,
		InwayAddresses:   make(map[string]bool),
	}
}

type Store struct {
	savePath string // path/to/file.json

	ServicesLock sync.RWMutex `json:"-"` // lock protects the services map
	Services     map[string]*StoredService
}

func NewStore(savePath string) *Store {
	saveFile, err := os.Open(savePath)
	if err != nil {
		if os.IsNotExist(err) {
			s := &Store{
				savePath: savePath,
				Services: make(map[string]*StoredService),
			}
			s.Save()
			return s
		}
		log.Fatalf("failed to open exising save file: %v", err)
	}
	s := &Store{
		savePath: savePath,
	}
	err = json.NewDecoder(saveFile).Decode(s)
	if err != nil {
		log.Fatalf("failed to decode store save contents: %v", err)
	}
	return s
}

// Save can be calld on a store to save it to disk.
// Note that the Save() call itself acuires a write-lock, so the caller must not keep any lock.
func (s *Store) Save() {
	s.ServicesLock.Lock()
	defer s.ServicesLock.Unlock()
	s.save()
}

// SaveAndUnlock can be called on a write-locked store and saves then unlocks the store.
func (s *Store) SaveAndUnlock() {
	defer s.ServicesLock.Unlock()
	s.save()
}

func (s *Store) save() {
	newsavePath := s.savePath + ".new"

	newsaveFile, err := os.Create(newsavePath)
	if err != nil {
		log.Fatalf("failed to create new store save file: %v", err)
	}

	err = json.NewEncoder(newsaveFile).Encode(s)
	if err != nil {
		log.Fatalf("failed to encode json to file: %v", err)
	}

	err = os.Rename(newsavePath, s.savePath)
	if err != nil {
		log.Fatalf("error saving store: %v", err)
	}
}
