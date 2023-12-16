package action

type actionable interface {
	Do() error
	Check() bool
}
