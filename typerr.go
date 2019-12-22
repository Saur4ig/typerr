package typerr

type typerr struct {
	msg string
}

func (t typerr) Error() string {
	return t.msg
}

func (t typerr) New(msg string) error {
	return &typerr{msg: msg}
}
