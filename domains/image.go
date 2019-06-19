package domains

type Image interface {
	GetName() string
	GetBody() []byte
}
