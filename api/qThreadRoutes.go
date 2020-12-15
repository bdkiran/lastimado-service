package api

import (
	"log"
	"strconv"

	"github.com/bdkiran/lastimado-service/db"
	"github.com/gofiber/fiber/v2"
)

func getAllQThreads() fiber.Handler {
	return func(c *fiber.Ctx) error {
		qThreads, err := db.GetQThreads()
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
			"data":   qThreads,
		})
	}
}

func getQThread() fiber.Handler {
	return func(c *fiber.Ctx) error {
		qThreadID, err := strconv.Atoi(c.Params("qThreadID"))
		if err != nil {
			c.Status(400)
			return c.JSON(&fiber.Map{
				"status": "fail",
				"data": map[string]interface{}{
					"qThreadID": "Invalid qThreadID submitted",
				},
			})
		}

		qThread, err := db.GetQThread(qThreadID)
		if err != nil {
			c.Status(400)
			return c.JSON(&fiber.Map{
				"status": "fail",
				"data": map[string]interface{}{
					"QThread": "Unable to retrieve user",
				},
			})
		}
		return c.JSON(&fiber.Map{
			"status": "success",
			"data":   qThread,
		})
	}
}

func createQThread() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var qThreadToCreate db.QThread
		err := c.BodyParser(&qThreadToCreate)
		if err != nil {
			log.Println(err)
			//This response should be able to define what fields are an issue?
			return c.JSON(&fiber.Map{
				"status": "fail",
				"data": map[string]interface{}{
					"qThread": "Invalid qThread parameters submitted",
				},
			})
		}
		err = db.InsertQThread(qThreadToCreate)
		if err != nil {
			log.Println(err)
			return c.JSON(&fiber.Map{
				"status": "fail",
				"data": map[string]interface{}{
					"user": "Unable to create qThread",
				},
			})
		}

		//We probabably want to change this api response.....
		return c.JSON(&fiber.Map{
			"status": "success",
			"data":   "qThread successfully created",
		})
	}
}

func updateQThread() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var qThreadToUpdated db.QThread
		err := c.BodyParser(&qThreadToUpdated)
		if err != nil {
			log.Println(err)
			return c.JSON(&fiber.Map{
				"status": "fail",
				"data": map[string]interface{}{
					"qThread": "Invalid qThread parameters submitted",
				},
			})
		}
		err = db.UpdateQThread(qThreadToUpdated)
		if err != nil {
			log.Println(err)
			return c.JSON(&fiber.Map{
				"status": "fail",
				"data": map[string]interface{}{
					"qThread": "Unable to update qThread",
				},
			})
		}
		return c.JSON(&fiber.Map{
			"status": "success",
			"data":   "QThread successfully updated",
		})
	}
}

func deleteQThread() fiber.Handler {
	return func(c *fiber.Ctx) error {
		qThreadIDToDelete, err := strconv.Atoi(c.Params("qThreadID"))
		if err != nil {
			log.Println(err)
			return c.JSON(&fiber.Map{
				"status": "fail",
				"data": map[string]interface{}{
					"qThread": "Invalid parameters sent",
				},
			})
		}

		err = db.DeleteQThread(qThreadIDToDelete)
		if err != nil {
			log.Println(err)
			return c.JSON(&fiber.Map{
				"status": "fail",
				"data": map[string]interface{}{
					"qThread": "Unable to delete qThread",
				},
			})
		}
		return c.JSON(&fiber.Map{
			"status": "success",
			"data":   "qThread successfully deleted.",
		})
	}
}
