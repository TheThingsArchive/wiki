## Prototype Network Setup

The current setup is temporary, but available for use.
We're finalising the initial network architecture, but
you can get a glimpse of what's coming by reading this page.
Of course, things will break and evolve over time, so
if you have questions please get in [contact](Contact)

![data flow overview](https://raw.githubusercontent.com/wiki/TheThingsNetwork/docs/imgs/ttn_prototype_data_flow.png)

For an overview of the current network, see
[[CurrentNetwork]]

For an overview of the (upcoming) network architecture, see
[[Architecture]]

## Description of distributed network setup
As we're planning it now, the gateways (open-source packet handler) will send packets to one or multiple routers (open source router software) routing to one or multiple application servers (open source or closed software). Any (set of) nodes can be linked to any set of application servers, and the routers will perform the mapping. We're bootstrapping this mapping with DNS, but eventually want to run something more distributed there as well.

So, there shouldn't be any "central" point in the near future, once more people start running routers and application servers. Of course we will provide a default router and application server with API, to get started, but everyone is free to add routers and app servers to the network.

Right now, though, there's one router and one application server running, and we're rewriting the proof-of-concept into a more stable version. Details will follow soon ;-)


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
		  "forward_crc_error": false,
		  "forward_crc_disabled": true
	  }
	}

Want to connect the Kerlink over gprs/3g using a simcard? Follow [these](gateways/kerlink/mobile-connection) instructions.



## Instructions Downstream (Server side)

### Getting the Data
By default, all packets sent by any node using our default key settings,
will be saved in a database (InfluxDB) for undetermined amount of time
(currently forever).

### API
There's a REST API available to query for the latest packets.
It's available at `http://thethingsnetwork.org/api/v0`. Visiting
with a browser will enable a webview.

Here's the endpoints:

| endpoint                       | GET parameters (optional)               | explanation                                  |
| ------------------------------ | --------------------------------------- | -------------------------------------------- |
| **`/nodes/`**                  | `time_span` e.g. 10m, 4h, 1w            | Last single packet for all available nodes   |
|                                |                                         | within given timeframe                       |
| **`/nodes/{node_eui}/`**       | `time_span` e.g. 10m, 4h, 1w            | Last packets for given node                  |
|                                | `limit` (int, def=20) per gateway       |                                              |
|                                | `offset` (int) per gateway              |                                              |
| **`/gateways/`**               | `time_span` e.g. 10m, 4h, 1w            | Last single status update for all gateways   |
|                                | `limit` (int, def=20)                   | within given timeframe                       |
|                                | `offset` (int)                          |                                              |
| **`/gateways/{eui}`**          | `time_span` e.g. 10m, 4h, 1w            | Last status updates for given gateway        |
|                                | `limit` (int, def=20)                   | within given timeframe                       |
|                                | `offset` (int)                          |                                              |

The node packets will include the following data fields:

  * `data_raw`: the unencrypted payload
  * `data`: base64-encoded decrypted data (if encrypted with standard key)
  * `data_plain`: ascii version of decrypted data (if `data` decode-able into ascii)
  * `data_json`: json object / dictionary (if `data_plain` contains json)


### Work In Progress
A note of caution:
The current setup is temporary and not as distributed as we would like it to be.
As we're rewriting the routing system we might change detauls like API endpoints
and data format, but we'll try to keep the system available in a more or less
stable format.

Upcoming: distributed routing, app registration API, publication subscriptions,
hosted Node-RED, example application dashboard; all available as open-source
software to run on your own server.

Bear with us while we're finalising the first version! Can't wait? Get in [[Contact]].

# Android SDK

Ready to develop Android apps using data from The Things Network? Grab the Android SDK and you're ready to go. Find it on GitHub (below).


## Github Repositories

### Nodes firmware
* [Netblocks XRange](https://github.com/TheThingsNetwork/XRange)
* [Arduino-based / SODAQ](https://github.com/TheThingsNetwork/loraduino)
* [Arduino + hoperf / semtech](https://github.com/tftelkamp/arduino-lmic-v1.5)

### Gateways firmware
* [Kerlink-based firmware](https://github.com/TheThingsNetwork/lora_gateway)
* [Kerlink-based packet forwarder](https://github.com/TheThingsNetwork/packet_forwarder)
* [Kerlink station firmware files](https://github.com/TheThingsNetwork/kerlink-station-firmware)
* [Supply fake packages to gateway](https://github.com/TheThingsNetwork/ghost_node)

### Server software
* [Docker / dev environment](https://github.com/TheThingsNetwork/server-devenv)
* ['server shared'](https://github.com/TheThingsNetwork/server-shared)
* [Croft -> gateway to queue](https://github.com/TheThingsNetwork/croft)
* [Jolie -> queue to storage and application endpoint](https://github.com/TheThingsNetwork/jolie)

### Application software
* [Android SDK](https://github.com/TheThingsNetwork/android-sdk)