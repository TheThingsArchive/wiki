It is possible to connect the Kerlink to a gprs/3g connection. This maybe eligible when LAN security is tight.
Sim card detection is only done at boot time. Insert the SIM card in the powered off LoRa station.

Set your APN settings in /etc/sysconfig/network:
># Selector operator APN
>GPRSAPN=m2minternet
># Enter pin code if activated
>GPRSPIN=
># Update /etc/resolv.conf to get dns facilities
>GPRSDNS=yes
># PAP authentication
>GPRSUSER=kerlink
>GPRSPASSWORD=password