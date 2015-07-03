package try

// Trial ...
type Trial struct {
	executions []func() error
	errored    error
}

// Begin invokes given func immediately!!
func Begin(f func() error) *Trial {
	tr := &Trial{}
	tr.errored = f()
	return tr
}

// Rescue invokes given func immediately if ERROR ALREADY OCCURED!!
func (tr *Trial) Rescue(f func() error) *Trial {
	if tr.errored == nil {
		return tr // abort func
	}
	tr.errored = f()
	return tr
}

// End returns finally errored or not.
func (tr *Trial) End() error {
	return tr.errored
}
