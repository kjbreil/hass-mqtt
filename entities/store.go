package entities

import "sync"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type StateStore struct {
	Light struct {
		Mutex          *sync.Mutex
		Availability   map[string]string
		JsonAttributes map[string]string
		State          map[string]string
	}
}

func initStore() StateStore {
	s := StateStore{}
	s.Light.Availability = make(map[string]string)
	s.Light.JsonAttributes = make(map[string]string)
	s.Light.State = make(map[string]string)
	s.Light.Mutex = &sync.Mutex{}
	return s
}
