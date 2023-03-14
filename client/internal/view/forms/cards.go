package forms

import (
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/rivo/tview"
)

type CardInfoForm struct {
	cardInfo *dto.CardInfo
	tview.Form
	itemService           AddSecretService
	retrieveUpdateService UpdateRetrieveCardService
	FormWithSaveAction
}

func NewCardInfoForm(service AddSecretService, updateService UpdateRetrieveCardService) *CardInfoForm {
	cardInfo := &dto.CardInfo{}
	form := tview.NewForm()
	return &CardInfoForm{cardInfo: cardInfo, Form: *form, itemService: service, retrieveUpdateService: updateService}
}

func (f *CardInfoForm) AddInputs() {
	f.AddInputField("Name", "", 64, nil, func(name string) {
		f.cardInfo.Name = name
	})
	f.AddInputField("Card number", "", 16, nil, func(number string) {
		f.cardInfo.Number = number
	})
	f.AddInputField("CVV", "", 3, nil, func(cvv string) {
		f.cardInfo.CVV = cvv
	})
	f.AddInputField("Expiration Month", "", 2, nil, func(monthStr string) {
		f.cardInfo.ExpirationMonth = monthStr
	})
	f.AddInputField("Expiration Year", "", 4, nil, func(yearStr string) {
		f.cardInfo.ExpirationYear = yearStr
	})
	f.AddInputField("Metadata", f.cardInfo.Metadata, 128, nil, func(metadata string) {
		f.cardInfo.Metadata = metadata
	})
}

func (f *CardInfoForm) Save() error {
	var err error
	switch f.saveAction {
	case UPDATE:
		err = f.itemService.AddCardInfo(f.cardInfo)
	case CREATE:
		err = f.itemService.AddCardInfo(f.cardInfo)
	}
	return err
}

func (f *CardInfoForm) AddBtn(label string, selected func()) {
	f.AddButton(label, selected)
}

func (f *CardInfoForm) SetSecret(id string) error {
	crd, err := f.retrieveUpdateService.GetSecretByID(id)
	f.cardInfo = &crd
	return err
}
