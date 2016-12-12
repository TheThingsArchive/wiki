# Frequently Asked Questions

## Can I run The Things Network on my own server?

Yes, if you know what you're doing, you can run a private deployment of our open source routing services. Check out [our Github](https://github.com/TheThingsNetwork/ttn) to read how. We recommend most people to just use the public community network, as this will save you a lot of time and allows you to build awesome applications instead.

## What is the difference between a "single-channel gateway" and a "real gateway"?

A real gateway is able to listen on 8 channels and all spreading factors at the same time. Single-channel gateways are fixed to one channel and spreading factor, so they will only receive about 2% of the messages unless you specifically configure your nodes to send at the exact same configuration as your single-channel gateway.

As LoRaWAN is a spread-spectrum radio protocol, single-channel gateways are not LoRaWAN-compatible. They can only be used for plain LoRa (without WAN) communication or for testing, and are not supported by The Things Network.

## My node says "successful transmission", but I'm not receiving anything in my application. What's wrong?

A "successful transmission" just means that the node transmitted the message. It does not mean that it was also received by a gateway. If you don't see the message in `ttnctl`, in the Console or in your application, your node might just be out of range of a gateway.

## My node says "join not accepted: denied". Why's that?

In LoRaWAN we only send "join accepted" messages. If a join is not accepted, no message will be sent back to the node. A "denied" can therefore be caused by a number of other reasons than the server denying the join. First make sure that you're in range of a gateway. 9 out of 10 times this is the reason for joins not being accepted. If you're sure that you're within gateway range, you should check if the settings (AppEUI, DevEUI and AppKey) match the ones in the backend and check if they are entered in the right format (msb, lsb, hex).

## I don't receive data from my ABP device anymore. What's wrong?

When messages from your ABP device stop arriving in the backend, your device probably restarted. When this happens, it resets the frame counters to 0 and starts sending again. The backend however will flag these messages as a possible replay attack, and will drop them. This behaviour can be disabled by disabling the frame counter check in your device's settings. This will make your device vulnerable to replay attacks, which is why we recommend to avoid ABP and prefer OTAA.

## What about all those different versions of the backend?

- `v0`, `croft`, `api/v0` and `api/v0.1` have been discontinued. 
- `v1`, `staging`, `v1-staging` is the most widely used version, although this version is not considered production-ready.
- `v2` will be released before the end of 2016, a preview is available as `v2-preview`