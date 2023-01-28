package builder

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type ManufacturingDirector struct {
	builder BuildProcess
}

func (m *ManufacturingDirector) SetBuilder(b BuildProcess) {
	m.builder = b
}

func (m *ManufacturingDirector) Construct() {
	m.builder.SetSeats().SetWheels().SetStructure()
}

type CarBuilder struct {
	vehicleProduct VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.vehicleProduct.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.vehicleProduct.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.vehicleProduct.Structure = "Car"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.vehicleProduct
}

type MotorbikeBuilder struct {
	vehicleProduct VehicleProduct
}

func (m *MotorbikeBuilder) SetWheels() BuildProcess {
	m.vehicleProduct.Wheels = 2
	return m
}

func (m *MotorbikeBuilder) SetSeats() BuildProcess {
	m.vehicleProduct.Seats = 2
	return m
}

func (m *MotorbikeBuilder) SetStructure() BuildProcess {
	m.vehicleProduct.Structure = "Motorbike"
	return m
}

func (m *MotorbikeBuilder) GetVehicle() VehicleProduct {
	return m.vehicleProduct
}
