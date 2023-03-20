package forms

import (
	"context"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/rivo/tview"
	"time"
)

// LoginPasswordForm - form to perform save operations with certain LoginPassword.
type LoginPasswordForm struct {
	loginPwd *dto.LoginPassword
	tview.Form
	addService            AddSecretService
	retrieveUpdateService UpdateRetrievePasswordService
	FormWithSaveAction
}

// NewLoginPasswordForm creates LoginPasswordForm object.
func NewLoginPasswordForm(service AddSecretService, updateService UpdateRetrievePasswordService) *LoginPasswordForm {
	pwd := &dto.LoginPassword{}
	form := tview.NewForm()
	return &LoginPasswordForm{loginPwd: pwd, Form: *form, addService: service, retrieveUpdateService: updateService}
}

// AddInputs adds input fields to form.
func (f *LoginPasswordForm) AddInputs() {
	f.AddInputField("Name", f.loginPwd.Name, 64, nil, func(name string) {
		f.loginPwd.Name = name
	})
	f.AddInputField("Login", f.loginPwd.Login, 32, nil, func(login string) {
		f.loginPwd.Login = login
	})
	f.AddInputField("Metadata", f.loginPwd.Metadata, 128, nil, func(metadata string) {
		f.loginPwd.Metadata = metadata
	})
	pwdInput := tview.NewInputField().
		SetLabel("Password").
		SetText(f.loginPwd.Password).
		SetFieldWidth(64).
		SetChangedFunc(func(pwd string) {
			f.loginPwd.Password = pwd
		})
	f.AddFormItem(pwdInput)
}

// Save performs operation with LoginPassword when the save button is clicked.
func (f *LoginPasswordForm) Save(ctx context.Context) error {
	f.loginPwd.UpdatedAt = time.Now().UTC()
	var err error
	switch f.saveAction {
	case UPDATE:
		err = f.retrieveUpdateService.UpdateSecret(ctx, *f.loginPwd)
	case CREATE:
		err = f.addService.AddCredentials(ctx, f.loginPwd)
	}
	return err
}

// AddBtn adds button with certain label and selected function.
func (f *LoginPasswordForm) AddBtn(label string, selected func()) {
	f.AddButton(label, selected)
}

// SetSecret associates specific LoginPassword with form.
func (f *LoginPasswordForm) SetSecret(ctx context.Context, id string) error {
	pwd, err := f.retrieveUpdateService.GetSecretByID(ctx, id)
	f.loginPwd = &pwd
	return err
}
