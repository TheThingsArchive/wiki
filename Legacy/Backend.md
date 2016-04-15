# Legacy Backend (croft)

[[/uploads/Croft.png]]

## Gateway-side communication

Use the following `local_conf.json` for the packet forwarder:

```json
{
  "gateway_conf": {
      "gateway_ID": *YOUR_OWN_SERIAL*,
      "serv_port_up": 1700,
      "serv_port_down": 1700,
      "server_address": "croft.thethings.girovito.nl",
      "forward_crc_valid": true,
      "forward_crc_error": false,
      "forward_crc_disabled": true
  }
}
```

## Application-side communication

### REST API `http://thethingsnetwork.org/api/v0/`

* `GET /nodes/`
* `GET /nodes/{node_EUI}`
* `GET /gateways/`
* `GET /nodes/{gateway_EUI}`

Parameters:

* `limit` (default `100`, max `250`)
* `offset` (default `0`)

### REST API `http://thethingsnetwork.org/api/v0.1/`

* `GET /nodes/`
* `GET /nodes/{node_EUI}`
* `GET /gateways/`
* `GET /nodes/{gateway_EUI}`

### MQTT `tcp://croft.thethings.girovito.nl:1883`

Topic for uplink messages from nodes: `nodes/{devAddr}/packets`
Topic for status updates from gateways: `gateways/{eui}/status`
