package bridge

type RemoteControl struct {
	device     Device
	volumeStep int
}

func NewRemoteControl(device Device, volumeStep int) *RemoteControl {
	return &RemoteControl{
		device:     device,
		volumeStep: volumeStep,
	}
}

func (r *RemoteControl) TogglePower() {

	if r.device.IsEnabled() {
		r.device.Disable()
	} else {
		r.device.Enable()
	}

}

func (r *RemoteControl) VolumeUp() error {
	return r.device.SetVolume(r.device.Volume() + r.volumeStep)
}

func (r *RemoteControl) VolumeDown() error {
	return r.device.SetVolume(r.device.Volume() - r.volumeStep)
}

func (r *RemoteControl) ChannelUp() error {
	return r.device.SetChannel(r.device.Channel() + 1)
}

func (r *RemoteControl) ChannelDown() error {
	return r.device.SetChannel(r.device.Channel() - 1)
}
