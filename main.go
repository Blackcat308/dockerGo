package main

import "github.com/kataras/iris/v12"

func main() {
	application := iris.New()

	booksApi := application.Party("/books")
	{
		booksApi.Use(iris.Compression)

		booksApi.Get("/", list)
		booksApi.Post("/", create)
	}

	application.Listen(":7799")
}

type Book struct {
	Title string `json:"title"`
}

func list(ctx iris.Context) {
	books := []Book{
		{"三体"},
		{"斗破苍穹"},
		{"大秦帝国"},
	}

	ctx.JSON(books)
}

func create(ctx iris.Context) {
	var b Book
	err := ctx.ReadJSON(&b)

	if err != nil {
		ctx.StopWithProblem(
			iris.StatusBadRequest,
			iris.NewProblem().Title("").DetailErr(err),
		)
		return
	}

	print("create success!")

	ctx.StatusCode(iris.StatusCreated)
}
