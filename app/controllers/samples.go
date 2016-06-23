package controllers

import "github.com/revel/revel"

type Samples struct {
	App
}

func (c *Samples) Index() revel.Result {
	c.RenderArgs["isSamples"] = true

	return c.Render()
}
