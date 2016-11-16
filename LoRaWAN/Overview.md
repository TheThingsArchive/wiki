# LoRaWAN

LoRaWAN is a media access control (MAC) protocol for wide area networks. It is designed to allow low-powered devices to communicate with Internet-connected applications over long range wireless connections. LoRaWAN can be mapped to the second and third layer of the OSI model. It is implemented on top of LoRa or FSK modulation in industrial, scientific and medical (ISM) radio bands. The LoRaWAN protocols are defined by the [LoRa Alliance](https://www.lora-alliance.org/) and formalized in the [LoRaWAN Specification](https://www.lora-alliance.org/portals/0/specs/LoRaWAN%20Specification%201R0.pdf). A LoRaWAN stack is displayed below:

<br>
<center>
[[/uploads/lorawan-overview.png]]

_Overview of The Things Network's LoRaWAN stack_ </center>

<br>

## Terminology
* **[Device, Node, Mote](https://www.thethingsnetwork.org/wiki/Hardware/Nodes/Overview)** - an object with an embedded low-power communication device

* **[Network](https://www.thethingsnetwork.org/wiki/Backend/Overview)** - route messages from Nodes to the right Application, and back.

* **[Gateway](https://www.thethingsnetwork.org/wiki/Hardware/Gateways/Overview)** - antennas that receive broadcasts from Nodes and send data back to Nodes

* **Application** - a piece of software, running on a server

* **Uplink Message** - a message from a Device to an Application

* **Downlink Message** - a message from an Application to a Device



## Devices
The LoRaWAN specification defines three device types. All LoRaWAN devices must implement Class A, whereas Class B and Class C are extensions to the specification of Class A devices.

* **Class A** devices support bi-directional communication between a device and a gateway. Uplink messages (from the device to the server) can be sent at any time (randomly). As depicted in the figure, the device then opens two receive windows at specified times (1s and 2s) after an uplink transmission. If the server does not respond in either of these receive windows (situation 1 in the figure), the next opportunity will be after the next uplink transmission from the device.  The server can respond either in the first receive window (situation 2 in the figure), or the second receive window (situation 3 in the figure).

<br>
<center>
[[/uploads/rx-class-c.png]]

_Receive window of Class A devices_ </center>

<br>

* **Class B** devices extend Class A by adding scheduled receive windows for downlink messages from the server. Using time-synchronized beacons transmitted by the gateway, the devices periodically open receive windows.

* **Class C** devices extend Class A by keeping the receive windows open unless they are transmitting, as shown in the figure below. This allows for low-latency communication but is many times more energy consuming than Class A devices.

<br>
<center>
[[/uploads/rx-class-c.png]]

_Receive window of Class C devices_ </center>

<br>
## Frequency Bands

LoRaWAN operates in unlicensed radio spectrum. This means that anyone can use the radio frequencies without having to pay million dollar fees for transmission rights. It is similar to WiFi, which uses the 2.4GHz and 5GHz ISM bands worldwide. Anyone is allowed to set up WiFi routers and transmit WiFi signals without the need for a license or permit.

LoRaWAN uses lower radio frequencies with a longer range. The fact that frequencies have a longer range also comes with more restrictions that are often country-specific. This poses a challenge for LoRaWAN, that tries to be as uniform as possible in all different regions of the world. As a result, LoRaWAN is specified for a number of bands for these regions. These bands are similar enough to support a region-agnostic protocol, but have a number of consequences for the implementation of the backend systems.

#### European 863-870 MHz and 433 MHz bands
Of the available ISM frequency bands, LoRaWAN uses the 863-870 MHz and 433 MHz bands. The former, which is usually referred to as the 868 MHz band, is currently supported by The Things Network, whereas the latter will be implemented later.

The LoRaWAN specification defines 3 common 125 kHz channels for the 868 MHz band (868.10, 868.30 and 868.50 MHz) that must be supported by all devices and networks, and that all gateways should always be receiving on. These three channels form a common set of channels that all devices can use to join with a network. During this join procedure, the network can instruct the devices to add additional channels to its channel set. These channels are used for both uplink and downlink messages.

##### Duty Cycle
The European frequency regulations impose specific duty-cycles on devices for each sub-band. These apply to each device that transmits on a certain frequency, so both gateways and devices have to respect these duty-cycles. Most channels used by LoRaWAN have a duty-cycle as low as 1% or even 0.1%. As a result, the network should be smart in scheduling messages on gateways that are less busy or on channels that have a higher duty-cycle.

#### US 902-928 MHz
In the United States, LoRaWAN operates in the 902-928 MHz frequency band. Unlike the European band, the US band has dedicated uplink and downlink channels. The band is divided into 8 sub-bands that each have 8x125 kHz uplink channels, 1x500 kHz uplink channel and 1x500 kHz downlink channel.

#### Australia 915-928 MHz
The specification of the Australian 915-928 MHz band is practically the same as the US 902- 928 MHz, except that its uplink frequencies are on higher frequencies than in the US band. Its downlink channels are the same as in the US 868 MHz band.

#### China 779-787 MHz and 470-510 MHz
The Chinese 779-787 MHz band behaves similar to the European bands. The 779-787 MHz band also has three common 125 kHz channels (779.5, 779.7 and 779.9 MHz). The Chinese 470-510 MHz band behaves similar to the US bands. There are 96 uplink channels and 48 downlink channels. In some regions, a subset of these channels is used by China Electric Power and can therefore not be used for LoRaWAN.

## Modulation and Data Rate
In most cases LoRaWAN uses LoRa modulation. LoRa modulation is based on Chirp spread- spectrum technology, which makes it work well with channel noise, multipath fading and the Doppler effect, even at low power.

The data rate depends on the used bandwidth and spreading factor. LoRaWAN can use channels with a bandwidth of either 125 kHz, 250 kHz or 500 kHz, depending on the region (like in the US band) or the frequency plan (in Europe). The spreading factor is chosen by the end-device and influences the time it takes to transmit a frame.

## Addressing

Devices and applications have a 64 bit unique identifier (`DevEUI` and `AppEUI`). When a device joins the network, it receives a dynamic (non-unique) 32-bit address (`DevAddr`). See the [[Addressing|LoRaWAN/Address-Space]] page for more information about the address space and device activation procedure in LoRaWAN.

## Security

LoRaWAN knows three distinct 128-bit security keys. The application key `AppKey` is only known by the device and by the application. When a device joins the network (this is called a join or activation), an application session key `AppSKey` and a network session key `NwkSKey` are generated. The `NwkSKey` is shared with the network, while the `AppSKey` is kept private. These session keys will be used for the duration of the session.

<br>
<center>
[[/uploads/keys.png]]

_Keys and Encryption in LoRaWAN_ </center>

<br>
The above figure shows how these keys are used. The `AppSKey` is used for end-to-end encryption of the frame payload. The algorithm used for this is AES-128, similar to the algorithm used in the 802.15.4 standard. The `NwkSKey` is known by the network and the device and is used to validate the integrity of each message by its Message Integrity Code (MIC). This MIC is similar to a checksum, except that it prevents intentional tampering with a message. For this, LoRaWAN uses AES-CMAC.

##### [Frame Counters](https://www.thethingsnetwork.org/wiki/LoRaWAN/Security#security-in-lorawan-and-ttn_frame-counters)
Frame Counters prevent replay attacks where an attacker re-transmits a previously recorded message. To prevent this, the network and the device must both reject messages that contain a frame counter that is lower than the expected frame counter.

## MAC Commands
The network server and device can perform network-related administration and management using MAC commands. The LoRaWAN specification specifies a number of commands that can be extended in future versions of LoRaWAN or extended with proprietary commands. There are currently commands for checking connectivity, requesting the status of a device, adapting the data rate of a device, and modifying channel settings.

## External Resources

* [LoRaWAN Specification 1.0 (LoRa Alliance)](https://www.lora-alliance.org/portals/0/specs/LoRaWAN%20Specification%201R0.pdf)
* [LoRa Modulation Basics (Semtech)](http://www.semtech.com/images/datasheet/an1200.22.pdf)
* [LoRa Modem Design Guide (Semtech)](http://www.semtech.com/images/datasheet/LoraDesignGuide_STD.pdf)
* [LoRa FAQs (Semtech)](http://www.semtech.com/wireless-rf/lora/LoRa-FAQs.pdf)
* [Reversing LoRa](http://static1.squarespace.com/static/54cecce7e4b054df1848b5f9/t/57489e6e07eaa0105215dc6c/1464376943218/Reversing-Lora-Knight.pdf)
