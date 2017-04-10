# Single Channel Gateways

A single-channel gateway is a LoRa device that acts as a gateway by forwarding LoRa packets to the network. As it's usually built using a SX1272/SX1276 instead of SX1301/SX1257, it's a lot cheaper than a full gateway, making it a favorite choice for people to start with LoRaWAN. However, single-channel gateways are extremely limited compared to a real gateway. Single-channel gateways often lead to undesired design choices for solutions that you might build and can hurt The Things Network as a whole.

## Compliance and Support

Single-Channel gateways are **not LoRaWAN compliant** and are currently **not supported** by The Things Network. 

## Channels and Spreading Factors

Single-channel gateways can only receive on **one channel** and **one spreading factor** at the same time, whereas a full gateway can receive on **eight channels** (some even ten) and **six spreading factors** at the same time. Some single-channel gateways can "hop" between frequencies and spreading factors to simulate a real gateway, but they still have less than 2% of the capacity of a real gateway.

## Downlink

Some single-channel gateways do not have downlink support. Downlink support is required for OTAA and ADR.
