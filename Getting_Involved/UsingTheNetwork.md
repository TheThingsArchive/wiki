## Introduction
What is a network without real-world use-cases? To get started with building Things using The Things Network, make sure there is a Gateway available nearby (2-15km). To find out, have a look at [this gateway map](http://thethingsnetwork.org/map), or search on the
[forum](http://forum.thethingsnetwork.org/) for nearby communities. If there isn't a gateway available yet, you or your community might have to [[buy (or build) one|Getting_Involved/Infrastructure]].

Like WiFi and 3G/4G, LoRa(WAN) connections need specific hardware support for connections. You'll need to buy special hardware supporting the LoRa(WAN) protocol. There is an increasing number of devices available supporting the protocol, as discussed below.


## Short note on hardware development
Below is technical information on various options for developing your own
Internet of Things node. If you've never done hardware development before,
you might want to start with some online resources first.
Have a look at the free [IoT introduction courses on coursera](https://www.coursera.org/specializations/iot).
(Note that the courses are freely accessible by searching for the course names
on Coursera.org.) Some background in programming both Arduino's and Raspberry
Pi's will greatly help prototyping your own IoT devices. There's many more
resources available on connecting particular hardware components to Arduino-compatible
microcontrollers, for example [this Sparkfun introduction](https://learn.sparkfun.com/tutorials/what-is-an-arduino),
[ladyada's tutorial](http://www.ladyada.net/learn/arduino/) and the
official practical [Arduino introductory booklet](https://store.arduino.cc/product/B000001).


## Nodes
There's many providers selling breakout boards for
LoRa(WAN) nodes. Most of them use the SX1272 or SX1276 chip,
produced by SemTech and providing the 'raw' LoRa hardware
protocol. The LoRa(WAN) software stack needs to run on a separate MCU, like an Arduino or ARM Cortex.
Some boards include a MCU, some don't and just provide the modem interface of the chip. 
There are bare breakout boards, Bee connectors, or all-included development boards. There are a few different implementations, and they're not yet all compatible with each other.

For an overview of node hardware available, have a look at [[our Node overview|Hardware/Nodes/Home]].


## Connect an application
Once the hardware is ready, you'll need to register your application (for free) and connect to the network to receive and send packages from and to your node(s). Currently the network supports the pub/sub MQTT protocol. Previously The Things Network foundation supported a REST API with data storage, but this has been deprecated. In the near future, there might be non-commercial and commercial applications available (made by you?).

For instructions, see [[Connect Your Application|Backend/Connect/Application]].