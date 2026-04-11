package store

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/SuryaKannan/eridian/internal/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Store struct {
	language    string
	eridianHome string
}

func checkDBExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("error checking file: %w", err)
	}
	return true, nil
}

func NewStore(language string) *Store {
	// we can assume this won't fail, as at startup this needs to pass anyway

	return &Store{
		language:    language,
		eridianHome: config.EridianHome(),
	}
}

func (s *Store) CreateLanguageDB() error {
	dbPath := filepath.Join(s.eridianHome, s.language+".db")

	dbExists, err := checkDBExists(dbPath)
	if err != nil {
		return fmt.Errorf("error checking language db: %w", err)
	}

	if dbExists {
		return fmt.Errorf("language %s already exists", s.language)
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error creating language db: %w", err)
	}

	err = db.AutoMigrate(&Entry{}, &Centroid{})
	if err != nil {
		return fmt.Errorf("error migrating language db: %w", err)
	}

	return nil
}

func (s *Store) DeleteLanguageDB() error {
	dbPath := filepath.Join(s.eridianHome, s.language+".db")

	err := os.Remove(dbPath)

	if err != nil {
		return fmt.Errorf("error deleting language db: %w", err)
	}

	return nil
}
