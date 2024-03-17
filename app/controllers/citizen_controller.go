package controllers

import (
	"encoding/json"
	"reflect"
	"strconv"
	"time"

	"github.com/bit-wiz/data-store-a/app/models"
	"github.com/bit-wiz/data-store-a/app/queries"
	"github.com/bit-wiz/data-store-a/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetFields(c *fiber.Ctx) error {
	fields := []string{}

	for i := 0; i < reflect.ValueOf(models.Citizen{}).NumField(); i++ {
		if (reflect.ValueOf(models.Citizen{}).Type().Field(i).Name == "ID") {
			continue
		}
		fields = append(fields, reflect.ValueOf(models.Citizen{}).Type().Field(i).Tag.Get("bson"))
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"data":    fields,
	})

}

func GetAllCitizens(c *fiber.Ctx) error {
	filter := utils.SearchQ(c.Query("s"))
	page, _ := strconv.Atoi(c.Query("page", "1"))
	findOptions := utils.FindOpts(c.Query("sort", "a-created_at"), page)

	data, err := queries.DB.GetUsers(filter, findOptions)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})

}

func GetCitizen(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   "no id given",
		})
	}

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	data, err := queries.DB.GetUser(objId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func CreateCitizen(c *fiber.Ctx) error {
	var data models.Citizen

	// parse the stringified json into the struct
	if err := json.Unmarshal(c.Body(), &data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   "invalid JSON data",
		})
	}
	c.BodyParser(&data)
	// data.ID = fmt.Sprintf("%s_%s", data.FirstName, uuid.New().String()[0:8])
	data.CreatedAt = time.Now()
	data.UpdatedAt = primitive.Timestamp{T: uint32(time.Now().Unix())}

	err := queries.DB.CreateUser(data)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func UpdateCitizen(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   "no id given",
		})
	}

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	var data models.Citizen
	// c.BodyParser(&data) can't use this need to parse stringified json
	if err := json.Unmarshal(c.Body(), &data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   "invalid JSON data",
		})
	}

	data.UpdatedAt = primitive.Timestamp{T: uint32(time.Now().Unix())}

	err = queries.DB.UpdateUser(objId, data)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func DeleteCitizen(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   "no id given",
		})
	}

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	err = queries.DB.DeleteUser(objId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"data":    "deleted",
	})
}
