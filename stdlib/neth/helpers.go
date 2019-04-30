package neth

import (
	"net"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	CIDRMaskKey = "CIDRMask"

	DialKey = "Dial"

	DialIPKey = "DialIP"

	DialTCPKey = "DialTCP"

	DialTimeoutKey = "DialTimeout"

	DialUDPKey = "DialUDP"

	DialUnixKey = "DialUnix"

	FileConnKey = "FileConn"

	FileListenerKey = "FileListener"

	FilePacketConnKey = "FilePacketConn"

	IPv4Key = "IPv4"

	IPv4MaskKey = "IPv4Mask"

	InterfaceAddrsKey = "InterfaceAddrs"

	InterfaceByIndexKey = "InterfaceByIndex"

	InterfaceByNameKey = "InterfaceByName"

	InterfacesKey = "Interfaces"

	JoinHostPortKey = "JoinHostPort"

	ListenKey = "Listen"

	ListenIPKey = "ListenIP"

	ListenMulticastUDPKey = "ListenMulticastUDP"

	ListenPacketKey = "ListenPacket"

	ListenTCPKey = "ListenTCP"

	ListenUDPKey = "ListenUDP"

	ListenUnixKey = "ListenUnix"

	ListenUnixgramKey = "ListenUnixgram"

	LookupAddrKey = "LookupAddr"

	LookupCNAMEKey = "LookupCNAME"

	LookupHostKey = "LookupHost"

	LookupIPKey = "LookupIP"

	LookupMXKey = "LookupMX"

	LookupNSKey = "LookupNS"

	LookupPortKey = "LookupPort"

	LookupSRVKey = "LookupSRV"

	LookupTXTKey = "LookupTXT"

	ParseCIDRKey = "ParseCIDR"

	ParseIPKey = "ParseIP"

	ParseMACKey = "ParseMAC"

	PipeKey = "Pipe"

	ResolveIPAddrKey = "ResolveIPAddr"

	ResolveTCPAddrKey = "ResolveTCPAddr"

	ResolveUDPAddrKey = "ResolveUDPAddr"

	ResolveUnixAddrKey = "ResolveUnixAddr"

	SplitHostPortKey = "SplitHostPort"
)

func New() hctx.Map {
	return hctx.Map{

		CIDRMaskKey: CIDRMask,

		DialKey: Dial,

		DialIPKey: DialIP,

		DialTCPKey: DialTCP,

		DialTimeoutKey: DialTimeout,

		DialUDPKey: DialUDP,

		DialUnixKey: DialUnix,

		FileConnKey: FileConn,

		FileListenerKey: FileListener,

		FilePacketConnKey: FilePacketConn,

		IPv4Key: IPv4,

		IPv4MaskKey: IPv4Mask,

		InterfaceAddrsKey: InterfaceAddrs,

		InterfaceByIndexKey: InterfaceByIndex,

		InterfaceByNameKey: InterfaceByName,

		InterfacesKey: Interfaces,

		JoinHostPortKey: JoinHostPort,

		ListenKey: Listen,

		ListenIPKey: ListenIP,

		ListenMulticastUDPKey: ListenMulticastUDP,

		ListenPacketKey: ListenPacket,

		ListenTCPKey: ListenTCP,

		ListenUDPKey: ListenUDP,

		ListenUnixKey: ListenUnix,

		ListenUnixgramKey: ListenUnixgram,

		LookupAddrKey: LookupAddr,

		LookupCNAMEKey: LookupCNAME,

		LookupHostKey: LookupHost,

		LookupIPKey: LookupIP,

		LookupMXKey: LookupMX,

		LookupNSKey: LookupNS,

		LookupPortKey: LookupPort,

		LookupSRVKey: LookupSRV,

		LookupTXTKey: LookupTXT,

		ParseCIDRKey: ParseCIDR,

		ParseIPKey: ParseIP,

		ParseMACKey: ParseMAC,

		PipeKey: Pipe,

		ResolveIPAddrKey: ResolveIPAddr,

		ResolveTCPAddrKey: ResolveTCPAddr,

		ResolveUDPAddrKey: ResolveUDPAddr,

		ResolveUnixAddrKey: ResolveUnixAddr,

		SplitHostPortKey: SplitHostPort,
	}
}

