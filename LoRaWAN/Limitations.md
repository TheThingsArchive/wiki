# LoRaWAN Limitations

LoRaWAN is not suitable for every use-case, so it is important that you understand the limitations. Here's a quick overview:

The LoRaWAN **yes** list:

* **Long range** - multiple kilometers
* **Low power** - can last months (hopefully years) on a battery
* **Low cost** - less than 20€ CAPEX per node, almost no OPEX
* **Low bandwidth** - something like 400 bytes per hour
* **Coverage everywhere** - just install your own gateways
* **Secure** - 128bit end-to-end encrypted

The LoRaWAN **no** list:

* **Live sensordata** - you can only send small packets every ±5 minutes
* **Phone calls** - you can do that with GPRS/3G/LTE
* **Controlling lights in your house** - look into ZigBee or BlueTooth
* **Sending photos, watching Netflix** - you might want to check out this technology called WiFi
* **Geolocation / Triangulation** - use GPS or wait a couple of years until LoRaWAN can do it

## Sending data from a Node to your Application (uplink)

* **Payload** should be as small as possible. This means that you **should not send JSON or plain (ASCII) text**, but instead encode your data as binary data.
* **Interval** between messages should be in the range of **several minutes**, so be smart with your data. You could for example transmit a `min|avg|max` every 5 minutes, or you could only transmit when you sensor value changed more than a certain threshold, etc.
* **Data Rate** should be as high as possible to minimize your airtime. `SF7BW125` is usually a good option. If you enable adaptive data rate (ADR), the network will be able to optimize your data rate.
* **Confirmed Uplink** eats a lot of your downlink airtime, so should only be used when you really need it.

## Sending responses from your Application to your Node (downlink)

We want to be able to handle as many Nodes as possible per Gateway. But as full-duplex radios are not widely available yet, a Gateway is not able to receive transmissions from Nodes while it is transmitting. Additionally, a Gateway can only transmit at one channel at a time (while it can receive on 8). As a result, the capacity for downlink messages is significantly lower than for uplink messages.

Until gateways are more widely deployed, we therefore encourage you to consider your fellow community members and try finding a solution that doesn't require downlink. If you do use downlink, be as efficient as possible.