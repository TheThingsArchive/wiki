# Connect a Gateway

[[/uploads/TTN-Simple-Deployment.png]]

## Server Addresses

Choose the Router instance depending on the frequency plan used in your region. Have a look at the [LoRaWAN Regional Parameters 1.0 Document](https://www.lora-alliance.org/For-Developers/LoRaWANDevelopers) and the wiki page on [country specific frequency plans](/wiki/LoRaWAN/Frequencies/By-Country).

```
router.eu.thethings.network # EU 433 and EU 863-870
router.us.thethings.network # US 902-928
router.cn.thethings.network # China 470-510 and 779-787
router.au.thethings.network # Australia 915-928 MHz
router.as.thethings.network # Southeast Asia 923 MHz
router.as1.thethings.network # Southeast Asia 920-923 MHz
router.as2.thethings.network # Southeast Asia 923-925 MHz
router.kr.thethings.network # Korea 920-923 MHz
```

The geographical location of a server and the supported frequency plans are two different things. As the LoRa Alliance publishes more region specifications, we will deploy Routers in datacenters near that region. Also, the yet to be implemented decentralized routing architecture enables community operators to provide routing infrastructure to The Things Network.

## Configuring Your Gateway

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

### Don't use `ping`
The routers are load balancers that do not respond to `ping` commands.

### Inspecting the port
Your gateway communicates with the routers via an UDP connection on port 1700. So if we listen to this port we should see packets being transferred.

```
sudo tcpdump -AUq port 1700
```
If your gateway is up and running this feed would (at least) give a stat packet every 30-60 seconds:

```
08:07:33.801265 IP 192.168.178.20.47497 > xx.xxx.xx.xxx.1700: UDP, length 221
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
