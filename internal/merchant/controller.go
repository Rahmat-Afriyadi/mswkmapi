package merchant

import (
	"fmt"
	"strconv"
	"wkm/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	repo := NewMerchantRepository(db)
	service := NewMerchantService(repo)
	handler := NewMerchantController(service)

	userRoutes := app.Group("/merchants")
	userRoutes.Get("/master-data", middleware.DeserializeUser, handler.MasterData)
	userRoutes.Get("/master-data/count", middleware.DeserializeUser, handler.MasterDataCount)
	userRoutes.Post("/create-merchant", middleware.DeserializeUser, handler.CreateMerchant)
	userRoutes.Post("/update-merchant", middleware.DeserializeUser, handler.UpdateMerchant)
}

type MerchantController interface {
	CreateMerchant(ctx *fiber.Ctx) error
	UpdateMerchant(ctx *fiber.Ctx) error
	MasterData(ctx *fiber.Ctx) error
	MasterDataCount(ctx *fiber.Ctx) error
	DetailMerchant(ctx *fiber.Ctx) error
}

type mstMtrController struct {
	mstMtrService MerchantService
}

func NewMerchantController(aS MerchantService) MerchantController {
	return &mstMtrController{
		mstMtrService: aS,
	}
}

func (tr *mstMtrController) DetailMerchant(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	return ctx.JSON(tr.mstMtrService.DetailMerchant(id))
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

func (tr *mstMtrController) UpdateMerchant(ctx *fiber.Ctx) error {
	var body Merchant
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	err = tr.mstMtrService.Update(body)
	if err != nil {
		return ctx.JSON(map[string]interface{}{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil Update"})
}
func (tr *mstMtrController) CreateMerchant(ctx *fiber.Ctx) error {
	var body Merchant
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	err = tr.mstMtrService.CreateMerchant(body)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil create"})
}
