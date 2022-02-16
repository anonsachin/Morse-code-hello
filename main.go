package main

import (
    "machine"
    "time"
)

// Basic main setup for arduino
func main() {
    led := machine.LED
    led.Configure(machine.PinConfig{Mode: machine.PinOutput})
    for {
        m := &MorseCode{
            Led: led,
            Duration: time.Millisecond * 1000,
        }

        HelloInMorseCode(m)
    }
}

/*
    Below this is all the code for setting up and Blinking out HELLO
*/

// HelloInMorseCode is to create a hello morsecode message
func HelloInMorseCode(m *MorseCode){
    // lets start of

    // H -> . . . .
    // .
    m.Dot()
    m.Wait()
    // .
    m.Dot()
    m.Wait()
    // .
    m.Dot()
    m.Wait()
    // .
    m.Dot()

    m.LetterWait()

    // E -> . 
    m.Dot()

    m.LetterWait()

    // L -> . --- . .
    // . 
    m.Dot()
    m.Wait()
    // ---
    m.Dash()
    m.Wait()
    // . 
    m.Dot()
    m.Wait()
    // . 
    m.Dot()

    m.LetterWait()

    // L -> . --- . .
    // . 
    m.Dot()
    m.Wait()
    // ---
    m.Dash()
    m.Wait()
    // . 
    m.Dot()
    m.Wait()
    // . 
    m.Dot()

    m.LetterWait()

    // O -> --- --- ---
    // ---
    m.Dash()
    m.Wait()
    // ---
    m.Dash()
    m.Wait()
    // ---
    m.Dash()

    m.WordWait()
}

// MorseCode is the place to define the settings of the system.
type MorseCode struct{
    Duration time.Duration
    Led machine.Pin
}

// Wait is the wait times between the dot's and dashes.
func (m *MorseCode) Wait() {
    time.Sleep(m.Duration)
}

// Dot is the turning on of the pin for the duration of the dot.
func (m *MorseCode) Dot() {
    m.Led.High()
    m.Wait()
    m.Led.Low()
}

// Dash is the turning on of the pin for the duration of a dash.
func (m *MorseCode) Dash() {
    // 3 times the dot suration
    m.Led.High()
    m.Wait()
    m.Wait()
    m.Wait()
    m.Led.Low()
}

// LetterWait is the wait duration between between letters.
func (m *MorseCode) LetterWait() {
    // 3 times the normal wait
    m.Wait()
    m.Wait()
    m.Wait()
}

// WordWait is the wait duration between between words.
func (m *MorseCode) WordWait() {
    // 7 times the wait of the default time
    m.LetterWait()
    m.LetterWait()
    m.Wait()
}