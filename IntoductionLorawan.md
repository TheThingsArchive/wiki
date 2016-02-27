Let's have a quick look at the technology enabling The
Things Network to connect thens of thousands of devices
within 15km to connect to the network.

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
For more information, have a look at [this FAQ from Semtech](http://www.semtech.com/wireless-rf/lora/LoRa-FAQs.pdf)
and [the Official LoRaWAN specification](https://www.lora-alliance.org/portals/0/specs/LoRaWAN%20Specification%201R0.pdf).


### Supported Hardware
Like WiFi and 3G/4G, LoRa connections need specific hardware
support for connections. Since the technology is only gaining
traction now, most devices won't support it (yet). You'll need
to build (or buy) devices yourself.

