# How to Migrate From Croft/Jolie

The demonstration setup at `croft.thethings.girovito.nl` has served us for a long time, but now it's time to say goodbye. We are doing our best to make sure the migration is smooth, so our new environment at `staging.thethingsnetwork.org` will temporarily forward a copy of all _uplink messages_ to the legacy `croft` environment.

To get started with our new **staging** environment, you'll have to make a couple of changes.

## Gateway Migration

Update your `local_conf.json` to look like the following:

```json
{
  "gateway_conf": {
      "gateway_ID": *YOUR_OWN_SERIAL*,
      "serv_port_up": 1700,
      "serv_port_down": 1700,
      "server_address": "staging.thethingsnetwork.org",
      "forward_crc_valid": true,
      "forward_crc_error": false,
      "forward_crc_disabled": true
  }
}
```

## Application Migration

If you have been using the MQTT API, the new setup will be quite similar. Uplink messages will be published to the `<app_eui>/devices/<dev_eui>/up` topic on our MQTT broker at `staging.thethingsnetwork.org:1883`.

> TODO: Explain or refer to message format.

You do have to register devices, which can be done using [[ttnctl|Backend/ttnctl/ttnctl]].

> TODO: Give example.

If you have been using the HTTP API (`thethingsnetwork.org/api/v0/` or `thethingsnetwork.org/api/v0.1`) you will have to migrate to a **storage handler**, as the **staging** environment does not store packets anymore. There is currently no storage handler available, but we'll let you know as soon as we have one.
