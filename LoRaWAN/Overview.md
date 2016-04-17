# LoRaWAN

LoRaWAN (**Lo**ng **Ra**nge **W**ide **A**rea **N**etwork) is a media access control (MAC) protocol for long range, low power networks (LPWAN) based on LoRa (Long Range) radio modulation technique. In the OSI Reference Model, LoRa represent layer 1 (PHY), and LoRaWAN roughly maps to layers 2 and 3.

LoRaWAN networks use a star-of-stars topology in which gateways forward messages between nodes and applications in the backend. Nodes use a single-hop wireless connection to one or more gateways.

Communication between nodes and gateways is spread out on different [frequency channels](Frequencies) and data rates. The selection of the data rate is a trade-off between communication range and message duration. LoRaWAN data rates range from 0.3 kbps to 50 kbps.

The LoRaWAN standard is steered by the [LoRa Alliance](https://www.lora-alliance.org/).

The communication between gateway and nodes is [secured by several layer of encryption](Security):
 - Unique Network key (EUI64) and ensure security on network level
 - Unique Application key (EUI64) ensure end to end security on application level
 - Device specific key (EUI128)

LoRaWAN defines three classes of nodes:
 - Class A: Nodes allow bi-directional communications, but downlink messages can only be transmited after an uplink transmission by the node. After sending an uplink message, the node opens 2 receive windows for downlink. This class is the lowest power devices.
 - Class B: Nodes allow bi-directional communication with scheduled receive slots.
 - Class C: Nodes with nearly continuously open receive windows, only closed when transmitting.

## External Resources

* [LoRaWAN Specification 1.0 (LoRa Alliance)](https://www.lora-alliance.org/portals/0/specs/LoRaWAN%20Specification%201R0.pdf)
* [LoRa Modulation Basics (Semtech)](http://www.semtech.com/images/datasheet/an1200.22.pdf)
* [LoRa Modem Design Guide (Semtech)](http://www.semtech.com/images/datasheet/LoraDesignGuide_STD.pdf)
* [LoRa FAQs (Semtech)](http://www.semtech.com/wireless-rf/lora/LoRa-FAQs.pdf)
