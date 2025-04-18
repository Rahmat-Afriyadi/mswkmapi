package controller

import (
	"fmt"
	"strconv"
	"time"
	"wkm/entity"
	"wkm/request"
	"wkm/service"
	"wkm/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type AuthController interface {
	SignInUser(c *fiber.Ctx) error
	SignInUserAdmin(c *fiber.Ctx) error
	RefreshAccessToken(c *fiber.Ctx) error
	RefreshAccessTokenAsuransi(c *fiber.Ctx) error
	LogoutUser(c *fiber.Ctx) error
	GeneratePassword(c *fiber.Ctx) error
	ResetPassword(c *fiber.Ctx) error
	ResetPasswordByOtp(c *fiber.Ctx) error
	SignUpUser(c *fiber.Ctx) error
	GenerateNewOtp(c *fiber.Ctx) error
	CheckOtp(c *fiber.Ctx) error
	CheckOtpReset(c *fiber.Ctx) error
}

type authController struct {
	aS service.AuthService
}

func NewAuthController(aS service.AuthService) AuthController {
	return &authController{
		aS,
	}
}

func (aC *authController) GeneratePassword(c *fiber.Ctx) error {
	return c.JSON(map[string]interface{}{"message": "Hallo guys"})
}

func (aC *authController) ResetPassword(c *fiber.Ctx) error {
	user := c.Locals("user")
	details, _ := user.(entity.User)
	var body request.ResetPassword
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	body.IdUser = details.ID
	response := aC.aS.ResetPassword(body)

	return c.Status(response.Status).JSON(map[string]interface{}{"message": response.Message})
}

func (aC *authController) ResetPasswordByOtp(c *fiber.Ctx) error {
	var body request.ResetPasswordOtp
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	response := aC.aS.ResetPasswordByOtp(body)

	return c.Status(response.Status).JSON(map[string]interface{}{"message": response.Message})
}

func (aC *authController) SignUpUser(c *fiber.Ctx) error {
	var body request.SignupRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	fmt.Println("ini body yaa ", body)
	user, err := aC.aS.SignUpUser(body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"message": err.Error()})
	}

	return c.JSON(map[string]interface{}{"user": user})
}

