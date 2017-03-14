Hello,


I'm having a bad data problem on the IC880A connected Rasp Pi3 board. Not only the data I send a lot but the results board IC880A received less.

Here is the config image on the IM880A node and the receiver data on the IC880a board.
Node send to ic880a 6 seconds once.

*************************************Configure node IM880a***************************************

#define RF_FREQUENCY                                868100000 // Hz
#define TX_OUTPUT_POWER                             20        // dBm


#define LORA_BANDWIDTH                              0         // [0: 125 kHz, //0
                                                              //  1: 250 kHz,
                                                              //  2: 500 kHz,
                                                              //  3: Reserved]
#define LORA_SPREADING_FACTOR                       7         // [SF7..SF12]
#define LORA_CODINGRATE                             1         // [1: 4/5,
                                                              //  2: 4/6,
                                                              //  3: 4/7,
                                                              //  4: 4/8]
#define LORA_PREAMBLE_LENGTH                        8         // Same for Tx and Rx
#define LORA_SYMBOL_TIMEOUT                         5         // Symbols
#define LORA_FIX_LENGTH_PAYLOAD_ON                  false
#define LORA_IQ_INVERSION_ON                        false

typedef enum
{
    LOWPOWER,
    RX,
    RX_TIMEOUT,
    RX_ERROR,
    TX,
    TX_TIMEOUT,
}States_t;

#define RX_TIMEOUT_VALUE                            100
#define BUFFER_SIZE                                 13 // Define the payload size here

**************************************************************************************

### 
Image reciever data board IC880a
https://github.com/VanXungUIT/Arduino_Esp8266/blob/master/bad.png

Help me!

Thanks.

