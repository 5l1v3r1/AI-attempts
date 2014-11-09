#include "neuron.hpp"
#include "link.hpp"

namespace nnn1 {

Neuron::~Neuron() {
}

Neuron::Neuron() {
}

unsigned int Neuron::CountMessages() {
  unsigned int count = 0;
  Link * link = firstInput;
  while (link) {
    if (link->GetSender().IsFiring()) {
      ++count;
    }
    link = link->GetReceiverNext();
  }
  return count;
}

}
