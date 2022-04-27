package api

import (
	"fmt"
	"rest-on-grpc-gateway/modules/payment/internal/filters"

	paymentpb "rest-on-grpc-gateway/api/proto/payment/v1"
)

// getSortForPaymentHistory get sort filter for payment history.
func getSortForPaymentHistory(in *paymentpb.GetPaymentsHistoryByAccountIDRequest) (_ filters.Filter, err error) {
	sort := filters.NewFilter()

	sortField := filters.NewPaymentFilter(in.SortBy)
	if err = sortField.IsValid(); err != nil {
		return filters.Filter{}, fmt.Errorf("sortField.IsValid: %w", err)
	}
	sort.SetSortingField(sortField)

	sortDirection := filters.NewSortDirection(in.Sort)
	if err = sortDirection.IsValid(); err != nil {
		return filters.Filter{}, fmt.Errorf("sortDirection.IsValid: %w", err)
	}
	sort.SetDirection(sortDirection)

	return *sort, nil
}

// getPaging get paging.
func getPaging(in *paymentpb.GetPaymentsHistoryByAccountIDRequest) (_ filters.Paging, err error) {
	p := filters.NewPaging()

	if err = p.SetOffset(in.Offset); err != nil {
		return filters.Paging{}, fmt.Errorf("p.SetOffset: %w", err)
	}

	if err = p.SetLimit(in.Limit); err != nil {
		return filters.Paging{}, fmt.Errorf("p.SetLimit: %w", err)
	}

	return *p, nil
}
