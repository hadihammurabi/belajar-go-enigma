package main

import (
	"bytes"
	"fmt"
)

type Enigma struct {
	plugboard *Plugboard
	rotors    [3]*Rotor
}

func NewEnigma() *Enigma {
	return &Enigma{
		plugboard: NewPlugboard(),
		rotors: [3]*Rotor{
			NewRotor(1),
			NewRotor(11),
			NewRotor(21),
		},
	}
}

func (e Enigma) Press(key byte) byte {
	out := e.plugboard.Press(key)
	for _, r := range e.rotors {
		out = r.Press(out)
	}
	for i := len(e.rotors) - 1; i >= 0; i-- {
		out = e.rotors[i].Press(out)
	}
	return out
}

func (e Enigma) Encrypt(in string) string {
	for _, r := range e.rotors {
		fmt.Printf("%d ", r.num)
	}
	fmt.Println()

	out := bytes.Buffer{}
	for _, c := range []byte(in) {
		if c == ' ' {
			out.WriteByte(' ')
			continue
		}
		o := e.Press(c)
		out.WriteByte(o)
	}

	for _, r := range e.rotors {
		fmt.Printf("%d ", r.num)
	}
	fmt.Println()

	return out.String()
}
