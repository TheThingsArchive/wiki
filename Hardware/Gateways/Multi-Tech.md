# Multi-Tech MultiConnect Conduit

> [Multi-Tech](http://www.multitech.com/brands/multiconnect-conduit)

[[/uploads/Multi-Tech.png]]

> TODO: Fact-checking

**Price:**
**Availability:**
**Capacity:**
Suitable for indoor use  
**Installation:**
**Hardware:** proprietary  
**Software:**

Tested by The Things Network Crew.   

Discussion: We have one, results of our test will follow asap.
First draft of an [installation guide](/wiki/Installing-your-Multitech-mLinux-Conduit)

## Prerequisites

> TODO: Add links

This guide assumes you have:

* MTCDT-210L-US-EU-GB Multitech mLinux Conduit (non-cellular)
* MTAC-LORA-868 accessory card which goes into the Conduit.
* MTKIT-LORA-868 dev kit to program the mDot card.

You will also need:

* An SSH terminal
* An ARM mbed user account so you can program your mDot. It's free, sign up here: [https://developer.mbed.org/explore/](https://developer.mbed.org/explore/)

## Setup

> TODO: Include external info in this page

Install the MTAC-LORA-868 card into the Conduit - the instructions are packed with the MTAC-LORA-868 card, or [online](http://www.multitech.net/developer/products/accessory-cards/installing-an-accessory-card/)

> TODO: Photo

Follow the instructions on the [multitech.net pages](http://www.multitech.net/developer/software/mlinux/getting-started-with-conduit-mlinux/) to get your Conduit set up.

Now get the latest versions of the various bits of software:

* [mLinux Firmware](http://www.multitech.net/developer/software/mlinux/using-mlinux/flashing-mlinux-firmware-for-conduit/)
* [LoRa Server & Packet Forwarder](http://www.multitech.net/developer/software/mlinux/using-mlinux/upgrade-lora-server/)

Verify that your mDot and Conduit can talk to each other: [Guide](http://www.multitech.net/developer/software/lora/getting-started-with-lora-conduit-mlinux/)

You may want to play a while with some sample mbed code at this point. There are plenty of samples on the [mbed site](https://developer.mbed.org/platforms/MTS-mDot-F411/) to familiarise yourself with.

When you are ready to connect up to The Things Network, Convert your Conduit to a packet forwarder. We have found [@kersing's `poly packet forwarder`](https://github.com/kersing/packet_forwarder/tree/master/poly_pkt_fwd) to be more reliable than the `gps_pkt_fwd` or `basic_pkt_fwd` that come as standard. Download the poly_pkt_fwd install packet onto your Conduit from [here](https://github.com/kersing/packet_forwarder/blob/master/multitech-bin/poly-packet-forwarder_2.1-r2_arm926ejste.ipk?raw=true).

Install `poly_pkt_fwd` on your Conduit using **opkg install poly-packet-forwarder_2.1-r2_arm926ejste.ipk**

Edit **/var/config/lora/local_conf.json**:

```js
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
}
```

You need to change the following:

* `gateway_ID` can be obtained by running `mts-io-sysfs show lora/eui`
* `ref_latitude`, `ref_logitude` and `ref_altitude` can all be got from Google
* `fake_gps` set to `true`
* `contact_email`
* `description`

Edit `/etc/init.d/lora-network-server`. Around line 17 you need to change from using `basic_pkt_fwd` to using `poly_pkt_fwd`: `pkt_fwd=/opt/lora/poly_pkt_fwd`

Go down to where the network server is started (around line 57). Comment it out:

```sh
#start network server
#start-stop-daemon --start --background --make-pidfile \
#--pidfile $net_server_pidfile --exec $net_server -- \
#-c $conf_file --lora-eui $lora_eui --lora-path $run_dir --db $conf_db \
#--noconsole -l $net_server_log
#sleep 1</code></pre>
```

Make sure the `-c` option for the new packet_forwarder under `start packet forwarder` is set to `-c $conf_dir` instead of `-c $run_dir/1`

Note: The `/1` is introduced in newer versions of MultiTech software to support multiple cards. If you do not remove the `/1`, the network server will crash instantly.

Similarly comment out the network server line in `do_stop()` a little further down:

```sh
do_stop() {
echo -n "Stopping $NAME: "
#start-stop-daemon --stop --quiet --oknodo --pidfile $net_server_pidfile --r
start-stop-daemon --stop --quiet --oknodo --pidfile $pkt_fwd_pidfile --retry
rm -f $net_server_pidfile $pkt_fwd_pidfile
echo "OK"
}
```

Restart the lora-network server: `/etc/init.d/lora-network-server restart`

If things have gone well you should be able to go to [TTN's gateway overview](http://thethingsnetwork.org/api/v0/gateways/) and see your gateway appear in the list. Congratulations!
