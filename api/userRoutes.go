package api

import (
	"log"
	"strconv"

	"github.com/bdkiran/lastimado-service/db"
	"github.com/gofiber/fiber/v2"
)

//Define how we want our api to look
//{
//status - [success, fail, error]
//data - {}
//message - When an api call fails due to server error
//}

func getAllUsers() fiber.Handler {
	return func(c *fiber.Ctx) error {
		users, err := db.GetUsers()
		if err != nil {
			c.Status(400)
			return c.JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}
		c.Status(200)
		return c.JSON(&fiber.Map{
			"status": "success",
			"data":   users,
		})
	}
}

func getUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, err := strconv.Atoi(c.Params("userID"))
		if err != nil {
			return c.JSON(&fiber.Map{
				"status": "fail",
				"data": map[string]interface{}{
					"userID": "Invalid userID submitted",
				},
			})
		}
		//GET user code
		user, err := db.GetUser(userID)
		if err != nil {
			return c.JSON(&fiber.Map{
				"status": "fail",
				"data": map[string]interface{}{
					"userID": "Unable to retrieve user",
				},
			})
		}
		return c.JSON(&fiber.Map{
			"status": "success",
			"data":   user,
		})
	}
}

func createUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var userToCreate db.User
		err := c.BodyParser(&userToCreate)
		if err != nil {
			log.Println(err)
			//This response should be able to define what fields are an issue?
			return c.JSON(&fiber.Map{
				"status": "fail",
				"data": map[string]interface{}{
					"user": "Invalid user parameters submitted",
				},
			})
		}
		err = db.InsertUser(userToCreate)
		if err != nil {
			log.Println(err)
			return c.JSON(&fiber.Map{
				"status": "fail",
				"data": map[string]interface{}{
					"user": "Unable to create user",
				},
			})
		}

		//We probabably want to change this api response.....
		return c.JSON(&fiber.Map{
			"status": "success",
			"data":   "User successfully created",
		})
	}
}

func updateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var userToUpdate db.User
		err := c.BodyParser(&userToUpdate)
		if err != nil {
			log.Println(err)
			return c.JSON(&fiber.Map{
				"status": "fail",
				"data": map[string]interface{}{
					"user": "Invalid user parameters submitted",
				},
			})
		}
		err = db.UpdateUser(userToUpdate)
		if err != nil {
			log.Println(err)
			return c.JSON(&fiber.Map{
				"status": "fail",
				"data": map[string]interface{}{
					"user": "Unable to create user",
				},
			})
		}
		return c.JSON(&fiber.Map{
			"status": "success",
			"data":   "User successfully updated",
		})
	}
}

func deleteUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userIDToDelete, err := strconv.Atoi(c.Params("userID"))
		if err != nil {
			log.Println(err)
			return c.JSON(&fiber.Map{
				"status": "fail",
				"data": map[string]interface{}{
					"user": "Invalid parameters sent",
				},
			})
		}

		err = db.DeleteUser(userIDToDelete)
		if err != nil {
			log.Println(err)
			return c.JSON(&fiber.Map{
				"status": "fail",
				"data": map[string]interface{}{
					"user": "Unable to delete user",
				},
			})
		}
		return c.JSON(&fiber.Map{
			"status": "success",
			"data":   "User successfully deleted.",
		})
	}
}
