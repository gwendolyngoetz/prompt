package color

import "testing"

func TestFormatFgColor(t *testing.T) {
	got := formatFgColor("1,2,3")
	want := "\033[38;2;1;2;3m"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestFormatFgColorBold(t *testing.T) {
	got := formatFgColorBold("1,2,3")
	want := "\033[1;38;2;1;2;3m"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestFormatBgColor(t *testing.T) {
	got := formatBgColor("1,2,3")
	want := "\033[48;2;1;2;3m"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestFormatRgb(t *testing.T) {
	got := formatRgb("1,2,3")
	want := "1;2;3"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestBuildFgColor(t *testing.T) {
	cb := ColorBuilder{}
	got := cb.SetFgColor("255,255,0").Build()
	want := "\033[38;2;255;255;0m\033[0m"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestBuildBgColor(t *testing.T) {
	cb := ColorBuilder{}
	got := cb.SetBgColor("255,255,0").Build()
	want := "\033[48;2;255;255;0m\033[0m"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestBuildText(t *testing.T) {
	cb := ColorBuilder{}
	got := cb.SetText("test").Build()
	want := "test"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestBuildFgBgText(t *testing.T) {
	cb := ColorBuilder{}
	got := cb.
		SetFgColor("255,255,0").
		SetBgColor("125,125,125").
		SetText("test").
		Build()

	want := "\033[48;2;125;125;125m\033[38;2;255;255;0mtest\033[0m\033[0m"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
