package entities

import "sync"

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type StateStore struct {
	AlarmControlPanel struct {
		Mutex          *sync.Mutex
		Availability   map[string]string
		JsonAttributes map[string]string
		State          map[string]string
	}
	BinarySensor struct {
		Mutex          *sync.Mutex
		Availability   map[string]string
		JsonAttributes map[string]string
		State          map[string]string
	}
	Button struct {
		Mutex          *sync.Mutex
		Availability   map[string]string
		JsonAttributes map[string]string
	}
	Camera struct {
		Mutex          *sync.Mutex
		Availability   map[string]string
		JsonAttributes map[string]string
		State          map[string]string
	}
	Climate struct {
		Mutex                    *sync.Mutex
		Action                   map[string]string
		Availability             map[string]string
		CurrentHumidity          map[string]string
		CurrentTemperature       map[string]string
		FanModeState             map[string]string
		JsonAttributes           map[string]string
		ModeState                map[string]string
		PresetModeState          map[string]string
		SwingHorizontalModeState map[string]string
		SwingModeState           map[string]string
		TargetHumidityState      map[string]string
		TemperatureHighState     map[string]string
		TemperatureLowState      map[string]string
		TemperatureState         map[string]string
	}
	Cover struct {
		Mutex          *sync.Mutex
		Availability   map[string]string
		JsonAttributes map[string]string
		Position       map[string]string
		State          map[string]string
		TiltStatus     map[string]string
	}
	DeviceTracker struct {
		Mutex          *sync.Mutex
		Availability   map[string]string
		JsonAttributes map[string]string
		State          map[string]string
	}
	DeviceTrigger struct {
		Mutex *sync.Mutex
		State map[string]string
	}
	Fan struct {
		Mutex            *sync.Mutex
		Availability     map[string]string
		DirectionState   map[string]string
		JsonAttributes   map[string]string
		OscillationState map[string]string
		PercentageState  map[string]string
		PresetModeState  map[string]string
		State            map[string]string
	}
	Humidifier struct {
		Mutex               *sync.Mutex
		Action              map[string]string
		Availability        map[string]string
		CurrentHumidity     map[string]string
		JsonAttributes      map[string]string
		ModeState           map[string]string
		State               map[string]string
		TargetHumidityState map[string]string
	}
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
	Lock struct {
		Mutex          *sync.Mutex
		Availability   map[string]string
		JsonAttributes map[string]string
		State          map[string]string
	}
	Number struct {
		Mutex          *sync.Mutex
		Availability   map[string]string
		JsonAttributes map[string]string
		State          map[string]string
	}
	Scene struct {
		Mutex          *sync.Mutex
		Availability   map[string]string
		JsonAttributes map[string]string
	}
	Select struct {
		Mutex          *sync.Mutex
		Availability   map[string]string
		JsonAttributes map[string]string
		State          map[string]string
	}
	Sensor struct {
		Mutex          *sync.Mutex
		Availability   map[string]string
		JsonAttributes map[string]string
		State          map[string]string
	}
	Siren struct {
		Mutex          *sync.Mutex
		Availability   map[string]string
		JsonAttributes map[string]string
		State          map[string]string
	}
	Switch struct {
		Mutex          *sync.Mutex
		Availability   map[string]string
		JsonAttributes map[string]string
		State          map[string]string
	}
	Tag struct {
		Mutex *sync.Mutex
		State map[string]string
	}
	Text struct {
		Mutex          *sync.Mutex
		Availability   map[string]string
		JsonAttributes map[string]string
		State          map[string]string
	}
	Update struct {
		Mutex          *sync.Mutex
		Availability   map[string]string
		JsonAttributes map[string]string
		State          map[string]string
	}
	Vacuum struct {
		Mutex          *sync.Mutex
		Availability   map[string]string
		JsonAttributes map[string]string
		State          map[string]string
	}
}

