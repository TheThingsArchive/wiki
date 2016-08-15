# Connect an Application

[[/uploads/TTN-Simple-Deployment.png]]

## Streaming Data

Use any MQTT client to connect to the Handler. For example, [Paho](http://www.eclipse.org/paho/) is an open source, feature-complete client for many programming languages. You can also use the builtin MQTT input node in [Node-RED](http://nodered.org).

### MQTT Broker

The MQTT broker can be reached at `staging.thethingsnetwork.org` (default port 1883).
TLS is supported (default port 8883), download the [staging mqtt-ca](http://staging.thethingsnetwork.org/mqtt-ca.pem) to connect.

### Authentication

Use the AppEUI as username and Access Key as password. [[More about Security|Backend/Security]]

You can find the AppEUI and Access Key with the `ttnctl applications` command, or in [The Things Network Dashboard](https://staging.thethingsnetwork.org) (in your application, click "learn how to get data").

In this example, the MQTT username is `0807060504030201` and the password is `I0f+e1W+CWgIiuIC4SjR5cpLxFZQfK2agDEpuCBpttI=`.

### Topics

The Handler publishes uplink messages from nodes to applications as well as activations of nodes, and subscribes to downlink messages from applications to nodes.

MQTT supports wildcards in topic paths: `+` for a single level and `#` for the remaining levels of the hierarchy. [More about MQTT topics](http://mosquitto.org/man/mqtt-7.html)

#### Uplink

Path: `<AppEUI>/devices/<DevEUI>/up` (use `+/devices/+/up` to get data from all devices)

Format:
```
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

The application payload `payload` is base64 encoded.

With the Mosquitto MQTT client you would connect and subscribe with the following command:

```
mosquitto_sub -h staging.thethingsnetwork.org -t '0807060504030201/devices/+/up' -u 0807060504030201 -P 'I0f+e1W+CWgIiuIC4SjR5cpLxFZQfK2agDEpuCBpttI='
```
To enable TLS you would use the following command:
```
mosquitto_sub --cafile <path>/mqtt-ca.pem -p 8883 -h staging.thethingsnetwork.org -t '0807060504030201/devices/+/up' -u 0807060504030201 -P 'I0f+e1W+CWgIiuIC4SjR5cpLxFZQfK2agDEpuCBpttI='
```

#### Downlink

Path: `<AppEUI>/devices/<DevEUI>/down`

Example message "Hello world":
```
{ "payload": "SGVsbG8gd29ybGQK=",
  "port": 1,
  "ttl": "1h" }
```

The application payload `payload` is base64 encoded.

The time-to-live (`ttl`) of the message specifies for how long the message should be queued for downlink before it expires. Depending on the class of the device (see [[LoRaWAN|LoRaWAN/Overview]]), the downlink message is a reply to the uplink message, it is sent on a schedule or it is sent immediately.

With the Mosquitto MQTT client you would connect and schedule a downlink with the following command:

```
mosquitto_pub -h staging.thethingsnetwork.org -t '0807060504030201/devices/0102030405060708/down' -u 0807060504030201 -P 'I0f+e1W+CWgIiuIC4SjR5cpLxFZQfK2agDEpuCBpttI=' -m '{ "payload":"SGVsbG8gd29ybGQK=","port":1,"ttl":"1h"}'
```
To enable TLS you would use the following command:
```
mosquitto_pub --cafile <path>/mqtt-ca.pem -p 8883 -h staging.thethingsnetwork.org -t '0807060504030201/devices/0102030405060708/down' -u 0807060504030201 -P 'I0f+e1W+CWgIiuIC4SjR5cpLxFZQfK2agDEpuCBpttI=' -m '{ "payload":"SGVsbG8gd29ybGQK=","port":1,"ttl":"1h"}'
```


**Note**: the `port` is not yet supported, and will always be `1`.

#### Activations

Path: `<AppEUI>/devices/<DevEUI>/activations` (use `+/devices/+/activations` to get activations from all devices)

The following topic will be returned on activation 
```
<EUI>/devices/0004A30B001B2F9C/activations
```
The payload contains metadata only
```
{
    "metadata" : [
    {
        "frequency" : 868.3,
        "datarate" : "SF7BW125",
        "codingrate" : "4/5",
        "gateway_timestamp" : 606962443,
        "gateway_time" : "2016-05-13T20:24:46.801449Z",
        "channel" : 1,
        "server_time" : "2016-05-13T20:24:46.836998337Z",
        "rssi" : -73,
        "lsnr" : 8.8,
        "rfchain" : 1,
        "crc" : 1,
        "modulation" : "LORA",
        "gateway_eui" : "1DEE192E3D82B8E4",
        "altitude" : 0,
        "longitude" : 4.79,
        "latitude" : 52.23985
    }
  ]
}

```

## Stored Data

The Things Network Handler does not store data. The [[Storage Handler|Backend/Storage Handler]] component is planned. This component stores data in a time-series database.
