package main

import (
	"fmt"
	"os"
	"wkm/config"
	"wkm/controller"
	"wkm/middleware"
	"wkm/repository"
	"wkm/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	connMain, sqlConnMain = config.GetConnectionMain()

	userRepository repository.UserRepository = repository.NewUserRepository(connMain)
	otpRepository  repository.OtpRepository  = repository.NewOtpRepository(connMain)
	authService    service.AuthService       = service.NewAuthService(userRepository, otpRepository)
	authController controller.AuthController = controller.NewAuthController(authService)

	memberRepository repository.MemberRepository = repository.NewMemberRepository(connMain)
	memberService    service.MemberService       = service.NewMemberService(memberRepository)
	memberController controller.MemberController = controller.NewMemberController(memberService)
)

func main() {

	defer sqlConnMain.Close()

	app := fiber.New(fiber.Config{})
	app.Static("/uploads", "./uploads")

	app.Use(logger.New(logger.Config{
		Format:     "${time} | ${status} | ${method} | ${path} | ${ip} | ${queryParams} |${latency} \n\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Indonesia/Jakarta",
		Done: func(c *fiber.Ctx, logString []byte) {
		},
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://gofiber.io, http://localhost:3002",
		AllowHeaders: "Access-Control-Allow-Headers, Origin, Content-Type, Accept, Authorization",
	}))

	auth := app.Group("/auth")
	auth.Post("/login", authController.SignInUser)
	auth.Post("/signup", authController.SignUpUser)
	auth.Post("/generate-new-otp", authController.GenerateNewOtp)
	auth.Post("/check-otp", authController.CheckOtp)
	auth.Post("/check-otp-reset", authController.CheckOtpReset)
	auth.Post("/refresh-token", authController.RefreshAccessTokenAsuransi)
	auth.Post("/reset-password", middleware.DeserializeUser, authController.ResetPassword)
	auth.Post("/reset-password-by-otp", authController.ResetPasswordByOtp)
	auth.Post("/logout", authController.LogoutUser)
	auth.Get("/generate-password", authController.GeneratePassword)

	member := app.Group("/member")
	member.Get("/my-card", middleware.DeserializeUser, memberController.Mine)
	member.Post("/add-card", middleware.DeserializeUser, memberController.AddCard)

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Current Directory:", dir)

	app.Listen(":3001")

}
