package flyweight

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTeamFlyweightFactory_GetTeam(t *testing.T) {
	factory := NewTeamFactory()

	teamA1, err := factory.GetTeam(TeamA)
	require.NoError(t, err)
	require.NotNil(t, teamA1)

	teamA2, err := factory.GetTeam(TeamB)
	require.NoError(t, err)
	require.NotNil(t, teamA2)

	require.NotEqual(t, teamA1, teamA2)
	require.Equal(t, 2, factory.GetNumberOfObjects())

}

func Test_HighVolume(t *testing.T) {
	factory := NewTeamFactory()

	teams := make([]*Team, 500000*2)
	var err error
	for i := 0; i < 500000; i++ {
		teams[i], err = factory.GetTeam(TeamA)
		require.NoError(t, err)
	}

	for i := 500000; i < 2*500000; i++ {
		teams[i], err = factory.GetTeam(TeamB)
		require.NoError(t, err)
	}

	require.Equal(t, 2, factory.GetNumberOfObjects())

	for i := 0; i < 3; i++ {
		fmt.Printf("Pointer %d points to %p and is located in %p\n", i, teams[i], &teams[i])
	}
}
