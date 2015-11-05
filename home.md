# Welcome to the The Things Network wiki!

testing this

We're a crowd-sourced long-distance wireless network,
that's free to use and owned by its community. You can
read more about us on http://thethingsnetwork.org and in
our [manifest](https://github.com/TheThingsNetwork/Manifest).

## Quick Index
Non-tech:

  * [[Getting-organized]]
  * [Media attention](Media)

Tech:

  * [[GettingStarted]]
  * [Gateways to buy](http://thethingsnetwork.org/wiki/Hardware/OverviewGateways)
  * [Connect your node](Hardware/OverviewNodes)
  * [Applications: overview + API + SDKs](http://thethingsnetwork.org/wiki/Software/Overview)

## Architecture
The Things Network uses the **LoRa(WAN)** technology for wireless
connectivity over long distances.
**Gateways** provide connectivity with a range of 2-25km.
**Nodes** (based on Arduino or other MCUs) equipped with a LoRa(WAN)
transceiver can send data to and receive data from nearby gateways.
The data will be forwarded through a **distributed network** via the
internet to any **server** you like. We provide a demo server as well.
See also [[CurrentNetwork]].

    node  <------->  gateway  <-->  /-----\
              __/                  | cloud | <--> your server
    node  <--/---->  gateway  <-->  \-----/
                /
    node  <----/

For more details on the architecture see [[Architecture]].

## Getting Started
To get started connecting to the The Things Network,
read [[Getting-organized]] and our technical **[[GettingStarted]]** manual.

Be sure to check out the [Examples](software/examples) pages and take an look on the [forum](http://forum.thethingsnetwork.org).

If you want to help building the network itself from a
non-technical perspective, get in [[contact]]!
We're looking for people to step up and take responsibility
for expanding the network in a specific area.

## Next Steps
[Github repositories](https://github.com/TheThingsNetwork)
coming soon..


## More Info
[Manifest](https://github.com/TheThingsNetwork/Manifest)

[Media about us](Media)

[Contact us](contact)
