package routes

import (
	"go-auth/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
	app.Get("/api/showall", controllers.ShowAll)
	app.Post("/api/addnewaccount", controllers.AddNewAccount)
	app.Post("/api/creditbalance", controllers.CreditBalance)
	app.Post("/api/debitbalance", controllers.DebitBalance)
	app.Post("/api/deleteaccount", controllers.DeleteAccount)

}
