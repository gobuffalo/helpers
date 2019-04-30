package nettesth

import (
	"internal/x/net/internal/nettest"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	IsMulticastCapableKey = "IsMulticastCapable"

	MaxOpenFilesKey = "MaxOpenFiles"

	NewLocalListenerKey = "NewLocalListener"

	NewLocalPacketListenerKey = "NewLocalPacketListener"

	ProtocolNotSupportedKey = "ProtocolNotSupported"

	RoutedInterfaceKey = "RoutedInterface"

	SupportsIPv4Key = "SupportsIPv4"

	SupportsIPv6Key = "SupportsIPv6"

	SupportsIPv6MulticastDeliveryOnLoopbackKey = "SupportsIPv6MulticastDeliveryOnLoopback"

	SupportsRawIPSocketKey = "SupportsRawIPSocket"

	TestableNetworkKey = "TestableNetwork"
)

func New() hctx.Map {
	return hctx.Map{

		IsMulticastCapableKey: IsMulticastCapable,

		MaxOpenFilesKey: MaxOpenFiles,

		NewLocalListenerKey: NewLocalListener,

		NewLocalPacketListenerKey: NewLocalPacketListener,

		ProtocolNotSupportedKey: ProtocolNotSupported,

		RoutedInterfaceKey: RoutedInterface,

		SupportsIPv4Key: SupportsIPv4,

		SupportsIPv6Key: SupportsIPv6,

		SupportsIPv6MulticastDeliveryOnLoopbackKey: SupportsIPv6MulticastDeliveryOnLoopback,

		SupportsRawIPSocketKey: SupportsRawIPSocket,

		TestableNetworkKey: TestableNetwork,
	}
}

// IsMulticastCapable reports whether ifi is an IP multicast-capable
// network interface. Network must be &#34;ip&#34;, &#34;ip4&#34; or &#34;ip6&#34;.
var IsMulticastCapable = nettest.IsMulticastCapable

// MaxOpenFiles returns the maximum number of open files for the
// caller&#39;s process.
var MaxOpenFiles = nettest.MaxOpenFiles

// NewLocalListener returns a listener which listens to a loopback IP
// address or local file system path.
// Network must be &#34;tcp&#34;, &#34;tcp4&#34;, &#34;tcp6&#34;, &#34;unix&#34; or &#34;unixpacket&#34;.
var NewLocalListener = nettest.NewLocalListener

// NewLocalPacketListener returns a packet listener which listens to a
// loopback IP address or local file system path.
// Network must be &#34;udp&#34;, &#34;udp4&#34;, &#34;udp6&#34; or &#34;unixgram&#34;.
var NewLocalPacketListener = nettest.NewLocalPacketListener

// ProtocolNotSupported reports whether err is a protocol not
// supported error.
var ProtocolNotSupported = nettest.ProtocolNotSupported

// RoutedInterface returns a network interface that can route IP
// traffic and satisfies flags. It returns nil when an appropriate
// network interface is not found. Network must be &#34;ip&#34;, &#34;ip4&#34; or
// &#34;ip6&#34;.
var RoutedInterface = nettest.RoutedInterface

// SupportsIPv4 reports whether the platform supports IPv4 networking
// functionality.
var SupportsIPv4 = nettest.SupportsIPv4

// SupportsIPv6 reports whether the platform supports IPv6 networking
// functionality.
var SupportsIPv6 = nettest.SupportsIPv6

// SupportsIPv6MulticastDeliveryOnLoopback reports whether the
// platform supports IPv6 multicast packet delivery on software
// loopback interface.
var SupportsIPv6MulticastDeliveryOnLoopback = nettest.SupportsIPv6MulticastDeliveryOnLoopback

// SupportsRawIPSocket reports whether the platform supports raw IP
// sockets.
var SupportsRawIPSocket = nettest.SupportsRawIPSocket

// TestableNetwork reports whether network is testable on the current
// platform configuration.
var TestableNetwork = nettest.TestableNetwork
