package timeh

import (
	"time"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	AfterKey = "After"

	AfterFuncKey = "AfterFunc"

	DateKey = "Date"

	FixedZoneKey = "FixedZone"

	LoadLocationKey = "LoadLocation"

	LoadLocationFromTZDataKey = "LoadLocationFromTZData"

	NewTickerKey = "NewTicker"

	NewTimerKey = "NewTimer"

	NowKey = "Now"

	ParseKey = "Parse"

	ParseDurationKey = "ParseDuration"

	ParseInLocationKey = "ParseInLocation"

	SinceKey = "Since"

	SleepKey = "Sleep"

	TickKey = "Tick"

	UnixKey = "Unix"

	UntilKey = "Until"
)

func New() hctx.Map {
	return hctx.Map{

		AfterKey: After,

		AfterFuncKey: AfterFunc,

		DateKey: Date,

		FixedZoneKey: FixedZone,

		LoadLocationKey: LoadLocation,

		LoadLocationFromTZDataKey: LoadLocationFromTZData,

		NewTickerKey: NewTicker,

		NewTimerKey: NewTimer,

		NowKey: Now,

		ParseKey: Parse,

		ParseDurationKey: ParseDuration,

		ParseInLocationKey: ParseInLocation,

		SinceKey: Since,

		SleepKey: Sleep,

		TickKey: Tick,

		UnixKey: Unix,

		UntilKey: Until,
	}
}

// After waits for the duration to elapse and then sends the current time
// on the returned channel.
// It is equivalent to NewTimer(d).C.
// The underlying Timer is not recovered by the garbage collector
// until the timer fires. If efficiency is a concern, use NewTimer
// instead and call Timer.Stop if the timer is no longer needed.
var After = time.After

// AfterFunc waits for the duration to elapse and then calls f
// in its own goroutine. It returns a Timer that can
// be used to cancel the call using its Stop method.
var AfterFunc = time.AfterFunc

// Date returns the Time corresponding to
// 	yyyy-mm-dd hh:mm:ss + nsec nanoseconds
// in the appropriate zone for that time in the given location.
//
// The month, day, hour, min, sec, and nsec values may be outside
// their usual ranges and will be normalized during the conversion.
// For example, October 32 converts to November 1.
//
// A daylight savings time transition skips or repeats times.
// For example, in the United States, March 13, 2011 2:15am never occurred,
// while November 6, 2011 1:15am occurred twice. In such cases, the
// choice of time zone, and therefore the time, is not well-defined.
// Date returns a time that is correct in one of the two zones involved
// in the transition, but it does not guarantee which.
//
// Date panics if loc is nil.
var Date = time.Date

// FixedZone returns a Location that always uses
// the given zone name and offset (seconds east of UTC).
var FixedZone = time.FixedZone

// LoadLocation returns the Location with the given name.
//
// If the name is &#34;&#34; or &#34;UTC&#34;, LoadLocation returns UTC.
// If the name is &#34;Local&#34;, LoadLocation returns Local.
//
// Otherwise, the name is taken to be a location name corresponding to a file
// in the IANA Time Zone database, such as &#34;America/New_York&#34;.
//
// The time zone database needed by LoadLocation may not be
// present on all systems, especially non-Unix systems.
// LoadLocation looks in the directory or uncompressed zip file
// named by the ZONEINFO environment variable, if any, then looks in
// known installation locations on Unix systems,
// and finally looks in $GOROOT/lib/time/zoneinfo.zip.
var LoadLocation = time.LoadLocation

// LoadLocationFromTZData returns a Location with the given name
// initialized from the IANA Time Zone database-formatted data.
// The data should be in the format of a standard IANA time zone file
// (for example, the content of /etc/localtime on Unix systems).
var LoadLocationFromTZData = time.LoadLocationFromTZData

// NewTicker returns a new Ticker containing a channel that will send the
// time with a period specified by the duration argument.
// It adjusts the intervals or drops ticks to make up for slow receivers.
// The duration d must be greater than zero; if not, NewTicker will panic.
// Stop the ticker to release associated resources.
var NewTicker = time.NewTicker

