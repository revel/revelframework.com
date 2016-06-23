package controllers

import "github.com/revel/revel"

type Docs struct {
	App
}

func (c *Docs) Index() revel.Result {
	c.RenderArgs["isDocs"] = true

	return c.Render()
}
