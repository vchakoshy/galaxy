package flex

//go:generate jsonenums -type=Status
type Status int

const (
	Pending Status = iota
	Sent
	Received
	Rejected
)
