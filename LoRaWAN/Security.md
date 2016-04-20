# Security in LoRaWAN and TTN

## Security Keys

LoRaWAN specifies a number of security keys: `NwkSKey`, `AppSKey` and `AppKey`. All keys have a length of 128 bits.

The network session key (`NwkSKey`) is used for interaction between the Node and the Network. This key is used to check the validity of messages. In the backend of The Things Network this validation is also used to map a non-unique device address (`DevAddr`) to a unique `DevEUI` and `AppEUI` (see the [[Address Space|LoRaWAN/Address-Space]] page).

The application session key (`AppSKey`) is used for encryption and decryption of the payload. The payload is fully encrypted between the Node and the Handler component of The Things Network (which you will be able to run on your own server). This means that nobody except you is able to read the contents of messages you send or receive.

These two session keys (`NwkSKey` and `AppSKey`) are unique per device, per session. If you dynamically activate your device ([[OTAA|LoRaWAN/Address-Space]]), these keys are re-generated on every activation. If you statically activate your device ([[ABP|LoRaWAN/Address-Space]]), these keys stay the same until you change them.

Dynamically activated devices ([[OTAA|LoRaWAN/Address-Space]]) use the application key (`AppKey`) to derive the two session keys during the activation procedure. In The Things Network you can have a _default_ `AppKey` which will be used to activate all devices, or customize the `AppKey` per device.

## Frame Counters

> TODO: Explain about Replay Attacks and how frame counters are used to detect them

## Spread Spectrum

> TODO: Explain why spread-spectrum communication is good for security (and reliability)

## Metadata

> TODO: Explain what kind of information can not be encrypted
