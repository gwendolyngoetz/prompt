package git

import "testing"

func TestWtIsRepo(t *testing.T) {
	got := IsRepo()
	want := true

	if got != want {
		t.Errorf("failed")
	}
}

func TestWtGetRepoName(t *testing.T) {
	got := GetRepoName()
	want := "prompt"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestWtGetRepoRelativePath(t *testing.T) {
	got := GetRepoRelativePath()
	want := ":pkg/git"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestWtHasStatusChanges(t *testing.T) {
	got := HasStatusChanges()
	want := true

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestWtHasStashes(t *testing.T) {
	got := HasStashes()
	want := true

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestWtHasChanges(t *testing.T) {
	got := HasChanges()
	want := true

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestWtGetBranchName(t *testing.T) {
	got := GetBranchName()
	want := "master"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