func initStore() StateStore {
	s := StateStore{}
	s.AlarmControlPanel.Availability = make(map[string]string)
	s.AlarmControlPanel.JsonAttributes = make(map[string]string)
	s.AlarmControlPanel.State = make(map[string]string)
	s.AlarmControlPanel.Mutex = &sync.Mutex{}
	s.BinarySensor.Availability = make(map[string]string)
	s.BinarySensor.JsonAttributes = make(map[string]string)
	s.BinarySensor.State = make(map[string]string)
	s.BinarySensor.Mutex = &sync.Mutex{}
	s.Button.Availability = make(map[string]string)
	s.Button.JsonAttributes = make(map[string]string)
	s.Button.Mutex = &sync.Mutex{}
	s.Camera.Availability = make(map[string]string)
	s.Camera.JsonAttributes = make(map[string]string)
	s.Camera.State = make(map[string]string)
	s.Camera.Mutex = &sync.Mutex{}
	s.Climate.Action = make(map[string]string)
	s.Climate.Availability = make(map[string]string)
	s.Climate.CurrentHumidity = make(map[string]string)
	s.Climate.CurrentTemperature = make(map[string]string)
	s.Climate.FanModeState = make(map[string]string)
	s.Climate.JsonAttributes = make(map[string]string)
	s.Climate.ModeState = make(map[string]string)
	s.Climate.PresetModeState = make(map[string]string)
	s.Climate.SwingHorizontalModeState = make(map[string]string)
	s.Climate.SwingModeState = make(map[string]string)
	s.Climate.TargetHumidityState = make(map[string]string)
	s.Climate.TemperatureHighState = make(map[string]string)
	s.Climate.TemperatureLowState = make(map[string]string)
	s.Climate.TemperatureState = make(map[string]string)
	s.Climate.Mutex = &sync.Mutex{}
	s.Cover.Availability = make(map[string]string)
	s.Cover.JsonAttributes = make(map[string]string)
	s.Cover.Position = make(map[string]string)
	s.Cover.State = make(map[string]string)
	s.Cover.TiltStatus = make(map[string]string)
	s.Cover.Mutex = &sync.Mutex{}
	s.DeviceTracker.Availability = make(map[string]string)
	s.DeviceTracker.JsonAttributes = make(map[string]string)
	s.DeviceTracker.State = make(map[string]string)
	s.DeviceTracker.Mutex = &sync.Mutex{}
	s.DeviceTrigger.State = make(map[string]string)
	s.DeviceTrigger.Mutex = &sync.Mutex{}
	s.Fan.Availability = make(map[string]string)
	s.Fan.DirectionState = make(map[string]string)
	s.Fan.JsonAttributes = make(map[string]string)
	s.Fan.OscillationState = make(map[string]string)
	s.Fan.PercentageState = make(map[string]string)
	s.Fan.PresetModeState = make(map[string]string)
	s.Fan.State = make(map[string]string)
	s.Fan.Mutex = &sync.Mutex{}
	s.Humidifier.Action = make(map[string]string)
	s.Humidifier.Availability = make(map[string]string)
	s.Humidifier.CurrentHumidity = make(map[string]string)
	s.Humidifier.JsonAttributes = make(map[string]string)
	s.Humidifier.ModeState = make(map[string]string)
	s.Humidifier.State = make(map[string]string)
	s.Humidifier.TargetHumidityState = make(map[string]string)
	s.Humidifier.Mutex = &sync.Mutex{}
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
	s.Lock.Availability = make(map[string]string)
	s.Lock.JsonAttributes = make(map[string]string)
	s.Lock.State = make(map[string]string)
	s.Lock.Mutex = &sync.Mutex{}
	s.Number.Availability = make(map[string]string)
	s.Number.JsonAttributes = make(map[string]string)
	s.Number.State = make(map[string]string)
	s.Number.Mutex = &sync.Mutex{}
	s.Scene.Availability = make(map[string]string)
	s.Scene.JsonAttributes = make(map[string]string)
	s.Scene.Mutex = &sync.Mutex{}
	s.Select.Availability = make(map[string]string)
	s.Select.JsonAttributes = make(map[string]string)
	s.Select.State = make(map[string]string)
	s.Select.Mutex = &sync.Mutex{}
	s.Sensor.Availability = make(map[string]string)
	s.Sensor.JsonAttributes = make(map[string]string)
	s.Sensor.State = make(map[string]string)
	s.Sensor.Mutex = &sync.Mutex{}
	s.Siren.Availability = make(map[string]string)
	s.Siren.JsonAttributes = make(map[string]string)
	s.Siren.State = make(map[string]string)
	s.Siren.Mutex = &sync.Mutex{}
	s.Switch.Availability = make(map[string]string)
	s.Switch.JsonAttributes = make(map[string]string)
	s.Switch.State = make(map[string]string)
	s.Switch.Mutex = &sync.Mutex{}
	s.Tag.State = make(map[string]string)
	s.Tag.Mutex = &sync.Mutex{}
	s.Text.Availability = make(map[string]string)
	s.Text.JsonAttributes = make(map[string]string)
	s.Text.State = make(map[string]string)
	s.Text.Mutex = &sync.Mutex{}
	s.Update.Availability = make(map[string]string)
	s.Update.JsonAttributes = make(map[string]string)
	s.Update.State = make(map[string]string)
	s.Update.Mutex = &sync.Mutex{}
	s.Vacuum.Availability = make(map[string]string)
	s.Vacuum.JsonAttributes = make(map[string]string)
	s.Vacuum.State = make(map[string]string)
	s.Vacuum.Mutex = &sync.Mutex{}
	return s
}
