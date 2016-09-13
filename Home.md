# Welcome!

The Things Network is a global, open, crowd-sourced Internet of Things data network.

To find out why we started all this, read our [Manifesto](https://github.com/TheThingsNetwork/Manifest). A lot of information is also available on our [website](http://thethingsnetwork.org), [forum](http://forum.thethingsnetwork.org/) and [Github](https://github.com/TheThingsNetwork).

>Looking for the old wiki? It's archived on [GitHub](https://github.com/TheThingsNetwork/wiki/tree/master).

# Overview

The Things Network uses the [[LoRaWAN|LoRaWAN/Overview]] network technology to provide **low power** wireless connectivity over **long range**.

If we look at The Things Network from a high level, we can distinguish the following components:

[[/uploads/TTN-Overview.jpg]]

* **[[Nodes:|Hardware/Nodes/Overview]]** Simple devices that are deployed "in the wild". They can do measurements, collect data or perform actions. Nodes can broadcast or receive small messages (usually about 1/10 the size of an SMS) either periodically (some devices broadcast every couple of minutes, other devices only once in a number of hours).
* **[[Gateways:|Hardware/Gateways/Overview]]** Antennas that receive broadcasts from Nodes and send data back to Nodes. Gateways are connected to the Internet and communicate with The Things Network's servers. Gateways have a long range, so they can provide connectivity to nodes that are multiple kilometers away.
* **[[The Things Network Backend|Backend/Overview]]** route messages from Nodes to the right Application, and back.
* **Your Application** connects to The Things Network's servers to receive messages from and send messages to your Nodes. What you do with it is entirely up to you!

# Getting Involved

There are many ways to get involved with The Things Network:

* [[Starting or joining a community|Getting_Involved/Community]] doesn't require any technical knowledge. Find an existing local community, or start a new community with other interested people.
* [[Expanding the network|Getting_Involved/Infrastructure]] by contributing infrastructure components, such as gateways (the network needs antennas), a good location for the gateways (maybe you have a nice high building) or servers (to route data between gateways and applications). Any contribution can help your local community and the global TTN community grow the network.
* [[Using the Network|Getting_Involved/UsingTheNetwork]]: **build things** to help solve real world issues with the network. You can help designing use cases, or developing them, or manufacturing them. If possible, share your knowledge and projects with others on any of the communication channels.
* [[Developing the network's backend|Getting_Involved/Development]] is extremely important, especially because the network is growing fast. We need efficient software to route data between gateways and applications. This does require some good programming skills.

**You are the network. Let's build this thing together.**

# Information Sources

Next to this wiki, we have a number of other places where you can find the information you're looking for.

* [The Labs](https://thethingsnetwork.org/labs/)
* [The Forum](https://www.thethingsnetwork.org/forum)
* [Github](https://github.com/TheThingsNetwork) ([backend](https://github.com/TheThingsNetwork/ttn), [examples](https://github.com/TheThingsNetwork/examples), [sdk](https://github.com/TheThingsNetwork/sdk))
* [Slack](http://slack.thethingsnetwork.org/)
* [Our story](https://medium.com/@wienke/the-things-network-building-a-global-iot-data-network-in-6-months-adc2c0b1ae9b)

>Beware of references to the deprecated REST API (sometimes referred to as "Croft/Jolie", using URLs starting with `thethingsnetwork.org/api`); see [[Connect an Application| Backend/Connect/Application]] to use MQTT instead.

# Social Media

* [Twitter](https://twitter.com/thethingsntwrk)
* [Facebook](https://www.facebook.com/thethingsnetwork)

Local communities often have their own social media pages, you can find these on their community pages.

# External Material

* [[Media Attention|External/Media]]
* [[Blog Posts|External/Blogs]]