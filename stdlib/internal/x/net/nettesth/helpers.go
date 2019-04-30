package nettesth

import (
	"internal/x/net/nettest"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	TestConnKey = "TestConn"
)

func New() hctx.Map {
	return hctx.Map{

		TestConnKey: TestConn,
	}
}

// TestConn tests that a net.Conn implementation properly satisfies the interface.
// The tests should not produce any false positives, but may experience
// false negatives. Thus, some issues may only be detected when the test is
// run multiple times. For maximal effectiveness, run the tests under the
// race detector.
var TestConn = nettest.TestConn
