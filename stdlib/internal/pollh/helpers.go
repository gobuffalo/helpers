package pollh

import (
	"internal/poll"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	DupCloseOnExecKey = "DupCloseOnExec"

	IsPollDescriptorKey = "IsPollDescriptor"

	IsPollDescriptorKey = "IsPollDescriptor"

	IsPollDescriptorKey = "IsPollDescriptor"

	SendFileKey = "SendFile"

	SendFileKey = "SendFile"

	SendFileKey = "SendFile"

	SendFileKey = "SendFile"

	SpliceKey = "Splice"
)

func New() hctx.Map {
	return hctx.Map{

		DupCloseOnExecKey: DupCloseOnExec,

		IsPollDescriptorKey: IsPollDescriptor,

		IsPollDescriptorKey: IsPollDescriptor,

		IsPollDescriptorKey: IsPollDescriptor,

		SendFileKey: SendFile,

		SendFileKey: SendFile,

		SendFileKey: SendFile,

		SendFileKey: SendFile,

		SpliceKey: Splice,
	}
}

// DupCloseOnExec dups fd and marks it close-on-exec.
var DupCloseOnExec = poll.DupCloseOnExec

// IsPollDescriptor reports whether fd is the descriptor being used by the poller.
// This is only used for testing.
var IsPollDescriptor = poll.IsPollDescriptor

// IsPollDescriptor reports whether fd is the descriptor being used by the poller.
// This is only used for testing.
var IsPollDescriptor = poll.IsPollDescriptor

// IsPollDescriptor reports whether fd is the descriptor being used by the poller.
// This is only used for testing.
var IsPollDescriptor = poll.IsPollDescriptor

// SendFile wraps the sendfile system call.
var SendFile = poll.SendFile

// SendFile wraps the sendfile system call.
var SendFile = poll.SendFile

// SendFile wraps the sendfile system call.
var SendFile = poll.SendFile

// SendFile wraps the TransmitFile call.
var SendFile = poll.SendFile

// Splice transfers at most remain bytes of data from src to dst, using the
// splice system call to minimize copies of data from and to userspace.
//
// Splice creates a temporary pipe, to serve as a buffer for the data transfer.
// src and dst must both be stream-oriented sockets.
//
// If err != nil, sc is the system call which caused the error.
var Splice = poll.Splice
