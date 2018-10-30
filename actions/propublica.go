package actions

import "github.com/gobuffalo/buffalo"

func PropublicaHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("propublica.html"))
}
