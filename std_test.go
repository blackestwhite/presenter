package presenter_test

import (
	"testing"

	"github.com/blackestwhite/presenter"
)

func TestString(t *testing.T) {
	s := presenter.Std{
		Ok:     true,
		Result: "result 1",
	}
	want := `{"ok":true,"result":"result 1"}`
	got := s.String()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}

	s = presenter.Std{
		Ok:               false,
		ErrorCode:        1,
		ErrorDescription: "error 1",
	}
	want = `{"ok":false,"error_code":1,"description":"error 1"}`
	got = s.String()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}
