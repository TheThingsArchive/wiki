# Connect an Application

[[/uploads/TTN-Simple-Deployment.png]]

## Streaming Data

Use any MQTT client to connect to the Handler. For example, [Paho](http://www.eclipse.org/paho/) is an open source, feature-complete client for many programming languages. You can also use the builtin MQTT input node in [Node-RED](http://nodered.org).

### Authentication

Use the AppEUI as username and Access Key as password. [[More about Security|Backend/Security]]

### Topics

The Handler publishes uplink messages from nodes to applications as well as activations of nodes, and subscribes to downlink messages from applications to nodes.

MQTT supports wildcards in topic paths: `+` for a single level and `#` for the remaining levels of the hierarchy. [More about MQTT topics](http://mosquitto.org/man/mqtt-7.html)

#### Uplink

Path: `<AppEUI>/devices/<DevEUI>/up`
Commonly used: `+/devices/+/up` to get data from all devices

Format (`payload` is base64 encoded):
```json
{ payload: 'DZYLXQCn',
  port: 1,
  counter: 3478,
  metadata:
   [ { frequency: 868.1,
       datarate: 'SF7BW125',
       codingrate: '4/5',
       gateway_timestamp: 2620279267,
       gateway_time: '2016-04-18T15:12:28.278421Z',
       channel: 0,
       server_time: '2016-04-18T15:10:50.4429229Z',
       rssi: -77,
       lsnr: 11.2,
       rfchain: 1,
       crc: 1,
       modulation: 'LORA',
       gateway_eui: '1DEE0DD778094AC5',
       altitude: 8,
       longitude: 6.55738,
       latitude: 53.18977 } ] }
```

#### Downlink

Path: `<AppEUI>/devices/<DevEUI>/down`

Format (payload is base64 encoded):
```json
{ payload: 'SGVsbG8gd29ybGQK',
  port: 1,
  ttl: '1h' }
```

#### Activations

Path: `<AppEUI>/devices/<DevEUI>/activations`
Commonly used: `+/devices/+/activations` to get activations from all devices

> TODO: Example

## Stored Data

The Things Network Handler does not store data. The [[Storage Handler|Backend/Storage Handler]] component is planned. This component stores data in a time-series database.
