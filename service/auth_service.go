package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"wkm/entity"
	"wkm/repository"
	"wkm/request"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	SignInUser(r request.SigninRequest) (entity.User, error)
	SignUpUser(r request.SignupRequest) (entity.User, error)
	GenerateNewOtp(r request.OtpCheck) (entity.Otp, error)
	RefreshToken(r string) (entity.User, error)
	RefreshTokenAsuransi(r string) (entity.User, error)
	ResetPassword(data request.ResetPassword) request.Response
	ResetPasswordByOtp(data request.ResetPasswordOtp) request.Response
	ConsumeFonnte(body request.OtpCheck) (map[string]interface{}, error)
	CheckOtp(r request.OtpCheck) error
	CheckOtpReset(r request.OtpCheck) (entity.Otp, error)
}

type authService struct {
	uR repository.UserRepository
	oR repository.OtpRepository
}

func NewAuthService(uR repository.UserRepository, oR repository.OtpRepository) AuthService {
	return &authService{
		uR,
		oR,
	}
}

func (s *authService) SignInUser(r request.SigninRequest) (entity.User, error) {
	user := s.uR.FindByPhoneNumber(r.Username)
	if !user.Active {
		otp := s.oR.CreateOtp(user.NoHp)

		data, err := s.ConsumeFonnte(request.OtpCheck{NoHp: user.NoHp, Otp: otp.Otp})
		if err != nil {
			return user, err
		}
		fmt.Println("fonnte res", data)
	}
	return user, nil
}

func (s *authService) ConsumeFonnte(body request.OtpCheck) (map[string]interface{}, error) {
	var client = &http.Client{}
	var data map[string]interface{}
	var param = url.Values{}
	param.Set("target", body.NoHp)
	param.Set("message", fmt.Sprintf("%s%d", "Berikut kode OTP ", body.Otp))
	param.Set("schedule", "0")
	param.Set("delay", "2")
	param.Set("countryCode", "62")
	var payload = bytes.NewBufferString(param.Encode())
	request, err := http.NewRequest("POST", "https://api.fonnte.com/send", payload)
	if err != nil {
		return map[string]interface{}{}, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", "k!ph_r+apphR8kJY@+gS")
	response, err := client.Do(request)
	if err != nil {
		return map[string]interface{}{}, err
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return map[string]interface{}{}, err
	}
	return data, nil
}

func (s *authService) SignUpUser(r request.SignupRequest) (entity.User, error) {
	password, _ := bcrypt.GenerateFromPassword([]byte(r.Password), 8)
	r.Password = string(password)
	user, err := s.uR.CreateUser(r)
	if err != nil {
		return user, err
	}

	otp := s.oR.CreateOtp(user.NoHp)

	data, err := s.ConsumeFonnte(request.OtpCheck{NoHp: user.NoHp, Otp: otp.Otp})
	if err != nil {
		return user, err
	}
	fmt.Println("Response Fonnte", data)

	return user, err
}

func (s *authService) GenerateNewOtp(r request.OtpCheck) (entity.Otp, error) {
	otp := s.oR.CreateOtp(r.NoHp)

	data, err := s.ConsumeFonnte(request.OtpCheck{NoHp: r.NoHp, Otp: otp.Otp})
	if err != nil {
		return otp, err
	}
	fmt.Println("Response Fonnte", data)

	return otp, err
}

func (s *authService) CheckOtp(r request.OtpCheck) error {
	return s.oR.Check(r)
}

func (s *authService) CheckOtpReset(r request.OtpCheck) (entity.Otp, error) {
	return s.oR.CheckOtpReset(r)
}

func (s *authService) RefreshToken(r string) (entity.User, error) {
	return s.uR.FindById(r), nil
}

func (s *authService) RefreshTokenAsuransi(r string) (entity.User, error) {
	return s.uR.FindById(r), nil
}

func (s *authService) ResetPassword(data request.ResetPassword) request.Response {
	user := s.uR.FindById(data.IdUser)
	if user.ID == "" {
		return request.Response{Status: 400, Message: "User tidak ditemukan"}
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.PasswordLama))
	if err != nil {
		return request.Response{Status: 400, Message: "Password salah"}
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 8)
	data.Password = string(password)
	s.uR.ResetPassword(data)
	return request.Response{Status: 201, Message: "Data berhasil diupdate"}
}

func (s *authService) ResetPasswordByOtp(data request.ResetPasswordOtp) request.Response {
	otp := s.oR.FindById(data.Token)
	if otp.Used {
		return request.Response{Status: 400, Message: "Kode OTP telah digunakan"}
	}
	user := s.uR.FindByPhoneNumber(otp.NoHp)
	if user.ID == "" {
		return request.Response{Status: 400, Message: "User tidak ditemukan"}
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 8)
	data.Password = string(password)
	otp.Used = true
	s.oR.Update(otp)
	s.uR.ResetPassword(request.ResetPassword{IdUser: user.ID, Password: data.Password})
	return request.Response{Status: 201, Message: "Data berhasil diupdate"}
}
