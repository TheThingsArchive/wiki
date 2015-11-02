**EU 863-870MHz ISM Band**

This page is intended to collect information on different channel plans
that are in use.

As long as TTN does not support MAC commands to communicate frequencies
from the router to the node, agreement on a frequency plan is required.

Also note that the MAC commands related to frequencies are processed by the network server (router), not by the gateway. 
This means that the router needs to know about the frequencies supported by every gateway connecting to it, or more practically,
all gateways on the same router need to adhere to the same freq. plan.

**LoRaWAN Specification 1R0**

LoRaWAN Default Channels:

868.10 Mhz (g1)

868.30 MHz (g1)

868.50 MHz (g1)

Those channels are the minimum set that all network gateways should
always be listening on.

The LoRaWAN enforces a per sub-band duty-cycle limitation. The ETSI regulations allow the choice of using either a duty-cycle limitation or a so-
called Listen Before Talk Adaptive Frequency Agility (LBT AFA) transmissions
management. The current LoRaWAN specification exclusively uses duty-cycled limited
transmissions to comply with the ETSI regulations.

Default radiated transmit output power: 14 dBm

The following table gives the list of frequencies that should be used by
end-devices to broadcast the JoinReq message.

864.10 Mhz (g)

864.30 Mhz (g)

864.50 Mhz (g)

868.10 Mhz (g1)

868.30 Mhz (g1)

868.50 Mhz (g1)

Class B beacon recommended frequency: 869.525 MHz (g3, 27dBm)

All unicast&multicastClass B downlinks use a single frequency channel
defined by the PingSlotChannelReq MAC command. The default frequency is
869.525MHz

**Bands & Regulations**

ERC Recommendation 70-03

g 865.0 – 868.0 MHz 1% or LBT+AFA, 25 mW (=14dBm)

g1 868.0 – 868.6 MHz 1% or LBT+AFA, 25 mW

g2 868.7 – 869.2 MHz 0.1% or LBT+AFA, 25 mW

g3 869.4 – 869.65 MHz 10% or LBT+AFA, 500 mW (=27dBm)

g4 869.7 – 870.0 MHz 1% or LBT+AFA, 25 mW (no dutcy-cyle requirement if power < 5 mW/7 dBm)

LBT+AFA: Listen Before Talk (LBT) with Adaptive Frequency Agility (AFA).

**Gateways: SX1301/SX1257 based**

The gateway receive bandwidth extends from 863MHz to 870 MHz. However
the 10 different channels cannot be dispatched arbitrarily inside this
entire range because of the way the radio front- end receiver is
actually implemented. The 10 channels must fall in two 1.2 MHz wide
intervals.

Default configuration Semtech packet forwarder:

Chan , Freq , Modulation , BW , SF , radio, band, duty-cyle limit

0, 868.1 MHz, LoRa, 125kHz, all SF, 1, g1

1, 868.3 MHz, LoRa, 125kHz, all SF, 1, g1

2, 868.5 MHz, LoRa, 125kHz, all SF, 1, g1

3, 867.1 MHz, LoRa, 125kHz, all SF, 0, g

4, 867.3 MHz, LoRa, 125kHz, all SF, 0, g

5, 867.5 MHz, LoRa, 125kHz, all SF, 0, g

6, 867.7 MHz, LoRa, 125kHz, all SF, 0, g

7, 867.9 MHz, LoRa, 125kHz, all SF, 0, g

8, 868.3 MHz, LoRa, 250kHz, SF7, 1, g1 (=10 kbps)

9, 868.8 MHz, FSK, 150kHz, 50kbps channel, 1, g2

**LoRaMac-Node (Semtech)**

// Band = { DutyCycle, TxMaxPower, LastTxDoneTime, TimeOff }

\#define BAND0 { 100 , TX\_POWER\_14\_DBM, 0, 0 } // 1.0 %

\#define BAND1 { 100 , TX\_POWER\_14\_DBM, 0, 0 } // 1.0 %

\#define BAND2 { 1000, TX\_POWER\_14\_DBM, 0, 0 } // 0.1 %

\#define BAND3 { 10 , TX\_POWER\_14\_DBM, 0, 0 } // 10.0 %

