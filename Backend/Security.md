# Security and Privacy

The Things Network is a highly secure public network that supports true end-to-end encryption, mitigations against various man-in-the-middle attacks and support for different 128-bit encryption keys for every single end device.

[[LoRaWAN enforces using AES 128-bit|LoRaWAN/Security]] message integrity check and payload encryption. Payload is encrypted and decrypted in your domain only: on the end device and in the Handler (see [[components|Backend/Overview]]). You can choose to operate your own Handler to keep your keys private. The Router and Broker route data based on public metadata and cannot decrypt payload.

## User Accounts

Users need an account to use The Things Network. User accounts are identified by an e-mail address and protected by a password. In private networking setups, users may use their own authentication and authorization mechanism.

Use [ttnctl](ttnctl/QuickStart.md) to create a user account.

## Applications

Users can create applications and they can authorize other users access to applications. Applications are identified by a globally unique AppEUI, a 64-bit number issued by The Things Network Foundation. Each Application has one or more Access Keys to access application data.

Use [ttnctl](ttnctl/QuickStart.md) to manage applications.
