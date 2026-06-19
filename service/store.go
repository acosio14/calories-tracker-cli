package service

import "github.com/acosio14/calories-tracker-cli/domain"

type StorageInterface interface {
	LoadUser() (*domain.User, error)
	SaveUser(*domain.User) error
}
