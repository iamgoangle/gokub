package required

import (
	"fmt"
	"os"
	"testing"
)

func TestPanicRequired(t *testing.T) {
	defer handlePanic()
	GetEnv("UNIT_TEST")
}

func TestRetrieveValueFromRequiredEnv(t *testing.T) {
	_ = os.Setenv("UNIT_TEST", "test")

	got := GetEnv("UNIT_TEST")
	want := "test"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func handlePanic() {
	r := recover()
	fmt.Println("recover message", r)
}


