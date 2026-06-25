package service

import "testing"

func TestEditDailyCalories(t *testing.T) {
	store := &FakeStore{user: newUser()} // starts at 2000
	goals := &GoalManager{Storage: store}

	if err := goals.EditDailyCalories(2500); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if store.user.Goal.DailyCalories != 2500 {
		t.Errorf("DailyCalories = %v, want 2500", store.user.Goal.DailyCalories)
	}
	if store.saveCount != 1 {
		t.Errorf("expected the change to be saved once, saveCount = %d", store.saveCount)
	}
}
