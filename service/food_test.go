package service

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/acosio14/calories-tracker-cli/domain"
)

func TestAddFoodItem_Success(t *testing.T) {
	store := &FakeStore{user: newUser()}
	tracker := &FoodTracker{Storage: store}

	err := tracker.AddFoodItem("oatmeal", 160, 35, 70, "breakfast")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(store.user.FoodJournal) != 1 {
		t.Fatalf("expected 1 food item, got %d", len(store.user.FoodJournal))
	}
	got := store.user.FoodJournal[0]
	if got.TotalCalories != 320 { // 160 * (70/35)
		t.Errorf("TotalCalories = %v, want 320", got.TotalCalories)
	}
	if got.ID != 0 {
		t.Errorf("ID = %d, want 0 for first item", got.ID)
	}
	if got.Name != "oatmeal" || got.Meal != "breakfast" {
		t.Errorf("Name/Meal = %q/%q, want oatmeal/breakfast", got.Name, got.Meal)
	}
}

func TestAddFoodItem_AssignsMaxIDPlusOne(t *testing.T) {
	// Journal already has an item with ID 5; the next ID must be 6 (max+1),
	// not 1 (len). This guards the ID-generation logic.
	u := newUser()
	u.FoodJournal = []domain.FoodItem{{ID: 5, Name: "eggs", ServingSize: 1, Quantity: 1}}
	store := &FakeStore{user: u}
	tracker := &FoodTracker{Storage: store}

	if err := tracker.AddFoodItem("toast", 100, 1, 1, "breakfast"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	last := store.user.FoodJournal[len(store.user.FoodJournal)-1]
	if last.ID != 6 {
		t.Errorf("new ID = %d, want 6", last.ID)
	}
}

func TestAddFoodItem_DefaultsMealByTime(t *testing.T) {
	store := &FakeStore{user: newUser()}
	tracker := &FoodTracker{Storage: store}

	if err := tracker.AddFoodItem("snack", 100, 50, 50, ""); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	meal := store.user.FoodJournal[0].Meal
	switch meal {
	case "breakfast", "lunch", "dinner":
		// ok
	default:
		t.Errorf("empty meal was not defaulted, got %q", meal)
	}
}

func TestAddFoodItem_ZeroServingSizeError(t *testing.T) {
	store := &FakeStore{user: newUser()}
	tracker := &FoodTracker{Storage: store}

	err := tracker.AddFoodItem("water", 0, 0, 1, "lunch")
	if err == nil {
		t.Fatal("expected an error for zero serving size, got nil")
	}
	if len(store.user.FoodJournal) != 0 {
		t.Errorf("nothing should have been added, got %d items", len(store.user.FoodJournal))
	}
}

func TestAddFoodItem_LoadErrorPropagates(t *testing.T) {
	wantErr := errors.New("boom")
	store := &FakeStore{loadErr: wantErr}
	tracker := &FoodTracker{Storage: store}

	err := tracker.AddFoodItem("oatmeal", 160, 35, 70, "breakfast")
	if !errors.Is(err, wantErr) {
		t.Errorf("error = %v, want it to wrap %v", err, wantErr)
	}
}

func TestTotalCalories(t *testing.T) {
	cases := []struct {
		name                    string
		cal, serving, qty, want float64
	}{
		{"double serving", 160, 35, 70, 320},
		{"single serving", 200, 100, 100, 200},
		{"half serving", 100, 50, 25, 50},
		{"quarter serving", 400, 100, 25, 100},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			store := &FakeStore{user: newUser()}
			tracker := &FoodTracker{Storage: store}

			if err := tracker.AddFoodItem("x", c.cal, c.serving, c.qty, "lunch"); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			got := store.user.FoodJournal[0].TotalCalories
			if got != c.want {
				t.Errorf("TotalCalories = %v, want %v", got, c.want)
			}
		})
	}
}

func TestEditFoodItem_UpdatesAndRecomputes(t *testing.T) {
	u := newUser()
	u.FoodJournal = []domain.FoodItem{{
		ID: 0, Name: "oatmeal", Meal: "breakfast",
		ServingSize: 35, CaloriesPerServing: 160, Quantity: 70, TotalCalories: 320,
	}}
	store := &FakeStore{user: u}
	tracker := &FoodTracker{Storage: store}

	// Change calories per serving to 200; total must recompute to 200*(70/35)=400.
	err := tracker.EditFoodItem(0, domain.FoodItemInput{Calories: ptr(200.0)})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := store.user.FoodJournal[0]
	if got.CaloriesPerServing != 200 {
		t.Errorf("CaloriesPerServing = %v, want 200", got.CaloriesPerServing)
	}
	if got.TotalCalories != 400 {
		t.Errorf("TotalCalories = %v, want 400 (should recompute)", got.TotalCalories)
	}
}