func (aC *authController) SignInUserAdmin(c *fiber.Ctx) error {
	var payload request.SigninRequest

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := utils.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})
	}

	message := "Invalid email or password"

	user, err := aC.aS.SignInUserAdmin(payload)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": err})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		fmt.Println("ini errornya ", err, user.Password, payload.Password)
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": message})
	}

	accessTokenDetails, err := utils.CreateToken(strconv.FormatUint(user.ID,10), 10, "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlDWEFJQkFBS0JnUUNOQkQwaDRPRGVTd3E0WnBIMmpBVVd5RFV1V01sb1ZlT3RSQm8yOW9TcVFJazVKSlJaCjR6SkZHZ2k0a1d4amJZemY2ZEMyRWJIdFIrbTYyM3RsV05iS3VJcklYQnZNTnVtS0RMeG9qN3VOVUVWclFVRVkKVmE5U25OUzFsM1U4S1FRbTFYaWFMamVIYWs3QzhYZi9wUEFNalhzcHdQNE1lNktpVE1RcU5ESE4zUUlEQVFBQgpBb0dBU2EwNlIzWVg1dXlzT0RZVzh3cXJLZ0VHa0NXQmJZcmFmcytESnM1YitCdnAxam8vYkV0aEcydUR2UEwxCi8yamdYcWpxREFab3dRRitvOHRDeUd2SEpLMWRURjRpcVRNcmQ3U1dlRk5WYkx3QkpLOXFiOW45NFBTZ1BuV1gKY2JFVXN4VEhLUHRwMlpOWURXcUsyWlFaVlpKRkJzTXZndVNyYlRjcFJia0drSjBDUVFEeEdiNVJRc1hNc3Z4YQpWOFFTUzZpT2J1M1BCVDRqazRJMnB6TzlMR2RwNDNVN0ZPd2QxdjVVbTRvdFpPaFVpWC9BdTRGTDdkS1A0RFdOCm5kRnBSR1lYQWtFQWxic2x3UlBtT3UxZ0VWOUtRakloWVZjMDB0ZlloSStERXdJMEJGVWtnRWNzZGFCRFhSM0YKYXdJL3o4Mlc0bklRN3ZpaXROSDZIYWVacXdid2RRR1lLd0pCQU00T1RXelAzNU5LS1lqZzE2ODNRRkN6RjhYRgoya3kzaGlORmxWK0pjcnk1N0hoWk1rOXlicDFLN2JaTU5wQUJqOURkcit4L3ptU3VuN1p2K2dpNHIzTUNRRVduCnduQ0g2VnNRZ3RpU0UrR25vSS9BR2ZyY0h3WE1IWllDT0dDcm0wZHgxT1VEb1ZMNFBwY0JmTjRYTGxJNTdsYTkKcERPcVcwamdaMFNBL2V2d3lmRUNRRW5WVnRmbThGdDZoSWFPOXNwY2VSRlhUeURBL2V1d1hhQVhubXk4RWZXZQpSaVZtSFhDVzlaSm9DbVpjbU5zY1o1Vm8zYzRMYzl2cTQyMXRKaXVRRGUwPQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQ==")
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	refreshTokenDetails, err := utils.CreateToken(strconv.FormatUint(user.ID,10), 30, "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlDWFFJQkFBS0JnUURSM0cyZkJYWDFiOGVRWXZKU2NwNSsrYzRRMThYTWRySVNFZmtSNVd5dzJUOS94eUQ5CjIxUUV2dmp2Y1BRdmlGd1hLb3VwYnFNdVFPRDFhUm9GdkJRN1J2MndUWjh5REE2YmYzcEo3NXJSYTA1UUpDaEYKL1pkR2FIZ2JYK0FuYXlzdUovNHR1OTdEcGFFMmpWVE55K05iMm1kb2ZSZ0RxbHMxbG1kY1BBY1B2UUlEQVFBQgpBb0dBT3Erd3ZCS1QzUkhvTmRsbHVHYXpLZ0VEZmpSSTdSZVlUbk5XT29uMDdqT2lqVUlMV05JMzJhZnFCMU9JCkJhN1ZTZWtzNnpHMFVsLzBTTXllYVZJaU9idktEdVdnb0MrN2ZOSUhWQUdWbkpwTXl5WXhyMXE4MC9vdWNiMlUKL1AybUs5UnNORmJORm8xOTlra1VoeG1rdjV5RUdJS0RsYmJjV3lYd0xHN1NXU0VDUVFEcERXRms5ZUVnTGFHQgpXWElQMG0yWHJQeFdRdFZ1WTZCRmRVSy9LS2FPL2NQMjVxV3JZL0R1MDhsaUdvY2NkSjBwZmtwU2tCb3hTUTVOClp4T0k5bldsQWtFQTVvWjFxV2VuS0YzbkdhRlBpdzR3RDVBMXNENVhYa3V0djl0cUVTcXQ1a0tuZG81OWdROUMKMFJ5aGQxZ3FoY0M0NC9TWmZEUURwTU5tTGxoRUljRUdPUUpBQmJZRU92c2poeXhYRnRwZ1J5NzY3SXFhckdwNgozSGVvaDhzMTFZVmpmNEdNZWRKeElPQVVHV1lyT3pJM09XVktMS2dobmlCVjQvdE1WRzFBTjAwQzJRSkJBSVJ5CmdLdmlhQUlqWWFJeU1sZDh3VlJQME9rQUNJYWZDS2NRMDdJbFNXRGdyd0xJLzRibFU4aDlvSy9ITWpkQzhYZlgKazAvdk9xQ3h1OFdvNVF4WHNORUNRUURHVFJtSElsS2g1QzJ6dUhlL1R6ajFjZjNKUkRsbVpQcXhNN1NuNTJpYwoxS0p3aHFrd1IxR0ZKZHpUVitpZEx1RlN6QUIvRHNjejFhTTRJUlBYSXVmNgotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQ==")
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "access_token": accessTokenDetails.Token, "name": user.Name, "refresh_token": refreshTokenDetails.Token, "permissions": user.Permissions, "role": user.RoleId, "role_name": user.Role.Name, "id": user.ID, "expired_at": time.Now().Add(time.Minute * 15).Format("2006-01-02 15:04:05")})
}


