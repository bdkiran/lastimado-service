package api

import (
	"github.com/gofiber/fiber/v2"
)

//InitializeRoutes initializes all new app routes
func InitializeRoutes(app *fiber.App) {
	app.Get("/users", getAllUsers())
	app.Get("/user/:userID", getUser())
	app.Post("/user", createUser())
	app.Put("/user", updateUser())
	app.Delete("/user/:userID", deleteUser())

	app.Get("/qthreads", getAllQThreads())
	app.Get("/qthread/:qThreadID", getQThread())
	app.Post("/qthread", createQThread())
	app.Put("/qthread", updateQThread())
	app.Delete("/qthread/:qThreadID", deleteQThread())

	app.Get("/qposts", getAllQPosts())
	app.Get("/qpost/:qPostID", getQPost())
	app.Post("/qpost", createQPost())
	app.Put("/qpost", updateQPost())
	app.Delete("/qpost/:qPostID", deleteQPost())
}
