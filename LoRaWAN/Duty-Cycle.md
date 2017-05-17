# Duty Cycle

**Duty Cycle indicates the fraction of time a resource is busy.**

When a single device transmits on a channel for 2 _time units_ every 10 _time units_, this device has a duty cycle of 20%.

[[/uploads/DutyCycleSingleChannel.png]]

However, if we also consider _channels_, things get a bit more complicated. When we have a device that transmits on 3 channels instead of one, each individual _channel_ is still occupied for 2 _time units_ every 10 _time units_ (so 20%). However, the _device_ is now transmitting for 6 _time units_ every 10 _time units_, giving it a duty cycle of 60%.

[[/uploads/DutyCycleMultiChannel.png]]

In our European frequency plan, we have channels in different _sub-bands_, so when considering the duty cycle, we also have to consider these. Let's say the 3 channels we used before, are in 2 different _sub-bands_. Each separate _channel_ still has a duty cycle of 20%, the _device_ still has a duty cycle of 60%, but we now see that _Band 1_ is in use for 2 _time units_ every 10 _time units_ (20%), while _Band 2_ is in use for 4 _time units_ every 10 _time units_ (40%).

[[/uploads/DutyCycleMultiBand.png]]