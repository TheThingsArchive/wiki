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

### Device Address Assignment

The Things Network Foundation has received a 7-bit device address prefix from the LoRa Alliance. This means that all TTN device addresses will start with `0x26` or `0x27` (although addresses that start with these might also belong to other networks with the same prefix). Within TTN, we assign device address prefixes to "regions" (for example, device addresses in the `eu` region start with `0x2601`). Within a region, the NetworkServer is responsible for assigning device addresses. We are using prefixes here too for different device classes (for example, ABP devices in the `eu` region start with `0x26011`) or to shard devices over different servers. 

The NetworkServer assigns device addresses to devices (based on configuration). For ABP devices you have to request an address from the NetworkServer (the console or `ttnctl` will do this for you). For OTAA devices, the NetworkServer will assign an address when the device joins.

It's good to keep in mind that device addresses are not unique. We can (and probably will) give hundreds of devices the same address. Finding the actual device that belongs to that address is done by matching the cryptographic signature (MIC) of the message to a device in the database.

## Applications

Applications in LoRaWAN and The Things Network have a 64 bit unique identifier (`AppEUI`). When you run the `ttnctl applications create` command, The Things Network's account server allocates an `AppEUI` from our MAC address block. This means that every `AppEUI` starts with `70B3D57ED`.

## Gateways

Although the `packet forwarder` configuration file suggests that one can just choose an EUI for the gateway, these also have a unique MAC address, which should also be used for identifying with The Things Network.

> TODO: Explain how you can find the Gateway MAC.
