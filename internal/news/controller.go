package news

import (
	"fmt"
	"strconv"
	"wkm/entity"
	"wkm/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	repo := NewNewsRepository(db)
	service := NewNewsService(repo)
	handler := NewNewsController(service)

	routes := app.Group("/news")
	routes.Get("/master-data/filter", handler.MasterData)
	routes.Get("/master-data/search", handler.MasterDataSearch)
	routes.Get("/master-data", middleware.DeserializeUser, handler.MasterData)
	routes.Get("/master-data/count", middleware.DeserializeUser, handler.MasterDataCount)
	routes.Get("/detail/:id", handler.DetailNews)
	routes.Get("/detail/free/:id", handler.DetailNews)
	routes.Get("/master-data/all", middleware.DeserializeUser, handler.MasterDataAll)

	
	routes.Delete("/delete/:id", middleware.DeserializeUserAdmin, handler.DeleteNews)
	routes.Post("/create-news", middleware.DeserializeUserAdmin, handler.CreateNews)
	routes.Post("/update-news", middleware.DeserializeUserAdmin, handler.UpdateNews)
}

type NewsController interface {
	CreateNews(ctx *fiber.Ctx) error
	UpdateNews(ctx *fiber.Ctx) error
	MasterDataAll(ctx *fiber.Ctx) error
	MasterDataSearch(ctx *fiber.Ctx) error
	MasterData(ctx *fiber.Ctx) error
	MasterDataCount(ctx *fiber.Ctx) error
	DetailNews(ctx *fiber.Ctx) error
	DeleteNews(ctx *fiber.Ctx) error
}

type newsController struct {
	newsService NewsService
}

func NewNewsController(aS NewsService) NewsController {
	return &newsController{
		newsService: aS,
	}
}

func (tr *newsController) MasterDataAll(ctx *fiber.Ctx) error {
	return ctx.JSON(map[string]interface{}{"data": tr.newsService.MasterDataAll(), "status": "success"})
}
func (tr *newsController) MasterDataSearch(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	return ctx.JSON(tr.newsService.MasterDataSearch(search))
}

func (tr *newsController) DeleteNews(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user := ctx.Locals("user")
	details, _ := user.(entity.UserS)
	err := tr.newsService.Delete(id, details.Name)
	if err != nil {
		return ctx.JSON(map[string]interface{}{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil Update"})
}

func (tr *newsController) DetailNews(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	return ctx.JSON(tr.newsService.DetailNews(id))
}

func (tr *newsController) MasterData(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	pageParams, _ := strconv.Atoi(ctx.Query("pageParams"))
	return ctx.JSON(tr.newsService.MasterData(search, limit, pageParams))
}

func (tr *newsController) MasterDataCount(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	return ctx.JSON(tr.newsService.MasterDataCount(search))
}

func (tr *newsController) UpdateNews(ctx *fiber.Ctx) error {
	var body entity.News
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	err = tr.newsService.Update(body)
	if err != nil {
		return ctx.JSON(map[string]interface{}{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil Update"})
}
func (tr *newsController) CreateNews(ctx *fiber.Ctx) error {
	var body entity.News
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	err = tr.newsService.CreateNews(body)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil create"})
}
