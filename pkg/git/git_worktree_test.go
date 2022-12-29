package git

import "testing"

func TestIsRepo(t *testing.T) {
	got := IsRepo()
	want := true

	if got != want {
		t.Errorf("failed")
	}
}

func TestGetRepoName(t *testing.T) {
	got := GetRepoName()
	want := "prompt"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestGetRepoRelativePath(t *testing.T) {
	got := GetRepoRelativePath()
	want := ":pkg/git"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestHasStatusChanges(t *testing.T) {
	got := HasStatusChanges()
	want := true

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestHasStashes(t *testing.T) {
	got := HasStashes()
	want := true

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestHasChanges(t *testing.T) {
	got := HasChanges()
	want := true

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestGetBranchName(t *testing.T) {
	got := GetBranchName()
	want := "master"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
