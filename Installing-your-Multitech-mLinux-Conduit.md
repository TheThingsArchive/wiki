This page assumes you have got:
* MTCDT-210L-US-EU-GB Multitech mLinux Conduit (non-cellular)
* MTAC-LORA-868 accessory card which goes into the Conduit.
* MTKIT-LORA-868 dev kit to program the mDot card.

You will also need:
An SSH terminal.  I use TeraTerm on Windows.
An ARM mbed user account so you can program your mDot. It's free, sign up here: [https://developer.mbed.org/explore/](https://developer.mbed.org/explore/)

1.  Install the MTAC-LORA-868 card into the Conduit - the instructions are packed with the MTAC-LORA-868 card, or here: [http://www.multitech.net/developer/products/accessory-cards/installing-an-accessory-card/](http://www.multitech.net/developer/products/accessory-cards/installing-an-accessory-card/)
2.  Follow the instructions on the multitech.net pages to get your Conduit set up: [http://www.multitech.net/developer/software/mlinux/getting-started-with-conduit-mlinux/](http://www.multitech.net/developer/software/mlinux/getting-started-with-conduit-mlinux/)
3.  Now get the latest versions of the various bits of software: [http://www.multitech.net/developer/software/mlinux/using-mlinux/flashing-mlinux-firmware-for-conduit/](http://www.multitech.net/developer/software/mlinux/using-mlinux/flashing-mlinux-firmware-for-conduit/) and [http://www.multitech.net/developer/software/mlinux/using-mlinux/upgrade-lora-server/](http://www.multitech.net/developer/software/mlinux/using-mlinux/upgrade-lora-server/)
4.  Verify that your mDot and Conduit can talk to each other: [http://www.multitech.net/developer/software/lora/getting-started-with-lora-conduit-mlinux/](http://www.multitech.net/developer/software/lora/getting-started-with-lora-conduit-mlinux/)

You may want to play a while with some sample mbed code at this point.  There are plenty of samples on the mbed site to familiarise yourself with: [https://developer.mbed.org/platforms/MTS-mDot-F411/](https://developer.mbed.org/platforms/MTS-mDot-F411/)

When you are ready to connect up to The Things Network:

5.  Convert your Conduit to a Basic packet forwarder: [http://www.multitech.net/developer/software/lora/conduit-mlinux-convert-to-basic-packet-forwarder/](http://www.multitech.net/developer/software/lora/conduit-mlinux-convert-to-basic-packet-forwarder/)
6.  Now we need to make some changes for TTN because the basic packet forwarder doesn't work too well with the TTN API because it doesn't send longitude and latitude figures, and your gateway won't appear in the TTN REST API.
7.  Edit /var/config/lora/global_conf.json  Go to the bottom and edit the gateway_conf section:

       "gateway_conf" :
        {
                "forward_crc_disabled" : false,
                "forward_crc_error" : true,
                "forward_crc_valid" : true,
                "fake_gps" : true,
                "ref_latitude" : 51.441026,
                "ref_longitude" : -0.967420,
                "ref_altitude" : 61,
                "gateway_ID" : "008000000000A052",
                "keepalive_interval" : 12,
                "push_timeout_ms" : 120,
                "serv_port_down" : 1700,
                "serv_port_up" : 1700,
                "server_address" : "croft.thethings.girovito.nl",
                "stat_interval" : 20,
                "synch_word" : 52
        }

    Your gateway ID can be obtained by running mts-io-sysfs show lora/eui
    Your latitude, logitude and altitude can all be got from Google
    The serv_port_down, serv_port_up and server_address are correct for The Things Network

8.  Edit /etc/init.d/lora-network-server  Around line 17 you need to change from using basic_pkt_fwd to using gps_pkt_fwd:
   
     pkt_fwd=/opt/lora/gps_pkt_fwd
   
    Go down to where the packet server is started.  The parameters for basic_pkt_fwd and gps_pkt_fwd are different so edit it:

	  # start packet forwarder
      start-stop-daemon --start --background --make-pidfile \
          --pidfile $pkt_fwd_pidfile --exec $pkt_fwd -- \
          -c $conf_dir

9.  Restart the lora-network server:   /etc/init.d/lora-network-server restart
10. If things have gone well you should be able to go to http://thethingsnetwork.org/api/v0/gateways/ and see your gateway appear in the list.  Congratulations!


   
   
   