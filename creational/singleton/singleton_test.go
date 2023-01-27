package singleton

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCounter(t *testing.T) {

	t.Run("test instance of counter", func(t *testing.T) {

		counter1 := NewCounter()

		require.NotNil(t, counter1)

		counter2 := NewCounter()

		require.Equal(t, counter1, counter2)

	})

	t.Run("increment must be increment the count to 1", func(t *testing.T) {

		c := NewCounter()

		c.Increment()

		require.Equal(t, c.GetCount(), 1)

	})

	t.Run("decrement must be decrement the count to 0", func(t *testing.T) {

		c := NewCounter()

		c.Decrement()

		require.Equal(t, c.GetCount(), 0)
	})

}
