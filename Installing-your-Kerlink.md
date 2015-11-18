
Kerlink Gateway installation guide
_With many thanks to Martin Aarts!_

1. Make sure you order the correct Kerlink IoT Station, it is dependent on your country wchich frequency it uses.
2. After unpacking the Kerlink IoT Station, open the case by putting a screwdriver in the top notch (where the antenna is located).
3. Connect a UTP network cable on the green connector, cable colors are noted next to the connector. You can use an existing cable by cutting of the connector of one side, or you need to make a new cable including attaching the connector (watch the coloring scheme).
4. Attach the UTP cable attached to the Kerlink IoT Station to the “data & power Out” port of the power adapter. Connect the “data IN” port of the power adapter to your existing network. If you use POE (Power Over Ethernet) switches, the power adapter is not needed.
5. After powering on, check your DHCP server which IP-address the gateway uses. The LEDs inside the gateway do not work by default, they only work for about a minute after shortly pressing the “Test” button. This includes the power LEDs.
6. Check if the gateway is only by directing a web browser to the IP-address of the gateway, for example http://10.1.0.117 (depending on the given IP-address by the DHCP server). The page will show “Hello World!” if the gateway is responding.
7. Next you need to install the DHCP patch and the Packet forwarder software to prepare the gateway for a LoRaWAN network connection. From the page http://iot.semtech.com/resources/Kerlink_IoT_station_page download “Kerlink_IoT_LoRa_update_DHCP.zip” and unzip the file to your harddrive.
8. Install the update with the DHCP patch in the following way: 
- Copy the content of the 'kerlink_IoT_LoRa_update_DHCP' directory onto an empty USB flash drive formatted in FAT-32. Make sure there is no .log file.
- Plug USB key into the gateway.
- Wait for 5 min.
- Unplug the key and check that a .log file has appeared. The file should contain a text like "WirmaV2 0x080XXXXX updated". This log file prevents any further installation on the gateways to avoid cyclic reboots. If you want to restart the update on same gateway, remove this log file from the key and proceed again. (it is not needed if you update another gateway).
9. Logon to the gateway by using the SSH protocol, on a Mac or Linux system just use the Terminal and run “ssh 10.1.0.117” (substitute by the correct IP-address).  On a Windows PC you can use the Putty.exe program, which can be downloaded at http://www.chiark.greenend.org.uk/~sgtatham/putty/download.html . Login with the user “root” and default password “root”.
10. The next steps require some knowledge of using the command-line in Linux or Mac . Change the following file:
>/mnt/fsuser-1/forwarder_network_demo/local_conf.json
to the example below. Make sure not to change the unique “gateway_ID” that is already in the file.
{
/* Put there parameters that are different for each gateway (eg. pointing one gateway to a test server while the others stay in production) */
/* Settings defined in global_conf will be overwritten by those in local_conf */	
"gateway_conf": {
	"gateway_ID": "AA555A00XXXXXXXX",
	"serv_port_up": 1700,
	"serv_port_down": 1700,
	"server_address": "croft.thethings.girovito.nl",
	"forward_crc_valid": true,
	"forward_crc_error": false,
	"forward_crc_disabled": true
	}
}
11. Restart the gateway by pressing the reset button inside the casing or issue the commands in your SSH terminal:
cd /mnt/fsuser-1/forwarder_network_demo
./pkt-fwd.sh
This will reload the packet forwarder software with the new settings. Resetting the gateway by disconnecting power does not work, as the gateway contains an internal power backup battery.
12. At your firewall system make sure the external IP-address used will map port 1700 to the internal IP-address of the gateway. This is needed when using NAT for internal IP-address translation. 
13. Use the following command on the gateway to check whether data is being sent and received:
tcpdump -i eth0 -n -vvvX host 54.229.214.112
The output must be somewhat similar to the following, check out if inbound as well as outbound traffic is shown:
>tcpdump -i eth0 -n -vvvX host 54.229.214.112
tcpdump: listening on eth0, link-type EN10MB (Ethernet), capture size 65535 bytes
19:35:07.292396 IP (tos 0x0, ttl 64, id 35878, offset 0, flags [DF], proto UDP (17), length 40)
    10.1.0.117.37763 > 54.229.214.112.1700: [udp sum ok] UDP, length 12
	0x0000:  4500 0028 8c26 4000 4011 96d3 0a01 0075  E..(.&@.@......u
	0x0010:  36e5 d670 9383 06a4 0014 c7e2 0169 7302  6..p.........is.
	0x0020:  aa55 5a00 0806 0529                      .UZ....)
19:35:07.321453 IP (tos 0x20, ttl 46, id 1404, offset 0, flags [DF], proto UDP (17), length 32)
    54.229.214.112.1700 > 10.1.0.117.37763: [udp sum ok] UDP, length 4
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

14. The gateway does not sent any data by itself to auto include it on the status pages of The Things Network. It only does this after a network node is sending data via the gateway.
