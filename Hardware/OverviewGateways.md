# Introduction
At the moment there are several options to get you own gateway. These fall into two categories:
* Ready to roll
* Do it yourself

There is also a third option, that is design your own hardware, but that is beyond the scope of this how-to.

## Ready to roll
Depending on how much money you want to spend and when the hardware is needed the following options are available (from most expensive to least expensive):

### [Kerlink](http://www.kerlink.fr/en/products/lora-iot-station/11-products-uk)
* Price: on request, around 1500 euro.
* Availability: direct available, (order in France)
* Capacity: 8 band, ~ 20K Nodes 
* Suitable for outdoor mounting (IP67)
* Installation: Requires a technician
* Hardware: proprietary
* Software: open source & proprietary
* Possibility to run on gprs/3g network
* Tested by The Things Network Crew.

Discussion: This is an industrial solution suitable for people who want to mount the gateway outside and have sufficient technical skills to connect, mount and maintain the device themselves. We have tested the device and although we have remarks about the somewhat older software that is being used, this device will do the job. A trained software engineer will be able to update the device using the software from The Things Network. 


### [MultiTech](http://www.multitech.com/brands/multiconnect-conduit)
* Price: on request, around 700 euro, may require additional boards to run
* Availability: unknown
* Capacity: ? 
* Suitable for indoor use
* Installation: Plug and play (?)
* Hardware: proprietary
* Software: proprietary (?)
* Tested by The Things Network Crew. 

Discussion: We have one, results of our test will follow asap. 


### [Lorank 8](http://www.hoperf.nl/LORANK-8)
* Price: 412 euro 
* Availability: pre-order, estimated delivery time 30 days
* Capacity: 8 band, ~ 20K Nodes 
* Suitable for indoor use
* Installation: Plug and play
* Hardware: radio proprietary, MCU board: open source 
* Software: open source
* Tested by The Things Network Crew.

Discussion: The device is build upon the radio board of IMST and the open
source hardware BeagleBone Black. Although the first version had some
issues, at this moment the device seems to work as advertised. Please consult
the online manual for further information.


### [The Things Gateway](http://thethingsnetwork.org/kickstarter-landing/kickstarter.html)
* Price: 200 euro 
* Availability: kickstarter will launch October
* Capacity: 1 band, ~ 3K Nodes (strech goal: 8 band ?) 
* Suitable for indoor use
* Installation: Plug and play
* Hardware: open source 
* Software: open source
* TO BE tested by The Things Network Crew.

Discussion: The device is specially developed for The Things Network and
we will make sure it will meet all requirements for a smooth experience on
The Things Network. Note that estimated first delivery is around May ~ June 2016.

## Do it yourself
Here's other people building their own gateway:

### DIY "Real" gateway (SX1301 + SX1257)
* [https://github.com/mirakonta/lora_gateway/wiki](https://github.com/mirakonta/lora_gateway/wiki)

### DIY "single-connection" gateway (SX1272, SX1276 etc)
* [http://www.daveakerman.com/?p=1719](http://www.daveakerman.com/?p=1719)

- interesting read, but for now does not give you anything useful for The Things Network
- to be completed


## Network
Once you've got a gateway ready to transmit and/or receive data,
start with our **[Software Overview](/wiki/Software/Overview)**.
