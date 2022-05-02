package filters

import (
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
)

// SortDirection type sort direction.
type SortDirection string

const (
	// SortDirectionASC sort ASC.
	SortDirectionASC SortDirection = `ASC`
	// SortDirectionDESC sort DESC.
	SortDirectionDESC SortDirection = `DESC`
)

func (d SortDirection) IsValid() error {
	switch d {
	case SortDirectionASC,
		SortDirectionDESC:
		return nil
	}

	// nolint:goerr113 // need dynamic error.
	return fmt.Errorf("%s - invalid sorting direction, should be ASC or DESC", d)
}

func NewSortDirection(value string) SortDirection {
	if value == "" {
		return SortDirectionDESC
	}

	return SortDirection(strings.ToUpper(value))
}

// PaymentFilter type for filter.
type PaymentFilter string

const (
	// PaymentFilterCreateAt field `create_at` from the database, by which the sorting will be done.
	PaymentFilterCreateAt PaymentFilter = "create_at"
	// PaymentFilterAmount field `amount` from the database, by which the sorting will be done.
	PaymentFilterAmount PaymentFilter = "amount"
)

// NewPaymentFilter return PaymentFilter to lower.
func NewPaymentFilter(value string) PaymentFilter {
	if value == "" {
		return PaymentFilterCreateAt
	}

	return PaymentFilter(strings.ToLower(value))
}

func (f PaymentFilter) IsValid() error {
	switch f {
	case PaymentFilterCreateAt,
		PaymentFilterAmount:
		return nil
	}

	// nolint:goerr113 // need dynamic error.
	return fmt.Errorf("%s - invalid sorting field, should be amount or created_at", f)
}

// Filter structure for sorting.
type Filter struct {
	direction SortDirection
	field     PaymentFilter
}

// NewFilter build and return Filter with default value.
func NewFilter() *Filter {
	return &Filter{
		direction: SortDirectionDESC,
		field:     PaymentFilterCreateAt,
	}
}

// SetDirection set field direction.
func (f *Filter) SetDirection(direction SortDirection) {
	f.direction = direction
}

// SetSortingField set sorting field.
func (f *Filter) SetSortingField(field PaymentFilter) {
	f.field = field
}

// ApplyTo implements FilterContract.
func (f *Filter) ApplyTo(builder squirrel.SelectBuilder) squirrel.SelectBuilder {
	return builder.OrderByClause(fmt.Sprintf("%s %s", f.field, f.direction))
}
