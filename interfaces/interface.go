package interfaces

type Humaner interface {
	Name() string
	Height() uint32
}

type Chinese struct {
	name string
	height uint32
}

func NewChinese(name string, h uint32) *Chinese {
	return &Chinese{name: name, height: h}
}

func (c *Chinese) Name() string {
	return c.name
}

func (c *Chinese) Height() uint32 {
	return c.height
}