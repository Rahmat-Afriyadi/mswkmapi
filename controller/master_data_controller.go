package controller

import (
	"wkm/service"

	"github.com/gofiber/fiber/v2"
)

type MasterDataController interface {
	KategoriMerchantAll(ctx *fiber.Ctx) error
	NewsKategoriAll(ctx *fiber.Ctx) error
	MediaPromosiAll(ctx *fiber.Ctx) error
	PicMroAll(ctx *fiber.Ctx) error
	KodeposAll(ctx *fiber.Ctx) error
}

type masterDataController struct {
	mS service.MasterDataService
}

func NewMasterDataController(mS service.MasterDataService) MasterDataController {
	return &masterDataController{
		mS: mS,
	}
}

func (c *masterDataController) KategoriMerchantAll(ctx *fiber.Ctx) error {
	data, _ := c.mS.KategoriMerchantAll()
	return ctx.JSON(map[string]interface{}{"data": data, "message": "Berhasil"})
}
func (c *masterDataController) NewsKategoriAll(ctx *fiber.Ctx) error {
	data, _ := c.mS.NewsKategoriAll()
	return ctx.JSON(map[string]interface{}{"data": data, "message": "Berhasil"})
}
func (c *masterDataController) KodeposAll(ctx *fiber.Ctx) error {
	data, _ := c.mS.KodeposAll()
	return ctx.JSON(map[string]interface{}{"data": data, "message": "Berhasil"})
}
func (c *masterDataController) MediaPromosiAll(ctx *fiber.Ctx) error {
	data, _ := c.mS.MediaPromosiAll()
	return ctx.JSON(map[string]interface{}{"data": data, "message": "Berhasil"})
}
func (c *masterDataController) PicMroAll(ctx *fiber.Ctx) error {
	data, _ := c.mS.PicMroAll()
	return ctx.JSON(map[string]interface{}{"data": data, "message": "Berhasil"})
}
