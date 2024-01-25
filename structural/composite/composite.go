package composite

import "fmt"

type Composite struct {
	Animal
	Athlete
	Swimmer
}

type Dog struct{}

func (d *Dog) Eat() error {
	fmt.Println("Eating!")
	return nil
}

type SampleAthlete struct{}

func (s *SampleAthlete) Train() error {
	fmt.Println("Training!")
	return nil
}

type Shark struct{}

func (s *Shark) Swim() error {
	fmt.Println("Swimming!")
	return nil
}

func Run() error {

	composite := Composite{
		Animal:  &Dog{},
		Athlete: &SampleAthlete{},
		Swimmer: &Shark{},
	}

	if err := composite.Swim(); err != nil {
		return err
	}
	if err := composite.Train(); err != nil {
		return err
	}
	if err := composite.Eat(); err != nil {
		return err
	}

	return nil
}