func (aC *authController) GenerateNewOtp(c *fiber.Ctx) error {
	var body request.OtpCheck
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	user, err := aC.aS.GenerateNewOtp(body)
	if err != nil {
		return c.JSON(map[string]interface{}{"message": err.Error()})
	}

	return c.JSON(map[string]interface{}{"user": user})
}

func (aC *authController) CheckOtp(c *fiber.Ctx) error {
	var body request.OtpCheck
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	err := aC.aS.CheckOtp(body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"message": err.Error()})
	}

	return c.JSON(map[string]interface{}{"status": true})
}

func (aC *authController) CheckOtpReset(c *fiber.Ctx) error {
	var body request.OtpCheck
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	otp, err := aC.aS.CheckOtpReset(body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"message": err.Error()})
	}
	return c.JSON(map[string]interface{}{"otp": otp, "status": true})
}

func (aC *authController) SignInUser(c *fiber.Ctx) error {
	var payload request.SigninRequest
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := utils.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})
	}

	message := "Invalid phone number or password"

	user, err := aC.aS.SignInUser(payload)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": err})
	}

	if payload.AutoLogin == "false" {
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
		if err != nil {
			fmt.Println("ini error login ",user.Password,payload.Password,  err)
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": message})
		}
	}

	if !user.Active {
		return c.JSON(fiber.Map{"status": "inactive", "message": "Akun belum aktif"})
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

	// c.Cookie(&fiber.Cookie{
	// 	Name:     "logged_in",
	// 	Value:    "true",
	// 	Path:     "/",
	// 	MaxAge:   15 * 60,
	// 	Secure:   false,
	// 	HTTPOnly: false,
	// 	Domain:   "localhost",
	// })

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "access_token": accessTokenDetails.Token, "name": user.Name, "refresh_token": refreshTokenDetails.Token, "is_admin": user.IsAdmin, "email": user.Email})
}



