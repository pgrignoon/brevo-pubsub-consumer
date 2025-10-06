package function

import "cloud.google.com/go/bigquery"

/*
Event is an interface that all event structs must implement.
It is used to convert the event struct to a bigquery event struct, for the generic DecodeAndSend function.
*/
type Event interface {
	ToBigquery() any
}

/*
Convert a string to a bigquery.NullString
*/
func toNullString(s *string) bigquery.NullString {
	if s == nil {
		return bigquery.NullString{}
	}
	return bigquery.NullString{StringVal: *s, Valid: true}
}

/*
Convert a int64 to a bigquery.NullInt64
*/
func toNullInt64(i *int64) bigquery.NullInt64 {
	if i == nil {
		return bigquery.NullInt64{}
	}
	return bigquery.NullInt64{Int64: *i, Valid: true}
}

/*
Convert a float64 to a bigquery.NullFloat64
*/
func toNullFloat64(f *float64) bigquery.NullFloat64 {
	if f == nil {
		return bigquery.NullFloat64{}
	}
	return bigquery.NullFloat64{Float64: *f, Valid: true}
}
