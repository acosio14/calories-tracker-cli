package service

import (
	"testing"
	"time"

	"github.com/acosio14/calories-tracker-cli/domain"
)

func TestAddWeight_Success(t *testing.T) {
	store := &FakeStore{user: newUser()} // seeded with one weight, ID 0
	tracker := &WeightTracker{Storage: store}

	date := time.Date(2026, 2, 1, 0, 0, 0, 0, time.UTC)
	if err := tracker.AddWeight(185, date); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(store.user.WeightTracker) != 2 {
		t.Fatalf("expected 2 weight entries, got %d", len(store.user.WeightTracker))
	}
	last := store.user.WeightTracker[1]
	if last.ID != 1 {
		t.Errorf("new ID = %d, want 1 (max+1)", last.ID)
	}
	if last.Value != 185 {
		t.Errorf("Value = %v, want 185", last.Value)
	}
	if !last.Date.Equal(date) {
		t.Errorf("Date = %v, want %v (flag must be honoured)", last.Date, date)
	}
}

func TestAddWeight_EmptyTracker(t *testing.T) {
	u := newUser()
	u.WeightTracker = []domain.Weight{} // empty
	store := &FakeStore{user: u}
	tracker := &WeightTracker{Storage: store}

	if err := tracker.AddWeight(200, time.Now()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(store.user.WeightTracker) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(store.user.WeightTracker))
	}
	if store.user.WeightTracker[0].ID != 0 {
		t.Errorf("first ID = %d, want 0", store.user.WeightTracker[0].ID)
	}
}

func TestDeleteWeightEntry_Success(t *testing.T) {
	u := newUser()
	u.WeightTracker = []domain.Weight{{ID: 0, Value: 180}, {ID: 1, Value: 178}}
	store := &FakeStore{user: u}
	tracker := &WeightTracker{Storage: store}

	if err := tracker.DeleteWeightEntry(0); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(store.user.WeightTracker) != 1 {
		t.Fatalf("expected 1 entry left, got %d", len(store.user.WeightTracker))
	}
	if store.user.WeightTracker[0].ID != 1 {
		t.Errorf("remaining ID = %d, want 1", store.user.WeightTracker[0].ID)
	}
}

func TestDeleteWeightEntry_NotFound(t *testing.T) {
	u := newUser()
	u.WeightTracker = []domain.Weight{{ID: 0, Value: 180}}
	store := &FakeStore{user: u}
	tracker := &WeightTracker{Storage: store}

	err := tracker.DeleteWeightEntry(999)
	if err == nil {
		t.Fatal("expected an error for unknown id, got nil")
	}
	if len(store.user.WeightTracker) != 1 {
		t.Errorf("nothing should have been deleted, got %d entries", len(store.user.WeightTracker))
	}
}
