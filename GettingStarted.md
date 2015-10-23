# Getting Started with The Things Network

## LoRa and LoRaWAN
Wireless connectivity on The Things Network uses the
**LoRa** technology. LoRa enables a serial data connection
between two devices. A smart filtering mechanism allows
connections over very long distances (2-25km) on a free
bandwidth (868MHz EU, 915MHz USA, 433MHz Asia), although
data throughput is limited (1-50 kbps). The connection is
best suited for short data packets.

**LoRaWAN** is a software stack on top of LoRa connections,
describing a network of multiple gateways connected to a
central server. The Things Network uses LoRaWAN, but replaces
the central server with a distributed network owned by the
community.
(todo: link to papers on LoRaWan)

## What you need minimal
###Gateway
Something in the neighborhoud that will listen to devices. This is similar with a wifi accesspoint. This can be of yourself of from someone else. Everybody is free to use eachother gateway.
###Node
This is the device that does something. for example a temperature sensor or a small blinking light.

(todo: insert image of LoRaWan architecture (incl technology)

## Supported Hardware
Like WiFi and 3G/4G, LoRa connections need specific hardware
support for connections. Since the technology is only gaining
traction now, most devices won't support it (yet). You'll need
to build (or buy) devices yourself.

###Nodes###
There's at least ten providers selling breakout boards for
LoRa(WAN) nodes. All of them use the SX1272 or SX1276 chip,
produced by SemTech and providing the 'raw' LoRa hardware
protocol. The LoRaWAN software stack needs to run on a
separate [MCU](https://www.quora.com/What-is-the-difference-between-a-microprocessor-and-microcontroller) like an Arduino or ARM Cortex. There are a
few different implementations, and they're not yet all
compatible with each other.

Some boards include a MCU, some don't and just provide the
modem interface of the chip. There's bare breakout boards,
Bee connectors, or all-included development boards.

To get started building a node, see our **[Hardware Nodes Overview](Hardware/OverviewNodes)**.

###Gateway###
On the gateway side there's also a few options. Since the
SX1272/SX1276 only allows for one connection at a time,
most of them use the SX1301 + SX1257 chips allowing up to
8 simultaneous connections (typically supporting 10000 ~ 20000
nodes).

To get started with gateways, see our **[Hardware Gateway Overview](Hardware/OverviewGateways)**.


## Network
Once you've got a node ready to transmit and/or receive data,
start with our **[Software Overview](Software/Overview)**.
