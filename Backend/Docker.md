# Running the Backend in Docker

The easiest way to run The Things Network's backend is using [Docker](https://www.docker.com/). We usually use a [docker-compose](https://docs.docker.com/compose/) file similar to the one displayed below.

Environment variables should be configured as explained in the page about [[ttn configuration|Backend/ttn/Configuration]].

```yaml
version: '2'
services:
  mosquitto:
    image: ansi/mosquitto
    ports:
      - "1883:1883" # Public access to MQTT
  router:
    image: thethingsnetwork/ttn
    command: router
    links:
      - broker
    volumes:
      - /home/dockeruser/ttn-data:/data
    environment:
      - TTN_DEBUG=true # Start the server in debug mode (extra logging)
      - TTN_ROUTER_DB_BROKERS=boltdb:/data/ttn_router_brokers.db
      - TTN_ROUTER_DB_GATEWAYS=boltdb:/data/ttn_router_gateways.db
      - TTN_ROUTER_DB_DUTY=boltdb:/data/ttn_router_duty.db
      - TTN_ROUTER_BROKERS=broker:1881 # The broker container
    ports:
      - "1700:1700/udp" # Public access to the Semtech port
      # - "1780:1780"   # No public access to the downlink RPC port
      - "10700:10700"   # Public access to the status port
  broker:
    image: thethingsnetwork/ttn
    command: broker
    volumes:
      - /home/dockeruser/ttn-data:/data
    environment:
      - TTN_DEBUG=true # Start the server in debug mode (extra logging)
      - TTN_BROKER_DB_APPS=boltdb:/data/ttn_broker_apps.db
      - TTN_BROKER_DB_DEVICES=boltdb:/data/ttn_broker_devices.db
      - TTN_BROKER_ACCOUNT_SERVER=https://account.thethingsnetwork.org
    ports:
      # - "1781:1781" # No public access to the downlink RPC port
      # - "1881:1881" # No public access to the uplink RPC port
      - "10701:10701" # Public access to the status port
  handler:
    image: thethingsnetwork/ttn
    command: handler
    links:
      - broker
      - mosquitto
    volumes:
      - /home/dockeruser/ttn-data:/data
    environment:
      - TTN_DEBUG=true # Start the server in debug mode (extra logging)
      - TTN_HANDLER_DB_DEVICES=boltdb:/data/ttn_handler_devices.db
      - TTN_HANDLER_DB_PACKETS=boltdb:/data/ttn_handler_packets.db
      - TTN_HANDLER_TTN_BROKER=broker:1781            # The broker container
      - TTN_HANDLER_MQTT_BROKER=mosquitto:1883        # The mosquitto container
      - TTN_HANDLER_MQTT_USERNAME=yourusername        # You don't need this if you use 
      - TTN_HANDLER_MQTT_PASSWORD=yourpassword        #   the default mosquitto server
      - TTN_HANDLER_INTERNAL_ADDRESS_ANNOUNCE=handler # Hostname where brokers can reach the handler
    ports:
      # - "1882:1882" # No public access to the private RPC port
      - "1782:1782"   # Public access to the public RPC port
      - "10702:10702" # Public access to the status port
```
