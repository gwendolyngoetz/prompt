package computer

import "testing"

func TestGetUserName(t *testing.T) {
	got := GetUserName()
	want := "gwendolyn"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestGetHostname(t *testing.T) {
	got := GetHostname()
	want := "thalia"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestGetOsIcon(t *testing.T) {
	got := GetOsIcon()
	want := ""

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestGetPwd(t *testing.T) {
	got := GetPwd()
	want := "~/src/github/gwendolyngoetz/prompt/pkg/computer"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestIsNarrowWindow(t *testing.T) {
	got := IsNarrowWindow()
	want := true

	if got != want {
		t.Errorf("got '%t' want '%t'", got, want)
	}
}

func TestIsRemote(t *testing.T) {
	got := IsRemote()
	want := false

	if got != want {
		t.Errorf("got '%t' want '%t'", got, want)
	}
}

func TestIsSudo(t *testing.T) {
	got := IsSudo()
	want := false

	if got != want {
		t.Errorf("got '%t' want '%t'", got, want)
	}
}
