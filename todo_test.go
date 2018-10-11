package main

import "testing"

func createTestTodo(title string) (*todoModel, error) {
	tm := &todoModel{
		Title:     title,
		Completed: 0,
	}
	return clientCreateTodo(tm)
}

var theID uint

func TestCreateTodo(t *testing.T) {
	t.Logf("Start TestCreateTodo")
	// Step 1: Create greeting
	tm, err := createTestTodo("test")
	if err != nil {
		t.Fatal(err)
	}
	theID = tm.ID
	t.Logf("Created todo %d", theID)
}

func TestGetTodo(t *testing.T) {
	t.Logf("Start TestGetTodo with ID: %d", theID)
	// Step 1: Get greeting
	tm, err := clientGetTodo(theID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Got todo %d", tm.ID)
}

func TestUpdateTodo(t *testing.T) {
	t.Logf("Start TestUpdateTodo with ID: %d", theID)
	// Step 1: First get todo
	tm, err := clientGetTodo(theID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Got todo with title: %s", tm.Title)

	// Step 2: Then update post with new title
	tm.Title = "Updated title"
	err = clientUpdateTodo(tm)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Updated todo with ID: %d", tm.ID)

	// Step 3: Finally, read post once more and check title
	tm2, err := clientGetTodo(theID)
	if err != nil {
		t.Fatal(err)
	}
	if tm.Title != tm2.Title {
		t.Errorf("Expected title %s but got %s", tm.Title, tm2.Title)
	}
	t.Logf("Got todo with title: %s", tm2.Title)
}

func TestDeleteTodo(t *testing.T) {
	t.Logf("Start TestDeleteTodo with ID: %d", theID)
	// Step 1: As a lst step: delete
	err := clientDeleteTodo(theID)
	if err != nil {
		t.Fatal(err)
	}
}
