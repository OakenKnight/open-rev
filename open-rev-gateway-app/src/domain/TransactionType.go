package domain

type Transaction int

const (
	Evaluate Transaction = iota + 1
	Submit
)

func (t Transaction) String() string {
	return [...]string{"Evaluate Transaction", "Submit Transaction"}[t-1]
}
