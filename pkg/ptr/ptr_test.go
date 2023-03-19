package ptr_test

import (
	"testing"

	"github.com/tirathawat/covid/pkg/ptr"
)

func TestAddr(t *testing.T) {
	t.Run("Should return pointer to the value when the value is int", func(t *testing.T) {
		i := 1
		actual := ptr.Addr(i)

		if *actual != 1 {
			t.Errorf("expected %v, got %v", 1, *actual)
		}

	})

	t.Run("Should return pointer to the value when the value is string", func(t *testing.T) {
		s := "hello"
		actual := ptr.Addr(s)

		if *actual != "hello" {
			t.Errorf("expected %v, got %v", "hello", *actual)
		}
	})
}
