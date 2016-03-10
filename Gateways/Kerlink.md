# Kerlink LoRa IoT Station

[[/uploads/Kerlink-LoRa-IoT.png]]

**Price:** on request, around 1.200 euro / 1.500 dollar.  
**Availability:** send inquiry to [Kerlink Sales](mailto:sales@kerlink.fr)  
**Capacity:** 8 band  
Suitable for outdoor mounting (IP67)  
**Installation:** requires a technician  
**Hardware:** proprietary  
**Software:** open source & proprietary  
Possibility to run on gprs/3g network

This is an industrial solution suitable for people who want to mount the gateway outside and who have sufficient technical skills to connect, mount and maintain the device themselves. We have tested the device and although we have remarks about the somewhat older software that is being used, this device will do the job. A trained software engineer will be able to update the device using the software from The Things Network.

## Setup

After unpacking the Kerlink IoT Station, open the case by putting a screwdriver in the top notch (where the antenna is located).

> TODO: Photo

Connect a UTP network cable on the green connector, cable colors are noted next to the connector. You can use an existing cable by cutting of the connector of one side, or you need to make a new cable including attaching the connector (watch the coloring scheme).

> TODO: Photo

Attach the UTP cable attached to the Kerlink IoT Station to the “data & power Out” port of the power adapter. Connect the “data IN” port of the power adapter to your existing network. If you use POE (Power Over Ethernet) switches, the power adapter is not needed.

> TODO: Photo

After powering on, check your DHCP server which IP-address the gateway uses. The LEDs inside the gateway do not work by default, they only work for about a minute after shortly pressing the “Test” button. This includes the power LEDs.

> TODO: Photo?

Check if the gateway is on by directing a web browser to the IP-address of the gateway, for example http://10.1.0.117 (depending on the given IP-address by the DHCP server). The page will show “Hello World!” if the gateway is responding.

## Configuration

To configure the gateway, you need to download The Things Network's [packet forwarder](https://github.com/TheThingsNetwork/kerlink-station-firmware/raw/master/dota/dota_thethingsnetwork_v1.1.tar.gz) and [produsb.zip](https://github.com/TheThingsNetwork/kerlink-station-firmware/raw/master/dota/produsb.zip) from our Github.

Install the update the following way:

- Copy the content of `produsb.zip` and the `dotafile` onto an empty USB flash drive formatted in `FAT-32`. Make sure there is no `.log` file.
- Plug the USB flash drive into the gateway.
- Wait for 5 min. During this time the gateway will reboot itself.
- Unplug the key and check that a `.log` file has appeared. The file should contain the following text `WirmaV2 0x080XXXXX updated`. This log file prevents any further installation on the gateways to avoid cyclic reboots.  
  To redo the update on same gateway, remove this log file from the flash drive reinsert it into the gateway USB. This is not needed if you update another gateway.

