package api

import (
	"log"
	"strconv"

	"github.com/bdkiran/lastimado-service/db"
	"github.com/gofiber/fiber/v2"
)

func getAllQPosts() fiber.Handler {
	return func(c *fiber.Ctx) error {
		qPosts, err := db.GetQPosts()
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
			"data":   qPosts,
		})
	}
}

func getQPost() fiber.Handler {
	return func(c *fiber.Ctx) error {
		qPostID, err := strconv.Atoi(c.Params("qPostID"))
		if err != nil {
			c.Status(400)
			return c.JSON(&fiber.Map{
				"status": "fail",
				"data": map[string]interface{}{
					"qPostID": "Invalid qPostID submitted",
				},
			})
		}

		qPost, err := db.GetQPost(qPostID)
		if err != nil {
			log.Println(err)
			c.Status(400)
			return c.JSON(&fiber.Map{
				"status": "fail",
				"data": map[string]interface{}{
					"QPost": "Unable to retrieve qPost",
				},
			})
		}
		c.Status(200)
		return c.JSON(&fiber.Map{
			"status": "success",
			"data":   qPost,
		})
	}
}

func createQPost() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var qPostToCreate db.QPost
		err := c.BodyParser(&qPostToCreate)
		if err != nil {
			log.Println(err)
			//This response should be able to define what fields are an issue?
			return c.JSON(&fiber.Map{
				"status": "fail",
				"data": map[string]interface{}{
					"qPost": "Invalid qPost parameters submitted",
				},
			})
		}
		err = db.InsertQPost(qPostToCreate)
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
			"data":   "qPost successfully created",
		})
	}
}

func updateQPost() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var qPostToUpdate db.QPost
		err := c.BodyParser(&qPostToUpdate)
		if err != nil {
			log.Println(err)
			return c.JSON(&fiber.Map{
				"status": "fail",
				"data": map[string]interface{}{
					"qPost": "Invalid qPost parameters submitted",
				},
			})
		}
		err = db.UpdateQPost(qPostToUpdate)
		if err != nil {
			log.Println(err)
			return c.JSON(&fiber.Map{
				"status": "fail",
				"data": map[string]interface{}{
					"qPost": "Unable to update qPost",
				},
			})
		}
		return c.JSON(&fiber.Map{
			"status": "success",
			"data":   "QPost successfully updated",
		})
	}
}

func deleteQPost() fiber.Handler {
	return func(c *fiber.Ctx) error {
		qPostIDToDelete, err := strconv.Atoi(c.Params("qPostID"))
		if err != nil {
			log.Println(err)
			return c.JSON(&fiber.Map{
				"status": "fail",
				"data": map[string]interface{}{
					"qPost": "Invalid parameters sent",
				},
			})
		}

		err = db.DeleteQPost(qPostIDToDelete)
		if err != nil {
			log.Println(err)
			return c.JSON(&fiber.Map{
				"status": "fail",
				"data": map[string]interface{}{
					"qPost": "Unable to delete qThread",
				},
			})
		}
		return c.JSON(&fiber.Map{
			"status": "success",
			"data":   "qPost successfully deleted.",
		})
	}
}
