# LoRaWAN Frequencies

This is a list of frequency plan definitions, based on what is specified in the LoRaWAN regional parameter document. To know which frequency plan to use for TTN in your country, see the [frequency plans by country list](By-Country).

## EU863-870

Uplink:

1. **868.1** - SF7BW125 to SF12BW125
2. **868.3** - SF7BW125 to SF12BW125 and SF7BW250
3. **868.5** - SF7BW125 to SF12BW125
4. **867.1** - SF7BW125 to SF12BW125
5. **867.3** - SF7BW125 to SF12BW125
6. **867.5** - SF7BW125 to SF12BW125
7. **867.7** - SF7BW125 to SF12BW125
8. **867.9** - SF7BW125 to SF12BW125
9. **868.8** - FSK

Downlink:

- Uplink channels 1-9 (RX1)
- **869.525** - SF9BW125 (RX2 downlink only)

## EU433

No frequency plan yet. Submit a proposal!

## US902-928

Used in USA, Canada and South America

### Frequency Plan

Uplink:

1. **903.9** - SF7BW125 to SF10BW125
2. **904.1** - SF7BW125 to SF10BW125
3. **904.3** - SF7BW125 to SF10BW125
4. **904.5** - SF7BW125 to SF10BW125
5. **904.7** - SF7BW125 to SF10BW125
6. **904.9** - SF7BW125 to SF10BW125
7. **905.1** - SF7BW125 to SF10BW125
8. **905.3** - SF7BW125 to SF10BW125
9. **904.6** - SF8BW500

Downlink:

1. **923.3** - SF7BW500 to SF12BW500
2. **923.9** - SF7BW500 to SF12BW500
3. **924.5** - SF7BW500 to SF12BW500
4. **925.1** - SF7BW500 to SF12BW500
5. **925.7** - SF7BW500 to SF12BW500
6. **926.3** - SF7BW500 to SF12BW500
7. **926.9** - SF7BW500 to SF12BW500
8. **927.5** - SF7BW500 to SF12BW500

## CN470-510
No frequency plans yet. Submit a proposal!

## CN779-787
No frequency plans yet. Submit a proposal!

## AU915-928

Uplink:

1. **916.8** - SF7BW125 to SF10BW125
2. **917.0** - SF7BW125 to SF10BW125
3. **917.2** - SF7BW125 to SF10BW125
4. **917.4** - SF7BW125 to SF10BW125
5. **917.6** - SF7BW125 to SF10BW125
6. **917.8** - SF7BW125 to SF10BW125
7. **918.0** - SF7BW125 to SF10BW125
8. **918.2** - SF7BW125 to SF10BW125
9. **917.5** SF8BW500

Downlink:

1. **923.3** - SF7BW500 to SF12BW500
2. **923.9** - SF7BW500 to SF12BW500
3. **924.5** - SF7BW500 to SF12BW500
4. **925.1** - SF7BW500 to SF12BW500
5. **925.7** - SF7BW500 to SF12BW500
6. **926.3** - SF7BW500 to SF12BW500
7. **926.9** - SF7BW500 to SF12BW500
8. **927.5** - SF7BW500 to SF12BW500

## AS923

We use two frequency plans, depending on the country. OTAA devices use two common channels: 923.2MHz and 923.4MHz. They will receive the additional channels on a successful join.

### AS920-923

Used in Japan, Malaysia, Singapore

Uplink:

1. **923.2** - SF7BW125 to SF12BW125
2. **923.4** - SF7BW125 to SF12BW125
3. **922.2** - SF7BW125 to SF12BW125
4. **922.4** - SF7BW125 to SF12BW125
5. **922.6** - SF7BW125 to SF12BW125
6. **922.8** - SF7BW125 to SF12BW125
7. **923.0** - SF7BW125 to SF12BW125
8. **922.0** - SF7BW125 to SF12BW125
9. **922.1** - SF7BW250
10. **921.8** - FSK

Downlink:

- Uplink channels 1-10 (RX1)
- **923.2** - SF10BW125 (RX2)

### AS923-925

Used in Brunei, Cambodia, Hong Kong, Indonesia, Laos, Taiwan, Thailand, Vietnam

Uplink:

1. **923.2** - SF7BW125 to SF12BW125
2. **923.4** - SF7BW125 to SF12BW125
3. **923.6** - SF7BW125 to SF12BW125
4. **923.8** - SF7BW125 to SF12BW125
5. **924.0** - SF7BW125 to SF12BW125
6. **924.2** - SF7BW125 to SF12BW125
7. **924.4** - SF7BW125 to SF12BW125
8. **924.6** - SF7BW125 to SF12BW125
9. **924.5** - SF7BW250
10. **924.8** - FSK

Downlink:

- Uplink channels 1-10 (RX1)
- **923.2** - SF10BW125 (RX2)

## KR920-923

Uplink:

1. **922.1** - SF7BW125 to SF12BW125
2. **922.3** - SF7BW125 to SF12BW125
3. **922.5** - SF7BW125 to SF12BW125
4. **922.7** - SF7BW125 to SF12BW125
5. **922.9** - SF7BW125 to SF12BW125
6. **923.1** - SF7BW125 to SF12BW125
7. **923.3** - SF7BW125 to SF12BW125
8. _none_

Downlink:

- Uplink channels 1-7
- **921.9** - SF12BW125 (RX2 downlink only; SF12BW125 might be changed to SF9BW125)

## IN865-867

_India uses the 865MHz to 867MHz band. Currently we support the three LoRaWAN default frequencies ([according to KotahiNet](http://kotahi.net/network/sending-data/)), but plan to add more channels later. Therefore consider this as a BETA version of the Indian frequency plan._

Uplink:

1. **865.0625** - SF7BW125 to SF12BW125
2. **865.4025** - SF7BW125 to SF12BW125
3. **865.9850** - SF7BW125 to SF12BW125

Downlink:

- Uplink channels 1-3 (RX1)
- **866.550** - SF10BW125 (RX2)
