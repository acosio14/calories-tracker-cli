package service

import (
	"io"
	"os"
	"testing"
	"time"

	"github.com/acosio14/calories-tracker-cli/domain"
)

// FakeStore is an in-memory Storage used by the service tests. It satisfies the
// Storage interface without touching the filesystem, so tests are fast and
// isolated. Set loadErr / saveErr to exercise the error paths.
type FakeStore struct {
	user      *domain.User
	loadErr   error
	saveErr   error
	saveCount int
}

func (s *FakeStore) LoadUser() (*domain.User, error) {
	if s.loadErr != nil {
		return nil, s.loadErr
	}
	return s.user, nil
}

func (s *FakeStore) SaveUser(u *domain.User) error {
	if s.saveErr != nil {
		return s.saveErr
	}
	s.user = u
	s.saveCount++
	return nil
}

// newUser returns a minimal valid user: 2000 daily calories, one seed weight
// entry, and an empty food journal. Tests customise it as needed.
func newUser() *domain.User {
	return &domain.User{
		Name:          "test",
		Goal:          domain.Goal{DailyCalories: 2000},
		WeightTracker: []domain.Weight{{ID: 0, Date: time.Now(), Value: 180}},
		FoodJournal:   []domain.FoodItem{},
	}
}

// ptr returns a pointer to v. Handy for building FoodItemInput, whose fields are
// pointers ("nil means not provided").
func ptr[T any](v T) *T { return &v }

// captureStdout runs f and returns whatever it wrote to os.Stdout. Used to test
// the View commands, which print rather than return their result.
func captureStdout(t *testing.T, f func()) string {
	t.Helper()
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("os.Pipe: %v", err)
	}
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old
	out, _ := io.ReadAll(r)
	return string(out)
}
