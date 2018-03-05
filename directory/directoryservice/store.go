package directoryservice

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"go.uber.org/zap"
)

// In the absence of a database, just store services in-memory and save to JSON file.
// This file is poorly documented as it is just a temporary store for PoC.

// StoredService contains metadata about a service and a list of inways which provide this service.
type StoredService struct {
	OrganizationName string
	ServiceName      string

	// InwayAddresses maps address to health state
	InwayAddresses     map[string]bool
	InwayAddressesLock sync.RWMutex `json:"-"` // lock protects the inway addresses map

	DocumentationURL string
}

// NewStoredService creates a new StoredService with provided names set and no inways attached.
func NewStoredService(orgName, serviceName string) *StoredService {
	return &StoredService{
		OrganizationName: orgName,
		ServiceName:      serviceName,
		InwayAddresses:   make(map[string]bool),
	}
}

// CanonicalServiceName returns the canonical service name within the NLX network
func (s *StoredService) CanonicalServiceName() string {
	return s.OrganizationName + `.` + s.ServiceName
}

// Store holds all services available in the directory. This is a PoC structure which will be replaced with a proper database post-poc.
type Store struct {
	savePath string // path/to/file.json
	logger   *zap.Logger

	// Services stores the list of services this directory is aware of
	Services     map[string]*StoredService
	ServicesLock sync.RWMutex `json:"-"` // lock protects the services map
}

// NewStore creates a new store from file. When no store file is available a new file is created.
func NewStore(logger *zap.Logger, savePath string) *Store {
	store := &Store{
		savePath: savePath,
		logger:   logger,

		Services: make(map[string]*StoredService),
	}

	saveFile, err := os.Open(savePath)
	if err != nil {
		if os.IsNotExist(err) {
			store.Save()
			return store
		}
		log.Fatalf("failed to open exising save file: %v", err)
	}
	err = json.NewDecoder(saveFile).Decode(store)
	if err != nil {
		log.Fatalf("failed to decode store save contents: %v", err)
	}
	return store
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
