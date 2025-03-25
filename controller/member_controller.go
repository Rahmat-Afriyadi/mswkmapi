package controller

import (
	"fmt"
	"wkm/entity"
	"wkm/request"
	"wkm/service"

	"github.com/gofiber/fiber/v2"
)

type MemberController interface {
	Mine(ctx *fiber.Ctx) error
	AddCard(ctx *fiber.Ctx) error
	CreateNewMemberCard(ctx *fiber.Ctx) error
}

type memberController struct {
	mS service.MemberService
}

func NewMemberController(mS service.MemberService) MemberController {
	return &memberController{
		mS: mS,
	}
}

func (c *memberController) Mine(ctx *fiber.Ctx) error {
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	data, _ := c.mS.Mine(details.ID)
	return ctx.JSON(map[string]interface{}{"data": data, "message": "Berhasil"})
}

func (c *memberController) AddCard(ctx *fiber.Ctx) error {
	body := request.AddMemberCard{}
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	body.UserID = details.ID
	result, err := c.mS.AddCard(body)
	fmt.Println("ini error gk sih ", err)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	return ctx.JSON(map[string]interface{}{"data": result, "message": "Berhasil"})
}
func (c *memberController) CreateNewMemberCard(ctx *fiber.Ctx) error {
	var body request.CreateNewMember
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	result, err := c.mS.CreateNewMemberCard(body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	return ctx.JSON(map[string]interface{}{"data": result, "message": "Berhasil"})
}
