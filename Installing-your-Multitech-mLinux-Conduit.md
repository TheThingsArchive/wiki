This page assumes you have got:

* MTCDT-210L-US-EU-GB Multitech mLinux Conduit (non-cellular)
* MTAC-LORA-868 accessory card which goes into the Conduit.
* MTKIT-LORA-868 dev kit to program the mDot card.

You will also need:

* An SSH terminal.  I use TeraTerm on Windows.
* An ARM mbed user account so you can program your mDot. It's free, sign up here: [https://developer.mbed.org/explore/](https://developer.mbed.org/explore/)

Instructions:

1.  Install the MTAC-LORA-868 card into the Conduit - the instructions are packed with the MTAC-LORA-868 card, or here: [http://www.multitech.net/developer/products/accessory-cards/installing-an-accessory-card/](http://www.multitech.net/developer/products/accessory-cards/installing-an-accessory-card/)
2.  Follow the instructions on the multitech.net pages to get your Conduit set up: [http://www.multitech.net/developer/software/mlinux/getting-started-with-conduit-mlinux/](http://www.multitech.net/developer/software/mlinux/getting-started-with-conduit-mlinux/)
3.  Now get the latest versions of the various bits of software: [http://www.multitech.net/developer/software/mlinux/using-mlinux/flashing-mlinux-firmware-for-conduit/](http://www.multitech.net/developer/software/mlinux/using-mlinux/flashing-mlinux-firmware-for-conduit/) and [http://www.multitech.net/developer/software/mlinux/using-mlinux/upgrade-lora-server/](http://www.multitech.net/developer/software/mlinux/using-mlinux/upgrade-lora-server/)
4.  Verify that your mDot and Conduit can talk to each other: [http://www.multitech.net/developer/software/lora/getting-started-with-lora-conduit-mlinux/](http://www.multitech.net/developer/software/lora/getting-started-with-lora-conduit-mlinux/)

    You may want to play a while with some sample mbed code at this point.  There are plenty of samples on the mbed site to familiarise yourself with: [https://developer.mbed.org/platforms/MTS-mDot-F411/](https://developer.mbed.org/platforms/MTS-mDot-F411/)

    When you are ready to connect up to The Things Network:

5.  Convert your Conduit to a packet forwarder.  We have found @kersig's poly packet forwarder to be more relaible than the gps_pkt_fwd or basic_pkt_fwd that come as standard. Download the poly_pkt_fwd install packet onto your Conduit from here: [https://github.com/kersing/packet_forwarder/blob/master/multitech-bin/poly-packet-forwarder_2.1-r1_arm926ejste.ipk](https://github.com/kersing/packet_forwarder/blob/master/multitech-bin/poly-packet-forwarder_2.1-r1_arm926ejste.ipk)
6.  Also download the sample global_conf.json file for an 868 card here: [https://github.com/kersing/packet_forwarder/blob/master/poly_pkt_fwd/global_conf_multitech-eu868.json](https://github.com/kersing/packet_forwarder/blob/master/poly_pkt_fwd/global_conf_multitech-eu868.json)
7.  Install poly_pkt_fwd on your Conduit using **opkg install poly-packet-forwarder_2.1-r1_arm926ejste.ipk**
8.  Copy the sample global_conf file to **/var/config/lora** and rename it **global_conf.json**
9.  Edit **/var/config/lora/global_conf.json**.  Go to the bottom and edit the gateway_conf section:
<pre><code>"gateway_conf": {
        /* change with default server address/ports, or overwrite in local_conf.
        "gateway_ID": "008000000000A052",
        /* Devices */
        "gps": true,
        "beacon": false,
        "monitor": false,
        "upstream": true,
        "downstream": true,
        "ghoststream": false,
        "radiostream": true,
        "statusstream": true,
        /* node server */
        "server_address": "127.0.0.1",
        "serv_port_up": 1680,
        "serv_port_down": 1681,
        /* node servers for poly packet server (max 4) */
        "servers":
        [ { "server_address": "127.0.0.1",
            "serv_port_up": 1680,
            "serv_port_down": 1681,
            "serv_enabled": false },
          { /* "server_address": "iot.semtech.com", */
            "server_address": "80.83.53.26",
            "serv_port_up": 1680,
            "serv_port_down": 1680,
            "serv_enabled": false },
          { /* "server_address": "croft.thethings.girovito.nl", */
            "server_address": "54.229.214.112",
            "serv_port_up" : 1700,
            "serv_port_down" : 1700,
            "serv_enabled": true } ],
        /* adjust the following parameters for your network */
        "keepalive_interval": 10,
        "stat_interval": 30,
        "push_timeout_ms": 100,
        /* forward only valid packets */
        "forward_crc_valid": true,
        "forward_crc_error": false,
        "forward_crc_disabled": true,
        /* GPS configuration */
        "gps_tty_path": "/dev/ttyAMA0",
        "fake_gps" : true,
        "ref_latitude" : 51.441026,
        "ref_longitude" : -0.967420,
        "ref_altitude" : 61,
        /* Ghost configuration */
        "ghost_address": "127.0.0.1",
        "ghost_port": 1918,
        /* Monitor configuration */
        "monitor_address": "127.0.0.1",
        "monitor_port": 2008,
        "ssh_path": "/usr/bin/ssh",
        "ssh_port": 22,
        "http_port": 80,
        "ngrok_path": "/usr/bin/ngrok",
        "system_calls": ["df -m","free -h","uptime","who -a","uname -a"],
        /* Platform definition, put a asterix here for the system value, max 24
        "platform": "*",
        /* Email of gateway operator, max 40 chars*/
        "contact_email": "mark.stanley@someconsultants.com",
        /* Public description of this device, max 64 chars */
        "description": "TTN Reading UK 001"
    }</code></pre>

    You need to change the following:
    
    * <em>gateway_ID</em> can be obtained by running **mts-io-sysfs show lora/eui**
    * <em>ref_latitude</em>, <em>ref_logitude</em> and <em>ref_altitude</em> can all be got from Google
    * <em>contact_email</em>
    * <em>description</em>
    
    The <em>serv_port_down</em>, <em>serv_port_up</em> and <em>server_address</em> are correct for The Things Network

8.  Edit **/etc/init.d/lora-network-server**.  Around line 17 you need to change from using <em>basic_pkt_fwd</em> to using <em>poly_pkt_fwd</em>:
	<pre><code>pkt_fwd=/opt/lora/poly_pkt_fwd</code></pre>

    Go down to where the network server is started (around line 57).  Comment it out:
    <pre><code># start network server
	#start-stop-daemon --start --background --make-pidfile \
    #    --pidfile $net_server_pidfile --exec $net_server -- \
    #    -c $conf_file --lora-eui $lora_eui --lora-path $run_dir --db $conf_db \
    #    --noconsole -l $net_server_log
    #sleep 1</code></pre>
    
    Similarly comment out the network server line in do_stop() a little further down:
	<pre><code>do_stop() {
    echo -n "Stopping $NAME: "
    #start-stop-daemon --stop --quiet --oknodo --pidfile $net_server_pidfile --r

    start-stop-daemon --stop --quiet --oknodo --pidfile $pkt_fwd_pidfile --retry
    rm -f $net_server_pidfile $pkt_fwd_pidfile
    echo "OK"
	}</code></pre>
    
9.  Restart the lora-network server:   <strong>/etc/init.d/lora-network-server restart </strong>
10. If things have gone well you should be able to go to [http://thethingsnetwork.org/api/v0/gateways/](http://thethingsnetwork.org/api/v0/gateways/) and see your gateway appear in the list.  Congratulations!


   
   
   