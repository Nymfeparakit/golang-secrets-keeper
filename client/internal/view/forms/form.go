package forms

import (
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/rivo/tview"
)

// UpdateRetrievePasswordService - service for retrieving/updating LoginPassword instance.
type UpdateRetrievePasswordService interface {
	GetSecretByID(id string) (dto.LoginPassword, error)
	UpdateSecret(secret dto.LoginPassword) error
}

// UpdateRetrieveCardService - service for retrieving/updating CardInfo instance.
type UpdateRetrieveCardService interface {
	GetSecretByID(id string) (dto.CardInfo, error)
	UpdateSecret(secret dto.CardInfo) error
}

// UpdateRetrieveTextService - service for retrieving/updating TextInfo instance.
type UpdateRetrieveTextService interface {
	GetSecretByID(id string) (dto.TextInfo, error)
	UpdateSecret(secret dto.TextInfo) error
}

// UpdateRetrieveBinaryService - service for retrieving/updating BinaryInfo instance.
type UpdateRetrieveBinaryService interface {
	GetSecretByID(id string) (dto.BinaryInfo, error)
	UpdateSecret(secret dto.BinaryInfo) error
}

// AddSecretService - service for adding new secrets.
type AddSecretService interface {
	AddCredentials(password *dto.LoginPassword) error
	AddTextInfo(text *dto.TextInfo) error
	AddCardInfo(card *dto.CardInfo) error
	AddBinaryInfo(bin *dto.BinaryInfo) error
}

// SaveAction type of save action in secret form.
type SaveAction int

const (
	UPDATE SaveAction = iota
	CREATE
)

// SaveItemForm - form to perform save operations with certain secret.
type SaveItemForm interface {
	tview.Primitive
	AddInputs()
	AddBtn(label string, selected func())
	Save() error
	SetSaveAction(action SaveAction)
	SetSecret(id string) error
}

// FillSaveItemForm adds inputs and save button to SaveItemForm.
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

// FormWithSaveAction - form with save action attribute.
type FormWithSaveAction struct {
	saveAction SaveAction
}

// SetSaveAction sets save action in form.
func (f *FormWithSaveAction) SetSaveAction(action SaveAction) {
	f.saveAction = action
}
