# Running the Backend in Docker

The easiest way to run The Things Network's backend is using [Docker](https://www.docker.com/). We usually use a [docker-compose](https://docs.docker.com/compose/) file similar to the one displayed below. Please see below for step by step instructions to install TTN in a particular operating system.

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

## Installing on Mint 17.3 (should also work for Mint 17.1, Mint 17.2, and Ubuntu 14.04)
Installation platform
- Mint 17.3 64 Bit (Mate) - based on Ubuntu 14.04 Trusty
- Docker 1.11.1, build 5604cbe
- Docker compose 1.7.1, build 0a9ab35
- Git 1.9.1
- TTN build 4bf8cbd67681859b1f5e117c0c9aeedc41edada1 on master  https://github.com/TheThingsNetwork/ttn/commits/master

### Docker installation
- Install Docker Engine

```
$sudo apt-get update
$sudo apt-get install apt-transport-https ca-certificates
$sudo apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys 58118E89F3A912897C070ADBF76221572C52609D
```

Open the /etc/apt/sources.list.d/docker.list file in your favorite editor. If the file doesn’t exist, create it.
Remove any existing entries and add
```
deb https://apt.dockerproject.org/repo ubuntu-trusty main
```
save and close
```
$sudo apt-get update
$sudo apt-get purge lxc-docker
```
 verify the repos are correct 
```
$apt-cache policy docker-engine
```

Install linux-image-extra package
```
$sudo apt-get install linux-image-extra-$(uname -r)
```
Install apparmor
```
$apt-get install apparmor
```
Finally install docker engine 
```
$sudo apt-get install docker-engine
```
start and test
```
$sudo service docker start
$sudo docker run hello-world
```

- Install Docker Compose
```
$sudo -i
$curl -L https://github.com/docker/compose/releases/download/1.7.1/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose
$chmod +x /usr/local/bin/docker-compose
$exit
```
Test the installation
```
$docker-compose --version
```

### Install TTN (router,broker,handler)

Create a folder e.g. TTN_Workspace
```
$mkdir TTN_Workspace & cd TTN_Workspace
```
Copy the configuration above (start of the page)
Create a file "docker-compose.yml" under the TTN_Workspace and paste the content in it

Compose the file using docker compose inside 'TTN_Workspace' write the following command
```
$sudo docker-compose up
```
Now you are running the local version of TTN. 

PS: The console will be inactive and will just printout logs. If you want to interact with TTN open another console and follow the next session.

### Install TTNCTL
-----------
TTNCTL is the controller that controls your local TTN. You will use TTN to communicate with TTN. First download and unzip the ttnctl from https://ttnreleases.blob.core.windows.net/release/src/github.com/TheThingsNetwork/ttn/release/branch/develop/ttnctl-linux-amd64.zip 

unzip the file under TTN_Workspace/ttnctl
rename extracted the executable to ttnctl

quickly create a user and application

Create an account: $ ttnctl user create [Your Email]
Sign in: $ ttnctl user login [Your Email]
Create an application: $ ttnctl applications create [Application Name]
List your applications: $ ttnctl applications
Choose an application to use from now on: $ ttnctl applications use [EUI shown in the previous output]
Create a new device: $ ttnctl devices register [DevEUI must be 16 digit e.g. 0102030405060708, also see 'ttnctl devices register personalized' for device with any size key]
List your devices: $ ttnctl devices
Get info about a specific device: $ ttnctl devices info [DevEUI]
Subscribe to incoming messages from this device: $ ttnctl subscribe [DevEUI] 
Send a message to TTN use http://staging.thethingsnetwork.org/wiki/Backend/ttnctl/ttnctl_uplink
Receive a message from a node use http://staging.thethingsnetwork.org/wiki/Backend/ttnctl/ttnctl_uplink

For more info see http://staging.thethingsnetwork.org/wiki/Backend/ttnctl/QuickStart
