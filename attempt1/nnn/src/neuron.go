package nnn

const (
	NEURON_XOR = iota
	NEURON_AND = iota
	NEURON_OR  = iota
)

type Neuron struct {
	Inputs   []*Link
	Outputs  []*Link
	Firing   bool
	willFire bool
	Function int
	Life     Lifetime
}

func NewNeuron(function int) *Neuron {
	return &Neuron{[]*Link{}, []*Link{}, false, false, function, NewLifetime()}
}

func NewOrNeuron() *Neuron {
	return NewNeuron(NEURON_OR)
}

func NewAndNeuron() *Neuron {
	return NewNeuron(NEURON_AND)
}

func NewXorNeuron() *Neuron {
	return NewNeuron(NEURON_XOR)
}

// InputCount returns the number of neurons which are actively firing.
func (self *Neuron) InputCount() uint {
	var count uint = 0
	for _, link := range self.Inputs {
		if link.Sender.Firing {
			count++
		}
	}
	return count
}

func (self *Neuron) NextCycle() bool {
	switch self.Function {
	case NEURON_XOR:
		return self.InputCount()%2 != 0
	case NEURON_AND:
		for _, link := range self.Inputs {
			if !link.Sender.Firing {
				return false
			}
		}
		return true
	case NEURON_OR:
		for _, link := range self.Inputs {
			if link.Sender.Firing {
				return true
			}
		}
		return false
	}
	return false
}

func (self *Neuron) RemoveLinks() {
	for len(self.Inputs) > 0 {
		self.Inputs[0].Remove()
	}
	for len(self.Outputs) > 0 {
		self.Outputs[0].Remove()
	}
}
