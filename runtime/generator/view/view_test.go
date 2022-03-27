package view_test

import (
	"context"
	"testing"

	"github.com/matryer/is"
	"gitlab.com/mnm/bud/internal/budtest"
)

func TestHello(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()
	dir := t.TempDir()
	bud := budtest.New(dir)
	bud.Files["action/action.go"] = `
		package action
		type Controller struct {}
		func (c *Controller) Index() string { return "" }
	`
	bud.Files["view/index.svelte"] = `<h1>hello</h1>`
	bud.NodeModules["svelte"] = "3.42.3"
	project, err := bud.Compile(ctx)
	is.NoErr(err)
	app, err := project.Build(ctx)
	is.NoErr(err)
	is.NoErr(app.Exists("bud/.app/view/view.go"))
	is.NoErr(app.Exists("bud/.app/main.go"))
	server, err := app.Start(ctx)
	is.NoErr(err)
	defer server.Close()
	res, err := server.Get("/")
	is.NoErr(err)
	// HTML response
	is.NoErr(res.ExpectHeaders(`
		HTTP/1.1 200 OK
		Content-Type: text/html
		Date: Fri, 31 Dec 2021 00:00:00 GMT
	`))
	is.NoErr(res.ContainsBody(`<h1>hello</h1>`))
}
