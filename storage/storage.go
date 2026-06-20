package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/acosio14/calories-tracker-cli/domain"
)

// Should contain save() and load() functions that service uses
// Do Later

type JSONStorage struct{}

func (s *JSONStorage) LoadUser() (*domain.User, error) {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("error finding home dir %v", err)
	}
	outputFolder := filepath.Join(homeDir, "Projects/calories-tracker-cli/output")

	entries, err := os.ReadDir(outputFolder)
	if err != nil {
		return nil, fmt.Errorf("Error with reading output directory %v", err)
	}

	var user domain.User
	if len(entries) < 1 {
		return nil, fmt.Errorf("User not created")
	} else {
		jsonFile := filepath.Join(outputFolder, string(entries[0].Name()))
		jsonContent, err := os.ReadFile(jsonFile)
		if err != nil {
			return nil, fmt.Errorf("read user %q: %w", jsonFile, err)
		}

		err = json.Unmarshal(jsonContent, &user)
		if err != nil {
			return nil, fmt.Errorf("parse user %q: %w", jsonFile, err)
		}
	}

	return &user, nil
}

func (s *JSONStorage) SaveUser(user *domain.User) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error finding home dir %v", err)
	}
	outputFolder := filepath.Join(homeDir, "Projects/calories-tracker-cli/output")

	userData, err := json.MarshalIndent(user, "", "	")
	if err != nil {
		return err
	}
	filename := fmt.Sprintf("%s.json", user.Name)
	outputPath := filepath.Join(outputFolder, filename)
	err = os.WriteFile(outputPath, userData, 0644)
	if err != nil {
		return err
	}
	return nil
}
