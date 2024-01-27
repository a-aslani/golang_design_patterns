package bridge

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTV(t *testing.T) {

	tv := TV{
		isEnabled: false,
		volume:    0,
		channel:   1,
	}

	remote := NewRemoteControl(&tv, 1)

	remote.TogglePower()
	require.Equal(t, remote.device.IsEnabled(), true)

	remote.TogglePower()
	require.Equal(t, remote.device.IsEnabled(), false)

	err := remote.VolumeUp()
	require.NoError(t, err)
	require.Equal(t, remote.device.Volume(), 1)

	err = remote.VolumeDown()
	require.NoError(t, err)
	require.Equal(t, remote.device.Volume(), 0)

	err = remote.VolumeDown()
	require.Error(t, err)

	err = remote.ChannelUp()
	require.NoError(t, err)
	require.Equal(t, remote.device.Channel(), 2)

	err = remote.ChannelDown()
	require.NoError(t, err)
	require.Equal(t, remote.device.Channel(), 1)

	err = remote.ChannelDown()
	require.Error(t, err)
}

func TestRadio(t *testing.T) {

	radio := Radio{
		isEnabled: false,
		volume:    0,
		channel:   1,
	}

	remote := NewRemoteControl(&radio, 1)

	remote.TogglePower()
	require.Equal(t, remote.device.IsEnabled(), true)

	remote.TogglePower()
	require.Equal(t, remote.device.IsEnabled(), false)

	err := remote.VolumeUp()
	require.NoError(t, err)
	require.Equal(t, remote.device.Volume(), 1)

	err = remote.VolumeDown()
	require.NoError(t, err)
	require.Equal(t, remote.device.Volume(), 0)

	err = remote.VolumeDown()
	require.Error(t, err)

	err = remote.ChannelUp()
	require.NoError(t, err)
	require.Equal(t, remote.device.Channel(), 2)

	err = remote.ChannelDown()
	require.NoError(t, err)
	require.Equal(t, remote.device.Channel(), 1)

	err = remote.ChannelDown()
	require.Error(t, err)
}
