package main

import (
	"fmt"
	"os"
	"wkm/config"
	"wkm/controller"
	"wkm/internal/merchant"
	"wkm/internal/outlet"
	"wkm/internal/token"
	"wkm/middleware"
	"wkm/repository"
	"wkm/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
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

	profileRepository repository.ProfileRepository = repository.NewProfileRepository(connMain)
	profileService    service.ProfileService       = service.NewProfileService(profileRepository)
	profileController controller.ProfileController = controller.NewProfileController(profileService)

	masterDataRepository repository.MasterDataRepository = repository.NewMasterDataRepository(connMain)
	masterDataService    service.MasterDataService       = service.NewMasterDataService(masterDataRepository)
	masterDataController controller.MasterDataController = controller.NewMasterDataController(masterDataService)
)

func main() {

	errEnv := godotenv.Load()
	if errEnv != nil {
		fmt.Println("ini errornya ", errEnv)
		panic("Failed to load env file")
	}

	defer sqlConnMain.Close()

	// connMain.AutoMigrate(&merchant.Merchant{})
	// connMain.AutoMigrate(&outlet.Outlet{})

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
		AllowOrigins: "https://gofiber.io, http://localhost:3002, http://192.168.70.17:3002, https://www.e-cardplus.co.id",
		AllowHeaders: "Access-Control-Allow-Headers, Origin, Content-Type, Accept, Authorization, Access-Control-Allow-Origin",
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

	profile := app.Group("/profile")
	profile.Get("/me", middleware.DeserializeUser, profileController.Me)
	profile.Post("/upload/image-profile", middleware.DeserializeUser, profileController.UploadImageProfile)
	profile.Post("/update/profile-data", middleware.DeserializeUser, profileController.Update)

	masterData := app.Group("/master-data")
	masterData.Get("/kategori-merchant", middleware.DeserializeUser, masterDataController.KategoriMerchantAll)
	masterData.Get("/media-promosi", middleware.DeserializeUser, masterDataController.MediaPromosiAll)
	masterData.Get("/pic-mro", middleware.DeserializeUser, masterDataController.PicMroAll)

	member := app.Group("/member")
	member.Get("/my-card", middleware.DeserializeUser, memberController.Mine)
	member.Post("/add-card", middleware.DeserializeUser, memberController.AddCard)
	member.Post("/create-new-member", middleware.AdminAccess, memberController.CreateNewMemberCard)


	// admin
	token.RegisterRoutes(app, connMain)
	merchant.RegisterRoutes(app, connMain)
	outlet.RegisterRoutes(app, connMain)

	app.Listen(":" + os.Getenv("PORT"))

}
