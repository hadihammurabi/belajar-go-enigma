package main

type Plugboard struct {
	chars map[byte]byte
}

func NewPlugboard() *Plugboard {
	return &Plugboard{
		chars: map[byte]byte{
			'c': 'n',
		},
	}
}

func (p Plugboard) Press(key byte) byte {
	if out, ok := p.chars[key]; ok {
		return out
	}

	return key
}
