package forms

import (
	"context"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/rivo/tview"
)

// UpdateRetrievePasswordService - service for retrieving/updating LoginPassword instance.
type UpdateRetrievePasswordService interface {
	GetSecretByID(ctx context.Context, id string) (dto.LoginPassword, error)
	UpdateSecret(ctx context.Context, secret dto.LoginPassword) error
}

// UpdateRetrieveCardService - service for retrieving/updating CardInfo instance.
type UpdateRetrieveCardService interface {
	GetSecretByID(ctx context.Context, id string) (dto.CardInfo, error)
	UpdateSecret(ctx context.Context, secret dto.CardInfo) error
}

// UpdateRetrieveTextService - service for retrieving/updating TextInfo instance.
type UpdateRetrieveTextService interface {
	GetSecretByID(ctx context.Context, id string) (dto.TextInfo, error)
	UpdateSecret(ctx context.Context, secret dto.TextInfo) error
}

// UpdateRetrieveBinaryService - service for retrieving/updating BinaryInfo instance.
type UpdateRetrieveBinaryService interface {
	GetSecretByID(ctx context.Context, id string) (dto.BinaryInfo, error)
	UpdateSecret(ctx context.Context, secret dto.BinaryInfo) error
}

// AddSecretService - service for adding new secrets.
type AddSecretService interface {
	AddCredentials(ctx context.Context, password *dto.LoginPassword) error
	AddTextInfo(ctx context.Context, text *dto.TextInfo) error
	AddCardInfo(ctx context.Context, card *dto.CardInfo) error
	AddBinaryInfo(ctx context.Context, bin *dto.BinaryInfo) error
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
	Save(ctx context.Context) error
	SetSaveAction(action SaveAction)
	SetSecret(ctx context.Context, id string) error
}

// FillSaveItemForm adds inputs and save button to SaveItemForm.
func FillSaveItemForm(
	ctx context.Context,
	c SaveItemForm,
	saveAction SaveAction,
	secretID string,
	processResultFunc func(SaveAction, error),
) (SaveItemForm, error) {
	c.SetSaveAction(saveAction)
	if saveAction == UPDATE {
		err := c.SetSecret(ctx, secretID)
		if err != nil {
			return nil, err
		}
	}
	c.AddInputs()
	c.AddBtn("Save", func() {
		err := c.Save(ctx)
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