Logon to the gateway by using the SSH protocol, on a Mac or Linux system just use the Terminal and run `ssh root@10.1.0.117` (substitute by the correct IP-address).  On a Windows PC you can use the [Putty.exe](http://www.chiark.greenend.org.uk/~sgtatham/putty/download.html) program. Login with the user `root` and default password `root`.

At your firewall system make sure the external IP-address used will map port 1700 to the internal IP-address of the gateway. This is needed when using NAT for internal IP-address translation.

Your are basically done now. For further checking continue to the next step. The following steps require some knowledge about how to use the command-line in Linux or Mac.

Use the following command on the gateway to check whether data is being sent and received:

`tcpdump -i eth0 -n -vvvX host 54.229.214.112`

The output must be somewhat similar to the following, check out if inbound as well as outbound traffic is shown:

```
tcpdump: listening on eth0, link-type EN10MB (Ethernet), capture size 65535 bytes
19:35:07.292396 IP (tos 0x0, ttl 64, id 35878, offset 0, flags [DF], proto UDP (17), length 40)
    <span style="background-color: #FFFF00">10.1.0.117.37763 > 54.229.214.112.1700</span>: [udp sum ok] UDP, length 12
	0x0000:  4500 0028 8c26 4000 4011 96d3 0a01 0075  E..(.&@.@......u
	0x0010:  36e5 d670 9383 06a4 0014 c7e2 0169 7302  6..p.........is.
	0x0020:  aa55 5a00 0806 0529                      .UZ....)
19:35:07.321453 IP (tos 0x20, ttl 46, id 1404, offset 0, flags [DF], proto UDP (17), length 32)
    <span style="background-color: #FFFF00">54.229.214.112.1700 > 10.1.0.117.37763</span>: [udp sum ok] UDP, length 4
	0x0000:  4520 0020 057c 4000 2e11 2f66 36e5 d670  E....|@.../f6..p
	0x0010:  0a01 0075 06a4 9383 000c d978 0169 7301  ...u.......x.is.
	0x0020:  0000 0000 0000 0000 0000 0000 0000 eca3  ................
	0x0030:  5d2a                                     ]*
19:35:17.322399 IP (tos 0x0, ttl 64, id 35879, offset 0, flags [DF], proto UDP (17), length 40)
    10.1.0.117.37763 > 54.229.214.112.1700: [udp sum ok] UDP, length 12
	0x0000:  4500 0028 8c27 4000 4011 96d2 0a01 0075  E..(.'@.@......u
	0x0010:  36e5 d670 9383 06a4 0014 3bfa 0151 ff02  6..p......;..Q..
	0x0020:  aa55 5a00 0806 0529                      .UZ....)
19:35:17.351482 IP (tos 0x20, ttl 46, id 2099, offset 0, flags [DF], proto UDP (17), length 32)
    54.229.214.112.1700 > 10.1.0.117.37763: [udp sum ok] UDP, length 4
	0x0000:  4520 0020 0833 4000 2e11 2caf 36e5 d670  E....3@...,.6..p
	0x0010:  0a01 0075 06a4 9383 000c 4d90 0151 ff01  ...u......M..Q..
	0x0020:  0000 0000 0000 0000 0000 0000 0000 1efe  ................
	0x0030:  8b8d                                     ..
^C
4 packets captured
5 packets received by filter
0 packets dropped by kernel
```

The gateway does not automatically send data to auto include itself on the status pages of The Things Network. This will only happen after a network node has sent or is sending data via the gateway.

## Connecting to GPRS or 3G

It is possible to connect the Kerlink to a GPRS/3G connection. This maybe eligible when LAN security is tight.

SIM card detection is only done at boot time. Insert the SIM card in the powered off LoRa station.

Set your APN settings in /etc/sysconfig/network (see [Provider Settings](#provider-settings)):

```sh
# Selector operator APN
GPRSAPN=m2minternet
# Enter pin code if activated
GPRSPIN=
# Update /etc/resolv.conf to get dns facilities
GPSDNS=yes
# PAP authentication
GPRSUSER=kerlink
GPRSPASSWORD=password

# Bearers priority order
BEARERS_PRIORITY="ppp0,eth0,eth1"
```

Configure the autoconnect in /knet/knetd.xml

```xml
<!-- ############## local device configuration ############## -->
<LOCAL_DEV role="KNONE"/>

<!-- ############## connection parameters ############## -->
<!-- enable the autoconnect feature (YES/NO) -->
<CONNECT auto_connection="YES" />
<!-- frequency of connection monitoring -ping- (in seconds) -->
<CONNECT link_timeout="30"/>
<!-- DNS servers will be pinged if commented or deleted. Some operators can block the ping on there DNS servers -->
<CONNECT ip_link="192.168.4.90"/>

<!-- ############## default area for connection policy ############## -->

<AREA id="default">
<ACCESS_POINT bearer="gprs" />
</AREA>    
```

> **Warning!** There is a bug in the software: When GPRSUSER and GPRSPASSWORD needs to stay empty the Kerlink does funny things and no connection is made.
To resolve this problem, please apply [this patch](https://github.com/TheThingsNetwork/kerlink-station-firmware/blob/master/dota/dota_update_gprs_script.tar.gz?raw=true).

### Provider Settings

#### KPN - NL

```sh
# Selector operator APN
GPRSAPN=portalmmm.nl
# Enter pin code if activated
GPRSPIN=
# Update /etc/resolv.conf to get dns facilities
GPSDNS=yes
# PAP authentication
GPRSUSER=internet
GPRSPASSWORD=internet
```
