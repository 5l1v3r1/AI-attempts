package nnn

import "math/rand"

// Generates a neuron with a random type.
// The new neuron will not be marked "permanent"
func RandomNeuron() *Neuron {
	// TODO: see if I should weigh this more heavily towards OR neurons...
	var result *Neuron
	switch rand.Intn(3) {
	case 0:
		result = NewOrNeuron()
	case 1:
		result = NewAndNeuron()
	default:
		result = NewXorNeuron()
	}
	result.Life.Permanent = false
	return result
}

// Adds a random neuron to the network with one or two inputs and one output.
// Both the inputs and the output will be chosen using WeightedChoose with the
// specified linkWeight function. The created links will not be permanent.
func Evolve(network *Network, linkWeight WeightFunc) *Neuron {
	if len(network.Neurons) < 2 {
		return nil
	}

	neuron := RandomNeuron()
	if len(network.Neurons) == 2 || rand.Intn(2) == 0 {
		// One input, one output
		conns := WeightedChoose(network, 2, linkWeight)
		NewLink(conns[0], neuron).Life.Permanent = false
		NewLink(neuron, conns[1]).Life.Permanent = false
	} else {
		// Two inputs, one output
		conns := WeightedChoose(network, 3, linkWeight)
		NewLink(conns[0], neuron).Life.Permanent = false
		NewLink(conns[1], neuron).Life.Permanent = false
		NewLink(neuron, conns[2]).Life.Permanent = false
	}
	network.AddNeuron(neuron)
	return neuron
}
