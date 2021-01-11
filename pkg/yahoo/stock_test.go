package yahoo

import (
	"fmt"
	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	//iter := options.GetStraddle("000725.SZ")
	//
	//fmt.Println(iter.Meta())
	//
	//for iter.Next() {
	//	fmt.Println(iter.Straddle().Strike)
	//}
	//if iter.Err() != nil {
	//	fmt.Println(iter.Err())
	//}

	// Basic quote example.
	// --------------------
	// q, err := quote.Get("SPY")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(q)
	// }

	// Basic chart example.
	// --------------------
	startDate := time.Date(2021, 1, 11, 00, 00, 00, 00, time.Local)
	endDate := time.Date(2021, 1, 11, 23, 59, 59, 00, time.Local)
	params := &chart.Params{
		Symbol:   "000725.SZ",
		Interval: datetime.OneDay,
		Start:    datetime.New(&startDate),
		End:      datetime.New(&endDate),
	}

	iter := chart.Get(params)

	for iter.Next() {
		b := iter.Bar()
		fmt.Println(datetime.FromUnix(b.Timestamp), b.Open.String(), b.Close.String())

	}
	if iter.Err() != nil {
		fmt.Println(iter.Err())
	}
}
