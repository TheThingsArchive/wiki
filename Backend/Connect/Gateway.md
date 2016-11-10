# Connect a Gateway

[[/uploads/TTN-Simple-Deployment.png]]

## Server Addresses

Choose the Router instance depending on the region as specified in the LoRaWAN 1.0.1 specification (still under NDA, [download 1.0](https://www.lora-alliance.org/portals/0/specs/LoRaWAN%20Specification%201R0.pdf)). This may not be the closest router geographically. For example, Brazil supports 915-928 MHz, but these are only specified by the LoRa Alliance for Australia, so Brazilian gateways should connect to the Australian router.

```
router.eu.thethings.network # EU 433 and EU 863-870
router.us.thethings.network # US 902-928
router.cn.thethings.network # China 470-510 and 779-787
router.au.thethings.network # Australia 915-928 MHz
```

Today, these DNS records refer to the staging environment, but we will update the records to the production environment when it comes available this summer. This will be announced timely in advance. Note: the staging environment currently only supports EU 863-870.

The geographical location of a server and the supported frequency plans are two different things. As the LoRa Alliance publishes more region specifications, we will deploy Routers in datacenters near that region. Also, the yet to be implemented decentralized routing architecture enables community operators to provide routing infrastructure to The Things Network.

## Staging Environment

It is highly recommended to update your gateway to the server addresses shown above. If you have test gateways that you want to force to use the staging environment and keep using the staging environment, use one of these server addresses. The staging environment will contain newer, but possibly less stable, builds of The Things Network back-end.

```
router.eu.staging.thethings.network # EU 433 and EU 863-870
router.us.staging.thethings.network # US 902-928
router.cn.staging.thethings.network # China 470-510 and 779-787
router.au.staging.thethings.network # Australia 915-928 MHz
```

## Configuring Your Gateway

*NOTE: Please use a packet forwarder based on version 2.2.x. The new `lora_pkt_fwd` version 3.0.0 is not supported yet.*

In the `local_conf.json` of the packet forwarder, update the fields `server_address` as follows:

```
"gateway_conf": {
        ...
        "servers": [{
            "server_address": "<insert server address here>",
            "serv_port_up": 1700,
            "serv_port_down": 1700,
            "serv_enabled": true,
            ...
        }]
}
```


## Checking connectivity
To see what happens under the hood, and make sure your gateway is sending its packets to the TTN-routers there are two approaches you can follow.

### Inspecting the port
Your gateway communicates with the routers via an UDP connection on port 1700; So if we listen to this port we should see packets being transferred.

```
sudo tcpdump -AUq port 1700
```
If your gateway is up and running this feed would (at least) give a stat packet every 30-60seconds:

```
08:07:33.801265 IP 192.168.178.20.47497 > 52.169.76.203.1700: UDP, length 221
E....Z@.@..h....4.L........'.6...'......{"stat":{"time":"2016-11-10 08:07:33 GMT","lati":51.1,"long":5.9,"alti":23,"rxnb":0,"rxok":0,"rxfw":0,"ackr":100.0,"dwnb":0,"txnb":0,"pfrm":"IMST + Rpi","mail":"info@example.nl","desc":"my-first-gateway"}}
```
If your gateway receives a transmission from a nearby node it wil issue a `rxpk` which also shows up in the feed. 

### Inspecting the log files
The packet forwarder periodically writes its status to a log file located at: `/var/log/syslog` You could tail this file to see what happens:

```
sudo tail -f /var/log/syslog
```

example log
```
##### 2016-11-10 08:10:33 GMT #####
### [UPSTREAM] ###
# RF packets received by concentrator: 0
# CRC_OK: 0.00%, CRC_FAIL: 0.00%, NO_CRC: 0.00%
# RF packets forwarded: 0 (0 bytes)
# PUSH_DATA datagrams sent: 2 (442 bytes)
# PUSH_DATA acknowledged: 100.00%
### [DOWNSTREAM] ###
# PULL_DATA sent: 6 (100.00% acknowledged)
# PULL_RESP(onse) datagrams received: 0 (0 bytes)
# RF packets sent to concentrator: 0 (0 bytes)
# TX errors: 0
### [GPS] ###
# Invalid gps time reference (age: 1478765433 sec)
# Manual GPS coordinates: latitude 51.1, longitude 5.9, altitude 23 m
##### END #####
```


