# hass-mqtt

Idea and generation code taken from https://github.com/W-Floyd/ha-mqtt-iot which is a very well done project but setup
for a singular use. I wanted to be able to just use Home Assistant Devices/Entities in other programs so extending what
was done there.

Massive work in progress right now.

## Generating Device Cache

To generate a new cache of Home Assistant MQTT device definitions:

1. From the project root directory, run: `go run ./helpers/`

This will:
- Fetch the latest device documentation from Home Assistant's GitHub repository
- Extract YAML configuration definitions for each device type  
- Cache both the raw markdown docs and processed YAML files in `./helpers/cache/`
- Generate Go code for entities based on the cached definitions

The cache includes device definitions for: alarm_control_panel, binary_sensor, button, camera, climate, cover, device_tracker, device_trigger, fan, humidifier, light, lock, number, scene, select, sensor, siren, switch, tag, text, update, and vacuum.

To force re-downloading of cached files, set `PullNew = true` in `devicepuller.go` before running.
