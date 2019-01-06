package puzh

import (
	"testing"
)

func TestInit(t *testing.T) {
	if err := It("test"); err != ErrNotInitialized {
		t.Errorf("should return error NotInitialized")
	}

	Init("token")
	if err := It("test"); err != nil {
		t.Errorf("should return no error")
	}
}

func TestIt(t *testing.T) {
	t.Skip()
	Init("")
	It("test")
}
