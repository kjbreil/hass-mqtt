package entities

import "sync"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type StateStore struct {
	Climate struct {
		Mutex                *sync.Mutex
		Action               map[string]string
		AuxState             map[string]string
		Availability         map[string]string
		CurrentHumidity      map[string]string
		CurrentTemperature   map[string]string
		FanModeState         map[string]string
		JsonAttributes       map[string]string
		ModeState            map[string]string
		PresetModeState      map[string]string
		SwingModeState       map[string]string
		TargetHumidityState  map[string]string
		TemperatureHighState map[string]string
		TemperatureLowState  map[string]string
		TemperatureState     map[string]string
	}
}

func initStore() StateStore {
	s := StateStore{}
	s.Climate.Action = make(map[string]string)
	s.Climate.AuxState = make(map[string]string)
	s.Climate.Availability = make(map[string]string)
	s.Climate.CurrentHumidity = make(map[string]string)
	s.Climate.CurrentTemperature = make(map[string]string)
	s.Climate.FanModeState = make(map[string]string)
	s.Climate.JsonAttributes = make(map[string]string)
	s.Climate.ModeState = make(map[string]string)
	s.Climate.PresetModeState = make(map[string]string)
	s.Climate.SwingModeState = make(map[string]string)
	s.Climate.TargetHumidityState = make(map[string]string)
	s.Climate.TemperatureHighState = make(map[string]string)
	s.Climate.TemperatureLowState = make(map[string]string)
	s.Climate.TemperatureState = make(map[string]string)
	s.Climate.Mutex = &sync.Mutex{}
	return s
}
