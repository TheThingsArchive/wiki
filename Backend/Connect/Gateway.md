# Connecting a Gateway

[[/uploads/TTN-Simple-Deployment.png]]

## Server Addresses

Choose the Router instance depending on the region as specified in the LoRaWAN 1.0.1 specification (still under NDA, [download 1.0](https://www.lora-alliance.org/portals/0/specs/LoRaWAN%20Specification%201R0.pdf)). This may not be the closest router geographically. For example, Brazil supports 915-928 MHz, but these are only specified by the LoRa Alliance for Australia, so Brazilian gateways should connect to the Australian router.

```
router.eu.thethings.network # EU 433 and EU 863-870
router.us.thethings.network # US 902-928
router.cn.thethings.network # China 470-510 and 779-787
router.au.thethings.network # Australia 915-928 MHz
```

Today, these DNS records refer to the staging environment, but we will update the records to the production environment when it comes available this summer. This will be announced timely in advance. Note: the staging environment currently only supports EU 863-870.

The geographical location of a server and the supported frequency plans are two different things. As the LoRa Alliance publishes more region specifications, we will deploy Routers in datacenters near that region. Also, the yet to be implemented decentralized routing architecture enables community operators to provide routing infrastructure to The Things Network.

## Staging Environment

It is highly recommended to update your gateway to the server addresses shown above. If you have test gateways that you want to force to use the staging environment and keep using the staging environment, use one of these server addresses. The staging environment will contain newer, but possibly less stable, builds of The Things Network back-end.

```
router.eu.staging.thethings.network # EU 433 and EU 863-870
router.us.staging.thethings.network # US 902-928
router.cn.staging.thethings.network # China 470-510 and 779-787
router.au.staging.thethings.network # Australia 915-928 MHz
```

## Configuring Your Gateway

In the `local_conf.json` of the packet forwarder, update the fields `server_address` as follows:

```
"gateway_conf": {
        ...
        "server_address": "<insert server address here>",
        "serv_port_up": 1700,
        "serv_port_down": 1700,
        ...
}
```
