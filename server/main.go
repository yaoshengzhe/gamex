package main

import "github.com/kataras/iris/v12"

type GameItem struct {
	Title string `json:"title"`
}

func main() {
	app := iris.New()

	api := app.Party("/games")
	{
		api.Use(iris.Compression)
		api.Get("/", list)
	}

	app.Listen(":8080")
}

func list(ctx iris.Context) {
	ctx.JSON([]GameItem{
		{"Pokemon"}, {"Final Fantasy CC"},
	})
}
