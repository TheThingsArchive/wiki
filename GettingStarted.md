# Getting Started with The Things Network

## Step 0: Get Informed
For an introduction into the communication technology used by The
Things Network, read through [this page](IntoductionLorawan).

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

