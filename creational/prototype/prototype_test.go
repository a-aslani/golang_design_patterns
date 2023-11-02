package prototype

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetShirtCloner(t *testing.T) {

	shirtCache := GetShirtCloner()
	require.NotNil(t, shirtCache)

	testCases := []struct {
		color     int
		prototype *Shirt
	}{
		{
			color:     White,
			prototype: whitePrototype,
		},
		{
			color:     Black,
			prototype: blackPrototype,
		},
		{
			color:     Blue,
			prototype: bluePrototype,
		},
	}

	for _, v := range testCases {

		clone, err := shirtCache.GetClone(v.color)
		require.Nil(t, err)
		require.NotNil(t, clone)
		require.Equal(t, clone, v.prototype)
		require.NotEmpty(t, clone.GetInfo())

		shirt := clone.(*Shirt)

		require.NotEqual(t, shirt.GetPrice(), 0)
	}

	clone, err := shirtCache.GetClone(-1)
	require.NotNil(t, err)
	require.Nil(t, clone)

}
