# LoRaWAN

LoRaWAN (Long Range Wide Area Network) is a media access control (MAC) protocol for long range, low power networks (LPWAN) based on LoRa (Long Range) radio modulation technique. In the OSI Reference Model, LoRa represent layer 1 (PHY), and LoRaWAN roughly maps to layers 2 and 3.

LoRaWAN networks use a star-of-stars topology in which gateways forward messages between nodes and applications in the backend. Nodes use a single-hop wireless connection to one or more gateways.

[[/uploads/TTN-Overview.jpg]]

Communication between nodes and gateways is spread out on different [[channels|LoRaWAN/Channels]] and data rates. The [[frequencies|LoRaWAN/Frequencies]] that LoRaWAN uses, depend the geographic region and its radio spectrum regulations. Selection of the data rate is a trade-off between communication range and message duration.

The LoRaWAN standard is steered by the [LoRa Alliance](https://www.lora-alliance.org/).

LoRaWAN defines three classes of nodes:
  - **Class A** nodes allow bi-directional communications, but downlink messages can only be transmitted only after an uplink transmission by the node. After sending an uplink message, the node opens 2 receive windows for downlink. This class is the lowest power devices and currently the only class of devices supported by The Things Network.
  - **Class B** nodes allow bi-directional communication with scheduled receive slots.
  - **Class C** nodes with nearly continuously open receive windows, only closed when transmitting.

## Addressing

Devices and applications have a 64 bit unique identifier (`DevEUI` and `AppEUI`). When a device joins the network, it receives a dynamic (non-unique) 32-bit address (`DevAddr`). See the [[Addressing|LoRaWAN/Address-Space]] page for more information about the address space and device activation procedure in LoRaWAN.

## Security

Message payloads are end-to-end encrypted (128 bit AES) with an 128 bit **application session key** (`AppSKey`) which is unique per device, per session. The network uses a similar (but different) 128 bit **network session key** (`NwkSKey`) to distinguish devices with the same address (`DevAddr`). Each time a device re-joins the network, these keys are re-generated. Finally, LoRaWAN uses frame counters to prevent replay attacks. See the [[Security|LoRaWAN/Security]] page for more information about security in LoRaWAN.

## External Resources

* [LoRaWAN Specification 1.0 (LoRa Alliance)](https://www.lora-alliance.org/portals/0/specs/LoRaWAN%20Specification%201R0.pdf)
* [LoRa Modulation Basics (Semtech)](http://www.semtech.com/images/datasheet/an1200.22.pdf)
* [LoRa Modem Design Guide (Semtech)](http://www.semtech.com/images/datasheet/LoraDesignGuide_STD.pdf)
* [LoRa FAQs (Semtech)](http://www.semtech.com/wireless-rf/lora/LoRa-FAQs.pdf)
