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

5.  Convert your Conduit to a packet forwarder.  We have found @kersing's poly packet forwarder to be more relaible than the gps_pkt_fwd or basic_pkt_fwd that come as standard. Download the poly_pkt_fwd install packet onto your Conduit from here: [https://github.com/kersing/packet_forwarder/blob/master/multitech-bin/poly-packet-forwarder_2.1-r3_arm926ejste.ipk](https://github.com/kersing/packet_forwarder/blob/master/multitech-bin/poly-packet-forwarder_2.1-r3_arm926ejste.ipk?raw=true)
6.  Install poly_pkt_fwd on your Conduit using **opkg install poly-packet-forwarder_2.1-r3_arm926ejste.ipk**

**NOTE**: The package requires DNS to be set up on the conduit to find the TTN backend!

7.  Edit **/var/config/lora/local_conf.json**.
<pre><code>
{
/* Put there parameters that are different for each gateway (eg. pointing one gateway to a test server while the others stay in production) */
/* Settings defined in global_conf will be overwritten by those in local_conf */
    "gateway_conf": {
	/* you must pick a unique 64b number for each gateway (represented by an hex string) */
        "gateway_ID": "AA555A000004BABA",
        /* Email of gateway operator, max 40 chars*/
        "contact_email": "operator@gateway.tst", 
        /* Public description of this device, max 64 chars */
        "description": "Update me",
	/* Enter VALID GPS coordinates below before enabling fake GPS */
        "fake_gps": false,
        "ref_latitude": 10,
        "ref_longitude": 20,
        "ref_altitude": -1
    }
}</code></pre>

    You need to change the following:
    
    * <em>gateway_ID</em> can be obtained by running **mts-io-sysfs show lora/eui**
    * <em>ref_latitude</em>, <em>ref_logitude</em> and <em>ref_altitude</em> can all be got from Google
    * <em>fake_gps</em> set to true
    * <em>contact_email</em>
    * <em>description</em>
    
8.  Edit **/etc/init.d/lora-network-server**.  Around line 17 you need to change from using <em>basic_pkt_fwd</em> to using <em>poly_pkt_fwd</em>:
	<pre><code>pkt_fwd=/opt/lora/poly_pkt_fwd</code></pre>

    Go down to where the network server is started (around line 57).  Comment it out:

	<pre><code>	   #start network server
		#start-stop-daemon --start --background --make-pidfile \
		#--pidfile $net_server_pidfile --exec $net_server -- \
		#-c $conf_file --lora-eui $lora_eui --lora-path $run_dir --db $conf_db \
		#--noconsole -l $net_server_log
		#sleep 1</code></pre>
    
    Make sure the -c option for the new packet_forwarder under <pre>start packet forwarder</pre> is set to 
    <pre><code>-c $conf_dir</code></pre> instead of <pre>-c $run_dir/1</pre>
	
    Note: The /1 is introduced in newer versions of Multitech software to support multiple cards.
    If you do not remove the /1, the network server will not start. 
	    

    Similarly comment out the network server line in <em>do_stop()</em> a little further down:
    
    <pre><code>   do_stop() {
		echo -n "Stopping $NAME: "
    	#start-stop-daemon --stop --quiet --oknodo --pidfile $net_server_pidfile --r
    	start-stop-daemon --stop --quiet --oknodo --pidfile $pkt_fwd_pidfile --retry
    	rm -f $net_server_pidfile $pkt_fwd_pidfile
    	echo "OK"
    	}</code></pre>
    
9.  Restart the lora-network server:   **/etc/init.d/lora-network-server restart**
10. If things have gone well you should be able to go to [http://thethingsnetwork.org/api/v0/gateways/](http://thethingsnetwork.org/api/v0/gateways/) and see your gateway appear in the list.  Congratulations!
