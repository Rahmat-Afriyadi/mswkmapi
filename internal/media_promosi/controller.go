package mediaPromosi

import (
	"fmt"
	"strconv"
	"wkm/entity"
	"wkm/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	repo := NewMediaPromosiRepository(db)
	service := NewMediaPromosiService(repo)
	handler := NewMediaPromosiController(service)

	routes := app.Group("/media-promosis")
	routes.Get("/master-data", middleware.DeserializeUser, handler.MasterData)
	routes.Get("/master-data/count", middleware.DeserializeUser, handler.MasterDataCount)
	routes.Get("/detail/:id", handler.DetailMediaPromosi)
	routes.Get("/detail/free/:id", handler.DetailMediaPromosi)
	routes.Get("/master-data/all", middleware.DeserializeUser, handler.MasterDataAll)

	
	routes.Delete("/delete/:id", middleware.DeserializeUserAdmin, handler.DeleteMediaPromosi)
	routes.Post("/create-mediaPromosi", middleware.DeserializeUserAdmin, handler.CreateMediaPromosi)
	routes.Post("/update-mediaPromosi", middleware.DeserializeUserAdmin, handler.UpdateMediaPromosi)
}

type MediaPromosiController interface {
	CreateMediaPromosi(ctx *fiber.Ctx) error
	UpdateMediaPromosi(ctx *fiber.Ctx) error
	MasterDataAll(ctx *fiber.Ctx) error
	MasterData(ctx *fiber.Ctx) error
	MasterDataCount(ctx *fiber.Ctx) error
	DetailMediaPromosi(ctx *fiber.Ctx) error
	DeleteMediaPromosi(ctx *fiber.Ctx) error
}

type mediaPromosiController struct {
	mediaPromosiService MediaPromosiService
}

func NewMediaPromosiController(aS MediaPromosiService) MediaPromosiController {
	return &mediaPromosiController{
		mediaPromosiService: aS,
	}
}

func (tr *mediaPromosiController) MasterDataAll(ctx *fiber.Ctx) error {
	return ctx.JSON(map[string]interface{}{"data": tr.mediaPromosiService.MasterDataAll(), "status": "success"})
}

func (tr *mediaPromosiController) DeleteMediaPromosi(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user := ctx.Locals("user")
	details, _ := user.(entity.UserS)
	err := tr.mediaPromosiService.Delete(id, details.Name)
	if err != nil {
		return ctx.JSON(map[string]interface{}{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil Update"})
}

func (tr *mediaPromosiController) DetailMediaPromosi(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	return ctx.JSON(tr.mediaPromosiService.DetailMediaPromosi(id))
}

func (tr *mediaPromosiController) MasterData(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	pageParams, _ := strconv.Atoi(ctx.Query("pageParams"))
	return ctx.JSON(tr.mediaPromosiService.MasterData(search, limit, pageParams))
}

func (tr *mediaPromosiController) MasterDataCount(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	return ctx.JSON(tr.mediaPromosiService.MasterDataCount(search))
}

func (tr *mediaPromosiController) UpdateMediaPromosi(ctx *fiber.Ctx) error {
	var body entity.MediaPromosi
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	err = tr.mediaPromosiService.Update(body)
	if err != nil {
		return ctx.JSON(map[string]interface{}{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil Update"})
}
func (tr *mediaPromosiController) CreateMediaPromosi(ctx *fiber.Ctx) error {
	var body entity.MediaPromosi
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	err = tr.mediaPromosiService.CreateMediaPromosi(body)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil create"})
}
