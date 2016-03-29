# TTN Configuration

The `ttn` program can be configured with command line options described in the [[ttn documentation|Backend/ttn/ttn]], with environment variables or with a configuration file.

## Environment

The router, broker and handler can be configured using environment variables.
The format of these variables is: `TTN_<component>_<option>`, where `<option>` is the underscored version of the command line options that can be found in the [[ttn documentation|Backend/ttn/ttn]]. Environment variables are uppercase. The environment variables shown below are the default values, if these work for you, you don't have to set them.

```sh
export TTN_DEBUG=false
export TTN_ROUTER_DB_BROKERS=boltdb:/tmp/ttn_router_brokers.db
export TTN_ROUTER_DB_GATEWAYS=boltdb:/tmp/ttn_router_gateways.db
export TTN_ROUTER_DB_DUTY=boltdb:/tmp/ttn_router_duty.db
export TTN_ROUTER_STATUS_ADDRESS=0.0.0.0
export TTN_ROUTER_STATUS_PORT=10700
export TTN_ROUTER_UPLINK_ADDRESS=0.0.0.0
export TTN_ROUTER_UPLINK_PORT=1700
export TTN_ROUTER_DOWNLINK_ADDRESS=0.0.0.0
export TTN_ROUTER_DOWNLINK_PORT=1780
export TTN_ROUTER_BROKERS=localhost:1881
export TTN_BROKER_DB_APPS=boltdb:/tmp/ttn_broker_apps.db
export TTN_BROKER_DB_DEVICES=boltdb:/tmp/ttn_broker_devices.db
export TTN_BROKER_STATUS_ADDRESS=0.0.0.0
export TTN_BROKER_STATUS_PORT=10701
export TTN_BROKER_UPLINK_ADDRESS=0.0.0.0
export TTN_BROKER_UPLINK_PORT=1881
export TTN_BROKER_DOWNLINK_ADDRESS=0.0.0.0
export TTN_BROKER_DOWNLINK_PORT=1781
export TTN_BROKER_OAUTH2_KEYFILE=~/.ttn/oauth2-token.pub
export TTN_HANDLER_DB_DEVICES=boltdb:/tmp/ttn_handler_devices.db
export TTN_HANDLER_DB_PACKETS=boltdb:/tmp/ttn_handler_packets.db
export TTN_HANDLER_STATUS_ADDRESS=0.0.0.0
export TTN_HANDLER_STATUS_PORT=10702
export TTN_HANDLER_INTERNAL_ADDRESS=0.0.0.0
export TTN_INTERNAL_ADDRESS_ANNOUNCE=localhost
export TTN_HANDLER_INTERNAL_PORT=1882
export TTN_HANDLER_PUBLIC_ADDRESS=0.0.0.0
export TTN_HANDLER_PUBLIC_PORT=1782
export TTN_HANDLER_MQTT_BROKER=localhost:1883
export TTN_HANDLER_TTN_BROKER=localhost:1781
```

###### Updated on 24-Mar-2016

## Configuration File

A configuration file can be specified using the `--config` option. By default, `ttn` looks for the file `~/.ttn.yaml` (in your home directory).
The configuration file shown below contains the default values, if these work for you, you don't have to set them.

```yaml
debug: false
router:
  db-brokers: "boltdb:/tmp/ttn_router_brokers.db"
  db-gateways: "boltdb:/tmp/ttn_router_gateways.db"
  db-duty: "boltdb:/tmp/ttn_router_duty.db"
  status-address: "0.0.0.0"
  status-port: 10700
  uplink-address: "0.0.0.0"
  uplink-port: 1700
  downlink-address: "0.0.0.0"
  downlink-port: 1780
  brokers: "localhost:1881"
broker:
  db-apps: "boltdb:/tmp/ttn_broker_apps.db"
  db-devices: "boltdb:/tmp/ttn_broker_devices.db"
  status-address: "0.0.0.0"
  status-port: 10701
  uplink-address: "0.0.0.0"
  uplink-port: 1881
  downlink-address: "0.0.0.0"
  downlink-port: 1781
  oauth2-keyfile: "~/.ttn/oauth2-token.pub"
handler:
  db-devices: "boltdb:/tmp/ttn_handler_devices.db"
  db-packets: "boltdb:/tmp/ttn_handler_packets.db"
  status-address: "0.0.0.0"
  status-port: 10702
  internal-address: "0.0.0.0"
  internal-address-announce: localhost
  internal-port: 1882
  public-address: "0.0.0.0"
  public-port: 1782
  mqtt-broker: "localhost:1883"
  ttn-broker: "localhost:1781"
```

###### Updated on 24-Mar-2016
