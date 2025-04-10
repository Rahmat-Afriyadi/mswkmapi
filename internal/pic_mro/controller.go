package picMro

import (
	"fmt"
	"strconv"
	"wkm/entity"
	"wkm/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	repo := NewPicMroRepository(db)
	service := NewPicMroService(repo)
	handler := NewPicMroController(service)

	routes := app.Group("/pic-mros")
	routes.Get("/master-data", middleware.DeserializeUser, handler.MasterData)
	routes.Get("/master-data/count", middleware.DeserializeUser, handler.MasterDataCount)
	routes.Get("/detail/:id", handler.DetailPicMro)
	routes.Get("/detail/free/:id", handler.DetailPicMro)
	routes.Get("/master-data/all", middleware.DeserializeUser, handler.MasterDataAll)

	
	routes.Delete("/delete/:id", middleware.DeserializeUserAdmin, handler.DeletePicMro)
	routes.Post("/create-picMro", middleware.DeserializeUserAdmin, handler.CreatePicMro)
	routes.Post("/update-picMro", middleware.DeserializeUserAdmin, handler.UpdatePicMro)
}

type PicMroController interface {
	CreatePicMro(ctx *fiber.Ctx) error
	UpdatePicMro(ctx *fiber.Ctx) error
	MasterDataAll(ctx *fiber.Ctx) error
	MasterData(ctx *fiber.Ctx) error
	MasterDataCount(ctx *fiber.Ctx) error
	DetailPicMro(ctx *fiber.Ctx) error
	DeletePicMro(ctx *fiber.Ctx) error
}

type picMroController struct {
	picMroService PicMroService
}

func NewPicMroController(aS PicMroService) PicMroController {
	return &picMroController{
		picMroService: aS,
	}
}

func (tr *picMroController) MasterDataAll(ctx *fiber.Ctx) error {
	return ctx.JSON(map[string]interface{}{"data": tr.picMroService.MasterDataAll(), "status": "success"})
}

func (tr *picMroController) DeletePicMro(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user := ctx.Locals("user")
	details, _ := user.(entity.UserS)
	err := tr.picMroService.Delete(id, details.Name)
	if err != nil {
		return ctx.JSON(map[string]interface{}{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil Update"})
}

func (tr *picMroController) DetailPicMro(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	return ctx.JSON(tr.picMroService.DetailPicMro(id))
}

func (tr *picMroController) MasterData(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	pageParams, _ := strconv.Atoi(ctx.Query("pageParams"))
	return ctx.JSON(tr.picMroService.MasterData(search, limit, pageParams))
}

func (tr *picMroController) MasterDataCount(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	return ctx.JSON(tr.picMroService.MasterDataCount(search))
}

func (tr *picMroController) UpdatePicMro(ctx *fiber.Ctx) error {
	var body entity.PicMro
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	err = tr.picMroService.Update(body)
	if err != nil {
		return ctx.JSON(map[string]interface{}{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil Update"})
}
func (tr *picMroController) CreatePicMro(ctx *fiber.Ctx) error {
	var body entity.PicMro
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	err = tr.picMroService.CreatePicMro(body)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil create"})
}
