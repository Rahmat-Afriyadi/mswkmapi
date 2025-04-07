package user

import (
	"fmt"
	"strconv"
	"wkm/entity"
	"wkm/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	repo := NewUserSRepository(db)
	service := NewUserSService(repo)
	handler := NewUserSController(service)

	routes := app.Group("/userses")
	routes.Get("/master-data", middleware.DeserializeUser, handler.MasterData)
	routes.Get("/master-data/count", middleware.DeserializeUser, handler.MasterDataCount)
	routes.Get("/detail/:id", handler.DetailUserS)
	routes.Get("/detail/free/:id", handler.DetailUserS)
	routes.Post("/create-user", middleware.DeserializeUser, handler.CreateUserS)
	routes.Post("/update-user", middleware.DeserializeUser, handler.UpdateUserS)
}

type UserSController interface {
	CreateUserS(ctx *fiber.Ctx) error
	UpdateUserS(ctx *fiber.Ctx) error
	MasterData(ctx *fiber.Ctx) error
	MasterDataCount(ctx *fiber.Ctx) error
	DetailUserS(ctx *fiber.Ctx) error
}

type userSController struct {
	userSService UserSService
}

func NewUserSController(aS UserSService) UserSController {
	return &userSController{
		userSService: aS,
	}
}

func (tr *userSController) DetailUserS(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	return ctx.JSON(tr.userSService.DetailUserS(id))
}

func (tr *userSController) MasterData(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	pageParams, _ := strconv.Atoi(ctx.Query("pageParams"))
	return ctx.JSON(tr.userSService.MasterData(search, limit, pageParams))
}

func (tr *userSController) MasterDataCount(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	return ctx.JSON(tr.userSService.MasterDataCount(search))
}

func (tr *userSController) UpdateUserS(ctx *fiber.Ctx) error {
	var body entity.UserS
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	body.UpdatedBy = &details.Name
	err = tr.userSService.Update(body)
	if err != nil {
		return ctx.JSON(map[string]interface{}{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil Update"})
}
func (tr *userSController) CreateUserS(ctx *fiber.Ctx) error {
	var body entity.UserS
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	body.CreatedBy = &details.Name
	err = tr.userSService.CreateUserS(body)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil create"})
}
