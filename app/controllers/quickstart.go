package controllers

import "github.com/revel/revel"

type QuickStart struct {
	App
}

func (c *QuickStart) Index() revel.Result {
	c.RenderArgs["isQuickStart"] = true

	return c.Render()
}