// CIDRMask returns an IPMask consisting of `ones&#39; 1 bits
// followed by 0s up to a total length of `bits&#39; bits.
// For a mask of this form, CIDRMask is the inverse of IPMask.Size.
var CIDRMask = net.CIDRMask

// Dial connects to the address on the named network.
//
// Known networks are &#34;tcp&#34;, &#34;tcp4&#34; (IPv4-only), &#34;tcp6&#34; (IPv6-only),
// &#34;udp&#34;, &#34;udp4&#34; (IPv4-only), &#34;udp6&#34; (IPv6-only), &#34;ip&#34;, &#34;ip4&#34;
// (IPv4-only), &#34;ip6&#34; (IPv6-only), &#34;unix&#34;, &#34;unixgram&#34; and
// &#34;unixpacket&#34;.
//
// For TCP and UDP networks, the address has the form &#34;host:port&#34;.
// The host must be a literal IP address, or a host name that can be
// resolved to IP addresses.
// The port must be a literal port number or a service name.
// If the host is a literal IPv6 address it must be enclosed in square
// brackets, as in &#34;[2001:db8::1]:80&#34; or &#34;[fe80::1%zone]:80&#34;.
// The zone specifies the scope of the literal IPv6 address as defined
// in RFC 4007.
// The functions JoinHostPort and SplitHostPort manipulate a pair of
// host and port in this form.
// When using TCP, and the host resolves to multiple IP addresses,
// Dial will try each IP address in order until one succeeds.
//
// Examples:
// 	Dial(&#34;tcp&#34;, &#34;golang.org:http&#34;)
// 	Dial(&#34;tcp&#34;, &#34;192.0.2.1:http&#34;)
// 	Dial(&#34;tcp&#34;, &#34;198.51.100.1:80&#34;)
// 	Dial(&#34;udp&#34;, &#34;[2001:db8::1]:domain&#34;)
// 	Dial(&#34;udp&#34;, &#34;[fe80::1%lo0]:53&#34;)
// 	Dial(&#34;tcp&#34;, &#34;:80&#34;)
//
// For IP networks, the network must be &#34;ip&#34;, &#34;ip4&#34; or &#34;ip6&#34; followed
// by a colon and a literal protocol number or a protocol name, and
// the address has the form &#34;host&#34;. The host must be a literal IP
// address or a literal IPv6 address with zone.
// It depends on each operating system how the operating system
// behaves with a non-well known protocol number such as &#34;0&#34; or &#34;255&#34;.
//
// Examples:
// 	Dial(&#34;ip4:1&#34;, &#34;192.0.2.1&#34;)
// 	Dial(&#34;ip6:ipv6-icmp&#34;, &#34;2001:db8::1&#34;)
// 	Dial(&#34;ip6:58&#34;, &#34;fe80::1%lo0&#34;)
//
// For TCP, UDP and IP networks, if the host is empty or a literal
// unspecified IP address, as in &#34;:80&#34;, &#34;0.0.0.0:80&#34; or &#34;[::]:80&#34; for
// TCP and UDP, &#34;&#34;, &#34;0.0.0.0&#34; or &#34;::&#34; for IP, the local system is
// assumed.
//
// For Unix networks, the address must be a file system path.
var Dial = net.Dial

// DialIP acts like Dial for IP networks.
//
// The network must be an IP network name; see func Dial for details.
//
// If laddr is nil, a local address is automatically chosen.
// If the IP field of raddr is nil or an unspecified IP address, the
// local system is assumed.
var DialIP = net.DialIP

// DialTCP acts like Dial for TCP networks.
//
// The network must be a TCP network name; see func Dial for details.
//
// If laddr is nil, a local address is automatically chosen.
// If the IP field of raddr is nil or an unspecified IP address, the
// local system is assumed.
var DialTCP = net.DialTCP

// DialTimeout acts like Dial but takes a timeout.
//
// The timeout includes name resolution, if required.
// When using TCP, and the host in the address parameter resolves to
// multiple IP addresses, the timeout is spread over each consecutive
// dial, such that each is given an appropriate fraction of the time
// to connect.
//
// See func Dial for a description of the network and address
// parameters.
var DialTimeout = net.DialTimeout

