package token

import (
	"fmt"
	"wkm/entity"
	"wkm/middleware"
	"wkm/request"
	"wkm/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	repo := NewTokenRepository(db)
	service := NewTokenService(repo)
	handler := NewTokenController(service)

	routes := app.Group("/tokens")
	routes.Post("/signin/by-token", handler.SignInUserByToken)
	routes.Get("/detail/:id", handler.DetailToken)
	routes.Post("/create-token", middleware.DeserializeUser, handler.CreateToken)
}

type TokenController interface {
	CreateToken(ctx *fiber.Ctx) error
	DetailToken(ctx *fiber.Ctx) error
	SignInUserByToken(c *fiber.Ctx) error
}

type tokenController struct {
	tokenService TokenService
}

func NewTokenController(aS TokenService) TokenController {
	return &tokenController{
		tokenService: aS,
	}
}

func (tr *tokenController) DetailToken(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	token, err := tr.tokenService.DetailToken(id)
	if err !=nil {
		return ctx.JSON(err)
	}
	return ctx.JSON(token)
}

func (aC *tokenController) SignInUserByToken(c *fiber.Ctx) error {
	var payload request.SigninTokenRequest
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	fmt.Println("ini payload ", payload)
	errors := utils.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})
	}
	token, err := aC.tokenService.DetailToken(payload.Token)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": err})
	}
	if token.Used {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "Token already expired"})
	}
	user, err := aC.tokenService.ActivateUser(token)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": err})
	}
	if token.NoHp != payload.Username {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "token yang kamu gunakan bukan punya kamu"})
	}

	accessTokenDetails, err := utils.CreateToken(user.ID, 5, "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlDWEFJQkFBS0JnUUNOQkQwaDRPRGVTd3E0WnBIMmpBVVd5RFV1V01sb1ZlT3RSQm8yOW9TcVFJazVKSlJaCjR6SkZHZ2k0a1d4amJZemY2ZEMyRWJIdFIrbTYyM3RsV05iS3VJcklYQnZNTnVtS0RMeG9qN3VOVUVWclFVRVkKVmE5U25OUzFsM1U4S1FRbTFYaWFMamVIYWs3QzhYZi9wUEFNalhzcHdQNE1lNktpVE1RcU5ESE4zUUlEQVFBQgpBb0dBU2EwNlIzWVg1dXlzT0RZVzh3cXJLZ0VHa0NXQmJZcmFmcytESnM1YitCdnAxam8vYkV0aEcydUR2UEwxCi8yamdYcWpxREFab3dRRitvOHRDeUd2SEpLMWRURjRpcVRNcmQ3U1dlRk5WYkx3QkpLOXFiOW45NFBTZ1BuV1gKY2JFVXN4VEhLUHRwMlpOWURXcUsyWlFaVlpKRkJzTXZndVNyYlRjcFJia0drSjBDUVFEeEdiNVJRc1hNc3Z4YQpWOFFTUzZpT2J1M1BCVDRqazRJMnB6TzlMR2RwNDNVN0ZPd2QxdjVVbTRvdFpPaFVpWC9BdTRGTDdkS1A0RFdOCm5kRnBSR1lYQWtFQWxic2x3UlBtT3UxZ0VWOUtRakloWVZjMDB0ZlloSStERXdJMEJGVWtnRWNzZGFCRFhSM0YKYXdJL3o4Mlc0bklRN3ZpaXROSDZIYWVacXdid2RRR1lLd0pCQU00T1RXelAzNU5LS1lqZzE2ODNRRkN6RjhYRgoya3kzaGlORmxWK0pjcnk1N0hoWk1rOXlicDFLN2JaTU5wQUJqOURkcit4L3ptU3VuN1p2K2dpNHIzTUNRRVduCnduQ0g2VnNRZ3RpU0UrR25vSS9BR2ZyY0h3WE1IWllDT0dDcm0wZHgxT1VEb1ZMNFBwY0JmTjRYTGxJNTdsYTkKcERPcVcwamdaMFNBL2V2d3lmRUNRRW5WVnRmbThGdDZoSWFPOXNwY2VSRlhUeURBL2V1d1hhQVhubXk4RWZXZQpSaVZtSFhDVzlaSm9DbVpjbU5zY1o1Vm8zYzRMYzl2cTQyMXRKaXVRRGUwPQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQ==")
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	refreshTokenDetails, err := utils.CreateToken(user.ID, 6, "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlDWFFJQkFBS0JnUURSM0cyZkJYWDFiOGVRWXZKU2NwNSsrYzRRMThYTWRySVNFZmtSNVd5dzJUOS94eUQ5CjIxUUV2dmp2Y1BRdmlGd1hLb3VwYnFNdVFPRDFhUm9GdkJRN1J2MndUWjh5REE2YmYzcEo3NXJSYTA1UUpDaEYKL1pkR2FIZ2JYK0FuYXlzdUovNHR1OTdEcGFFMmpWVE55K05iMm1kb2ZSZ0RxbHMxbG1kY1BBY1B2UUlEQVFBQgpBb0dBT3Erd3ZCS1QzUkhvTmRsbHVHYXpLZ0VEZmpSSTdSZVlUbk5XT29uMDdqT2lqVUlMV05JMzJhZnFCMU9JCkJhN1ZTZWtzNnpHMFVsLzBTTXllYVZJaU9idktEdVdnb0MrN2ZOSUhWQUdWbkpwTXl5WXhyMXE4MC9vdWNiMlUKL1AybUs5UnNORmJORm8xOTlra1VoeG1rdjV5RUdJS0RsYmJjV3lYd0xHN1NXU0VDUVFEcERXRms5ZUVnTGFHQgpXWElQMG0yWHJQeFdRdFZ1WTZCRmRVSy9LS2FPL2NQMjVxV3JZL0R1MDhsaUdvY2NkSjBwZmtwU2tCb3hTUTVOClp4T0k5bldsQWtFQTVvWjFxV2VuS0YzbkdhRlBpdzR3RDVBMXNENVhYa3V0djl0cUVTcXQ1a0tuZG81OWdROUMKMFJ5aGQxZ3FoY0M0NC9TWmZEUURwTU5tTGxoRUljRUdPUUpBQmJZRU92c2poeXhYRnRwZ1J5NzY3SXFhckdwNgozSGVvaDhzMTFZVmpmNEdNZWRKeElPQVVHV1lyT3pJM09XVktMS2dobmlCVjQvdE1WRzFBTjAwQzJRSkJBSVJ5CmdLdmlhQUlqWWFJeU1sZDh3VlJQME9rQUNJYWZDS2NRMDdJbFNXRGdyd0xJLzRibFU4aDlvSy9ITWpkQzhYZlgKazAvdk9xQ3h1OFdvNVF4WHNORUNRUURHVFJtSElsS2g1QzJ6dUhlL1R6ajFjZjNKUkRsbVpQcXhNN1NuNTJpYwoxS0p3aHFrd1IxR0ZKZHpUVitpZEx1RlN6QUIvRHNjejFhTTRJUlBYSXVmNgotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQ==")
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    *accessTokenDetails.Token,
		Path:     "/",
		MaxAge:   15 * 60,
		Secure:   false,
		HTTPOnly: true,
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    *refreshTokenDetails.Token,
		Path:     "/",
		MaxAge:   60 * 60,
		Secure:   false,
		HTTPOnly: true,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "access_token": accessTokenDetails.Token, "name": user.Name, "refresh_token": refreshTokenDetails.Token, "is_admin": user.IsAdmin, "email": user.Email})
}

func (tr *tokenController) CreateToken(ctx *fiber.Ctx) error {
	var body entity.Token
	err := ctx.BodyParser(&body)
	if err != nil {
		fmt.Println("error body parser ", err)
	}
	err = tr.tokenService.CreateToken(body)
	if err != nil {
		return ctx.JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(map[string]string{"message": "Berhasil create"})
}
