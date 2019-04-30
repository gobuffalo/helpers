package sqlh

import (
	"database/sql"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	DriversKey = "Drivers"

	NamedKey = "Named"

	OpenKey = "Open"

	OpenDBKey = "OpenDB"

	RegisterKey = "Register"
)

func New() hctx.Map {
	return hctx.Map{

		DriversKey: Drivers,

		NamedKey: Named,

		OpenKey: Open,

		OpenDBKey: OpenDB,

		RegisterKey: Register,
	}
}

// Drivers returns a sorted list of the names of the registered drivers.
var Drivers = sql.Drivers

// Named provides a more concise way to create NamedArg values.
//
// Example usage:
//
//     db.ExecContext(ctx, `
//         delete from Invoice
//         where
//             TimeCreated &lt; @end
//             and TimeCreated &gt;= @start;`,
//         sql.Named(&#34;start&#34;, startTime),
//         sql.Named(&#34;end&#34;, endTime),
//     )
var Named = sql.Named

// Open opens a database specified by its database driver name and a
// driver-specific data source name, usually consisting of at least a
// database name and connection information.
//
// Most users will open a database via a driver-specific connection
// helper function that returns a *DB. No database drivers are included
// in the Go standard library. See https://golang.org/s/sqldrivers for
// a list of third-party drivers.
//
// Open may just validate its arguments without creating a connection
// to the database. To verify that the data source name is valid, call
// Ping.
//
// The returned DB is safe for concurrent use by multiple goroutines
// and maintains its own pool of idle connections. Thus, the Open
// function should be called just once. It is rarely necessary to
// close a DB.
var Open = sql.Open

// OpenDB opens a database using a Connector, allowing drivers to
// bypass a string based data source name.
//
// Most users will open a database via a driver-specific connection
// helper function that returns a *DB. No database drivers are included
// in the Go standard library. See https://golang.org/s/sqldrivers for
// a list of third-party drivers.
//
// OpenDB may just validate its arguments without creating a connection
// to the database. To verify that the data source name is valid, call
// Ping.
//
// The returned DB is safe for concurrent use by multiple goroutines
// and maintains its own pool of idle connections. Thus, the OpenDB
// function should be called just once. It is rarely necessary to
// close a DB.
var OpenDB = sql.OpenDB

// Register makes a database driver available by the provided name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
var Register = sql.Register
