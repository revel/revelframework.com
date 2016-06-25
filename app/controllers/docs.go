package controllers

import (
	"github.com/revel/revel"
	"github.com/revel/revelframework.com/app/meta"
)

type Docs struct {
	App
}

func (c *Docs) Index() revel.Result {
	c.RenderArgs["isDocs"] = true
	c.RenderArgs["docsMeta"] = meta.Docs

	// pathPrefix := c.RenderArgs["revelDocsPathPrefix"].(string)

	return c.Render()
}
