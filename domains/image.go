package domains

type Image interface {
	getName() string
	getBody() []byte
}
