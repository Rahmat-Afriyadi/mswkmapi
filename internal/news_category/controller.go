package newsKategori

import (
	"fmt"
	"strconv"
	"wkm/entity"
	"wkm/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	repo := NewNewsKategoriRepository(db)
	service := NewNewsKategoriService(repo)
	handler := NewNewsKategoriController(service)

	routes := app.Group("/news-kategoris")
	routes.Get("/master-data", middleware.DeserializeUser, handler.MasterData)
	routes.Get("/master-data/count", middleware.DeserializeUser, handler.MasterDataCount)
	routes.Get("/detail/:id", handler.DetailNewsKategori)
	routes.Get("/detail/free/:id", handler.DetailNewsKategori)
	routes.Get("/master-data/all", middleware.DeserializeUser, handler.MasterDataAll)

	
	routes.Delete("/delete/:id", middleware.DeserializeUserAdmin, handler.DeleteNewsKategori)
	routes.Post("/create-news-kategoris", middleware.DeserializeUserAdmin, handler.CreateNewsKategori)
	routes.Post("/update-news-kategoris", middleware.DeserializeUserAdmin, handler.UpdateNewsKategori)
}

type NewsKategoriController interface {
	CreateNewsKategori(ctx *fiber.Ctx) error
	UpdateNewsKategori(ctx *fiber.Ctx) error
	MasterDataAll(ctx *fiber.Ctx) error
	MasterData(ctx *fiber.Ctx) error
	MasterDataCount(ctx *fiber.Ctx) error
	DetailNewsKategori(ctx *fiber.Ctx) error
	DeleteNewsKategori(ctx *fiber.Ctx) error
}

type kategoriController struct {
	kategoriService NewsKategoriService
}

func NewNewsKategoriController(aS NewsKategoriService) NewsKategoriController {
	return &kategoriController{
		kategoriService: aS,
	}
}

func (tr *kategoriController) MasterDataAll(ctx *fiber.Ctx) error {
	return ctx.JSON(map[string]interface{}{"data": tr.kategoriService.MasterDataAll(), "status": "success"})
}

func (tr *kategoriController) DeleteNewsKategori(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user := ctx.Locals("user")
	details, _ := user.(entity.UserS)
	err := tr.kategoriService.Delete(id, details.Name)
	if err != nil {
		return ctx.JSON(map[string]interface{}{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil Update"})
}

func (tr *kategoriController) DetailNewsKategori(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	return ctx.JSON(tr.kategoriService.DetailNewsKategori(id))
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

func (tr *kategoriController) UpdateNewsKategori(ctx *fiber.Ctx) error {
	var body entity.NewsKategori
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
func (tr *kategoriController) CreateNewsKategori(ctx *fiber.Ctx) error {
	var body entity.NewsKategori
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	err = tr.kategoriService.CreateNewsKategori(body)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil create"})
}
