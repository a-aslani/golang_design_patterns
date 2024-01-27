package bridge

import (
	"fmt"
)

type Radio struct {
	isEnabled bool
	volume    int
	channel   int
}

func (r *Radio) IsEnabled() bool {
	fmt.Println("IsEnabled")
	return r.isEnabled
}

func (r *Radio) Enable() {
	fmt.Println("Enable")
	r.isEnabled = true
}

func (r *Radio) Disable() {
	fmt.Println("Disable")
	r.isEnabled = false
}

func (r *Radio) Volume() int {
	fmt.Println("Volume")
	return r.volume
}

func (r *Radio) SetVolume(volume int) error {
	fmt.Println("SetVolume")

	if volume < 0 {
		return fmt.Errorf("invalid volume: %d", volume)
	}

	r.volume = volume
	return nil
}

func (r *Radio) Channel() int {
	fmt.Println("Channel")

	return r.channel
}

func (r *Radio) SetChannel(channel int) error {
	fmt.Println("SetChannel")

	if channel <= 0 {
		return fmt.Errorf("invalid channel: %d", channel)
	}

	r.channel = channel
	return nil
}
