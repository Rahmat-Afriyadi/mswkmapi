package merchant

import (
	"fmt"
	"strconv"
	"wkm/entity"
	"wkm/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	repo := NewMerchantRepository(db)
	service := NewMerchantService(repo)
	handler := NewMerchantController(service)

	app.Use(func(ctx *fiber.Ctx) error {
		ctx.Set("Access-Control-Allow-Origin", "*")
		ctx.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		return ctx.Next()
	})

	routes := app.Group("/merchants")
	routes.Get("/master-data/filter", handler.MasterData)
	routes.Get("/master-data", middleware.DeserializeUser, handler.MasterData)
	routes.Get("/master-data/search", handler.MasterDataSearch)
	routes.Get("/master-data/count", middleware.DeserializeUser, handler.MasterDataCount)
	routes.Get("/master-data/all", middleware.DeserializeUser, handler.MasterDataAll)
	routes.Get("/detail/:id", middleware.DeserializeUser, handler.DetailMerchant)
	routes.Get("/detail/free/:id", handler.DetailMerchant)
	routes.Post("/create-merchant", middleware.DeserializeUser, handler.CreateMerchant)
	routes.Post("/update-merchant", middleware.DeserializeUser, handler.UpdateMerchant)
}

type MerchantController interface {
	CreateMerchant(ctx *fiber.Ctx) error
	UpdateMerchant(ctx *fiber.Ctx) error
	MasterData(ctx *fiber.Ctx) error
	MasterDataSearch(ctx *fiber.Ctx) error
	MasterDataCount(ctx *fiber.Ctx) error
	MasterDataAll(ctx *fiber.Ctx) error
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
	lokasi := ctx.Query("lokasi")
	return ctx.JSON(tr.mstMtrService.DetailMerchant(id, lokasi))
}

func (tr *mstMtrController) MasterDataAll(ctx *fiber.Ctx) error {
	return ctx.JSON(map[string]interface{}{"data": tr.mstMtrService.MasterDataAll(), "status": "success"})
}
func (tr *mstMtrController) MasterData(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	kategori := ctx.Query("kategori")
	lokasi := ctx.Query("lokasi")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	pageParams, _ := strconv.Atoi(ctx.Query("pageParams"))
	return ctx.JSON(tr.mstMtrService.MasterData(search, kategori, lokasi, limit, pageParams))
}

func (tr *mstMtrController) MasterDataSearch(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	return ctx.JSON(tr.mstMtrService.MasterDataSearch(search))
}

func (tr *mstMtrController) MasterDataCount(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	kategori := ctx.Query("kategori")
	lokasi := ctx.Query("lokasi")
	return ctx.JSON(tr.mstMtrService.MasterDataCount(search, kategori, lokasi))
}

func (tr *mstMtrController) UpdateMerchant(ctx *fiber.Ctx) error {
	var body Merchant
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
func (tr *mstMtrController) CreateMerchant(ctx *fiber.Ctx) error {
	var body Merchant
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	body.CreatedBy = details.Name
	err = tr.mstMtrService.CreateMerchant(body)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil create"})
}
