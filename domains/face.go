package domains

type Face struct {
	name string
	body []byte
}

func GetFace(name string, body []byte) *Face {
	return &Face{
		name: name,
		body: body,
	}
}

func (f *Face) GetName() string {
	return f.name
}

func (f *Face) GetBody() []byte {
	return f.body
}
