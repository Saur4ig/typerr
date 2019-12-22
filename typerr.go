package typerr

type Typerr struct {
	Err string
}

func (t Typerr) Error() string {
	return t.Err
}
