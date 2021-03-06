package puzh

import (
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	if err := It("test"); err != errNotInitialized {
		t.Errorf("should return error ErrNotInitialized")
	}

	Init("my-token")
	if err := It("test"); err != nil {
		t.Errorf("should return no error")
	}
}

func TestIt(t *testing.T) {
	t.Skip()
	Init("")
	It("test")
}

func ExampleItf() {
	Init("my-token")

	Itf("Hi, *%s*", os.Getenv("USER"))
}
