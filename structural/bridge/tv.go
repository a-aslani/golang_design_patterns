package bridge

import (
	"fmt"
)

type TV struct {
	isEnabled bool
	volume    int
	channel   int
}

func (r *TV) IsEnabled() bool {
	fmt.Println("IsEnabled")
	return r.isEnabled
}

func (r *TV) Enable() {
	fmt.Println("Enable")
	r.isEnabled = true
}

func (r *TV) Disable() {
	fmt.Println("Disable")
	r.isEnabled = false
}

func (r *TV) Volume() int {
	fmt.Println("Volume")
	return r.volume
}

func (r *TV) SetVolume(volume int) error {
	fmt.Println("SetVolume")

	if volume < 0 {
		return fmt.Errorf("invalid volume: %d", volume)
	}

	r.volume = volume
	return nil
}

func (r *TV) Channel() int {
	fmt.Println("Channel")

	return r.channel
}

func (r *TV) SetChannel(channel int) error {
	fmt.Println("SetChannel")

	if channel <= 0 {
		return fmt.Errorf("invalid channel: %d", channel)
	}

	r.channel = channel
	return nil
}