\#define BAND4 { 100 , TX\_POWER\_14\_DBM, 0, 0 } // 1.0 %

// Channel = { Frequency \[Hz\], { ( ( DrMax &lt;&lt; 4 ) | DrMin ) },
Band }

\#define LC1 { 868100000, { ( ( DR\_5 &lt;&lt; 4 ) | DR\_0 ) }, 1 }

\#define LC2 { 868300000, { ( ( DR\_6 &lt;&lt; 4 ) | DR\_0 ) }, 1 } //
DR6 = SF7BW250

\#define LC3 { 868500000, { ( ( DR\_5 &lt;&lt; 4 ) | DR\_0 ) }, 1 }

\#define LC4 { 867100000, { ( ( DR\_5 &lt;&lt; 4 ) | DR\_0 ) }, 0 }

\#define LC5 { 867300000, { ( ( DR\_5 &lt;&lt; 4 ) | DR\_0 ) }, 0 }

\#define LC6 { 867500000, { ( ( DR\_5 &lt;&lt; 4 ) | DR\_0 ) }, 0 }

\#define LC7 { 867700000, { ( ( DR\_5 &lt;&lt; 4 ) | DR\_0 ) }, 0 }

\#define LC8 { 867900000, { ( ( DR\_5 &lt;&lt; 4 ) | DR\_0 ) }, 0 }

\#define LC9 { 868800000, { ( ( DR\_7 &lt;&lt; 4 ) | DR\_7 ) }, 2 } //
FSK

Note: LC2 is two channels (BW125 and BW250)

**IBM LMIC**

Note: in the Arduino adaptation of LMIC v1.5 we have modified these default frequencies to match the Semtech packet forwarder defaults.

// Default frequency plan for EU 868MHz ISM band

// Bands:

// g1 : 1% 14dBm

// g2 : 0.1% 14dBm

// g3 : 10% 27dBm

enum { EU868\_F1 = 868100000, // g1 SF7-12

EU868\_F2 = 868300000, // g1 SF7-12 FSK SF7/250

EU868\_F3 = 868500000, // g1 SF7-12

EU868\_F4 = 868850000, // g2 SF7-12

EU868\_F5 = 869050000, // g2 SF7-12

EU868\_F6 = 869525000, // g3 SF7-12

EU868\_J4 = 864100000, // g2 SF7-12 used during join

EU868\_J5 = 864300000, // g2 SF7-12 ditto

EU868\_J6 = 864500000, // g2 SF7-12 ditto

**KPN**

The KPN setup has been validated with the Loramote default software and
the Semtech code from Github (https://github.com/Lora-net/LoRaMac-node)
it is important to adjust the used frequency setup in the
LoRaMac-board.h file as follows:

// Channel = { Frequency \[Hz\], { ( ( DrMax &lt;&lt; 4 ) | DrMin ) },
Band }

\#define LC1 { 868100000, { ( ( DR\_SF7 &lt;&lt; 4 ) | DR\_SF12 ) }, 1 }

\#define LC2 { 868300000, { ( ( DR\_SF7 &lt;&lt; 4 ) | DR\_SF12 ) }, 1 }

\#define LC3 { 868500000, { ( ( DR\_SF7 &lt;&lt; 4 ) | DR\_SF12 ) }, 1 }

\#define LC4 { 868850000, { ( ( DR\_SF7 &lt;&lt; 4 ) | DR\_SF12 ) }, 2 }

\#define LC5 { 869050000, { ( ( DR\_SF7 &lt;&lt; 4 ) | DR\_SF12 ) }, 2 }

\#define LC6 { 869525000, { ( ( DR\_SF7 &lt;&lt; 4 ) | DR\_SF12 ) }, 3 }

**Actility ThingPark Wireless**

LC1 868.10 g1 SF7-SF12

LC2 868.30 g1 SF7-SF12

LC3 868.50 g1 SF7-SF12

LC4 868.85 g2 SF7-SF12

LC5 869.05 g2 SF7-SF12

LC6 869.525 g3 SF7-SF12

LC7 868.3 g1 SF7BW250

FC1 868.30 g1 FSK 250Khz 100 kbps
