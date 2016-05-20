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
  redis:
    image: redis
    command: redis-server --appendonly yes
    ports:
      #- "6379:6379" # No public access to Redis
    volumes:
      - /var/lib/docker/redis:/data
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
      - TTN_HANDLER_REDIS_ADDR=redis:6379             # The redis server
      - TTN_HANDLER_INTERNAL_ADDRESS_ANNOUNCE=handler # Hostname where brokers can reach the handler
    ports:
      # - "1882:1882" # No public access to the private RPC port
      - "1782:1782"   # Public access to the public RPC port
      - "10702:10702" # Public access to the status port
```

## Installing Docker on Ubuntu 14.04, Mint 17.x and similar

This guide was made with the following configuration:

- [Linux Mint](https://linuxmint.com/) version 17.3 64 Bit (Mate) - based on Ubuntu 14.04 Trusty
- [Docker](https://www.docker.com/) version 1.11.1, build 5604cbe
- [Docker compose](https://docs.docker.com/compose/) version 1.7.1, build 0a9ab35
- Git 1.9.1
- TTN [master branch](https://github.com/TheThingsNetwork/ttn/commits/master), build 4bf8cbd

### Docker installation

First you should install Docker Engine. See the [Docker documentation](https://docs.docker.com/engine/installation/linux/ubuntulinux/) for more details.

```
$ sudo apt-get update
$ sudo apt-get install apt-transport-https ca-certificates
$ sudo apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys 58118E89F3A912897C070ADBF76221572C52609D
```

Open `/etc/apt/sources.list.d/docker.list` in your favorite editor. If the file doesn’t exist, create it. Remove any existing entries and add:

```
deb https://apt.dockerproject.org/repo ubuntu-trusty main
```

Save and close the file.

```
$ sudo apt-get update
```

Next, we will remove old versions of Docker, if they happen to be installed.

```
$ sudo apt-get purge lxc-docker
```

You can now verify that the repos are correct:

```
$ apt-cache policy docker-engine
```

Install linux-image-extra package:

```
$ sudo apt-get install linux-image-extra-$(uname -r)
```

Install apparmor:

```
$ apt-get install apparmor
```

Finally install Docker Engine:

```
$ sudo apt-get install docker-engine
```

You can now start Docker and test with the `hello-world` image:

```
$ sudo service docker start
$ sudo docker run hello-world
```

### Docker Compose Installation

```
$ sudo -i
$ curl -L https://github.com/docker/compose/releases/download/1.7.1/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose
$ chmod +x /usr/local/bin/docker-compose
$ exit
```

Test the installation:

```
$ docker-compose --version
```

### The Things Network Backend Installation

Create a folder e.g. `TTN_Workspace`:

```
$ mkdir TTN_Workspace & cd TTN_Workspace
```

Copy the configuration [above](#running-the-backend-in-docker) and paste it into a file called `docker-compose.yml` in the TTN_Workspace folder.

With your console in the `TTN_Workspace` folder, start the containers:

```
$ sudo docker-compose up
```

Now you are running the local version of TTN.

The console will print out logs, to interact with tools like `ttnctl` you should open another terminal window.
