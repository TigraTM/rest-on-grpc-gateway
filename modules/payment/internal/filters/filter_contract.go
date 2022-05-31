// Package filters contains data for building sort in database query.
package filters

import "github.com/Masterminds/squirrel"

// FilterContract interface for sql builder.
type FilterContract interface {
	ApplyTo(builder squirrel.SelectBuilder) squirrel.SelectBuilder
}
