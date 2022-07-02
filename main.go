package main

func main() {
	emulationInterNeuronalSpaceIonicConcentration := IonicConcentration{
		Sodium:    0,
		Potassium: 0,
		Chloride:  0,
		Calcium:   0,
	}

	emulationInterNeuronalSpace := MembraneEnclosedSpace{
		Volume:             0,
		VolumeUnit:         0,
		HasNucleus:         false,
		Gates:              []Gate{},
		IonicConcentration: emulationInterNeuronalSpaceIonicConcentration,
	}

	emulationNeuronPopulation := []MembraneEnclosedSpace{}

	newEmulation := Emulation{
		HasIterationLimit:     true,
		HasRootEmulation:      false,
		EmulationIsDistibuted: false,
		IsConcluded:           false,
		DateCreated:           "",
		CurrentIteration:      0,
		MaxIteration:          10,
		Neurons:               emulationNeuronPopulation,
		InterNeuronalSpace:    emulationInterNeuronalSpace,
	}

	err := newEmulation.InitiateEmulationRuntime()
	if err != nil {

	}
}

/*
	EmulationVolume is a quantity of
*/
type Emulation struct {
	HasIterationLimit     bool
	HasRootEmulation      bool
	EmulationIsDistibuted bool
	IsConcluded           bool
	DateCreated           string
	CurrentIteration      int
	MaxIteration          int
	Neurons               []MembraneEnclosedSpace
	InterNeuronalSpace    MembraneEnclosedSpace
}

func (e *Emulation) InitiateEmulationRuntime() (error) {
	if e.MaxIteration != -1 {
		for i := 0; i < e.MaxIteration; i++ {
			err := e.RuntimeCycle()
			if err != nil {
				break
			}
		}
	} else {
		for !e.IsConcluded {
			err := e.RuntimeCycle()
			if err != nil {
				break 
			}
		}
	}
	return nil
}

func (e *Emulation) RuntimeCycle() (error) {
	// iterate over Emulated Neurons 
	for i := 0; i < len(e.Neurons); i ++ {
		// iterate over Current Neurons gates 
		for j
	}
	return nil
}

type MembraneEnclosedSpace struct {
	Volume             int
	VolumeUnit         int
	HasNucleus         bool
	Gates              []Gate
	IonicConcentration IonicConcentration
}

type IonicConcentration struct {
	Potassium int
	Sodium    int
	Chloride  int
	Calcium   int
}

type Gate struct {
	TransportIsActive bool
}
