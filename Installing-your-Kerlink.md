
Kerlink Gateway installation guide
_With many thanks to Martin Aarts!_

1. Make sure you order the correct Kerlink IoT Station, as it is dependent on your country which frequency are allowed and will therefore determine which frequencies the gateway uses.
2. After unpacking the Kerlink IoT Station, open the case by putting a screwdriver in the top notch (where the antenna is located).
3. Connect a UTP network cable on the green connector, cable colors are noted next to the connector. You can use an existing cable by cutting of the connector of one side, or you need to make a new cable including attaching the connector (watch the coloring scheme).
4. Attach the UTP cable attached to the Kerlink IoT Station to the “data & power Out” port of the power adapter. Connect the “data IN” port of the power adapter to your existing network. If you use POE (Power Over Ethernet) switches, the power adapter is not needed.
5. After powering on, check your DHCP server which IP-address the gateway uses. The LEDs inside the gateway do not work by default, they only work for about a minute after shortly pressing the “Test” button. This includes the power LEDs.
6. Check if the gateway is on by directing a web browser to the IP-address of the gateway, for example http://10.1.0.117 (depending on the given IP-address by the DHCP server). The page will show “Hello World!” if the gateway is responding.
7. Next you need to download [thethingsnetwork packetforwarder](https://github.com/TheThingsNetwork/kerlink-station-firmware/raw/master/dota/dota_thethingsnetwork_v1.1.tar.gz) and [produsb.zip](https://github.com/TheThingsNetwork/kerlink-station-firmware/raw/master/dota/produsb.zip) from our github. 
8. Install the update the following way: 
- Copy the content of produsb.zip and the dotafile onto an empty USB flash drive formatted in FAT-32. Make sure there is no .log file.
- Plug USB key into the gateway.
- Wait for 5 min. During this time the gateway will reboot itself.
- Unplug the key and check that a .log file has appeared. The file should contain the following text "WirmaV2 0x080XXXXX updated". This log file prevents any further installation on the gateways to avoid cyclic reboots. To redo the update on same gateway, remove this log file from the flash drive reinsert it into the gateway USB. (it is not needed if you update another gateway).
9. Logon to the gateway by using the SSH protocol, on a Mac or Linux system just use the Terminal and run “ssh 10.1.0.117” (substitute by the correct IP-address).  On a Windows PC you can use the Putty.exe program, which can be downloaded at http://www.chiark.greenend.org.uk/~sgtatham/putty/download.html . Login with the user “root” and default password “root”.
10. At your firewall system make sure the external IP-address used will map port 1700 to the internal IP-address of the gateway. This is needed when using NAT for internal IP-address translation. 
11. Your are basically done now. For further checking continue to the next step. The following steps require some knowledge about how to use the command-line in Linux or Mac.
12. Use the following command on the gateway to check whether data is being sent and received:<br>
<strong>tcpdump -i eth0 -n -vvvX host 54.229.214.112</strong><br>
The output must be somewhat similar to the following, check out if inbound as well as outbound traffic is shown:
<pre><code> >tcpdump -i eth0 -n -vvvX host 54.229.214.112
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
</code></pre>
13. The gateway does not send any data by itself to auto include itself on the status pages of The Things Network. This will only happen after a network node has sent or is sending data via the gateway.



Want to connect the Kerlink over gprs/3g using a simcard? Follow [these steps](Software/gateways/kerlink/mobile-connection) instructions.


