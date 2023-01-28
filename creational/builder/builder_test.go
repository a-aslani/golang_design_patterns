package builder

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBuilderPattern(t *testing.T) {

	manufacturingComplex := ManufacturingDirector{}

	t.Run("test a car builder", func(t *testing.T) {

		builder := &CarBuilder{}

		manufacturingComplex.SetBuilder(builder)
		manufacturingComplex.Construct()

		car := builder.GetVehicle()

		require.NotNil(t, car)
		require.Equal(t, car.Structure, "Car")
		require.Equal(t, car.Seats, 5)
		require.Equal(t, car.Wheels, 4)
	})

	t.Run("test a motorbike builder", func(t *testing.T) {

		builder := &MotorbikeBuilder{}

		manufacturingComplex.SetBuilder(builder)
		manufacturingComplex.Construct()

		motorbike := builder.GetVehicle()

		require.NotNil(t, motorbike)
		require.Equal(t, motorbike.Structure, "Motorbike")
		require.Equal(t, motorbike.Seats, 2)
		require.Equal(t, motorbike.Wheels, 2)
	})

}
