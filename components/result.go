package components

type Result struct {
	Err   error
	Value string `json:"value"`
}

func (r *Result) Ok() bool {
	return r.Err == nil
}
