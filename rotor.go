package main

type Rotor struct {
	num       int
	usedCount int
}

func NewRotor(num int) *Rotor {
	return &Rotor{
		num: num,
	}
}

func (r *Rotor) Press(key byte) byte {
	out := key + byte(r.num)
	if out > 122 {
		out -= 26
	}
	r.Shift()
	return out
}

func (r Rotor) IsLast() bool {
	return r.num >= 25
}

func (r *Rotor) Shift() {
	if r.usedCount >= 1 {
		r.usedCount = 0
	} else {
		r.usedCount++
		return
	}

	if r.IsLast() {
		r.num = 0
	} else {
		r.num++
	}
}
