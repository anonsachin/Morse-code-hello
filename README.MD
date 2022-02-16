MORSE CODE HELLO
================

Blinking **Hello** using morse code on an arduino uno. 
The development is done on windows 11 machine, not tested on linux or mac.

Set the `PORT` env variable to the one where your device is connected and then run.
```
    make flash
```

Requirements
------------

[Blog to learn to setup tinygo for arduino uno on windows](https://create.arduino.cc/projecthub/alankrantas/tinygo-on-arduino-uno-an-introduction-6130f6)

- tinygo
- avrdude
- make

`Note - Also set the environment variables for system and your user for the tinygo and avrdude bins`