package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/acosio14/calories-tracker-cli/domain"
)

type JSONStorage struct{}

func OutputFolderPath() (*string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("error finding home dir %w", err)
	}
	outputFolder := filepath.Join(homeDir, "calories-tracker-output")

	return &outputFolder, nil
}

func userExists(outputFilePath string) error {

	_, err := os.Stat(outputFilePath)
	if os.IsNotExist(err) { //if err doesn't exist, then file exist, then user file already exists
		return fmt.Errorf("user file already exist: %s\n", outputFilePath)
	}

	return nil
}

func (s *JSONStorage) LoadUser() (*domain.User, error) {
	outputFolder, err := OutputFolderPath()
	if err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(*outputFolder)
	if err != nil {
		return nil, fmt.Errorf("error with reading output directory %w", err)
	}

	var user domain.User
	if len(entries) < 1 {
		return nil, fmt.Errorf("user not created")
	} else {
		jsonFile := filepath.Join(*outputFolder, string(entries[0].Name()))
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

	outputFolder, err := OutputFolderPath()
	if err != nil {
		return err
	}

	err = os.MkdirAll(*outputFolder, 0755)
	if err != nil {
		return fmt.Errorf("error creating output folder %w", err)
	}

	userData, err := json.MarshalIndent(user, "", "	")
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%s.json", user.Name)
	outputFilePath := filepath.Join(*outputFolder, filename)
	err = os.WriteFile(outputFilePath, userData, 0644)
	if err != nil {
		return err
	}
	return nil
}
