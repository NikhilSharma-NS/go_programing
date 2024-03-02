package main

type Time struct {
	sec  int64
	nsec int32
}

func Now() Time {
	return Time{sec: 1, nsec: 1}
}

func (t Time) Add(d int) Time {
	return Time{}
}

func (t *Time) UnmarshalBinary(data []byte) error {
	return nil
}
