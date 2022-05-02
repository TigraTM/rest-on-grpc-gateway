package filters

import (
	"errors"

	"github.com/Masterminds/squirrel"
)

const (
	defaultOffset = 0
	defaultLimit  = 100
)

// Errors.
var (
	ErrInvalidLimit  = errors.New("limit must be > 0")
	ErrInvalidOffset = errors.New("offset must be >= 0")
)

// Paging structure for building paging.
type Paging struct {
	limit  uint64
	offset uint64
}

// NewPaging build and return paging with default value.
func NewPaging() *Paging {
	return &Paging{
		limit:  defaultLimit,
		offset: defaultOffset,
	}
}

// SetLimit set limit.
func (p *Paging) SetLimit(limit int64) error {
	if limit <= 0 {
		p.limit = uint64(defaultLimit)

		return nil
	}

	p.limit = uint64(limit)

	return nil
}

// SetOffset set offset.
func (p *Paging) SetOffset(offset int64) error {
	if offset < 0 {
		p.offset = uint64(defaultOffset)

		return nil
	}

	p.offset = uint64(offset)

	return nil
}

// ApplyTo implements FilterContract.
func (p *Paging) ApplyTo(builder squirrel.SelectBuilder) squirrel.SelectBuilder {
	return builder.Limit(p.limit).Offset(p.offset)
}
