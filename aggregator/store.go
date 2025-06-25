package main

import "github.com/vsespontanno/calculate-toll/types"

type MemoryStore struct {
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}

func (s *MemoryStore) Insert(d types.Distance) error {
	return nil
}
