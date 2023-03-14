package forms

import (
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/rivo/tview"
)

type UpdateRetrievePasswordService interface {
	GetSecretByID(id string) (dto.LoginPassword, error)
	UpdateSecret(secret dto.LoginPassword) error
}

type UpdateRetrieveCardService interface {
	GetSecretByID(id string) (dto.CardInfo, error)
	UpdateSecret(secret dto.CardInfo) error
}

type UpdateRetrieveTextService interface {
	GetSecretByID(id string) (dto.TextInfo, error)
	UpdateSecret(secret dto.TextInfo) error
}

type UpdateRetrieveBinaryService interface {
	GetSecretByID(id string) (dto.BinaryInfo, error)
	UpdateSecret(secret dto.BinaryInfo) error
}

type AddSecretService interface {
	AddCredentials(password *dto.LoginPassword) error
	AddTextInfo(text *dto.TextInfo) error
	AddCardInfo(card *dto.CardInfo) error
	AddBinaryInfo(bin *dto.BinaryInfo) error
}

type SaveAction int

const (
	UPDATE SaveAction = iota
	CREATE
)

type SaveItemForm interface {
	tview.Primitive
	AddInputs()
	AddBtn(label string, selected func())
	Save() error
	SetSaveAction(action SaveAction)
	SetSecret(id string) error
}

func FillSaveItemForm(
	c SaveItemForm,
	saveAction SaveAction,
	secretID string,
	processResultFunc func(SaveAction, error),
) (SaveItemForm, error) {
	c.SetSaveAction(saveAction)
	if saveAction == UPDATE {
		err := c.SetSecret(secretID)
		if err != nil {
			return nil, err
		}
	}
	c.AddInputs()
	c.AddBtn("Save", func() {
		err := c.Save()
		processResultFunc(saveAction, err)
	})
	return c, nil
}

type FormWithSaveAction struct {
	saveAction SaveAction
}

func (f *FormWithSaveAction) SetSaveAction(action SaveAction) {
	f.saveAction = action
}