// DialUDP acts like Dial for UDP networks.
//
// The network must be a UDP network name; see func Dial for details.
//
// If laddr is nil, a local address is automatically chosen.
// If the IP field of raddr is nil or an unspecified IP address, the
// local system is assumed.
var DialUDP = net.DialUDP

// DialUnix acts like Dial for Unix networks.
//
// The network must be a Unix network name; see func Dial for details.
//
// If laddr is non-nil, it is used as the local address for the
// connection.
var DialUnix = net.DialUnix

// FileConn returns a copy of the network connection corresponding to
// the open file f.
// It is the caller&#39;s responsibility to close f when finished.
// Closing c does not affect f, and closing f does not affect c.
var FileConn = net.FileConn

// FileListener returns a copy of the network listener corresponding
// to the open file f.
// It is the caller&#39;s responsibility to close ln when finished.
// Closing ln does not affect f, and closing f does not affect ln.
var FileListener = net.FileListener

// FilePacketConn returns a copy of the packet network connection
// corresponding to the open file f.
// It is the caller&#39;s responsibility to close f when finished.
// Closing c does not affect f, and closing f does not affect c.
var FilePacketConn = net.FilePacketConn

// IPv4 returns the IP address (in 16-byte form) of the
// IPv4 address a.b.c.d.
var IPv4 = net.IPv4

// IPv4Mask returns the IP mask (in 4-byte form) of the
// IPv4 mask a.b.c.d.
var IPv4Mask = net.IPv4Mask

// InterfaceAddrs returns a list of the system&#39;s unicast interface
// addresses.
//
// The returned list does not identify the associated interface; use
// Interfaces and Interface.Addrs for more detail.
var InterfaceAddrs = net.InterfaceAddrs

// InterfaceByIndex returns the interface specified by index.
//
// On Solaris, it returns one of the logical network interfaces
// sharing the logical data link; for more precision use
// InterfaceByName.
var InterfaceByIndex = net.InterfaceByIndex

// InterfaceByName returns the interface specified by name.
var InterfaceByName = net.InterfaceByName

// Interfaces returns a list of the system&#39;s network interfaces.
var Interfaces = net.Interfaces

// JoinHostPort combines host and port into a network address of the
// form &#34;host:port&#34;. If host contains a colon, as found in literal
// IPv6 addresses, then JoinHostPort returns &#34;[host]:port&#34;.
//
// See func Dial for a description of the host and port parameters.
var JoinHostPort = net.JoinHostPort

// Listen announces on the local network address.
//
// The network must be &#34;tcp&#34;, &#34;tcp4&#34;, &#34;tcp6&#34;, &#34;unix&#34; or &#34;unixpacket&#34;.
//
// For TCP networks, if the host in the address parameter is empty or
// a literal unspecified IP address, Listen listens on all available
// unicast and anycast IP addresses of the local system.
// To only use IPv4, use network &#34;tcp4&#34;.
// The address can use a host name, but this is not recommended,
// because it will create a listener for at most one of the host&#39;s IP
// addresses.
// If the port in the address parameter is empty or &#34;0&#34;, as in
// &#34;127.0.0.1:&#34; or &#34;[::1]:0&#34;, a port number is automatically chosen.
// The Addr method of Listener can be used to discover the chosen
// port.
//
// See func Dial for a description of the network and address
// parameters.
var Listen = net.Listen

// ListenIP acts like ListenPacket for IP networks.
//
// The network must be an IP network name; see func Dial for details.
//
// If the IP field of laddr is nil or an unspecified IP address,
// ListenIP listens on all available IP addresses of the local system
// except multicast IP addresses.
var ListenIP = net.ListenIP

// ListenMulticastUDP acts like ListenPacket for UDP networks but
// takes a group address on a specific network interface.
//
// The network must be a UDP network name; see func Dial for details.
//
// ListenMulticastUDP listens on all available IP addresses of the
// local system including the group, multicast IP address.
// If ifi is nil, ListenMulticastUDP uses the system-assigned
// multicast interface, although this is not recommended because the
// assignment depends on platforms and sometimes it might require
// routing configuration.
// If the Port field of gaddr is 0, a port number is automatically
// chosen.
//
// ListenMulticastUDP is just for convenience of simple, small
// applications. There are golang.org/x/net/ipv4 and
// golang.org/x/net/ipv6 packages for general purpose uses.
var ListenMulticastUDP = net.ListenMulticastUDP

