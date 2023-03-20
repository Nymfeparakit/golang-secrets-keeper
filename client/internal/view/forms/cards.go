package forms

import (
	"context"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/rivo/tview"
	"time"
)

// CardInfoForm - form to perform save operations with certain CardInfo.
type CardInfoForm struct {
	cardInfo *dto.CardInfo
	tview.Form
	itemService           AddSecretService
	retrieveUpdateService UpdateRetrieveCardService
	FormWithSaveAction
}

// NewCardInfoForm creates CardInfoForm object.
func NewCardInfoForm(service AddSecretService, updateService UpdateRetrieveCardService) *CardInfoForm {
	cardInfo := &dto.CardInfo{}
	form := tview.NewForm()
	return &CardInfoForm{cardInfo: cardInfo, Form: *form, itemService: service, retrieveUpdateService: updateService}
}

// AddInputs adds input fields to form.
func (f *CardInfoForm) AddInputs() {
	f.AddInputField("Name", f.cardInfo.Name, 64, nil, func(name string) {
		f.cardInfo.Name = name
	})
	f.AddInputField("Card number", f.cardInfo.Number, 16, nil, func(number string) {
		f.cardInfo.Number = number
	})
	f.AddInputField("CVV", f.cardInfo.CVV, 3, nil, func(cvv string) {
		f.cardInfo.CVV = cvv
	})
	f.AddInputField("Expiration Month", f.cardInfo.ExpirationYear, 2, nil, func(monthStr string) {
		f.cardInfo.ExpirationMonth = monthStr
	})
	f.AddInputField("Expiration Year", f.cardInfo.ExpirationYear, 4, nil, func(yearStr string) {
		f.cardInfo.ExpirationYear = yearStr
	})
	f.AddInputField("Metadata", f.cardInfo.Metadata, 128, nil, func(metadata string) {
		f.cardInfo.Metadata = metadata
	})
}

// Save performs operation with CardInfo when the save button is clicked.
func (f *CardInfoForm) Save(ctx context.Context) error {
	f.cardInfo.UpdatedAt = time.Now().UTC()
	var err error
	switch f.saveAction {
	case UPDATE:
		err = f.retrieveUpdateService.UpdateSecret(ctx, *f.cardInfo)
	case CREATE:
		err = f.itemService.AddCardInfo(ctx, f.cardInfo)
	}
	return err
}

// AddBtn adds button with certain label and selected function.
func (f *CardInfoForm) AddBtn(label string, selected func()) {
	f.AddButton(label, selected)
}

// SetSecret associates specific CardInfo with form.
func (f *CardInfoForm) SetSecret(ctx context.Context, id string) error {
	crd, err := f.retrieveUpdateService.GetSecretByID(ctx, id)
	f.cardInfo = &crd
	return err
}
