package kategori

import (
	"fmt"
	"strconv"
	"wkm/entity"
	"wkm/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	repo := NewKategoriRepository(db)
	service := NewKategoriService(repo)
	handler := NewKategoriController(service)

	routes := app.Group("/kategoris")
	routes.Get("/master-data", middleware.DeserializeUser, handler.MasterData)
	routes.Get("/master-data/count", middleware.DeserializeUser, handler.MasterDataCount)
	routes.Get("/detail/:id", handler.DetailKategori)
	routes.Get("/detail/free/:id", handler.DetailKategori)
	routes.Get("/master-data/all", middleware.DeserializeUser, handler.MasterDataAll)

	
	routes.Delete("/delete/:id", middleware.DeserializeUserAdmin, handler.DeleteKategori)
	routes.Post("/create-kategori", middleware.DeserializeUserAdmin, handler.CreateKategori)
	routes.Post("/update-kategori", middleware.DeserializeUserAdmin, handler.UpdateKategori)
}

type KategoriController interface {
	CreateKategori(ctx *fiber.Ctx) error
	UpdateKategori(ctx *fiber.Ctx) error
	MasterDataAll(ctx *fiber.Ctx) error
	MasterData(ctx *fiber.Ctx) error
	MasterDataCount(ctx *fiber.Ctx) error
	DetailKategori(ctx *fiber.Ctx) error
	DeleteKategori(ctx *fiber.Ctx) error
}

type kategoriController struct {
	kategoriService KategoriService
}

func NewKategoriController(aS KategoriService) KategoriController {
	return &kategoriController{
		kategoriService: aS,
	}
}

func (tr *kategoriController) MasterDataAll(ctx *fiber.Ctx) error {
	return ctx.JSON(map[string]interface{}{"data": tr.kategoriService.MasterDataAll(), "status": "success"})
}

func (tr *kategoriController) DeleteKategori(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user := ctx.Locals("user")
	details, _ := user.(entity.UserS)
	err := tr.kategoriService.Delete(id, details.Name)
	if err != nil {
		return ctx.JSON(map[string]interface{}{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil Update"})
}

func (tr *kategoriController) DetailKategori(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	return ctx.JSON(tr.kategoriService.DetailKategori(id))
}

func (tr *kategoriController) MasterData(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	pageParams, _ := strconv.Atoi(ctx.Query("pageParams"))
	return ctx.JSON(tr.kategoriService.MasterData(search, limit, pageParams))
}

func (tr *kategoriController) MasterDataCount(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	return ctx.JSON(tr.kategoriService.MasterDataCount(search))
}

func (tr *kategoriController) UpdateKategori(ctx *fiber.Ctx) error {
	var body entity.Kategori
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	err = tr.kategoriService.Update(body)
	if err != nil {
		return ctx.JSON(map[string]interface{}{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil Update"})
}
func (tr *kategoriController) CreateKategori(ctx *fiber.Ctx) error {
	var body entity.Kategori
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	err = tr.kategoriService.CreateKategori(body)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil create"})
}
