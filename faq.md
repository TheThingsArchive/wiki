test 
### Are you a network operator?
No, we're not a network operater, at least not in the traditional way.
We're a community of individuals and organisations (both non-profit and
commercial) owning our own network equipment and developing shared open-source
software for everyone to use. We aim at building a global network together,
but don't have any central power nor usage fees.

Have a look at our [manifesto](https://github.com/TheThingsNetwork/Manifest).


### Do I have to pay money to use the network?
No, the network is free to use. You do need special hardware to connect to the
network (see below). If there is no coverage in your area yet, you might want
to contribute by buying and setting up your own [gateway](http://thethingsnetwork.org/wiki/Hardware/OverviewGateways)
for yourself and others to use.


### Do I need to buy a gateway / router to get started?
If there is coverage in your area already, you don't need a gateway. You can
search for local communities at [http://thethingsnetwork.org](http://thethingsnetwork.org/#communities)
or on the [forum](http://forum.thethingsnetwork.org/).
If there's no gateway nearby yet, you might want to contribute by buying one!


### Is it secure (for our company network) to connect gateway to it?
..


### Can any gateway owner read packages sent via TTN?
No, since the data is encrypted.


### Who has to buy a server (Network Server - NS)?
In general, there's no need for servers. The Things Network Foundation will
provide basic network routing infrastructure to get started. If you're a heavy
network user or are running a local community, you might want to setup your own
network server and routing instances to help support the load.
Also, if you want to build your own (web)applications you might (obviously) need
a server or cloud instance to run your application.


### How fast is the network? What data-rate can we use?
The network uses LoRaWAN, which was developed for long-range data-networks that
don't need a constant connection. The protocol allows for different speeds and
data packet sizes, all relatively small and depending on various instances.
Maximum packet sizes vary between **55 and 222 bytes**. Air time is often
restricted to 1%; therefore a node may need to wait **5 seconds up to 4 minutes** before sending another packet on the same sub-band,
depending on packet size, distance from the nearest gateway and other factors. Furthermore, the TTN fair access 
policy will limit this even more, allowing an average of **30 seconds per day** time on air **per device**.

For more info, see [this forum post](http://forum.thethingsnetwork.org/t/parse-data-with-node-red/835/2).


### How do I connect to the network? Does this work on my smartphone?
LoRaWAN is a new technology that - like wifi or bluetooth - needs special
hardware in order to connect. It's meant for `things` rather than human-operated
devices, so it won't work on your smartphone. You'll need to buy `things` or
nodes supporting LoRaWAN.

-> link to wiki page


### What about Security? Can the NSA control my Things?
The network protocol uses end-to-end AES encryption by default. The key will be
known by your devices and the application server. If you're using our example
application server and keys (or one hosted by the NSA), others will be able to
see your data, but if you're running your own handler and application the data
will be unreadable for gateway owners, routers and anyone in between.

As with everything on the internet, you might want to consider adding additional
layers of encryption on top of the existing connection.


### Why is X not decentralised?
Truly decentralising a network and community is a hard problem; there's always
a trade-off between decentralisation and practicality. We try to build a truly
decentralised network, but most succesful community projects still have some
semi-centralised parts to get things going.

All [software](https://github.com/TheThingsNetwork) is open-source, all hardware
designs developed by TTN are open-source, the network is (semi-)decentralised
without central authority, The Things Network is a non-profit Foundation, and we
try to be as much community-driven as is practical.
We also try to make much of the (public) data available
([website source-code](https://github.com/TheThingsNetwork/TheThingsNetwork.org),
[live gateways](http://ttnstatus.org/gateways),
[wiki data](https://github.com/TheThingsNetwork/wiki)).
However, we do have a core team pushing forward on software development and
community management, and we do host a few websites ourselves.


### I want to help out! What can I do?
..

-> start local community (link), start building, start contributing software



### I have another question.
Search on [the Forum](http://forum.thethingsnetwork.org/) for related questions. If you can't find the answer, post a new question so everyone can learn and contribute.


## Note to FAQ writers
Please keep answers short (1 paragraph) and link to related wiki/forum topics. Only add answers to FREQUENTLY asked questions, keep the list capped at 20 items max. Thanks for contributing!*
