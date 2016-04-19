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
* **Phone calls**
* **Controlling lights in your house**
* **Sending photos, watching Netflix**
* **Geolocation / Triangulation** - use GPS or wait a couple of years until LoRaWAN can do it

## Sending data from a Node to your Application (uplink)

* **Payload** should be as small as possible. A good goal is to keep it **under 12 bytes**. This means that you should not send your sensor values as JSON or plain text, but encode them as binary values.
* **Interval** between messages should be in the range of **several minutes**, so be smart with your data. You could transmit `min|avg|max` every 5 minutes, you could only transmit when you sensor value changed more than a certain threshold, etc.
* **Data Rate** should be as high as possible to minimize your airtime. `SF7BW125` is usually a good option. If you enable adaptive data rate (ADR), the network will be able to optimize your data rate.
* **Confirmed Uplink** eats a lot of your downlink airtime, so should only be used when you really need it.

## Sending responses from your Application to your Node (downlink)

The capacity for downlink messages is even lower than for uplink messages, so don't waste it.
