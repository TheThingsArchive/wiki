# Welcome to the The Things Network wiki!

**You have reached the OLD wiki. Please consult [the NEW wiki](http://staging.thethingsnetwork.org/wiki/Home) for up to date information!**


We're a crowd-sourced long-distance wireless network,
that's free to use and owned by its community. You can
read more about us on [thethingsnetwork.org](http://thethingsnetwork.org) and in
our [manifest](https://github.com/TheThingsNetwork/Manifest).

## WORK IN PROGRESS!

You're reading the official public wiki, but The Things Network is in heavy
development: in March 2016 [a new staging environment of TTN Back-end 1.0 was announced](http://forum.thethingsnetwork.org/t/announcing-staging-environment-of-ttn-back-end-1-0/1852).

And with that staging environment comes **[an updated wiki](http://staging.thethingsnetwork.org/wiki)** as well.


## Quick Index

  * [Frequently Asked Questions](faq)
  
Non-tech:

  * [[Getting-organized]]
  * [Media attention](Media)
  * [Live Usecase examples](CurrentUsecases)

Tech:

  * **[[GettingStarted]]** <- start here to start using the network
  * [Gateways to buy](http://thethingsnetwork.org/wiki/Hardware/OverviewGateways)
  * [Nodes to buy & Connect your node](Hardware/OverviewNodes)
  * [Applications: overview + API + SDKs](http://thethingsnetwork.org/wiki/Software/Overview)
  * [Semtech PDF LoRA FAQ](http://www.semtech.com/wireless-rf/lora/LoRa-FAQs.pdf)

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

[Current network](CurrentNetwork)

[Contact us](contact)