func (aC *authController) RefreshAccessToken(c *fiber.Ctx) error {
	message := "could not refresh access token"
	var body map[string]string

	c.BodyParser(&body)

	if body["refresh_token"] == "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": message})
	}

	tokenClaims, err := utils.ValidateToken(body["refresh_token"], "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FEUjNHMmZCWFgxYjhlUVl2SlNjcDUrK2M0UQoxOFhNZHJJU0Vma1I1V3l3MlQ5L3h5RDkyMVFFdnZqdmNQUXZpRndYS291cGJxTXVRT0QxYVJvRnZCUTdSdjJ3ClRaOHlEQTZiZjNwSjc1clJhMDVRSkNoRi9aZEdhSGdiWCtBbmF5c3VKLzR0dTk3RHBhRTJqVlROeStOYjJtZG8KZlJnRHFsczFsbWRjUEFjUHZRSURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQ==")
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	user, err := aC.aS.RefreshToken(tokenClaims.UserID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": err})
	}

	accessTokenDetails, err := utils.CreateToken(user.ID, 5, "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlDWEFJQkFBS0JnUUNOQkQwaDRPRGVTd3E0WnBIMmpBVVd5RFV1V01sb1ZlT3RSQm8yOW9TcVFJazVKSlJaCjR6SkZHZ2k0a1d4amJZemY2ZEMyRWJIdFIrbTYyM3RsV05iS3VJcklYQnZNTnVtS0RMeG9qN3VOVUVWclFVRVkKVmE5U25OUzFsM1U4S1FRbTFYaWFMamVIYWs3QzhYZi9wUEFNalhzcHdQNE1lNktpVE1RcU5ESE4zUUlEQVFBQgpBb0dBU2EwNlIzWVg1dXlzT0RZVzh3cXJLZ0VHa0NXQmJZcmFmcytESnM1YitCdnAxam8vYkV0aEcydUR2UEwxCi8yamdYcWpxREFab3dRRitvOHRDeUd2SEpLMWRURjRpcVRNcmQ3U1dlRk5WYkx3QkpLOXFiOW45NFBTZ1BuV1gKY2JFVXN4VEhLUHRwMlpOWURXcUsyWlFaVlpKRkJzTXZndVNyYlRjcFJia0drSjBDUVFEeEdiNVJRc1hNc3Z4YQpWOFFTUzZpT2J1M1BCVDRqazRJMnB6TzlMR2RwNDNVN0ZPd2QxdjVVbTRvdFpPaFVpWC9BdTRGTDdkS1A0RFdOCm5kRnBSR1lYQWtFQWxic2x3UlBtT3UxZ0VWOUtRakloWVZjMDB0ZlloSStERXdJMEJGVWtnRWNzZGFCRFhSM0YKYXdJL3o4Mlc0bklRN3ZpaXROSDZIYWVacXdid2RRR1lLd0pCQU00T1RXelAzNU5LS1lqZzE2ODNRRkN6RjhYRgoya3kzaGlORmxWK0pjcnk1N0hoWk1rOXlicDFLN2JaTU5wQUJqOURkcit4L3ptU3VuN1p2K2dpNHIzTUNRRVduCnduQ0g2VnNRZ3RpU0UrR25vSS9BR2ZyY0h3WE1IWllDT0dDcm0wZHgxT1VEb1ZMNFBwY0JmTjRYTGxJNTdsYTkKcERPcVcwamdaMFNBL2V2d3lmRUNRRW5WVnRmbThGdDZoSWFPOXNwY2VSRlhUeURBL2V1d1hhQVhubXk4RWZXZQpSaVZtSFhDVzlaSm9DbVpjbU5zY1o1Vm8zYzRMYzl2cTQyMXRKaXVRRGUwPQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQ==")
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	refreshTokenDetails, err := utils.CreateToken(user.ID, 6, "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlDWFFJQkFBS0JnUURSM0cyZkJYWDFiOGVRWXZKU2NwNSsrYzRRMThYTWRySVNFZmtSNVd5dzJUOS94eUQ5CjIxUUV2dmp2Y1BRdmlGd1hLb3VwYnFNdVFPRDFhUm9GdkJRN1J2MndUWjh5REE2YmYzcEo3NXJSYTA1UUpDaEYKL1pkR2FIZ2JYK0FuYXlzdUovNHR1OTdEcGFFMmpWVE55K05iMm1kb2ZSZ0RxbHMxbG1kY1BBY1B2UUlEQVFBQgpBb0dBT3Erd3ZCS1QzUkhvTmRsbHVHYXpLZ0VEZmpSSTdSZVlUbk5XT29uMDdqT2lqVUlMV05JMzJhZnFCMU9JCkJhN1ZTZWtzNnpHMFVsLzBTTXllYVZJaU9idktEdVdnb0MrN2ZOSUhWQUdWbkpwTXl5WXhyMXE4MC9vdWNiMlUKL1AybUs5UnNORmJORm8xOTlra1VoeG1rdjV5RUdJS0RsYmJjV3lYd0xHN1NXU0VDUVFEcERXRms5ZUVnTGFHQgpXWElQMG0yWHJQeFdRdFZ1WTZCRmRVSy9LS2FPL2NQMjVxV3JZL0R1MDhsaUdvY2NkSjBwZmtwU2tCb3hTUTVOClp4T0k5bldsQWtFQTVvWjFxV2VuS0YzbkdhRlBpdzR3RDVBMXNENVhYa3V0djl0cUVTcXQ1a0tuZG81OWdROUMKMFJ5aGQxZ3FoY0M0NC9TWmZEUURwTU5tTGxoRUljRUdPUUpBQmJZRU92c2poeXhYRnRwZ1J5NzY3SXFhckdwNgozSGVvaDhzMTFZVmpmNEdNZWRKeElPQVVHV1lyT3pJM09XVktMS2dobmlCVjQvdE1WRzFBTjAwQzJRSkJBSVJ5CmdLdmlhQUlqWWFJeU1sZDh3VlJQME9rQUNJYWZDS2NRMDdJbFNXRGdyd0xJLzRibFU4aDlvSy9ITWpkQzhYZlgKazAvdk9xQ3h1OFdvNVF4WHNORUNRUURHVFJtSElsS2g1QzJ6dUhlL1R6ajFjZjNKUkRsbVpQcXhNN1NuNTJpYwoxS0p3aHFrd1IxR0ZKZHpUVitpZEx1RlN6QUIvRHNjejFhTTRJUlBYSXVmNgotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQ==")
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	// c.Cookie(&fiber.Cookie{
	// 	Name:     "access_token",
	// 	Value:    *accessTokenDetails.Token,
	// 	Path:     "/",
	// 	MaxAge:   15 * 60,
	// 	Secure:   false,
	// 	HTTPOnly: true,
	// 	Domain:   "localhost",
	// })

	// c.Cookie(&fiber.Cookie{
	// 	Name:     "logged_in",
	// 	Value:    "true",
	// 	Path:     "/",
	// 	MaxAge:   15 * 60,
	// 	Secure:   false,
	// 	HTTPOnly: false,
	// 	Domain:   "localhost",
	// })

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "access_token": accessTokenDetails.Token, "refresh_token": refreshTokenDetails.Token, "is_admin": user.IsAdmin})
}

