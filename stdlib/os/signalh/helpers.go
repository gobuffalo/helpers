package signalh

import (
	"os/signal"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	IgnoreKey = "Ignore"

	NotifyKey = "Notify"

	ResetKey = "Reset"

	StopKey = "Stop"
)

func New() hctx.Map {
	return hctx.Map{

		IgnoreKey: Ignore,

		NotifyKey: Notify,

		ResetKey: Reset,

		StopKey: Stop,
	}
}

// Ignore causes the provided signals to be ignored. If they are received by
// the program, nothing will happen. Ignore undoes the effect of any prior
// calls to Notify for the provided signals.
// If no signals are provided, all incoming signals will be ignored.
var Ignore = signal.Ignore

// Notify causes package signal to relay incoming signals to c.
// If no signals are provided, all incoming signals will be relayed to c.
// Otherwise, just the provided signals will.
//
// Package signal will not block sending to c: the caller must ensure
// that c has sufficient buffer space to keep up with the expected
// signal rate. For a channel used for notification of just one signal value,
// a buffer of size 1 is sufficient.
//
// It is allowed to call Notify multiple times with the same channel:
// each call expands the set of signals sent to that channel.
// The only way to remove signals from the set is to call Stop.
//
// It is allowed to call Notify multiple times with different channels
// and the same signals: each channel receives copies of incoming
// signals independently.
var Notify = signal.Notify

// Reset undoes the effect of any prior calls to Notify for the provided
// signals.
// If no signals are provided, all signal handlers will be reset.
var Reset = signal.Reset

// Stop causes package signal to stop relaying incoming signals to c.
// It undoes the effect of all prior calls to Notify using c.
// When Stop returns, it is guaranteed that c will receive no more signals.
var Stop = signal.Stop
