package abstract_factory

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetVehicleFactory(t *testing.T) {

	t.Run("error check", func(t *testing.T) {
		f := -1
		factory, err := GetVehicleFactory(f)
		require.NotNil(t, err)
		require.Nil(t, factory)
		require.EqualError(t, err, fmt.Sprintf("Factory with id %d not recognized", f))

		t.Run("car", func(t *testing.T) {
			factory, err = GetVehicleFactory(CarFactoryType)
			require.Nil(t, err)
			require.NotNil(t, factory)

			vehicle, err := factory.GetVehicle(f)
			require.NotNil(t, err)
			require.Nil(t, vehicle)
			require.EqualError(t, err, fmt.Sprintf("Vehicle of type %d not recognized", f))
		})

		t.Run("motorbike", func(t *testing.T) {
			factory, err = GetVehicleFactory(MotorbikeFactoryType)
			require.Nil(t, err)
			require.NotNil(t, factory)

			vehicle, err := factory.GetVehicle(f)
			require.NotNil(t, err)
			require.Nil(t, vehicle)
			require.EqualError(t, err, fmt.Sprintf("Vehicle of type %d not recognized", f))
		})
	})

	t.Run("car factory", func(t *testing.T) {

		factory, err := GetVehicleFactory(CarFactoryType)
		require.Nil(t, err)
		require.NotNil(t, factory)

		t.Run("luxury car", func(t *testing.T) {

			vehicle, err := factory.GetVehicle(LuxuryCarType)
			require.Nil(t, err)
			require.NotNil(t, vehicle)

			wheels := vehicle.NumWheels()
			require.Equal(t, wheels, 4)

			seats := vehicle.NumSeats()
			require.Equal(t, seats, 4)

			car, ok := vehicle.(Car)
			require.True(t, ok)
			require.NotNil(t, car)

			doors := car.NumDoors()
			require.Equal(t, doors, 5)
		})

		t.Run("family car", func(t *testing.T) {

			vehicle, err := factory.GetVehicle(FamilyCarType)
			require.Nil(t, err)
			require.NotNil(t, vehicle)

			wheels := vehicle.NumWheels()
			require.Equal(t, wheels, 4)

			seats := vehicle.NumSeats()
			require.Equal(t, seats, 7)

			car, ok := vehicle.(Car)
			require.True(t, ok)
			require.NotNil(t, car)

			doors := car.NumDoors()
			require.Equal(t, doors, 5)
		})
	})

	t.Run("motorbike factory", func(t *testing.T) {

		factory, err := GetVehicleFactory(MotorbikeFactoryType)
		require.Nil(t, err)
		require.NotNil(t, factory)

		t.Run("sport motorbike", func(t *testing.T) {

			vehicle, err := factory.GetVehicle(SportMotorbikeType)
			require.Nil(t, err)
			require.NotNil(t, vehicle)

			wheels := vehicle.NumWheels()
			require.Equal(t, wheels, 2)

			seats := vehicle.NumSeats()
			require.Equal(t, seats, 1)

			car, ok := vehicle.(Motorbike)
			require.True(t, ok)
			require.NotNil(t, car)

			motorbikeType := car.GetMotorbikeType()
			require.Equal(t, motorbikeType, SportMotorbikeType)
		})

		t.Run("cruise motorbike", func(t *testing.T) {

			vehicle, err := factory.GetVehicle(CruiseMotorbikeType)
			require.Nil(t, err)
			require.NotNil(t, vehicle)

			wheels := vehicle.NumWheels()
			require.Equal(t, wheels, 2)

			seats := vehicle.NumSeats()
			require.Equal(t, seats, 2)

			car, ok := vehicle.(Motorbike)
			require.True(t, ok)
			require.NotNil(t, car)

			motorbikeType := car.GetMotorbikeType()
			require.Equal(t, motorbikeType, CruiseMotorbikeType)
		})

	})
}