func (aC *authController) RefreshAccessTokenAsuransi(c *fiber.Ctx) error {
	message := "could not refresh access token"
	var body map[string]string

	c.BodyParser(&body)

	if body["refresh_token"] == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": message})
	}

	tokenClaims, err := utils.ValidateToken(body["refresh_token"], "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FEUjNHMmZCWFgxYjhlUVl2SlNjcDUrK2M0UQoxOFhNZHJJU0Vma1I1V3l3MlQ5L3h5RDkyMVFFdnZqdmNQUXZpRndYS291cGJxTXVRT0QxYVJvRnZCUTdSdjJ3ClRaOHlEQTZiZjNwSjc1clJhMDVRSkNoRi9aZEdhSGdiWCtBbmF5c3VKLzR0dTk3RHBhRTJqVlROeStOYjJtZG8KZlJnRHFsczFsbWRjUEFjUHZRSURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQ==")
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	user, err := aC.aS.RefreshTokenAsuransi(tokenClaims.UserID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": err})
	}

	accessTokenDetails, err := utils.CreateToken(user.ID, 5, "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlDWEFJQkFBS0JnUUNOQkQwaDRPRGVTd3E0WnBIMmpBVVd5RFV1V01sb1ZlT3RSQm8yOW9TcVFJazVKSlJaCjR6SkZHZ2k0a1d4amJZemY2ZEMyRWJIdFIrbTYyM3RsV05iS3VJcklYQnZNTnVtS0RMeG9qN3VOVUVWclFVRVkKVmE5U25OUzFsM1U4S1FRbTFYaWFMamVIYWs3QzhYZi9wUEFNalhzcHdQNE1lNktpVE1RcU5ESE4zUUlEQVFBQgpBb0dBU2EwNlIzWVg1dXlzT0RZVzh3cXJLZ0VHa0NXQmJZcmFmcytESnM1YitCdnAxam8vYkV0aEcydUR2UEwxCi8yamdYcWpxREFab3dRRitvOHRDeUd2SEpLMWRURjRpcVRNcmQ3U1dlRk5WYkx3QkpLOXFiOW45NFBTZ1BuV1gKY2JFVXN4VEhLUHRwMlpOWURXcUsyWlFaVlpKRkJzTXZndVNyYlRjcFJia0drSjBDUVFEeEdiNVJRc1hNc3Z4YQpWOFFTUzZpT2J1M1BCVDRqazRJMnB6TzlMR2RwNDNVN0ZPd2QxdjVVbTRvdFpPaFVpWC9BdTRGTDdkS1A0RFdOCm5kRnBSR1lYQWtFQWxic2x3UlBtT3UxZ0VWOUtRakloWVZjMDB0ZlloSStERXdJMEJGVWtnRWNzZGFCRFhSM0YKYXdJL3o4Mlc0bklRN3ZpaXROSDZIYWVacXdid2RRR1lLd0pCQU00T1RXelAzNU5LS1lqZzE2ODNRRkN6RjhYRgoya3kzaGlORmxWK0pjcnk1N0hoWk1rOXlicDFLN2JaTU5wQUJqOURkcit4L3ptU3VuN1p2K2dpNHIzTUNRRVduCnduQ0g2VnNRZ3RpU0UrR25vSS9BR2ZyY0h3WE1IWllDT0dDcm0wZHgxT1VEb1ZMNFBwY0JmTjRYTGxJNTdsYTkKcERPcVcwamdaMFNBL2V2d3lmRUNRRW5WVnRmbThGdDZoSWFPOXNwY2VSRlhUeURBL2V1d1hhQVhubXk4RWZXZQpSaVZtSFhDVzlaSm9DbVpjbU5zY1o1Vm8zYzRMYzl2cTQyMXRKaXVRRGUwPQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQ==")
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	refreshTokenDetails, err := utils.CreateToken(user.ID, 6, "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlDWFFJQkFBS0JnUURSM0cyZkJYWDFiOGVRWXZKU2NwNSsrYzRRMThYTWRySVNFZmtSNVd5dzJUOS94eUQ5CjIxUUV2dmp2Y1BRdmlGd1hLb3VwYnFNdVFPRDFhUm9GdkJRN1J2MndUWjh5REE2YmYzcEo3NXJSYTA1UUpDaEYKL1pkR2FIZ2JYK0FuYXlzdUovNHR1OTdEcGFFMmpWVE55K05iMm1kb2ZSZ0RxbHMxbG1kY1BBY1B2UUlEQVFBQgpBb0dBT3Erd3ZCS1QzUkhvTmRsbHVHYXpLZ0VEZmpSSTdSZVlUbk5XT29uMDdqT2lqVUlMV05JMzJhZnFCMU9JCkJhN1ZTZWtzNnpHMFVsLzBTTXllYVZJaU9idktEdVdnb0MrN2ZOSUhWQUdWbkpwTXl5WXhyMXE4MC9vdWNiMlUKL1AybUs5UnNORmJORm8xOTlra1VoeG1rdjV5RUdJS0RsYmJjV3lYd0xHN1NXU0VDUVFEcERXRms5ZUVnTGFHQgpXWElQMG0yWHJQeFdRdFZ1WTZCRmRVSy9LS2FPL2NQMjVxV3JZL0R1MDhsaUdvY2NkSjBwZmtwU2tCb3hTUTVOClp4T0k5bldsQWtFQTVvWjFxV2VuS0YzbkdhRlBpdzR3RDVBMXNENVhYa3V0djl0cUVTcXQ1a0tuZG81OWdROUMKMFJ5aGQxZ3FoY0M0NC9TWmZEUURwTU5tTGxoRUljRUdPUUpBQmJZRU92c2poeXhYRnRwZ1J5NzY3SXFhckdwNgozSGVvaDhzMTFZVmpmNEdNZWRKeElPQVVHV1lyT3pJM09XVktMS2dobmlCVjQvdE1WRzFBTjAwQzJRSkJBSVJ5CmdLdmlhQUlqWWFJeU1sZDh3VlJQME9rQUNJYWZDS2NRMDdJbFNXRGdyd0xJLzRibFU4aDlvSy9ITWpkQzhYZlgKazAvdk9xQ3h1OFdvNVF4WHNORUNRUURHVFJtSElsS2g1QzJ6dUhlL1R6ajFjZjNKUkRsbVpQcXhNN1NuNTJpYwoxS0p3aHFrd1IxR0ZKZHpUVitpZEx1RlN6QUIvRHNjejFhTTRJUlBYSXVmNgotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQ==")
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	// c.Cookie(&fiber.Cookie{
	// 	Name:     "access_token",
	// 	Value:    *accessTokenDetails.Token,
	// 	Path:     "/",
	// 	MaxAge:   15 * 60,
	// 	Secure:   false,
	// 	HTTPOnly: true,
	// 	Domain:   "localhost",
	// })

	// c.Cookie(&fiber.Cookie{
	// 	Name:     "logged_in",
	// 	Value:    "true",
	// 	Path:     "/",
	// 	MaxAge:   15 * 60,
	// 	Secure:   false,
	// 	HTTPOnly: false,
	// 	Domain:   "localhost",
	// })

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "access_token": accessTokenDetails.Token, "refresh_token": refreshTokenDetails.Token})
}

func (aC *authController) LogoutUser(c *fiber.Ctx) error {
	message := "Token is invalid or session has expired"

	refresh_token := c.Cookies("refresh_token")

	if refresh_token == "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": message})
	}

	expired := time.Now().Add(-time.Hour * 24)
	c.Cookie(&fiber.Cookie{
		Name:    "access_token",
		Value:   "",
		Expires: expired,
	})
	c.Cookie(&fiber.Cookie{
		Name:    "refresh_token",
		Value:   "",
		Expires: expired,
	})
	c.Cookie(&fiber.Cookie{
		Name:    "logged_in",
		Value:   "",
		Expires: expired,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}
