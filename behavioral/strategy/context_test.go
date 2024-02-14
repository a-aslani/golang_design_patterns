package strategy

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestContextExecute(t *testing.T) {

	t.Run("add operation", func(t *testing.T) {
		addContext := NewContext(&OperationAdd{})
		addResult := addContext.execute(2, 2)
		require.Equal(t, addResult, 4)
	})

	t.Run("substract operation", func(t *testing.T) {
		substractContext := NewContext(&OperationSubstract{})
		substractResult := substractContext.execute(2, 2)
		require.Equal(t, substractResult, 0)
	})

	t.Run("multiply operation", func(t *testing.T) {
		multiplyContext := NewContext(&OperationMultiply{})
		multiplyResult := multiplyContext.execute(2, 2)
		require.Equal(t, multiplyResult, 4)
	})

}
