package view

import (
	"fmt"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/rivo/tview"
	"github.com/rs/zerolog/log"
)

// AuthService - service to register and login users.
type AuthService interface {
	Register(user *dto.User) error
	Login(email string, pwd string) error
}

// AuthView - view with pages for authentication operations.
type AuthView struct {
	PagesView
	authService AuthService
}

// NewAuthView creates new AuthView object.
func NewAuthView(authService AuthService) *AuthView {
	pagesView := NewPagesView()
	return &AuthView{authService: authService, PagesView: *pagesView}
}

// RegisterUserPage shows page to perform user registration.
func (v *AuthView) RegisterUserPage() {
	var user dto.User
	form := tview.NewForm()
	emailInput := v.newEmailInput().SetChangedFunc(func(email string) {
		user.Email = email
	})
	form.AddFormItem(emailInput)
	pwdInput := v.newPasswordInput().SetChangedFunc(func(pwd string) {
		user.Password = pwd
	})
	form.AddFormItem(pwdInput)
	form.AddButton("Sign up", func() {
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

// LoginUserPage shows page to perform user login.
func (v *AuthView) LoginUserPage() {
	var userEmail string
	var userPwd string
	form := tview.NewForm()
	emailInput := v.newEmailInput().SetChangedFunc(func(email string) {
		userEmail = email
	})
	form.AddFormItem(emailInput)
	pwdInput := v.newPasswordInput().SetChangedFunc(func(pwd string) {
		userPwd = pwd
	})
	form.AddFormItem(pwdInput)
	form.AddButton("Login", func() {
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

func (v *AuthView) newEmailInput() *tview.InputField {
	return tview.NewInputField().SetLabel("Email").SetFieldWidth(32)
}

func (v *AuthView) newPasswordInput() *tview.InputField {
	return tview.NewInputField().SetLabel("Password").SetFieldWidth(64).SetMaskCharacter('*')
}
