package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/simonireilly/coke/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
