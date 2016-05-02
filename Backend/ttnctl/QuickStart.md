# Quick Start Guide for `ttnctl`

This document gives a quick introduction to `ttnctl`.

## Installation

You can download the latest version of `ttnctl` using the following links:

* [Mac (64 bit, zip)](https://ttnreleases.blob.core.windows.net/release/src/github.com/TheThingsNetwork/ttn/release/branch/develop/ttnctl-darwin-amd64.zip)
* [Linux (64 bit, zip)](https://ttnreleases.blob.core.windows.net/release/src/github.com/TheThingsNetwork/ttn/release/branch/develop/ttnctl-linux-amd64.zip)
* [Linux (32 bit, zip)](https://ttnreleases.blob.core.windows.net/release/src/github.com/TheThingsNetwork/ttn/release/branch/develop/ttnctl-linux-386.zip)
* [Windows (64 bit, zip)](https://ttnreleases.blob.core.windows.net/release/src/github.com/TheThingsNetwork/ttn/release/branch/develop/ttnctl-windows-amd64.zip)
* [Windows (32 bit, zip)](https://ttnreleases.blob.core.windows.net/release/src/github.com/TheThingsNetwork/ttn/release/branch/develop/ttnctl-windows-386.zip)

> _Downloads are also available as `.tar.gz` and `.tar.xz` archives._

We often update `ttnctl`, so re-download once in a while.

## Configuration

If you don't like the default configuration options of `ttnctl`, you can configure them with environment variables or with a configuration file. This is explained [[here|Backend/ttnctl/Configuration]]. In most cases the default options should be fine. Later in this guide we'll add our application ID to the config file to save us some typing.

## Register and Login

To register as a new user, use the command [`ttnctl user create`](ttnctl_user_create):

```
❯ ttnctl user create demo@thethingsnetwork.org
Password: <enter password>
  INFO User created
```

You can now login with [`ttnctl user login`](ttnctl_user_login):

```
❯ ttnctl user login demo@thethingsnetwork.org
Password: <enter password>
  INFO Logged in as demo@thethingsnetwork.org and persisted token in $HOME/.ttnctl/auths.json
```

## Application Management

Now create a new application with [`ttnctl applications create`](ttnctl_applications_create). In this example we will create an application named `Hello World App`. The

```
❯ ttnctl applications create 'Hello World App'
  INFO Application created successfully
```

Show your applications along with their assigned EUIs with `ttnctl applications` command.

If you work in a group, you can authorize your colleagues to manage the application with the [`ttnctl applications authorize`](ttnctl_applications_authorize) command:

```
❯ ttnctl applications authorize 0807060504030201 john@doe.org
  INFO User authorized successfully
```

## Setting the new application as "active"

We are going to use the application we just created for the rest of this guide. We could use the `--app-eui 0807060504030201` flag in every command, but it's a lot easier to just configure it once and have `ttnctl` use this application from now on, so let's call [`ttnctl applications use`](ttnctl_applications_use):

```
❯ ttnctl applications use 0807060504030201
  INFO You are now using application 0807060504030201.
```

## Device Management

To register an Over-the-Air activated device (OTAA) you can use the [`ttnctl devices register`](ttnctl_devices_register) command. In the following example we register a device with DevEUI `0102030405060708` and AppKey `01020304050607080102030405060708`:

```
❯ ttnctl devices register 0102030405060708 01020304050607080102030405060708
  INFO Ok
```

For registering a personalized device (ABP) we have the command [`ttnctl devices register personalized`](ttnctl_devices_register_personalized). In the following example we register a personalized device with DevAddr `01020304`, NwkSKey `01020304050607080102030405060708` and AppSKey `08070605040302010807060504030201`:

```
❯ ttnctl devices register personalized 01020304 01020304050607080102030405060708 08070605040302010807060504030201
  INFO Ok
```

You can see a list of registered devices with [`ttnctl devices`](ttnctl_devices):

```
❯ ttnctl devices
  INFO Found 1 personalized devices (ABP)

DevAddr 	FCntUp	FCntDown
01020304	0     	0

  INFO Found 1 dynamic devices (OTAA)

DevEUI          	DevAddr	FCntUp	FCntDown
0102030405060708	       	0     	0

  INFO Run 'ttnctl devices info [DevAddr|DevEUI]' for more information about a specific device
```

You can also get information about a specific device with [`ttnctl devices info`](ttnctl_devices_info), for example for the personalized device:

```
❯ ttnctl devices info 01020304
Personalized device:

  DevAddr: 01020304
           {0x01, 0x02, 0x03, 0x04}

  NwkSKey: 01020304050607080102030405060708
           {0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}

  AppSKey: 08070605040302010807060504030201
           {0x08, 0x07, 0x06, 0x05, 0x04, 0x03, 0x02, 0x01, 0x08, 0x07, 0x06, 0x05, 0x04, 0x03, 0x02, 0x01}

  FCntUp:  0
  FCntDn:  0
```

## Receiving Uplink Messages

The [`ttnctl subscribe`](ttnctl_subscribe) command subscribes to all uplink messages that are sent to your application:

```
❯ ttnctl subscribe
  INFO Subscribing to uplink messages from all devices in application 0807060504030201
  INFO Subscribed. Waiting for messages...
  INFO Hello world!                             DevEUI=0000000001020304
```

You can also supply a DevEUI if you only want to see messages from a specific device. For ABP devices, the DevEUI is the DevAddr prepended with 8 zero's. In the following example we subscribe to an ABP device with DevAddr `01020304`:

```
❯ ttnctl subscribe 0000000001020304
  INFO Subscribing uplink messages from device
  INFO Subscribed. Waiting for messages...
  INFO Hello world!                             DevEUI=0000000001020304
```

## Scheduling Downlink Messages

Scheduling downlink messages is done with the [`ttnctl downlink`](ttnctl_downlink) command. In the following example we plan a downlink message for an ABP device with DevAddr `01020304` with a TTL of 2 hours:

```
❯ ttnctl downlink 0000000001020304 'Hello back!' 2h
```

> If the node doesn't send an uplink message within those 2 hours, the scheduled downlink message is dropped.
