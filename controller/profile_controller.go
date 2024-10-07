package controller

import (
	"fmt"
	"time"
	"wkm/entity"
	"wkm/service"

	"github.com/gofiber/fiber/v2"
)

type ProfileController interface {
	Me(ctx *fiber.Ctx) error
	UploadImageProfile(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
}

type profileController struct {
	pS service.ProfileService
}

func NewProfileController(pS service.ProfileService) ProfileController {
	return &profileController{
		pS: pS,
	}
}

func (c *profileController) Me(ctx *fiber.Ctx) error {
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	data := c.pS.Me(details.ID)
	return ctx.JSON(map[string]interface{}{"data": data, "message": "Berhasil"})
}
func (c *profileController) UploadImageProfile(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		fmt.Println("error file ", err)
		return err
	}
	destination := fmt.Sprintf("./uploads/%s", file.Filename)
	if err := ctx.SaveFile(file, destination); err != nil {
		return err
	}
	return ctx.JSON(map[string]interface{}{"data": "/uploads/" + file.Filename, "message": "Berhasil"})
}
func (c *profileController) Update(ctx *fiber.Ctx) error {
	var body entity.Profile
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	parseTimezone := body.TglLahir.Add(7 * time.Hour)
	body.TglLahir = &parseTimezone
	profile := c.pS.Update(body)
	if profile.Id == "" {
		return ctx.Status(400).JSON(map[string]interface{}{"data": "Data Gagal di update", "message": "Gagal"})
	}
	return ctx.JSON(map[string]interface{}{"data": "Data Berhasil di update", "message": "Berhasil"})
}
