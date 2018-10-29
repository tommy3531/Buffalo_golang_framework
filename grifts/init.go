package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/tommarler/bottle2/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
