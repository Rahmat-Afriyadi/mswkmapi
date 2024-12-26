package outlet

import (
	"fmt"
	"strconv"
	"wkm/entity"
	"wkm/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	repo := NewOutletRepository(db)
	service := NewOutletService(repo)
	handler := NewOutletController(service)

	routes := app.Group("/outlets")
	routes.Get("/master-data", middleware.DeserializeUser, handler.MasterData)
	routes.Get("/master-data/count", middleware.DeserializeUser, handler.MasterDataCount)
	routes.Get("/detail/:id", handler.DetailOutlet)
	routes.Get("/detail/free/:id", handler.DetailOutlet)
	routes.Post("/create-outlet", middleware.DeserializeUser, handler.CreateOutlet)
	routes.Post("/update-outlet", middleware.DeserializeUser, handler.UpdateOutlet)
}

type OutletController interface {
	CreateOutlet(ctx *fiber.Ctx) error
	UpdateOutlet(ctx *fiber.Ctx) error
	MasterData(ctx *fiber.Ctx) error
	MasterDataCount(ctx *fiber.Ctx) error
	DetailOutlet(ctx *fiber.Ctx) error
}

type mstMtrController struct {
	mstMtrService OutletService
}

func NewOutletController(aS OutletService) OutletController {
	return &mstMtrController{
		mstMtrService: aS,
	}
}

func (tr *mstMtrController) DetailOutlet(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	return ctx.JSON(tr.mstMtrService.DetailOutlet(id))
}

func (tr *mstMtrController) MasterData(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	pageParams, _ := strconv.Atoi(ctx.Query("pageParams"))
	return ctx.JSON(tr.mstMtrService.MasterData(search, limit, pageParams))
}

func (tr *mstMtrController) MasterDataCount(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	return ctx.JSON(tr.mstMtrService.MasterDataCount(search))
}

func (tr *mstMtrController) UpdateOutlet(ctx *fiber.Ctx) error {
	var body Outlet
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	body.UpdatedBy = details.Name
	err = tr.mstMtrService.Update(body)
	if err != nil {
		return ctx.JSON(map[string]interface{}{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil Update"})
}
func (tr *mstMtrController) CreateOutlet(ctx *fiber.Ctx) error {
	var body Outlet
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	body.CreatedBy = details.Name
	err = tr.mstMtrService.CreateOutlet(body)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil create"})
}
