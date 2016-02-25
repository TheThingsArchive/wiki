**Wanting to Get started with your data?**

To start developing applications using the network **API** or **MQTT**,
please [see below](#getting-the-data).

## Prototype Network Setup

The current setup is temporary, but available for use.
We're finalising the initial network architecture, but
you can get a glimpse of what's coming by reading this page.
Of course, things will break and evolve over time, so
if you have questions please get in [contact](Contact)

For an overview of the current network, see
[[CurrentNetwork]]

For an overview of the (upcoming) network architecture, see
[[Architecture]]

## Description of distributed network setup
The gateways (open-source packet handler) will send packets to one or multiple routers (open source router software) routing to one or multiple application servers (open source or closed software). Any (set of) nodes can be linked to any set of application servers, and the routers will perform the mapping. We might bootstrap this mapping with DNS, but eventually want to run something more distributed there as well.

So, there shouldn't be any "central" point in the near future, once more people start running routers and application servers. Of course we will provide a default router and application server with API, to get started, but everyone is free to add routers and app servers to the network.

Right now, though, there's one router and one application server running, and we're rewriting the proof-of-concept into a more stable version. Details will follow soon ;-)


## Instructions Upstream (Things / Node side)
### Kerlink
* An installation guide [Installation guide](../Installing-your-Kerlink)
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

### Other gateways
If you have another gateway, you might find help on the [TTN forum](http://forum.thethingsnetwork.org/).
There's many different gateways out there already, so you'll probably be able to
find a thread for your gateway already. If so, please add a link to the thread here.


## Instructions Downstream (Server side: application development)

### Getting the Data
Currently there are two different ways of getting the data:

  - pull from REST API: for batched or semi-realtime data gathering (based on storage)
  - connect to MQTT stream: for realtime pub/sub data gathering (live data only)

By default, all packets sent by any node that are encrypted using our default
key settings, will be both saved in a database (MongoDB) for undetermined amount
of time (currently forever), and made available via the MQTT broker.

### API
There's a REST API available to query for the latest packets.
It's available at `http://thethingsnetwork.org/api/v0`. Visiting
with a browser will enable a webview.

Here's the endpoints:

| endpoint                       | GET parameters (optional)               | explanation                                  |
| ------------------------------ | --------------------------------------- | -------------------------------------------- |
| **`/nodes/`**                  | `limit` (int, def=100, max=250)         | Aggregated info on nodes (sorted last seen)  |
|                                | `offset` (int)                          |                                              |
| **`/nodes/{node_eui}/`**       | `limit` (int, def=100, max=250)         | Last packets for given node                  |
|                                | `offset` (int)                          |                                              |
| **`/gateways/`**               | `limit` (int, def=100, max=250)         | Aggregated info on all gateways (sorted last |
|                                | `offset` (int)                          | status message received)                     |
| **`/gateways/{eui}`**          | `limit` (int, def=100, max=250)         | Last status messages for given gateway       |
|                                | `offset` (int)                          |                                              |

The node packets will include the following data fields:

  * `data_raw`: the unencrypted payload
  * `data`: base64-encoded decrypted data (if encrypted with standard key)
  * `data_plain`: ascii version of decrypted data (if `data` decode-able into ascii)
  * `data_json`: json object / dictionary (if `data_plain` contains json)

A second, simpler REST API (`v0.1`) exposes the same data and is available at `http://thethingsnetwork.org/api/v0.1`. 

Endpoints:

| endpoint                  | explanation                                  |
| ------------------------- | -------------------------------------------- |
| **`/nodes/`**             | Aggregated info on nodes that have been      |
|                           | active in the past 7 days (sorted last seen) |
|                           |                                              |
| **`/nodes/{node_eui}/`**  | Last packets for given node                  |
|                           |                                              |
| **`/gateways/`**          | Aggregated info on all gateways that have    |
|                           | been active in the past 7 days (sorted last  |
|                           | status message received)                     |
| **`/gateways/{eui}`**     | Last status messages for given gateway       |
|                           |                                              |

The node packets will include the following data fields:

  * `data_raw`: the unencrypted payload
  * `data`: base64-encoded decrypted data (if encrypted with standard key)




### MQTT and NodeRED
MQTT is gaining momentum as the defacto protocol used for Internet of Things
applications. It allows for real-time data communication between any combination
of `publishers` (sending data) and `subscribers` (receiving data) via a semi-distributed
network. There's client libraries available for most programming languages.

For real-time publish-subscribe (Pub/Sub) connections we're hosting an MQTT broker at:

    tcp://croft.thethings.girovito.nl:1883
    
The topics you can subscribe to are:

* **Messages from nodes**: `nodes/{devAddr}/packets`
* **Status updates from gateways**: `gateways/{eui}/status`
    
    
For demo purposes, there's also a hosted version of [Node-RED](http://nodered.org/) here:

    http://croft.thethings.girovito.nl:1880/

For long-term usage, it's better to host your [own version of Node-RED](http://nodered.org/docs/), as there's currently no authentication for the web interface. Node-RED comes pre-installed in Rasbian-Jessie (November 2015) for RaspberryPi
    

### Work In Progress
A note of caution:
The current setup is temporary and not as distributed as we would like it to be.
As we're rewriting the routing system we might change details like API endpoints
and data format, but we'll try to keep the system available in a more or less
stable format.

Upcoming: distributed routing, app registration API, example application dashboard, etc; all available as open-source
software to run on your own server.

Bear with us while we're finalising the first version! Can't wait? Get in [[Contact]].


### Android SDK

Ready to develop Android apps using data from The Things Network?
Grab the Android SDK (which is implementing the REST API) and you're ready to
go. Find it on GitHub (below).

Example Android app:
[The Things Network SDK Sample](https://play.google.com/store/apps/details?id=org.ttn.android.sample)

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