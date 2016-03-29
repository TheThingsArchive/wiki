# TTNCTL Configuration

The `ttnctl` program can be configured with command line options described in the [[ttnctl documentation|Backend/ttnctl/ttnctl]], with environment variables or with a configuration file.

## Environment

`ttnctl` can be configured using environment variables.
The format of these variables is the underscored version of the command line options that can be found in the [[ttnctl documentation|Backend/ttnctl/ttnctl]], prefixed with `TTNCTL_`. Environment variables are uppercase. The environment variables shown below are the default values, if these work for you, you don't have to set them.

```sh
export TTNCTL_DEBUG=true
export TTNCTL_TTN_ROUTER=localhost:1700
export TTNCTL_APP_EUI=0102030405060708
export TTNCTL_MQTT_BROKER=localhost:1883
export TTNCTL_TTN_HANDLER=0.0.0.0:1782
export TTNCTL_TTN_ACCOUNT_SERVER=https://account.thethingsnetwork.org
```

###### Updated on 24-Mar-2016

## Configuration File

A configuration file can be specified using the `--config` option. By default, `ttnctl` looks for the file `~/.ttnctl.yaml` (in your home directory).
The configuration file shown below contains the default values, if these work for you, you don't have to set them.

```yaml
debug: true
ttn-router: "localhost:1700"
app-eui: "0102030405060708"
mqtt-broker: "localhost:1883"
ttn-handler: "0.0.0.0:1782"
ttn-account-server: "https://account.thethingsnetwork.org"

```

###### Updated on 24-Mar-2016
