package bridge

type Device interface {
	IsEnabled() bool
	Enable()
	Disable()
	Volume() int
	SetVolume(volume int) error
	Channel() int
	SetChannel(channel int) error
}
