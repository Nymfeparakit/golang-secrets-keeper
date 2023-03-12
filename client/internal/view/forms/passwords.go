package forms

import (
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/rivo/tview"
)

type LoginPasswordForm struct {
	loginPwd *dto.LoginPassword
	tview.Form
	itemService ItemsService
}

func NewLoginPasswordForm(service ItemsService) *LoginPasswordForm {
	pwd := &dto.LoginPassword{}
	form := tview.NewForm()
	return &LoginPasswordForm{loginPwd: pwd, Form: *form, itemService: service}
}

func (f *LoginPasswordForm) AddInputs() {
	f.AddInputField("Name", "", 64, nil, func(name string) {
		f.loginPwd.Name = name
	})
	f.AddInputField("Login", "", 32, nil, func(login string) {
		f.loginPwd.Login = login
	})
	pwdInput := tview.NewInputField().
		SetLabel("Password").
		SetFieldWidth(64).
		SetChangedFunc(func(pwd string) {
			f.loginPwd.Password = pwd
		})
	f.AddFormItem(pwdInput)
}

func (f *LoginPasswordForm) Save(saveAction SaveAction) error {
	var err error
	switch saveAction {
	case UPDATE:
		err = f.itemService.AddCredentials(f.loginPwd)
	case CREATE:
		err = f.itemService.AddCredentials(f.loginPwd)
	}
	return err
}

func (f *LoginPasswordForm) AddBtn(label string, selected func()) {
	f.AddButton(label, selected)
}
