# LoRa hardware - Nodes

## Overview
Here's a table of a few hardware boards supporting LoRa.
Specific instructions can be found below. If you're unsure
what to buy, send us an [email](Contact).

| Device            | Chip   | MCU       | Connector   | Cost (EUR)   | Comments                   |
| ----------------- | ------ | --------- | ----------- | ------------:| -------------------------- |
| Sodaq Tatu + Bee  | ?      | AVR       | Bee         | 35,- + 35,-  | Arduino-IDE compatible     |
| Kickst. TTN Uno   | SX1276 |           | IO          | tbd (35-45)  | Arduino-IDE compatible
| Netblocks         | SX1272 | STM32L151 | IO          | 35,-         | program with ST-link       |
| HopeRF RFM92W     | RFM92W | -         | (raw board) | 10,-         | cheap              |
| Kerlink loramote  | SX1272 | yes       | IO          | 150,-        | professional; has GPS      |
| ....              |        |           |             |              |                            |

There's many more; for a bigger list, [contact us](Contact).

## SODAQ Tatu
The SODAQ Tatu is based on the ATmega1284P and Arduino compatible.

Get them:

* Tatu: [http://shop.sodaq.com/en/sodaq-tatu.html](http://shop.sodaq.com/en/sodaq-tatu.html)
  OR Mbili (bigger): [http://shop.sodaq.com/en/sodaq-mbili.html](http://shop.sodaq.com/en/sodaq-mbili.html)
* LoRa bee: [http://shop.sodaq.com/en/lorabee.html](http://shop.sodaq.com/en/lorabee.html)

Program them:

1. Download the latest Arduino IDE: [https://www.arduino.cc/en/Main/Software](https://www.arduino.cc/en/Main/Software)
2. Install hardware support for Tatu: [https://github.com/SodaqMoja/HardwareTatu](https://github.com/SodaqMoja/HardwareTatu).
   Ignore the part on bootloaders, just extract the zip within the Arduino hardware folder and restart the IDE.
3. Install dependencies: [https://github.com/bblanchon/ArduinoJson](https://github.com/bblanchon/ArduinoJson)
4. Clone our demo repository: [https://github.com/TheThingsNetwork/loraduino](https://github.com/TheThingsNetwork/loraduino)
5. Add sensors, add custom logic in the .ino Arduino sketch file:
    * uncomment the right `DEVICE_` line, optionally change other define statements on top
    * you may want to change the `measure()` and `send_measurement()` functions for your particular sensors
    (for help, read the [https://www.google.nl/search?q=arduino+getting+started+pdf](Arduino - Getting Started) book).
    * set a unique device ID in `EmbitBee.h` (search for `NETWORK_ADDRESS` and update `uint8_t Address[4]={0xAA, 0xFF, 0x01, 0x10};`).   
6. Compile, Upload, boot.
   The device will start sending measurements about once every two minutes.



## Kickstarter TTN Arduino Uno
We'll launch a kickstarter early October, including an Arduino Uno with build-in
LoRa chip and open-source libraries to connect to The Things Network gateways.
More info: [http://thethingsnetwork.org/landing/kickstarter](http://thethingsnetwork.org/landing/kickstarter)

As we're getting closer to shipping this page will be updated with instructions.


## Netblocks
We're working on support for the Netblocks, but not everything's working yet.

Get them:

* Choose one of the kits: [http://www.netblocks.eu/shop/](http://www.netblocks.eu/shop/)

Program them:

1. [https://github.com/TheThingsNetwork/XRange](https://github.com/TheThingsNetwork/XRange)
2. ask us what to do next ;-)




## HopeRF
This board works with Arduino (tested on Teensy 3.1/LC). Does not work on Uno (ATmega 328) due to memory limitations.

Here is howto connect it to an Arduino: [https://github.com/matthijskooijman/arduino-lmic/wiki](https://github.com/matthijskooijman/arduino-lmic/wiki)

The github is based on IBM LMIC 1.4 (LoraMAC-in-C), a working LMIC 1.5 based implementation will be available soon.



## Kerlink Loramote
This device should work out-of-the-box. It will start sending GPS, temperature and battery data as soon as it is powered.

TODO: custom applications compile. Kerlink has a nice wiki with instructions.



## More devices will be contributed soon (by you?)



## General settings instructions
There are a couple of settings in LoRaWAN related to The Things Network
that all devices (no matter the LoRaWAN implementation) should set:

  * Network key / NwkSKey: a fixed network key for all The Things Network devices.
    It is `2B7E151628AED2A6ABF7158809CF4F3C` or in bytes (c++): `{0x2B, 0x7E, 0x15, 0x16, 0x28, 0xAE, 0xD2, 0xA6, 0xAB, 0xF7, 0x15, 0x88, 0x09, 0xCF, 0x4F, 0x3C}`.
    This key is used for calculating message integrity (MIC). It is only used for payload encryption if Port = 0 (MAC commands, not yet implemented in The Things Network).
  * Application session key / AppSKey: used for encryption of the payload (Port > 0). To get started, use the default one,
    which is identical to the Network key (note this will make all your testdata publicly available through our API).
  * DeviceID / Device Address: this should be unique per node.
    In the final routing implementation we expect to have blocks of devices
    being routed to application servers. But for now, you can pick anything
    you like.
  * Application key / AppKey: a key specific for an application. Used for over-the-air activation, not yet implemented in The Things Network.
  * Frequencies: The Kerlink gateways are listening to all SF on 868.1, 868.3, 868.5, 867.1, 867.3, 867.5, 867.7, 867.9 Mhz, SF7/250Khz on 868.3 MHz, and FSK (50kbps) on 868.8 MHz. See: [LoRaWAN Frequencies](http://thethingsnetwork.org/wiki/LoRaWAN-Frequencies)


# Software (where's my data?)
[Software Overview](http://thethingsnetwork.org/wiki/Software/Overview)