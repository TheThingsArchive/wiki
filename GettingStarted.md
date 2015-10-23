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

##Minimal Setup
###Gateway
This is something, when you are lucky is already in the neighborhood available. Otherwise this is something you have to build. The gateway is the portal between your IoT Device to the internet & thethingsnetwork.
###Node
This is the IoT device that does something. for example a temperature sensor or a small blinking light. It will communicate through the Gateway to the things network. An IoT device you configure once (or the supplier does this) and then you can use it everywhere on the world where an TTN gateway is.
(todo: insert image of LoRaWan architecture (incl technology)

    node  <------->  gateway  <-->  /-----\
              __/                  | cloud | <--> your server
    node  <--/---->  gateway  <-->  \-----/
                /
    node  <----/

## Supported Hardware
Like WiFi and 3G/4G, LoRa connections need specific hardware
support for connections. Since the technology is only gaining
traction now, most devices won't support it (yet). You'll need
to build (or buy) devices yourself.

###Nodes###
There's at least ten providers selling breakout boards for
LoRa(WAN) nodes. All of them use the SX1272 or SX1276 chip,
produced by SemTech and providing the 'raw' LoRa hardware
protocol. 

A **node** has two building blocks. a communication part and a calculation part. 

**Communication** is done via the SX1272 or SX1276 chip this is the modem part of the node. 
**Calculation** is done by a separate [MCU](https://www.quora.com/What-is-the-difference-between-a-microprocessor-and-microcontroller) like  an Arduino or ARM Cortex.
The LoRaWAN software stack needs to run on the MCU (the calculation part)
Some boards include a MCU, some don't and just provide the modem interface of the chip. 
There are bare breakout boards, Bee connectors, or all-included development boards. There are a few different implementations, and they're not yet all compatible with each other.

To get started building a node, see our **[Hardware Nodes Overview](Hardware/OverviewNodes)**.
This includes hardware, software and how to buy or build it.

###Gateway###
On the gateway side there's also a few options. 
The Gateway also has a communication part and a calculation part. 
The gateway supports many nodes thus the calculation part should be bigger then the node.

For **Communication** there are two options for the chip the SX1272/SX1276 option and the SX1301 + SX1257 option.
*the SX1272/SX1276 only allows for one (1) connection at a time,
*the SX1301 + SX1257 chips allows up to eight (8() simultaneous connections (typically supporting 10000 ~ 20000 nodes).
For the **Calculation** part a device that supports linux is usually used. These are provided by many many suppliers, more on this in ouw Hardware Gateway Overview.

To get started with gateways, see our **[Hardware Gateway Overview](Hardware/OverviewGateways)**.
This includes hardware, software and how to buy or build it.

## Network
Once you've got a node ready to transmit and/or receive data via a Gateway,
start with our **[Software Overview](Software/Overview)**.
