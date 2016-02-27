# Getting Started with The Things Network

## Step 0: Get Informed
... TODO move to separate page (info below)

## Step 1: Get A Gateway
The gateway is the portal between your IoT Device to the internet
& thethingsnetwork. It will connect all your nodes within a range
of 2-15km with applications on the internet.

There might be a gateway already available in your neighborhood.
To find out, have a look at [this gateway map](http://thethingsnetwork.org/map),
one of the maps [mentioned here](CurrentNetwork), and search on the
[forum](http://forum.thethingsnetwork.org/) for nearby communities.

If there isn't a gateway available yet, you or your community has to buy
(or build) one. Have a look at [Which Gateway to buy](http://forum.thethingsnetwork.org/).

## Step 2: Prototype a Node
You'll need a IoT device that does something and uses the network.
For example, a temperature sensor or a small blinking light. It will
communicate through the Gateway to the things network. An IoT device
you configure once (by you or the supplier) and then you can use it
everywhere on the world while in range of a local TTN gateway.

    node  <------->  gateway  <-->  /-----\
              __/                  | cloud | <--> your server
    node  <--/---->  gateway  <-->  \-----/
                /
    node  <----/
    
To prototype your first node, have a look at [this page](Hardware/OverviewNodes).
The page will point to online introductions in hardware prototyping,
and features information on many different hardware devices that can connect
to The Things Network Gateways over the LoRaWAN protocol.
    

## Step 3: Connect an Application
Once you've got a node ready to transmit and/or receive data via a Gateway,
start with our **[Software Overview](Software/Overview)**.


... TODO


## Background information

### LoRa and LoRaWAN
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



### Supported Hardware
Like WiFi and 3G/4G, LoRa connections need specific hardware
support for connections. Since the technology is only gaining
traction now, most devices won't support it (yet). You'll need
to build (or buy) devices yourself.