func TestEditFoodItem_PartialLeavesOtherFields(t *testing.T) {
	u := newUser()
	u.FoodJournal = []domain.FoodItem{{
		ID: 0, Name: "oatmeal", Meal: "breakfast",
		ServingSize: 35, CaloriesPerServing: 160, Quantity: 70, TotalCalories: 320,
	}}
	store := &FakeStore{user: u}
	tracker := &FoodTracker{Storage: store}

	// Only change the meal; name and quantity must be untouched.
	if err := tracker.EditFoodItem(0, domain.FoodItemInput{Meal: ptr("lunch")}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := store.user.FoodJournal[0]
	if got.Meal != "lunch" {
		t.Errorf("Meal = %q, want lunch", got.Meal)
	}
	if got.Name != "oatmeal" {
		t.Errorf("Name = %q, want oatmeal (should be unchanged)", got.Name)
	}
	if got.Quantity != 70 {
		t.Errorf("Quantity = %v, want 70 (should be unchanged)", got.Quantity)
	}
}

func TestEditFoodItem_NotFound(t *testing.T) {
	u := newUser()
	u.FoodJournal = []domain.FoodItem{{ID: 0, Name: "oatmeal", ServingSize: 1, Quantity: 1}}
	store := &FakeStore{user: u}
	tracker := &FoodTracker{Storage: store}

	err := tracker.EditFoodItem(999, domain.FoodItemInput{Meal: ptr("lunch")})
	if err == nil {
		t.Fatal("expected an error for unknown id, got nil")
	}
	if store.user.FoodJournal[0].Name != "oatmeal" {
		t.Error("journal should be unchanged on a not-found edit")
	}
}

func TestDeleteFoodItem_Success(t *testing.T) {
	u := newUser()
	u.FoodJournal = []domain.FoodItem{
		{ID: 0, Name: "eggs"},
		{ID: 1, Name: "toast"},
	}
	store := &FakeStore{user: u}
	tracker := &FoodTracker{Storage: store}

	if err := tracker.DeleteFoodItem(0); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(store.user.FoodJournal) != 1 {
		t.Fatalf("expected 1 item left, got %d", len(store.user.FoodJournal))
	}
	if store.user.FoodJournal[0].ID != 1 {
		t.Errorf("remaining item ID = %d, want 1", store.user.FoodJournal[0].ID)
	}
}

func TestDeleteFoodItem_NotFound(t *testing.T) {
	u := newUser()
	u.FoodJournal = []domain.FoodItem{{ID: 0, Name: "eggs"}}
	store := &FakeStore{user: u}
	tracker := &FoodTracker{Storage: store}

	err := tracker.DeleteFoodItem(999)
	if err == nil {
		t.Fatal("expected an error for unknown id, got nil")
	}
	if len(store.user.FoodJournal) != 1 {
		t.Errorf("nothing should have been deleted, got %d items", len(store.user.FoodJournal))
	}
}

func TestViewRemainingCalories(t *testing.T) {
	day := time.Date(2026, 1, 15, 8, 0, 0, 0, time.UTC)
	u := newUser() // DailyCalories: 2000
	u.FoodJournal = []domain.FoodItem{
		{ID: 0, Name: "oatmeal", Date: day, Quantity: 2, TotalCalories: 320},
		{ID: 1, Name: "old", Date: day.AddDate(0, 0, -1), Quantity: 1, TotalCalories: 999}, // different day
	}
	store := &FakeStore{user: u}
	tracker := &FoodTracker{Storage: store}

	// Same calendar day (different hour) -> only the 320 item counts: 2000-320=1680.
	out := captureStdout(t, func() {
		query := time.Date(2026, 1, 15, 20, 0, 0, 0, time.UTC)
		if err := tracker.ViewRemainingCalories(query); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	if !strings.Contains(out, "1680.00") {
		t.Errorf("output should report 1680.00 remaining, got:\n%s", out)
	}
	if strings.Contains(out, "999") {
		t.Errorf("item from a different day should be excluded, got:\n%s", out)
	}
}
