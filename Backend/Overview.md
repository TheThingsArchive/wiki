# The Things Network Backend

In the backend of The Things Network, we consider 5 internet-connected components:
**Gateways**, **[[Routers|Backend/Router]]**, **[[Brokers|Backend/Broker]]**, **[[Handlers|Backend/Handler]]** and **Applications**.

In private networking setups, the components are statically connected:

[[/uploads/TTN-Simple-Deployment.png]]

In public networking setups, components are dynamically connected:

[[/uploads/TTN-Distributed-Deployment.png]]

The connections between the components can be local connections on the same machine, connections within a private network or connections over the internet. Therefore, The Things Network back-end components can deployed for different scenarios:

1. **Public community network**: Gateways are provided by the community, Routers, Brokers and Handlers (with limited capability) are operated by The Things Network Foundation and other community operators. Users may use their own Handler instance to keep control over encryption/decryption of messages;
2. **Private connected network**: Gateways, Router, Broker and Handler are operated by the user. Packets that cannot be handled by the private network are sent to the community network;
3. **Private offline network**: Gateways, Router, Broker and Handler are operated by the user and is not connected to the public community network. This is useful in firewalled scenarios or when using The Things Network back-end in an offline setup, for example on a Raspberry Pi.

# Connect to the Network

In order to interact with The Things Network, you need an user account. Learn more about [[Security|Backend/Security]].

Read more about the key scenarios of connecting to The Things Network:

1. [[Connect a gateway|Backend/Connect/Gateway]]
2. [[Connect an application|Backend/Connect/Application]]
3. [[Connect as operator|Backend/Connect/Operator]]

# Network Operators

In terms of operations, we consider a number of operators that own and maintain these components.

* **Gateway Operators** have one or more gateways that are registered with one or more **Router Operators**.
* **Router Operators** run the [ttn router](ttn/ttn_router) component and are responsible for gateway registrations. They could also inform **Gateway Operators** when a problem with their gateway is detected.
* **Broker Operators** run the [ttn broker](ttn/ttn_broker) component and are responsible for application registrations and device address registrations within their assigned address space.
* **Handler Operators** run the [ttn handler](ttn/ttn_handler) component and are responsible for encryption/decryption and for communication with applications. Handlers can be public or private, or can even be integrated in an application.

# Backend Components

[[/uploads/TTN-Backend-Components.png]]
