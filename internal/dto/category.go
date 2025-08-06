package dto

type Category struct {
	id         uint16
	name       string
	desciption string
}

func (c *Category) Id() uint16 {
	return c.id
}

func (c *Category) Name() string {
	return c.name
}

func (c *Category) Desciption() string {
	return c.desciption
}

func (c *Category) SetName(name string) {
	c.name = name
}

func (c *Category) SetDescription(description string) {
	c.desciption = description
}
