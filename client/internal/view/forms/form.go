package forms

import (
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/rivo/tview"
)

type ItemsService interface {
	AddPassword(password *dto.LoginPassword) error
	AddTextInfo(text *dto.TextInfo) error
	AddCardInfo(card *dto.CardInfo) error
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
	Save(saveAction SaveAction) error
}

func FillSaveItemForm(
	c SaveItemForm,
	saveAction SaveAction,
	processResultFunc func(SaveAction, error),
) SaveItemForm {
	c.AddInputs()
	c.AddBtn("Save", func() {
		err := c.Save(saveAction)
		processResultFunc(saveAction, err)
	})
	return c
}
