package semaphore

import (
	"testing"

	"golang.org/x/net/context"
)

func TestMemLockClientInit(t *testing.T) {
	for i, tt := range []struct {
		ee error
	}{
		{nil},
	} {
		_, got := NewMemLockClient(context.Background())
		if got != tt.ee {
			t.Errorf("case %d: unexpected error state initializing Client: got %v", i, got)
			continue
		}
	}
}