// ListenPacket announces on the local network address.
//
// The network must be &#34;udp&#34;, &#34;udp4&#34;, &#34;udp6&#34;, &#34;unixgram&#34;, or an IP
// transport. The IP transports are &#34;ip&#34;, &#34;ip4&#34;, or &#34;ip6&#34; followed by
// a colon and a literal protocol number or a protocol name, as in
// &#34;ip:1&#34; or &#34;ip:icmp&#34;.
//
// For UDP and IP networks, if the host in the address parameter is
// empty or a literal unspecified IP address, ListenPacket listens on
// all available IP addresses of the local system except multicast IP
// addresses.
// To only use IPv4, use network &#34;udp4&#34; or &#34;ip4:proto&#34;.
// The address can use a host name, but this is not recommended,
// because it will create a listener for at most one of the host&#39;s IP
// addresses.
// If the port in the address parameter is empty or &#34;0&#34;, as in
// &#34;127.0.0.1:&#34; or &#34;[::1]:0&#34;, a port number is automatically chosen.
// The LocalAddr method of PacketConn can be used to discover the
// chosen port.
//
// See func Dial for a description of the network and address
// parameters.
var ListenPacket = net.ListenPacket

// ListenTCP acts like Listen for TCP networks.
//
// The network must be a TCP network name; see func Dial for details.
//
// If the IP field of laddr is nil or an unspecified IP address,
// ListenTCP listens on all available unicast and anycast IP addresses
// of the local system.
// If the Port field of laddr is 0, a port number is automatically
// chosen.
var ListenTCP = net.ListenTCP

// ListenUDP acts like ListenPacket for UDP networks.
//
// The network must be a UDP network name; see func Dial for details.
//
// If the IP field of laddr is nil or an unspecified IP address,
// ListenUDP listens on all available IP addresses of the local system
// except multicast IP addresses.
// If the Port field of laddr is 0, a port number is automatically
// chosen.
var ListenUDP = net.ListenUDP

// ListenUnix acts like Listen for Unix networks.
//
// The network must be &#34;unix&#34; or &#34;unixpacket&#34;.
var ListenUnix = net.ListenUnix

// ListenUnixgram acts like ListenPacket for Unix networks.
//
// The network must be &#34;unixgram&#34;.
var ListenUnixgram = net.ListenUnixgram

// LookupAddr performs a reverse lookup for the given address, returning a list
// of names mapping to that address.
//
// When using the host C library resolver, at most one result will be
// returned. To bypass the host resolver, use a custom Resolver.
var LookupAddr = net.LookupAddr

// LookupCNAME returns the canonical name for the given host.
// Callers that do not care about the canonical name can call
// LookupHost or LookupIP directly; both take care of resolving
// the canonical name as part of the lookup.
//
// A canonical name is the final name after following zero
// or more CNAME records.
// LookupCNAME does not return an error if host does not
// contain DNS &#34;CNAME&#34; records, as long as host resolves to
// address records.
var LookupCNAME = net.LookupCNAME

// LookupHost looks up the given host using the local resolver.
// It returns a slice of that host&#39;s addresses.
var LookupHost = net.LookupHost

// LookupIP looks up host using the local resolver.
// It returns a slice of that host&#39;s IPv4 and IPv6 addresses.
var LookupIP = net.LookupIP

// LookupMX returns the DNS MX records for the given domain name sorted by preference.
var LookupMX = net.LookupMX

// LookupNS returns the DNS NS records for the given domain name.
var LookupNS = net.LookupNS

// LookupPort looks up the port for the given network and service.
var LookupPort = net.LookupPort

// LookupSRV tries to resolve an SRV query of the given service,
// protocol, and domain name. The proto is &#34;tcp&#34; or &#34;udp&#34;.
// The returned records are sorted by priority and randomized
// by weight within a priority.
//
// LookupSRV constructs the DNS name to look up following RFC 2782.
// That is, it looks up _service._proto.name. To accommodate services
// publishing SRV records under non-standard names, if both service
// and proto are empty strings, LookupSRV looks up name directly.
var LookupSRV = net.LookupSRV

// LookupTXT returns the DNS TXT records for the given domain name.
var LookupTXT = net.LookupTXT

