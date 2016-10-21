# The Things Network Backend

[[/uploads/Infrastructure-Overview.png]]

## Architecture

The backend systems of The Things Network are responsible for routing Internet of Things data between devices and applications. A typical Internet of Things network requires gateways as a bridge between specific radio protocols and the Internet. In cases where the devices themselves support the IP stack, these gateways only have to forward packets to the Internet. Non-IP protocols such as LoRaWAN require some form of routing and processing before messages can be delivered to an application. The Things Network is positioned between the gateways and the applications (see the figure below) and takes care of these routing and processing steps.

[[/uploads/Routing-Services.png]]

<center>_The Things Network’s different routing service components:  
Gateway, Router, Broker, NetworkServer, Handler and Application_ </center>
</p>

The Things Network's vision is to perform all these routing functions in a decentralized and distributed way. Any interested party should be able to set up their own network and their own part of the backend, allowing them to participate in the global community network. In order to decentralize the network, it was split up into a number of components shown in the figure above. To simplify, components are only shown once, even though it is possible to have one-to-many or many-to-many relations between components. These kind of relations are indicated with asterisks.

Nodes (N in the top Figure) broadcast LoRaWAN messages over the LoRa radio protocol. These messages are received by a number of Gateways (G). The Gateway is a piece of hardware that forwards radio transmissions to the backend. It is connected to one Router (R). The Router is responsible for managing the gateway's status and for scheduling transmissions. Each Router is connected to one or more Brokers (B). Brokers are the central part of The Things Network. Their responsibility is to map a device to an application, to forward uplink messages to the correct application and to forward downlink messages to the correct Router (which forwards them to a Gateway). The Network Server (NS) is responsible for functionality that is specific for LoRaWAN. A Handler (H) is responsible for handling the data of one or more Applications (A). To do so, it connects to a Broker where it registers applications and devices. The Handler is also the point where data is encrypted or decrypted.

The goal of The Things Network is to be very flexible in terms of deployment options. The preferred option is to connect to the public community network hosted by The Things Network Foundation. In this case the Application connects to The Things Network Foundation's Handler as shown in Figure.

[[/uploads/Public-Network.png]]

<center>_Public Community Network_ </center>

It should also be possible to deploy private networks, by running all previously components in a private environment. This way, all data will remain within the private environment as shown in Figure.

[[/uploads/Private-Network.png]]

<center>_Private Network_ </center>

Hybrid deployments will be possible in the future. The most simple option for this, is for someone to run his own Handler, allowing him to handle the encryption and decryption of messages. A more complicated option is a private network that exchanges data with the public network. For this to work, private Routers will have to connect to public Brokers and vice versa. In this case the private network can offload public traffic to the community network and use the public community network as back-up.

[[/uploads/Private-Handler.png]]

<center>_Private Handler_ </center>

[[/uploads/Private-Exchange.png]]

<center>_Private Network with Community Exchange_ </center>

## Core Functionality

Although the goal of The Things Network is to support for any protocol that can be useful for the community, the focus is currently on LoRaWAN. LoRaWAN is a ''network-intensive'' protocol, _intensive_ in the sense that due to the simple and minimalistic approach for devices, the backend systems are responsible for most of the logic. LoRaWAN was designed for the centralized architecture of telecom operators. In order to run on a distributed infrastructure like The Things Network some steps had to be added. In The Things Network's backend we now distinguish a number of different core functions.

