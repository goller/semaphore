package semaphore

import (
	"errors"

	"golang.org/x/net/context"
)

type MemLockClient struct {
	sem Semaphore
}

// Creates a new in-memory client. Useful for testing in place of etcd.
func NewMemLockClient(ctx context.Context) (*MemLockClient, error) {
	return &MemLockClient{}, nil
}

func (m *MemLockClient) Init(ctx context.Context) error {
	m.sem = Semaphore{
		Index:     0,
		Semaphore: 1,
		Max:       1,
		Holders:   nil,
	}
	return nil
}

// Returns internal semaphore member
func (m *MemLockClient) Get(ctx context.Context) (*Semaphore, error) {
	return &m.sem, nil
}

// Set sets a Semaphore in memory.
func (m *MemLockClient) Set(ctx context.Context, sem *Semaphore) error {
	if sem == nil {
		return errors.New("cannot set nil semaphore")
	}
	m.sem = *sem
	return nil
}
