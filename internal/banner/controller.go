package banner

import (
	"fmt"
	"strconv"
	"wkm/entity"
	"wkm/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	repo := NewBannerRepository(db)
	service := NewBannerService(repo)
	handler := NewBannerController(service)

	routes := app.Group("/banners")
	routes.Get("/master-data/pin", handler.MasterDataPin)

	routes.Get("/master-data/filter", handler.MasterData)
	routes.Get("/master-data/search", handler.MasterDataSearch)
	routes.Get("/master-data", middleware.DeserializeUser, handler.MasterData)
	routes.Get("/master-data/count", middleware.DeserializeUser, handler.MasterDataCount)
	routes.Get("/detail/:id", handler.DetailBanner)
	routes.Get("/detail/free/:id", handler.DetailBanner)
	routes.Get("/master-data/all", middleware.DeserializeUser, handler.MasterDataAll)

	
	routes.Delete("/delete/:id", middleware.DeserializeUserAdmin, handler.DeleteBanner)
	routes.Post("/create-banner", middleware.DeserializeUserAdmin, handler.CreateBanner)
	routes.Post("/update-banner", middleware.DeserializeUserAdmin, handler.UpdateBanner)
}

type BannerController interface {
	MasterDataPin(ctx *fiber.Ctx) error
	CreateBanner(ctx *fiber.Ctx) error
	UpdateBanner(ctx *fiber.Ctx) error
	MasterDataAll(ctx *fiber.Ctx) error
	MasterDataSearch(ctx *fiber.Ctx) error
	MasterData(ctx *fiber.Ctx) error
	MasterDataCount(ctx *fiber.Ctx) error
	DetailBanner(ctx *fiber.Ctx) error
	DeleteBanner(ctx *fiber.Ctx) error
}

type bannerController struct {
	bannerService BannerService
}

func NewBannerController(aS BannerService) BannerController {
	return &bannerController{
		bannerService: aS,
	}
}

func (tr *bannerController) MasterDataPin(ctx *fiber.Ctx) error {
	data, err := tr.bannerService.MasterDataPin()
	if err != nil {
		return ctx.Status(500).JSON(map[string]interface{}{"message": err.Error()}) 
	}
	return ctx.JSON(data)
}

func (tr *bannerController) MasterDataAll(ctx *fiber.Ctx) error {
	return ctx.JSON(map[string]interface{}{"data": tr.bannerService.MasterDataAll(), "status": "success"})
}
func (tr *bannerController) MasterDataSearch(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	return ctx.JSON(tr.bannerService.MasterDataSearch(search))
}

func (tr *bannerController) DeleteBanner(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user := ctx.Locals("user")
	details, _ := user.(entity.UserS)
	err := tr.bannerService.Delete(id, details.Name)
	if err != nil {
		return ctx.JSON(map[string]interface{}{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil Update"})
}

func (tr *bannerController) DetailBanner(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	return ctx.JSON(tr.bannerService.DetailBanner(id))
}

func (tr *bannerController) MasterData(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	pageParams, _ := strconv.Atoi(ctx.Query("pageParams"))
	return ctx.JSON(tr.bannerService.MasterData(search, limit, pageParams))
}

func (tr *bannerController) MasterDataCount(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	return ctx.JSON(tr.bannerService.MasterDataCount(search))
}

func (tr *bannerController) UpdateBanner(ctx *fiber.Ctx) error {
	var body entity.Banner
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	err = tr.bannerService.Update(body)
	if err != nil {
		return ctx.JSON(map[string]interface{}{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil Update"})
}
func (tr *bannerController) CreateBanner(ctx *fiber.Ctx) error {
	var body entity.Banner
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	err = tr.bannerService.CreateBanner(body)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil create"})
}
