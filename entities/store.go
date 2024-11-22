package entities

import "sync"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type StateStore struct {
	Cover struct {
		Mutex          *sync.Mutex
		Availability   map[string]string
		JsonAttributes map[string]string
		Position       map[string]string
		State          map[string]string
		TiltStatus     map[string]string
	}
}

func initStore() StateStore {
	s := StateStore{}
	s.Cover.Availability = make(map[string]string)
	s.Cover.JsonAttributes = make(map[string]string)
	s.Cover.Position = make(map[string]string)
	s.Cover.State = make(map[string]string)
	s.Cover.TiltStatus = make(map[string]string)
	s.Cover.Mutex = &sync.Mutex{}
	return s
}