// NewTimer creates a new Timer that will send
// the current time on its channel after at least duration d.
var NewTimer = time.NewTimer

// Now returns the current local time.
var Now = time.Now

// Parse parses a formatted string and returns the time value it represents.
// The layout defines the format by showing how the reference time,
// defined to be
// 	Mon Jan 2 15:04:05 -0700 MST 2006
// would be interpreted if it were the value; it serves as an example of
// the input format. The same interpretation will then be made to the
// input string.
//
// Predefined layouts ANSIC, UnixDate, RFC3339 and others describe standard
// and convenient representations of the reference time. For more information
// about the formats and the definition of the reference time, see the
// documentation for ANSIC and the other constants defined by this package.
// Also, the executable example for Time.Format demonstrates the working
// of the layout string in detail and is a good reference.
//
// Elements omitted from the value are assumed to be zero or, when
// zero is impossible, one, so parsing &#34;3:04pm&#34; returns the time
// corresponding to Jan 1, year 0, 15:04:00 UTC (note that because the year is
// 0, this time is before the zero Time).
// Years must be in the range 0000..9999. The day of the week is checked
// for syntax but it is otherwise ignored.
//
// In the absence of a time zone indicator, Parse returns a time in UTC.
//
// When parsing a time with a zone offset like -0700, if the offset corresponds
// to a time zone used by the current location (Local), then Parse uses that
// location and zone in the returned time. Otherwise it records the time as
// being in a fabricated location with time fixed at the given zone offset.
//
// When parsing a time with a zone abbreviation like MST, if the zone abbreviation
// has a defined offset in the current location, then that offset is used.
// The zone abbreviation &#34;UTC&#34; is recognized as UTC regardless of location.
// If the zone abbreviation is unknown, Parse records the time as being
// in a fabricated location with the given zone abbreviation and a zero offset.
// This choice means that such a time can be parsed and reformatted with the
// same layout losslessly, but the exact instant used in the representation will
// differ by the actual zone offset. To avoid such problems, prefer time layouts
// that use a numeric zone offset, or use ParseInLocation.
var Parse = time.Parse

// ParseDuration parses a duration string.
// A duration string is a possibly signed sequence of
// decimal numbers, each with optional fraction and a unit suffix,
// such as &#34;300ms&#34;, &#34;-1.5h&#34; or &#34;2h45m&#34;.
// Valid time units are &#34;ns&#34;, &#34;us&#34; (or &#34;µs&#34;), &#34;ms&#34;, &#34;s&#34;, &#34;m&#34;, &#34;h&#34;.
var ParseDuration = time.ParseDuration

// ParseInLocation is like Parse but differs in two important ways.
// First, in the absence of time zone information, Parse interprets a time as UTC;
// ParseInLocation interprets the time as in the given location.
// Second, when given a zone offset or abbreviation, Parse tries to match it
// against the Local location; ParseInLocation uses the given location.
var ParseInLocation = time.ParseInLocation

// Since returns the time elapsed since t.
// It is shorthand for time.Now().Sub(t).
var Since = time.Since

// Sleep pauses the current goroutine for at least the duration d.
// A negative or zero duration causes Sleep to return immediately.
var Sleep = time.Sleep

// Tick is a convenience wrapper for NewTicker providing access to the ticking
// channel only. While Tick is useful for clients that have no need to shut down
// the Ticker, be aware that without a way to shut it down the underlying
// Ticker cannot be recovered by the garbage collector; it &#34;leaks&#34;.
// Unlike NewTicker, Tick will return nil if d &lt;= 0.
var Tick = time.Tick

// Unix returns the local Time corresponding to the given Unix time,
// sec seconds and nsec nanoseconds since January 1, 1970 UTC.
// It is valid to pass nsec outside the range [0, 999999999].
// Not all sec values have a corresponding time value. One such
// value is 1&lt;&lt;63-1 (the largest int64 value).
var Unix = time.Unix

// Until returns the duration until t.
// It is shorthand for t.Sub(time.Now()).
var Until = time.Until
