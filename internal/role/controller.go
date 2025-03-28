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
	repo := NewRoleRepository(db)
	service := NewRoleService(repo)
	handler := NewRoleController(service)

	routes := app.Group("/userses")
	routes.Get("/master-data", middleware.DeserializeUser, handler.MasterData)
	routes.Get("/master-data/count", middleware.DeserializeUser, handler.MasterDataCount)
	routes.Get("/detail/:id", handler.DetailRole)
	routes.Get("/detail/free/:id", handler.DetailRole)
	routes.Post("/create-user", middleware.DeserializeUser, handler.CreateRole)
	routes.Post("/update-user", middleware.DeserializeUser, handler.UpdateRole)
}

type RoleController interface {
	CreateRole(ctx *fiber.Ctx) error
	UpdateRole(ctx *fiber.Ctx) error
	MasterData(ctx *fiber.Ctx) error
	MasterDataCount(ctx *fiber.Ctx) error
	DetailRole(ctx *fiber.Ctx) error
}

type roleController struct {
	roleService RoleService
}

func NewRoleController(aS RoleService) RoleController {
	return &roleController{
		roleService: aS,
	}
}

func (tr *roleController) DetailRole(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	return ctx.JSON(tr.roleService.DetailRole(id))
}

func (tr *roleController) MasterData(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	pageParams, _ := strconv.Atoi(ctx.Query("pageParams"))
	return ctx.JSON(tr.roleService.MasterData(search, limit, pageParams))
}

func (tr *roleController) MasterDataCount(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	return ctx.JSON(tr.roleService.MasterDataCount(search))
}

func (tr *roleController) UpdateRole(ctx *fiber.Ctx) error {
	var body entity.Role
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	err = tr.roleService.Update(body)
	if err != nil {
		return ctx.JSON(map[string]interface{}{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil Update"})
}
func (tr *roleController) CreateRole(ctx *fiber.Ctx) error {
	var body entity.Role
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	err = tr.roleService.CreateRole(body)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil create"})
}