// ParseCIDR parses s as a CIDR notation IP address and prefix length,
// like &#34;192.0.2.0/24&#34; or &#34;2001:db8::/32&#34;, as defined in
// RFC 4632 and RFC 4291.
//
// It returns the IP address and the network implied by the IP and
// prefix length.
// For example, ParseCIDR(&#34;192.0.2.1/24&#34;) returns the IP address
// 192.0.2.1 and the network 192.0.2.0/24.
var ParseCIDR = net.ParseCIDR

// ParseIP parses s as an IP address, returning the result.
// The string s can be in dotted decimal (&#34;192.0.2.1&#34;)
// or IPv6 (&#34;2001:db8::68&#34;) form.
// If s is not a valid textual representation of an IP address,
// ParseIP returns nil.
var ParseIP = net.ParseIP

// ParseMAC parses s as an IEEE 802 MAC-48, EUI-48, EUI-64, or a 20-octet
// IP over InfiniBand link-layer address using one of the following formats:
//   01:23:45:67:89:ab
//   01:23:45:67:89:ab:cd:ef
//   01:23:45:67:89:ab:cd:ef:00:00:01:23:45:67:89:ab:cd:ef:00:00
//   01-23-45-67-89-ab
//   01-23-45-67-89-ab-cd-ef
//   01-23-45-67-89-ab-cd-ef-00-00-01-23-45-67-89-ab-cd-ef-00-00
//   0123.4567.89ab
//   0123.4567.89ab.cdef
//   0123.4567.89ab.cdef.0000.0123.4567.89ab.cdef.0000
var ParseMAC = net.ParseMAC

// Pipe creates a synchronous, in-memory, full duplex
// network connection; both ends implement the Conn interface.
// Reads on one end are matched with writes on the other,
// copying data directly between the two; there is no internal
// buffering.
var Pipe = net.Pipe

// ResolveIPAddr returns an address of IP end point.
//
// The network must be an IP network name.
//
// If the host in the address parameter is not a literal IP address,
// ResolveIPAddr resolves the address to an address of IP end point.
// Otherwise, it parses the address as a literal IP address.
// The address parameter can use a host name, but this is not
// recommended, because it will return at most one of the host name&#39;s
// IP addresses.
//
// See func Dial for a description of the network and address
// parameters.
var ResolveIPAddr = net.ResolveIPAddr

// ResolveTCPAddr returns an address of TCP end point.
//
// The network must be a TCP network name.
//
// If the host in the address parameter is not a literal IP address or
// the port is not a literal port number, ResolveTCPAddr resolves the
// address to an address of TCP end point.
// Otherwise, it parses the address as a pair of literal IP address
// and port number.
// The address parameter can use a host name, but this is not
// recommended, because it will return at most one of the host name&#39;s
// IP addresses.
//
// See func Dial for a description of the network and address
// parameters.
var ResolveTCPAddr = net.ResolveTCPAddr

// ResolveUDPAddr returns an address of UDP end point.
//
// The network must be a UDP network name.
//
// If the host in the address parameter is not a literal IP address or
// the port is not a literal port number, ResolveUDPAddr resolves the
// address to an address of UDP end point.
// Otherwise, it parses the address as a pair of literal IP address
// and port number.
// The address parameter can use a host name, but this is not
// recommended, because it will return at most one of the host name&#39;s
// IP addresses.
//
// See func Dial for a description of the network and address
// parameters.
var ResolveUDPAddr = net.ResolveUDPAddr

// ResolveUnixAddr returns an address of Unix domain socket end point.
//
// The network must be a Unix network name.
//
// See func Dial for a description of the network and address
// parameters.
var ResolveUnixAddr = net.ResolveUnixAddr

// SplitHostPort splits a network address of the form &#34;host:port&#34;,
// &#34;host%zone:port&#34;, &#34;[host]:port&#34; or &#34;[host%zone]:port&#34; into host or
// host%zone and port.
//
// A literal IPv6 address in hostport must be enclosed in square
// brackets, as in &#34;[::1]:80&#34;, &#34;[::1%lo0]:80&#34;.
//
// See func Dial for a description of the hostport parameter, and host
// and port results.
var SplitHostPort = net.SplitHostPort
