package main

import "fmt"

func main() {
	emulationInterNeuronalSpaceIonicConcentration := IonicConcentration{
		Sodium:    10,
		Potassium: 10,
		Chloride:  0,
		Calcium:   0,
	}

	neuronalSpace0IonicConcentration := IonicConcentration{
		Sodium:    10,
		Potassium: 10,
		Chloride:  0,
		Calcium:   0,
	}

	emulationInterNeuronalSpace := MembraneEnclosedSpace{
		Volume:     0,
		VolumeUnit: 0,
		HasNucleus: false,
		Gates: []Gate{
			{
				true,
			},
		},
		IonicConcentration: emulationInterNeuronalSpaceIonicConcentration,
	}

	emulationNeuronPopulation := []MembraneEnclosedSpace{
		{
			Volume:     0,
			VolumeUnit: 0,
			HasNucleus: false,
			Gates: []Gate{
				{
					true,
				},
			},
			IonicConcentration: neuronalSpace0IonicConcentration,
		},
	}

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

func (e *Emulation) InitiateEmulationRuntime() error {
	if e.MaxIteration != -1 {
		for i := 0; i < e.MaxIteration; i++ {
			fmt.Println(fmt.Sprintf("The state of the emulation at cycle number %d of the emulation runtime: is the Neuronal space posesses %d potassium ions and %d sodium ions, while the InterNeuronal space posesses %d potassium ions and %d sodium ions", e.CurrentIteration, e.Neurons[0].IonicConcentration.Potassium, e.Neurons[0].IonicConcentration.Sodium, e.InterNeuronalSpace.IonicConcentration.Potassium, e.InterNeuronalSpace.IonicConcentration.Sodium))
			e.CurrentIteration += 1

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

func (e *Emulation) RuntimeCycle() error {
	fmt.Println("initiating runtime cycle")
	// iterate over Emulated Neurons
	for i := 0; i < len(e.Neurons); i++ {
		fmt.Println("Iterating over neuron #%d of this emulations neurons, which ha %d gates", i, len(e.Neurons[i].Gates))
		// iterate over Current Neurons gates
		for j := 0; j < len(e.Neurons[i].Gates); j++ {
			fmt.Println("iterating over Neuronal Gates")
			err := e.Neurons[i].Gates[j].OperateGate(&e.Neurons[i], &e.InterNeuronalSpace)
			if err != nil {
				fmt.Println("failure to operate gate")
			} else {
				fmt.Println("gate successfully operated")
			}
		}
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

// how should we factor for the diffusion of solutes within MembraneEnclosed spaces?

// how will we factor graded potential?

// Membrane Potential(transmembrane potential or membrane voltage): is the difference in electrical potential between the interior and the exterior of the cell

// returns MembranePotential of the Neuron
func (ms *MembraneEnclosedSpace) MembranePotential() error {
	// how is membrane potential calculated? how will we calculate it for this simplified environment?
	return nil
}

// returns ActionPotential of the Neuron
func (ms *MembraneEnclosedSpace) ActionPotential() error {
	// how is action potential calculated? how will we calculate it for this simplified environment?
	return nil
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

func (g *Gate) OperateGate(interiorMembraneEnclosedSpace *MembraneEnclosedSpace, exteriorMembraneEnclosedSpace *MembraneEnclosedSpace) error {
	// logic for determining availability of ions(keys) for the gate
	fmt.Println("Operate gate invoked")
	if true {
		// sodium potassium pump logic

		if interiorMembraneEnclosedSpace.IonicConcentration.Sodium >= 3 && exteriorMembraneEnclosedSpace.IonicConcentration.Potassium >= 2 {
			fmt.Println("Gate is operable")
			// should we factor for ATP aquisition and consumption in the future?
			exteriorMembraneEnclosedSpace.IonicConcentration.Potassium -= 2
			interiorMembraneEnclosedSpace.IonicConcentration.Potassium += 2
			interiorMembraneEnclosedSpace.IonicConcentration.Sodium -= 3
			exteriorMembraneEnclosedSpace.IonicConcentration.Sodium += 3
		} else {

		}
	}

	return nil
}
