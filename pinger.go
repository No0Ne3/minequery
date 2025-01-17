package minequery

import (
	"net"
	"time"
)

// PingerOption is a configuring function that applies certain changes to Pinger.
type PingerOption func(*Pinger)

// WithDialer sets Pinger Dialer used on Ping* function calls.
//goland:noinspection GoUnusedExportedFunction
func WithDialer(dialer *net.Dialer) PingerOption {
	return func(p *Pinger) {
		p.Dialer = dialer
	}
}

// WithTimeout sets Pinger Dialer timeout to the provided value.
//goland:noinspection GoUnusedExportedFunction
func WithTimeout(timeout time.Duration) PingerOption {
	return func(p *Pinger) {
		p.Timeout = timeout
		p.Dialer.Timeout = timeout
	}
}

// WithUseStrict sets Pinger UseStrict to the provided value.
//goland:noinspection GoUnusedExportedFunction
func WithUseStrict(useStrict bool) PingerOption {
	return func(p *Pinger) { p.UseStrict = useStrict }
}

// WithProtocolVersion16 sets Pinger ProtocolVersion16 value.
//goland:noinspection GoUnusedExportedFunction
func WithProtocolVersion16(version byte) PingerOption {
	return func(p *Pinger) {
		p.ProtocolVersion16 = version
	}
}

// WithProtocolVersion17 sets Pinger ProtocolVersion17 value.
//goland:noinspection GoUnusedExportedFunction
func WithProtocolVersion17(version int32) PingerOption {
	return func(p *Pinger) {
		p.ProtocolVersion17 = version
	}
}

// defaultPinger is a default (zero-value) Pinger used in functions
// that don't have Pinger as receiver. The default Pinger has timeout set to 15 seconds.
var defaultPinger = NewPinger(WithTimeout(15 * time.Second))

// Pinger contains options to ping and query Minecraft servers.
type Pinger struct {
	// Dialer used to establish and maintain connection with servers.
	Dialer *net.Dialer

	// Timeout is used to set TCP/UDP connection timeout on call of Ping* and Query* functions.
	Timeout time.Duration

	// UseStrict is a configuration value that defines if tolerable errors (in server ping/query responses)
	// that are by default silently ignored should be actually returned as errors.
	UseStrict bool

	// ProtocolVersion16 is protocol version to use when pinging with Ping16 function.
	// By default, Ping16ProtocolVersion162 (=74) will be used.
	// See ping_16.go for full list of built-in constants.
	ProtocolVersion16 byte

	// ProtocolVersion17 is protocol version to use when pinging with Ping17 function.
	// By default, Ping17ProtocolVersionUndefined (=-1) will be used.
	// See ping_17.go for full list of built-in constants.
	ProtocolVersion17 int32
}

// NewPinger constructs new Pinger instance optionally with additional options.
func NewPinger(options ...PingerOption) *Pinger {
	pinger := &Pinger{Dialer: &net.Dialer{}}
	for _, configure := range options {
		configure(pinger)
	}
	return pinger
}
