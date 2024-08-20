package repository

import (
	"crypto/rand"
	"errors"
	"io"
	"strconv"
	"time"
	"wkm/entity"
	"wkm/request"

	"gorm.io/gorm"
)

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func EncodeToString(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	c := string(b)
	if len(c) < 6 {
		EncodeToString(6)
	}
	return c
}

type OtpRepository interface {
	FindById(username string) entity.Otp
	FindByPhoneNumber(username string) entity.User
	Check(data request.OtpCheck) error
	CreateOtp(data string) entity.Otp
	CheckOtpReset(data request.OtpCheck) (entity.Otp, error)
	Update(data entity.Otp)
}

type otpRepository struct {
	connUser *gorm.DB
}

func NewOtpRepository(connUser *gorm.DB) OtpRepository {
	return &otpRepository{
		connUser: connUser,
	}
}

func (lR *otpRepository) Check(data request.OtpCheck) error {
	otp := entity.Otp{}
	lR.connUser.Where("otp", data.Otp).Where("no_hp", data.NoHp).Where("used", 0).First(&otp)
	now := time.Now()
	if otp.ID == "" {
		return errors.New("kode OTP Salah")
	}
	diff := now.Sub(*otp.CreatedAt)
	if diff.Minutes() > 1 {
		return errors.New("kode OTP telah expired, silahkan generate ulang kode OTP")
	}
	user := entity.User{}
	lR.connUser.Where("no_hp", otp.NoHp).First(&user)
	if user.ID == "" {
		return errors.New("kode OTP ini tidak cocok dengan nomor telepon anda!")
	}
	user.Active = true
	lR.connUser.Save(&user)
	otp.Used = true
	lR.connUser.Save(&otp)
	return nil
}

func (lR *otpRepository) CheckOtpReset(data request.OtpCheck) (entity.Otp, error) {
	otp := entity.Otp{}
	lR.connUser.Where("otp", data.Otp).Where("no_hp", data.NoHp).Where("used", 0).First(&otp)
	now := time.Now()
	if otp.ID == "" {
		return otp, errors.New("kode OTP salah!")
	}
	diff := now.Sub(*otp.CreatedAt)
	if diff.Minutes() > 1 {
		return otp, errors.New("kode OTP telah expired, silahkan generate ulang kode OTP!")
	}
	return otp, nil
}

func (lR *otpRepository) CreateOtp(noHp string) entity.Otp {
	code, _ := strconv.Atoi(EncodeToString(6))
	lR.connUser.Where("no_hp", noHp).Delete(&entity.Otp{})
	otp := entity.Otp{NoHp: noHp, Otp: code}
	lR.connUser.Create(&otp)
	return otp
}

func (lR *otpRepository) ResetPassword(data request.ResetPassword) {
	user := entity.User{ID: data.IdUser}
	lR.connUser.Find(&user)
	user.Password = data.Password
	lR.connUser.Save(&user)
}

func (lR *otpRepository) FindByPhoneNumber(username string) entity.User {
	var user entity.User
	lR.connUser.Where("no_hp", username).First(&user)

	// var permissions []entity.Permission
	// lR.connUser.Where("role_id", user.RoleId).Find(&permissions)
	// for _, v := range permissions {
	// 	user.Permissions = append(user.Permissions, v.Name)
	// }

	return user
}

func (lR *otpRepository) FindById(id string) entity.Otp {
	otp := entity.Otp{ID: id}
	lR.connUser.Find(&otp)

	return otp
}

func (lR *otpRepository) Update(otp entity.Otp) {
	lR.connUser.Save(&otp)
}
