package transport

import "tumblr/circuit/use/circuit"

var (
	ErrEnd            = circuit.NewError("end")
	ErrAlreadyClosed  = circuit.NewError("already closed")
	errCollision      = circuit.NewError("conn id collision")
	ErrNotSupported   = circuit.NewError("not supported")
)