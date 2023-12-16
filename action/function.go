package action

type (
	fn     string
	action []any
)

func NewFunction() {
	return &Function{}
}

type Function struct {
}

func (f Function) Do() error {
}

func (f Function) Check() bool {
}