Firstly, there are some Gateway-related functions such as scheduling and managing the utilization of the gateways. Scheduling is needed because a gateway can only do one transmission at the same time. The utilization information is used to evenly distribute load over different gateways and to be compliant with the [European duty cycles](https://www.thethingsnetwork.org/wiki/LoRaWAN/Overview#lorawan_frequency-bands_european-863-870-mhz-and-433-mhz-bands) . Another important feature is monitoring the status of each gateway.

Secondly, we need device-related functions that manage the state of devices in the network. As explained in [Addressing](https://www.thethingsnetwork.org/wiki/LoRaWAN/Overview#lorawan_addressing), device address are non-unique, so the network has to keep track of which addresses are used by which devices in order to map a message to the correct device and application. Other things the network must keep track of are the security keys and [frame counters](https://www.thethingsnetwork.org/wiki/LoRaWAN/Overview#lorawan_security_frame-counters) . In the future we will also need to keep track of the network utilization of each node.

Thirdly there is some functionality related to applications. For example, the Brokers and Handlers need to know to which server traffic for a specific application needs to be forwarded. This functionality also includes translation to other higher-layer protocols.

Finally, and most importantly, as The Things Network will be a distributed network, there has to be functionality that supports this distribution. Service discovery functionality helps components to determine where traffic should be routed to. Currently, this is implemented as a centralized Discovery server, giving The Things Network Foundation control over which components are allowed to announce specific services. For the future, the goal is to implement this discovery as a decentralized, peer-to-peer exchanged service.

## Separation of Concerns

In order to make the different backend components as decoupled as possible, we make a clear separation of the responsibilities of each component. An overview of these responsibilities will be published soon. The general idea is that the Router is responsible for all gateway-related functionality and region-specific details. A Broker handles a range of device addresses and is responsible for finding the right Handler to forward each message to. The Network Server is responsible for keeping the state of all individual devices. The Handler is responsible for encryption and decryption of messages and for forwarding messages to applications.

## Processing Flow of Uplink Messages

Based on this separation of concerns we implemented The Things Network's backend. As each component component has a number of high-level responsibilities, it has to execute a number of tasks when processing uplink and downlink messages. An overview of this flow is depicted in the Figure and is discussed in detail in the rest of this section.

<br>
![flow.png](https://s18.postimg.org/6eg028e2x/flow.png)

<center>_Processing Flow_ </center>

<br>

#### Gateway Protocol Translation (Router/Bridge)

When a gateway receives a message that was transmitted over LoRa, it is encapsulated and forwarded to The Things Network over the Internet (see figure below). Many gateways use the same reference gateway protocol, but alternative protocols have been developed for specific backends. The Things Network is also developing its own gateway protocol that is more suitable for The Things Network than the reference protocol in terms of security and access control.

<br>
![packet-forwarding.png](https://s16.postimg.org/whzdi2i4l/packet_forwarding.png)
<center>_Forwarding LoRaWAN Packets to the Backend_ </center>

<br>
Most gateway protocols have the same structure. When one or more messages are received, their binary data (usually base64-encoded) is forwarded to the backend, together with some metadata such as signal strength (RSSI) and signal-to-noise ratio (SNR). Periodically the gateway also sends some status information about the gateway itself, such as GPS coordinates and the number of packets received and transmitted.

We expect to connect gateways from different vendors, running different protocols. In order to keep the backend of The Things Network as generic as possible, we will implement a number of _bridges_ that translate from each vendor-specific gateway protocol to the common [communication protocol](https://www.thethingsnetwork.org/wiki/Backend/Overview#the-things-network-backend_communication-between-components) used in the backend of The Things Network.

#### Gateway Status and Metadata (Router)

The GPS coordinates of a gateway can especially be relevant to the application. Therefore the backend stores the latest status message sent by the gateway and injects the GPS information into the metadata of each uplink message. The implementation allows for two strategies: local storage and shared storage. When the local storage option is used, the backend stores the gateway statuses in-memory to keep the lookup times as short as possible. For a distributed system it is often desirable to keep as little state as possible. However, in this specific case it can be acceptable, because gateways already have a persistent connection to one backend server. When some load balancing features are added that might make this impossible, it is easy to switch to the shared storage (cache) implemented with [Redis](http://redis.io/) as both options have the same interface:

```c
type StatusStore interface {
    // Insert or Update the status
    Update(status *gateway.Status) error
    // Get the last status
    Get() (*gateway.Status, error)
}
```

#### Downlink Configuration (Router)

In LoRaWAN the downlink response to an uplink message is highly dependent on the geographic region of the gateway.  [Frequency bands](https://www.thethingsnetwork.org/wiki/LoRaWAN/Overview#lorawan_frequency-bands) section describes the different channel plans and requirements. As the Router is responsible for all gateway-related and region-specific details, the Router has to determine how a downlink response can be sent to a device. As described in [LoRaWAN devices](https://www.thethingsnetwork.org/wiki/LoRaWAN/Overview#lorawan_devices) section, each has two receive windows, one at exactly 1 second after the uplink, the other after 2 seconds. Therefore, for each gateway that received the uplink message, a Router builds two downlink configurations.

In order to select the best option later, the Router additionally has to calculate a _score_ for each option. This score is influenced by a number of factors. At the moment we consider airtime, signal strength, gateway utilization and already scheduled transmissions. The latter is quite obvious, as a gateway can not do two transmissions at the same time. Scheduling a downlink message on a gateway that had a better signal strength (signal-to-noise ratio) also makes it more likely that a node will receive the downlink correctly.

The combination of the airtime of a message and the utilization of a gateway is used to optimize the network as a whole. As each transmission blocks the receivers of a gateway for some time, it is better to send messages in a shorter time. Therefore, downlink messages at a higher data rate are preferred over messages with a lower data rate. The utilization of a gateway indicates the percentage of time a gateway is receiving messages. Gateways with a higher utilization (because of being positioned at a good location) are therefore more valuable to the network than gateways with a lower utilization. Therefore the latter should be preferred for downlink messages, so that the former can continue receiving uplink messages.

| g | 863.0 - 868.0 Mhz | 1% duty-cycle | 25 mW (= 14dBm) |
|---|---|---|---|
| g1 | 868.0 - 868.6 Mhz | 1% duty-cycle | 25 mW  |
| g2 | 868.7 - 869.2 Mhz | 0.1% duty-cycle | 25 mW  |
| g3 | 869.4 - 869.65 Mhz | 10% duty-cycle | 500 mW (= 27dBm) |
| g4 | 869.7 - 870.0 Mhz | 1% duty-cycle | 25 mW  |

_Table: Sub-bands in the European 863.0 - 870.0 Mhz band_

<br>
The gateway utilization is additionally used to comply with the spectrum regulations of the [European frequencies](https://www.thethingsnetwork.org/wiki/LoRaWAN/Overview#lorawan_frequency-bands_european-863-870-mhz-and-433-mhz-bands). The current implementation of The Things Network operates in the 863.0 - 870.0 MHz band, which is divided into a number of sub-bands shown in the Table above. Each of these sub-bands has a different duty-cycle. For example on the 868.1 MHz frequency, the duty-cycle is 1%. This means that a gateway should not transmit for 99% of the time, so after a transmission with an airtime of 50 ms, the transmitter is not allowed to transmit on this frequency for 4950 ms. In order to enforce this behaviour in the backend, the backend will mark downlink options as impossible as long as the utilization of a gateway is higher than the duty-cycle.

#### Device Address Extraction (Router)

The first step in routing a packet is based on the device address. As mentioned in  [Addressing](https://www.thethingsnetwork.org/wiki/LoRaWAN/Overview#lorawan_addressing), `DevAddr` is a non-unique 32-bit address of which 25 bits can be assigned by the network operator. The fact that these addresses can be assigned by the operator (The Things Network) makes this a very suitable point for distributing the traffic.

The Things Network has chosen to distribute the traffic based on device address prefix. Each Broker announces a device address prefix with a discovery service. These prefixes are similar to how IP address ranges are announced in  [BGP](http://www.cisco.com/c/en/us/about/open-source/open-standards/border-gateway-protocol-bgp.html). Network components can periodically retrieve the list of Brokers and their announced prefixes. Messages are forwarded to all Brokers that announce a prefix matching the device address of the message.

#### De-Duplication (Broker)

LoRaWAN is a long-range radio protocol, making it likely that the message is received by more than one gateway. This means that the backend has to perform some kind of de-duplication in order to deliver a message only once to the application. This does not mean that the duplicates are unimportant. The metadata of these messages might be valuable as well. For example, when combining the locations of the gateways that received the message with the reception time and the signal strength, it might be possible to determine the location of the device that sent the message.

De-duplication is not easy to do in a distributed system, but can be very simple when there is only one server. Therefore, we decided to make a an abstraction on top of the de-duplication implementation. This makes it easy to replace the actual implementation with one that is most suitable for the use-case, while still keeping the same interface:

```c
type Deduplicator interface {
 // Deduplicate a value based on a given key
 Deduplicate(key string, value interface{}) []interface{}
}
```

The current implementation of the backend locally de-duplicates uplink messages based on the md5 sum of the payload. The payload of the message will be the same for all duplicates, and the chance that during the de-duplication period (usually a couple of seconds) a different message with the same hash arrives, is extremely low.

Gateways can be connected to any access network. Some are connected to wired ethernet, others use WiFi or even GPRS connections to the Internet. So even though radio waves travel at the speed of light, the network delay makes that duplicate messages do not arrive at the Broker at the same time. In order to collect the metadata added to the message by each gateway, the Broker has to buffer duplicates for some amount of time. That time should be long enough to gather as many duplicates as possible, but short enough to give the application enough time to reply to a message in the receive window that will be opened 1 second after the transmission.

Our measurements in the current deployment of The Things Network have shown that the average delay between the first and last duplicate is just under 100 ms, with the maximum delay at about 300 ms. We do realize that these measurements are done in a network where gateways are not yet densely deployed, so we need more data in order to draw conclusions. However, these values give a decent indication for the required de-duplication time, which is currently set at 200 ms. When more data is gathered, this time can be further optimized.

#### Device and Application lookup (Broker/Network Server)

Because device addresses are non-unique it is necessary to determine the exact device that sent the message, and the application it belongs to. To do this, the backend has to perform a series of cryptographic message integrity code (MIC) checks, one for each device that uses the same device address. To do so, the Broker requests a list of devices with the given device address from the Network Server and checks if the MIC can be validated using network session key. If no match is found, the message is dropped.

#### Frame Counter Check (Broker)

The [frame counter](https://www.thethingsnetwork.org/wiki/LoRaWAN/Overview#lorawan_security_frame-counters) we discussed is a security measure used to detect replay attacks. After validating the MIC, the Broker checks if the Frame counter is valid. As frame counters can only increase, a message with a frame counter that is lower than the last known frame counter should be dropped. Additionally, the Broker has to verify that the gap between the last known frame counter and the counter in the message is not too big. According to the LoRaWAN specification, the maximum gap is 16384.

LoRaWAN supports both 16-bit and 32-bit frame counters. However, only the 16 least significant bits of the counter are included in the message header. Therefore, the backend has to keep track of the full 32-bit frame counter and use this instead of the 16-bit counter that is included in the message.

#### Metadata Collection (Broker)

When all checks have succeeded, the Broker can continue processing the message. First, it merges the duplicates received from all different routers and gateways. In this step it is important to differentiate between metadata that is the same for every gateway that received the message and metadata that is specific to each reception. For example, the frequency, modulation and data rate will be the same for all gateways, so it only needs to be forwarded once. On the other hand, the signal strength, reception time and GPS coordinates of each gateway should all be included when forwarding the message. In this step also the different downlink configurations are collected in order to select the best option in the next step.

#### Selecting Best Downlink Option (Broker)

The Broker has to select the best option for a downlink response to a message. As the Broker does not have any information about the gateway that received a message, it is very difficult to do this. Therefore the Router already calculated a _score_ for each [downlink configuration](https://www.thethingsnetwork.org/wiki/Backend/Overview#the-things-network-backend_processing-flow-of-uplink-messages_downlink-configuration-router). If this _score_ calculation is done in a standard way, the Broker now only has to sort all possible downlink options and use the best option.

#### Device State and MAC Commands (Network Server)

Before forwarding the uplink message to the Handler, it is first sent to the Network Server so that the device's state can be updated. The Network Server also adds a _downlink template_ to the message. This template can be used by the Handler to send a downlink message back to the device. It contains all necessary values (such as the frame counter, message type and option flags) so that the Handler only has to add the application payload to the message. Additionally, this gives the Network Server a chance to add MAC commands to the message. For example, based on the number of gateways that received a message and their signal strength, the Network Server may add MAC commands that instruct the device to transmit at a higher data rate.

#### Message Decruption (Handler)

As messages are end-to-end encrypted, the backend is also responsible for decrypting messages. However, not in all cases the application owner might want The Things Network to be responsible for that. Therefore message decryption is placed in a separate component (the Handler), allowing an application owner to run this Handler in his own private environment as shown in [Figure](https://www.thethingsnetwork.org/wiki/Backend/Overview#the-things-network-backend_architecture).

After decrypting the message payload, the Handler can pass the message up to the application. However, for many applications, some simple decoding and conversion is required.

#### Payload Conversion (Handler)

After decryption, the Handler is able to decode and convert the payload into a format that is easily accessible by the application. The implementation of the default Handler therefore includes so-called _payload functions_. These functions are simple JavaScript functions that can be used to decode, convert and validate data. The _decoder_ is used to decode the binary payload into a format that is more appropriate. For example, the following function decodes a temperature value sent as two bytes to a JSON object:

```c
function (bytes) {
    var data = (bytes[0] << 8) | bytes[1];
    return { temperature: data / 100.0 };
}    
```

The optional _converter_ can convert values in the decoded JSON object. This could for example be a conversion from a voltage to an actual value, or from a temperature in Celsius to a temperature in Fahrenheit. The optional _validator_ can be used to check the validity of the data and drop outliers. In the future these payload functions can adapt their behaviour to the `FPort` field of the message, allowing the community to define standard data encoding formats for each `FPort`. An example could be a standard format for sending weather station data.

#### MQTT (Handler)

The default Handler implementation simply publishes a JSON representation of uplink messages to a topic `<app_eui>/devices/<dev_eui>/up` on an MQTT broker. This allows applications to simply subscribe to the same MQTT topic and process the data in any way. From the following message, the application could for example see that the temperature measured by device `001122334455667788` was 12.86 degrees:

```
Topic: 0102030405060708/devices/001122334455667788/up

{ payload: 'BQY=',
  fields:{temperature: 12.86 },
  port: 14,
  counter: 1234,
  metadata:
  [ { frequency: 868.1,
      datarate: 'SF7BW125',
      codingrate: '4/5',
      ...
      longitude: 6.55738,
      latitude: 53.18977 } ] }
```
The public community network will probably stick with this API and format, but this behaviour can be easily adapted to other use cases. The weather station example mentioned earlier could for example publish the actual measurements to MQTT topics like `<location>/temperature, <location>/wind_speed` etc. It would even be possible to completely remove MQTT and use other protocols. The implementation of The Things Network's backend components is open source, so it is very easy to build a custom Handler that has different behaviour than the default implementation.

#### Downlink (Handler)

After publishing the uplink message to MQTT, the Handler will determine whether it is necessary to reply to the device with a downlink message. There are three situations when a downlink message needs to be sent. The first and most obvious is when the application has a payload available to send to the device. In this case the payload is added to the response template that was generated by the Network Server. The second case is when the uplink message requires confirmation. Regardless of a downlink payload being available or not, an acknowledgement has to be sent. The third situation is when the Network Server needs to send MAC commands to the device. If this is detected, the handler could decide whether to obey the Network Server or not, although the current implementation always follows the Network Server's request.

Similar to with uplink messages, the Handler is responsible for encrypting the payload of the message.

If no downlink payload is available, the Handler may choose to wait for a short time to let the application prepare a downlink message based on the uplink message it just received. After this deadline is expired, the Handler must send the downlink message back to the Broker.

#### Device State (Network Server)

After the Broker receives a downlink message from a Handler, it sends the message to the Network Server, which will update the device's state (specifically, the frame counters) in the database and generate the MIC of the message. After this, the Broker forwards the downlink message to the router that is responsible for the gateway that has to transmit the the downlink message.

#### Downlink Scheduling (Router)

As mentioned in the beginning of this chapter, the Router is responsible for managing the schedule of the Gateway. As most gateways only have a buffer of 1 downlink message, the Router has to buffer scheduled messages until the last moment, and then send each message just in time to the gateway.

## Communication Between Components

In a distributed network like The Things Network it is important that communication between components is fast. We have looked at different options, each having their upsides and downsides.

##### Initial Solution

The first prototype of The Things Network's backend used a Message Queue to facilitate communication between the two components of the system, as shown in the figure. In this prototype the component that is responsible for translation of the gateway protocol (Croft), published messages to the queue, while the other component (Jolie) subscribed to those messages and stored them in a database. This method worked well in the prototype, but when the network will consist of hundreds of distributed components, setting up message queue servers between all components will have too much overhead.

<br>
![croft-jolie.png](https://s12.postimg.org/xbpfcrpd9/croft_jolie.png)

<center>_Message Queue between components_ </center>

<br>
Many cloud-based systems communicate using web APIs. It is very common for a service to communicate with the REST API or JSON endpoint of another service. The second implementation of The Things Network's backend used a similar strategy as shown in the figure. We soon realized that this would not scale, as each uplink message would lead to a large number of HTTP connections being opened for a really short time.

<br>
![webhooks.png](https://s21.postimg.org/z62s7d2hz/webhooks.png)

<center>_HTTP request/response communication between components_ </center>

<br>
Re-using TCP connections already significantly improves the performance, but it is still inefficient to scale this.

##### Current Solution: gRPC

We decided to look at how large services handle this problem and found out that Google developed [gRPC](http://www.grpc.io). It is a Remote Procedure Call system that allows one to define _service interfaces_ that describe functions that can be executed. These functions can be implemented by a server, allowing a client to execute these functions on the server as if they were local functions.

gRPC is language-neutral, which means that it does not make any assumptions on what programming language or environment is used. This allows developers to implement both servers and clients in any language. gRPC leverages HTTP/2 to multiplex requests over a single TCP connection and save bandwidth with header compression.

gRPC uses Protocol Buffers for data serialization. Protocol Buffers are high-level specifications of data structures that can be marshaled to very compact binary wire formats. For example, the following code defines an uplink message in The Things Network:

```
message UplinkMessage {
  bytes                 payload             = 1;
  protocol.RxMetadata   protocol_metadata   = 11;
  gateway.RxMetadata    gateway_metadata    = 12;
 }
```

After specifying the format of a message, the specification can be _compiled_ to a programming language. The compiled code provides functions for marshaling the data structure to a binary format and unmarshaling from that binary format back to the data structure. Protocol buffers use default values for each field, so that it is possible to add fields to the specification while remaining backwards compatible.

gRPC supports bidirectional streaming, which is very suitable for the particular flow of traffic in The Things Network's backend. These streams have a very low overhead and are almost as simple as a `for` loop. The Things Network's backend extensively uses gRPC streams, for example in the Broker (simplified):

```
service Broker {
  // Router initiates an Association with the Broker.
  rpc Associate(stream UplinkMessage) returns (stream DownlinkMessage);
 }
```

This service definition allows a Router to simply connect to the Broker, call the Associate method and stream uplink messages to the Broker:

```
// Connection Setup
conn, _ :=grpc.Dial("ttn-broker.com:1234")
client := broker.NewBrokerClient(conn)
association, _ := client.Associate()

// Sending uplink messages to the stream
for uplink := range <- uplinks {
  association.Send(&broker.UplinkMessage{
  ...
 })
}
```

## Component Scalability

The described architecture is one step in the process of making The Things Network's backend fully distributed and scalable. We will now look at the scalability of individual components. The goal of this is to scale a component horizontally by increasing the number of instances, while to other components it still looks as one component.

When more gateways connect to the network, it is easy to scale horizontally by increasing the number of Routers. A load balancer can then make sure that each Router handles an (approximately) equal amount of connections from gateways. The gateway will keep a persistent connection to one specific Router behind the load balancer while new gateways will connect to a different Router. This concept is shown in the figure.

<br>
![router-scale.png](https://s18.postimg.org/tna994dm1/router_scale.png)

<center>_Router instances behind a load balancer, separate Brokers_ </center>

<br>
When the network traffic (uplink messages) increases, the Brokers will be the next bottleneck as they have to handle de-duplication and do a series of MIC checks. The de-duplication step makes it difficult to scale Brokers horizontally. The first step in scaling Brokers is making each broker responsible for a smaller subset of devices by adapting the prefix it announces as shown in the figure. However, this means that a Handler needs to connect to multiple Brokers in order to receive messages that are sent by different devices that happen to have a different prefix in their address. Further research is necessary to find an efficient solution for this.

The Network Server might remain a bottleneck, as it is responsible for keeping the state of devices. However most transactions are reads (a Broker requesting a list of devices for an address) or atomic writes (incrementing frame counters), it will be relatively easy to scale the Network Server as well.
