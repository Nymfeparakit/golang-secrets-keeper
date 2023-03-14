package forms

import (
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/rivo/tview"
)

type LoginPasswordForm struct {
	loginPwd *dto.LoginPassword
	tview.Form
	addService            AddSecretService
	retrieveUpdateService UpdateRetrievePasswordService
	FormWithSaveAction
}

func NewLoginPasswordForm(service AddSecretService, updateService UpdateRetrievePasswordService) *LoginPasswordForm {
	pwd := &dto.LoginPassword{}
	form := tview.NewForm()
	return &LoginPasswordForm{loginPwd: pwd, Form: *form, addService: service, retrieveUpdateService: updateService}
}

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

func (f *LoginPasswordForm) Save() error {
	var err error
	switch f.saveAction {
	case UPDATE:
		err = f.retrieveUpdateService.UpdateSecret(*f.loginPwd)
	case CREATE:
		err = f.addService.AddCredentials(f.loginPwd)
	}
	return err
}

func (f *LoginPasswordForm) AddBtn(label string, selected func()) {
	f.AddButton(label, selected)
}

func (f *LoginPasswordForm) SetSecret(id string) error {
	pwd, err := f.retrieveUpdateService.GetSecretByID(id)
	f.loginPwd = &pwd
	return err
}
