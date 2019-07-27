package actions

import "github.com/gobuffalo/buffalo"

// CokeHandler as a new handler
// A handler is the action part
// handlers receive a buffalo context
func CokeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}
