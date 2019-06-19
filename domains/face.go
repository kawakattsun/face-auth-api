package domains

type Face struct {
	name string
	body []byte
}

func (f *Face) getName() string {
	return f.name
}

func (f *Face) getBody() []byte {
	return f.body
}
