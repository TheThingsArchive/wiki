# The Things Network Backend

In the backend of The Things Network, we consider 5 internet-connected components:
**Gateways**, **[[Routers|Backend/Router]]**, **[[Brokers|Backend/Broker]]**, **[[Handlers|Backend/Handler]]** and **Applications**.

[[/uploads/TTN-Private-Deployment.png]]

This is the most simple setup, mainly suitable for private deployments. The Things Network looks more like this:

[[/uploads/TTN-Public-Deployment.png]]

# Network Operators

In terms of operations, we consider a number of operators that own and maintain these components.

* **Gateway Operators** have one or more gateways that are registered with one or more **Router Operators**.
* **Router Operators** run the [ttn router](ttn/ttn_router) component and are responsible for gateway registrations. They could also inform **Gateway Operators** when a problem with their gateway is detected.
* **Broker Operators** run the [ttn broker](ttn/ttn_broker) component and are responsible for application registrations and device address registrations within their assigned address space.
* **Handler Operators** run the [ttn handler](ttn/ttn_handler) component and are responsible for encryption/decryption and for communication with applications. Handlers can be public or private, or can even be integrated in an application.
