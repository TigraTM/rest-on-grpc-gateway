package api

import (
	"fmt"

	paymentpb "rest-on-grpc-gateway/api/proto/payment/v1"
	"rest-on-grpc-gateway/modules/payment/internal/filters"
)

// getSortForPaymentHistory get sort filter for payment history.
func getSortForPaymentHistory(in *paymentpb.GetPaymentsHistoryByUserIDRequest) (_ filters.Filter, err error) {
	sort := filters.NewFilter()

	sortField := filters.NewPaymentFilter(in.SortBy)
	if err = sortField.IsValid(); err != nil {
		return filters.Filter{}, fmt.Errorf("%w: %s", errUncorrectedFilter, err.Error())
	}
	sort.SetSortingField(sortField)

	sortDirection := filters.NewSortDirection(in.Sort)
	if err = sortDirection.IsValid(); err != nil {
		return filters.Filter{}, fmt.Errorf("%w: %s", errUncorrectedFilter, err.Error())
	}
	sort.SetDirection(sortDirection)

	return *sort, nil
}

// getPaging get paging.
func getPaging(in *paymentpb.GetPaymentsHistoryByUserIDRequest) (_ filters.Paging, err error) {
	p := filters.NewPaging()

	if err = p.SetOffset(in.Offset); err != nil {
		return filters.Paging{}, fmt.Errorf("%w: %s", errUncorrectedPaging, err.Error())
	}

	if err = p.SetLimit(in.Limit); err != nil {
		return filters.Paging{}, fmt.Errorf("%w: %s", errUncorrectedPaging, err.Error())
	}

	return *p, nil
}
