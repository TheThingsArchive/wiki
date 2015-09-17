The handler sits between the router and the application server.
The handler has five key functions:

1. Message integrity checking and decryption of payload
1. Deduplication
1. Transformation from binary data to key/value store
1. Storage with a REST API to access historical data
1. Stream data over websockets to applications and to MQTT brokers

> We bootstrap the router with DNS look-up to route data to the handler
> that belongs to the application. For example, the handler of the
> application with key 0xA0FF31 is IP 123.45.67.89 by resolving
> a0ff31.app.thethings.network. The router posts a message to the
> RabbitMQ of the handler on this URL, providing the application key.

## See also

* [Architecture Overview](Architecture)