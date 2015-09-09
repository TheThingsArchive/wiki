## Prototype Network Setup

The current setup is temporary, but available for use.
We're finalising the initial network architecture, but
you can get a glimpse of what's coming by reading this page.
Of course, things will break and evolve over time, so
if you have questions please get in [contact](Contact)

![data flow overview](http://thethingsnetwork.org/wiki/Software/Overview/attachment/3/ttn_prototype_data_flow.png)

For an overview of the current network, see
[[CurrentNetwork]]


## Instructions Upstream (Things side)
### Kerlink
* Install Demo LoRa Packet forwarder as described [Here](http://wikikerlink.fr/lora-station/doku.php?id=wiki:semtech#demo_lora_packet_forwarderready-to-use_package).
* Change and add the following lines in /mnt/fsuser-1/demo_gps_loramote/local_conf.json
* Don't change your gateway_ID, it's linked to your serial.

local_conf.json:

	{
	  "gateway_conf": {
	      "gateway_ID": *YOUR_OWN_SERIAL*,
		  "serv_port_up": 1700,
		  "serv_port_down": 1700,
		  "server_address": "croft.thethings.girovito.nl",
		  "forward_crc_valid": true,
		  "forward_crc_error": true,
		  "forward_crc_disabled": true
	  }
	}

Want to connect the Kerlink over gprs/3g using a simcard? Follow [these](gateways/kerlink/mobile-connection) instructions.

## Instructions Downstream (Server side)
### Getting the Data
We're releasing instructions on connecting to our router(s) soon. The data will be available with various endpoints (REST API, publication subscriptions, database login, example dashboard), all available as open-source software to run on your own server.

Bare with us while we're finalising the first version! Can't wait? Get in [[Contact]].


## Github Repositories

### Nodes firmware
* Netblocks XRange: [[https://github.com/TheThingsNetwork/XRange]]
* Arduino-based / SODAQ: [[https://github.com/TheThingsNetwork/loraduino]]

### Gateways firmware
* Kerlink-based firmware: [[https://github.com/TheThingsNetwork/lora_gateway]]
* Kerlink-based packet forwarder: [[https://github.com/TheThingsNetwork/packet_forwarder]]
* Kerlink station firmware files: [[https://github.com/TheThingsNetwork/kerlink-station-firmware]]
* Supply fake packages to gateway: [[https://github.com/TheThingsNetwork/ghost_node]]

### Server software
* Docker / dev environment: [[https://github.com/TheThingsNetwork/server-devenv]]
* 'server shared': [[https://github.com/TheThingsNetwork/server-shared]]
* Croft -> gateway to queue: [[https://github.com/TheThingsNetwork/croft]]
* Jolie -> queue to storage and application endpoint: [[https://github.com/TheThingsNetwork/jolie]]