package view

import (
	"fmt"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/rivo/tview"
	"github.com/rs/zerolog/log"
)

type AuthService interface {
	Register(user *dto.User) error
	Login(email string, pwd string) error
}

type AuthView struct {
	PagesView
	authService AuthService
}

func NewAuthView(authService AuthService) *AuthView {
	pagesView := NewPagesView()
	return &AuthView{authService: authService, PagesView: *pagesView}
}

func (v *AuthView) RegisterUserPage() {
	var user dto.User
	form := tview.NewForm().
		AddInputField("Email", "", 32, nil, func(email string) {
			user.Email = email
		}).
		AddInputField("Password", "", 128, nil, func(pwd string) {
			user.Password = pwd
		}).
		AddButton("Save", func() {
			err := v.authService.Register(&user)
			resultMsg := "New user has been successfully registered"
			if err != nil {
				resultMsg = fmt.Sprintf("An error occurred during registration: %v", err)
			}
			v.ResultPage(resultMsg)
		})
	v.pages.AddPage("Register user", form, true, true)
	err := v.app.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
}

func (v *AuthView) LoginUserPage() {
	var userEmail string
	var userPwd string
	form := tview.NewForm().
		AddInputField("Email", "", 32, nil, func(email string) {
			userEmail = email
		}).
		AddInputField("Password", "", 128, nil, func(pwd string) {
			userPwd = pwd
		}).
		AddButton("Save", func() {
			err := v.authService.Login(userEmail, userPwd)
			resultMsg := "User credentials saved"
			if err != nil {
				resultMsg = fmt.Sprintf("An error occurred during login: %v", err)
			}
			v.ResultPage(resultMsg)
		})
	v.pages.AddPage("Login user", form, true, true)
	err := v.app.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
