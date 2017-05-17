# Duty Cycle

**Duty Cycle indicates the fraction of time a resource is busy.**

When a single device transmits on a channel for 2 _time units_ every 10 _time units_, this device has a duty cycle of 20%.

[[/uploads/DutyCycleSingleChannel.png]]

However, if we also consider _channels_, things get a bit more complicated. When we have a device that transmits on 3 channels instead of one, each individual _channel_ is still occupied for 2 _time units_ every 10 _time units_ (so 20%). However, the _device_ is now transmitting for 6 _time units_ every 10 _time units_, giving it a duty cycle of 60%.

[[/uploads/DutyCycleMultiChannel.png]]

In our European frequency plan, we have channels in different _sub-bands_, so when considering the duty cycle, we also have to consider these. Let's say the 3 channels we used before, are in 2 different _sub-bands_. Each separate _channel_ still has a duty cycle of 20%, the _device_ still has a duty cycle of 60%, but we now see that _Band 1_ is in use for 2 _time units_ every 10 _time units_ (20%), while _Band 2_ is in use for 4 _time units_ every 10 _time units_ (40%).

[[/uploads/DutyCycleMultiBand.png]]

## Maximum Duty Cycle

The duty cycle of radio devices is often regulated by government. If this is the case, the duty cycle is commonly set to 1%, but make sure to check the regulations of your local government to be sure.

In Europe, duty cycles are regulated by section 7.2.3 of the ETSI EN300.220 standard. This standard defines the following sub-bands and their duty cycles:

* **g** (863.0 – 868.0 MHz): 1%
* **g1** (868.0 – 868.6 MHz): 1%
* **g2** (868.7 – 869.2 MHz): 0.1%
* **g3** (869.4 – 869.65 MHz): 10%
* **g4** (869.7 – 870.0 MHz): 1%

Additionally, the LoRaWAN specification dictates duty cycles for the _join frequencies_, the frequencies devices of all LoRaWAN-compliant networks use for over-the-air activations (OTAA) of devices. In most regions this duty cycle is set to **1%**.

Finally, on The Things Network's public community network, we have a **Fair Access Policy** that limits the **uplink airtime** to **30 seconds per day (24 hours) per node** and the **downlink messages** to **10 messages per day (24 hours) per node**. If you use a private network, these limits does not apply, but you still have to be compliant with the governmental and LoRaWAN limits.

## Compliance

Every radio device must be compliant with the regulated duty cycle limits. This applies to **both nodes and gateways**.

In practice, this means that you should program your nodes in such that they stay within the limits. The easiest way to do this, is to calculate how much _airtime_ each message consumes using one of the many [airtime calculators](https://docs.google.com/spreadsheets/d/1QvcKsGeTTPpr9icj4XkKXq4r2zTc2j0gsHLrnplzM3I/edit) and use that information to choose a good transmit interval.

Some radio modules (such as the RN2483) also enforce the duty cycle limits. If you exceed the limits, the module will complain with a message `no_free_channel`. Specifically, the RN2483 limits the duty cycle on a per-channel basis. This means that if you only have 1 channel configured, the module will start enforcing the duty cycle after the first message.

[[/uploads/DutyCycleSingleChannelOffAir.png]]

In the European band, a transmission on a channel within a frequency band, also influences the other frequencies in that band. 

[[/uploads/DutyCycleMultiBandOffAir.png]]

As a per-channel duty cycle limit is easier to implement, you can also divide the sub-band duty cycle over the number of channels in that sub-band. So for example, in a sub-band with 8 channels and a duty cycle of 1%, each channel has a duty cycle of 1/8% (that's 0.125%). 

This method is also implemented by the RN2483 module, and as a result, instead of seeing the `no_free_channel` when you send too quickly after the first message you can send multiple messages before all 8 channels are "blocked" and the duty cycle is enforced.

[[/uploads/DutyCycleMultiChannelOffAir.png]]