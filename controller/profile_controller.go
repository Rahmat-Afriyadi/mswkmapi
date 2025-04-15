package controller

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
	"wkm/entity"
	"wkm/service"

	"github.com/gofiber/fiber/v2"
)

type ProfileController interface {
	Me(ctx *fiber.Ctx) error
	UploadImageProfile(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
}

type profileController struct {
	pS service.ProfileService
}

func NewProfileController(pS service.ProfileService) ProfileController {
	return &profileController{
		pS: pS,
	}
}

func (c *profileController) Me(ctx *fiber.Ctx) error {
	user := ctx.Locals("user")
	details, _ := user.(entity.User)
	data := c.pS.Me(details.ID)
	return ctx.JSON(map[string]interface{}{"data": data, "message": "Berhasil"})
}

func compressAndSaveImage(file multipart.File, outputPath string, quality int) (string,error) {
    buf, err := io.ReadAll(file)
    if err != nil {
        return "",err
    }

    img, format, err := image.Decode(bytes.NewReader(buf))
    if err != nil {
        return "",fmt.Errorf("decode failed: %w", err)
    }
    fmt.Println("Image format:", format)
	
	destination := fmt.Sprintf("%s.%s",outputPath, format)

    os.MkdirAll(filepath.Dir(fmt.Sprintf(".%s",destination)), os.ModePerm)
    out, err := os.Create(fmt.Sprintf(".%s",destination))
    if err != nil {
        return "",err
    }
    defer out.Close()

    // Cek jenis gambar (JPEG, PNG, dll)
    switch format {
    case "jpeg", "jpg":
        return destination,jpeg.Encode(out, img, &jpeg.Options{Quality: quality})
    case "png":
        return destination,png.Encode(out, img)
    default:
        return destination,fmt.Errorf("unsupported image format: %s", format)
    }
}

func (c *profileController) UploadImageProfile(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		fmt.Println("error file ", err)
		return err
	}
	
	fileOpen, _ := file.Open()
	filename, err := compressAndSaveImage(fileOpen, fmt.Sprintf("/uploads/%d", time.Now().Unix()), 95)
	if err != nil {
		fmt.Println("ini error yaa ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	fmt.Println("filename ", filename)
	return ctx.JSON(map[string]interface{}{"data": filename, "message": "Berhasil"})
}
func (c *profileController) Update(ctx *fiber.Ctx) error {
	var body entity.Profile
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	parseTimezone := body.TglLahir.Add(7 * time.Hour)
	body.TglLahir = &parseTimezone
	profile := c.pS.Update(body)
	if profile.Id == "" {
		return ctx.Status(400).JSON(map[string]interface{}{"data": "Data Gagal di update", "message": "Gagal"})
	}
	return ctx.JSON(map[string]interface{}{"data": "Data Berhasil di update", "message": "Berhasil"})
}


// destination := fmt.Sprintf("./uploads/%s", file.Filename)
	// if err := ctx.SaveFile(file, destination); err != nil {
	// 	return err
	// }