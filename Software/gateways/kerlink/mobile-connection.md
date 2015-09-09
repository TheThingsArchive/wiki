It is possible to connect the Kerlink to a gprs/3g connection. This maybe eligible when LAN security is tight.
Sim card detection is only done at boot time. Insert the SIM card in the powered off LoRa station.

Set your APN settings in /etc/sysconfig/network:
	
    # Selector operator APN
	GPRSAPN=m2minternet
	# Enter pin code if activated
	GPRSPIN=
	# Update /etc/resolv.conf to get dns facilities
	GPSDNS=yes
	# PAP authentication
	GPRSUSER=kerlink
	GPRSPASSWORD=password
    
---
**Warning!** There is a bug in the software: When GPRSUSER and GPRSPASSWORD needs to stay empty the Kerlink does funny things and no connection is made. 
To resolve this problem, please apply [this patch](Software/gateways/kerlink/mobile-connection/attachment/1/dota_update_gprs_script.tar.gz).
---