# Address Space in LoRaWAN

LoRaWAN knows a number of identifiers for devices, applications and gateways.

* `DevEUI` - 64 bit end-device identifier, EUI-64 (unique)
* `DevAddr` - 32 bit device address (non-unique)
* `AppEUI` - 64 bit application identifier, EUI-64 (unique)
* `GatewayEUI` - 64 bit gateway identifier, EUI-64 (unique)

## Devices

LoRaWAN devices have a 64 bit unique identifier (`DevEUI`) that is assigned to the device by the chip manufacturer. However, all communication is done with a dynamic 32 bit device address (`DevAddr`) of which 7 bits are fixed for The Things Network, leaving 25 bits that can be assigned to individual devices, a procedure called **Activation**.

### Over-the-Air Activation (OTAA)

Over-the-Air Activation (OTAA) is the preferred and most secure way to connect with The Things Network. Devices perform a join-procedure with the network, during which a dynamic `DevAddr` is assigned and security keys are negotiated with the device.

### Activation by Personalization (ABP)

In some cases you might need to hardcode the `DevAddr` as well as the security keys in the device. This means activating a device by personalization (ABP). This strategy might seem simpler, because you skip the join procedure, but it has some downsides related to security. See the [[security page|LoRaWAN/Security]] for more information about LoRaWAN security.

## Applications

Applications in LoRaWAN and The Things Network have a 64 bit unique identifier (`AppEUI`). When you run the `ttnctl applications create` command, The Things Network's account server allocates an `AppEUI` from our MAC address block. This means that every `AppEUI` starts with `70B3D57ED`.

## Gateways

Although the `packet forwarder` configuration file suggests that one can just choose an EUI for the gateway, these also have a unique MAC address, which should also be used for identifying with The Things Network.

> TODO: Explain how you can find the Gateway MAC.
