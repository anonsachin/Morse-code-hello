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

[Blog to learn to setup tinygo for arduino uno](https://create.arduino.cc/projecthub/alankrantas/tinygo-on-arduino-uno-an-introduction-6130f6)

- tinygo
- avrdude
- make

`Note - Also set the environment variables for system and your user for the tinygo and avrdude bins`

Versions
--------
**Tinygo**
```
    tinygo version 0.21.0 windows/amd64 (using go version go1.17.6 and LLVM version 11.0.0)
```

**AVRDUDE**
```
    avrdude version 6.3
```

**GIT**
```
    GNU Make 4.2.1
```