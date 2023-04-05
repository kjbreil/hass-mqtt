package entities

import "sync"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type StateStore struct {
	Light struct {
		Mutex           *sync.Mutex
		Availability    map[string]string
		BrightnessState map[string]string
		ColorModeState  map[string]string
		ColorTempState  map[string]string
		EffectState     map[string]string
		HsState         map[string]string
		JsonAttributes  map[string]string
		RgbState        map[string]string
		RgbwState       map[string]string
		RgbwwState      map[string]string
		State           map[string]string
		XyState         map[string]string
	}
}

func initStore() StateStore {
	s := StateStore{}
	s.Light.Availability = make(map[string]string)
	s.Light.BrightnessState = make(map[string]string)
	s.Light.ColorModeState = make(map[string]string)
	s.Light.ColorTempState = make(map[string]string)
	s.Light.EffectState = make(map[string]string)
	s.Light.HsState = make(map[string]string)
	s.Light.JsonAttributes = make(map[string]string)
	s.Light.RgbState = make(map[string]string)
	s.Light.RgbwState = make(map[string]string)
	s.Light.RgbwwState = make(map[string]string)
	s.Light.State = make(map[string]string)
	s.Light.XyState = make(map[string]string)
	s.Light.Mutex = &sync.Mutex{}
	return s
}
