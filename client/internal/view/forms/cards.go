package forms

import (
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/rivo/tview"
	"github.com/rs/zerolog/log"
	"strconv"
)

type CardInfoForm struct {
	cardInfo *dto.CardInfo
	tview.Form
	itemService ItemsService
}

func NewCardInfoForm(service ItemsService) *CardInfoForm {
	cardInfo := &dto.CardInfo{}
	form := tview.NewForm()
	return &CardInfoForm{cardInfo: cardInfo, Form: *form, itemService: service}
}

func (f *CardInfoForm) AddInputs() {
	f.AddInputField("Name", "", 64, nil, func(name string) {
		f.cardInfo.Name = name
	})
	f.AddInputField("Card number", "", 16, nil, func(number string) {
		f.cardInfo.CardNumber = number
	})
	// todo: how to validate value correctly?
	f.AddInputField("CVV", "", 3, nil, func(cvv string) {
		cvvNum, err := strconv.Atoi(cvv)
		if err != nil {
			log.Fatal().Err(err).Msg("wrong value for cvv")
		}
		f.cardInfo.CVV = int32(cvvNum)
	})
	f.AddInputField("Expiration Month", "", 2, nil, func(monthStr string) {
		month, err := strconv.Atoi(monthStr)
		if err != nil {
			log.Fatal().Err(err).Msg("wrong value for expiration month")
		}
		f.cardInfo.ExpirationMonth = int32(month)
	})
	f.AddInputField("Expiration Year", "", 4, nil, func(yearStr string) {
		year, err := strconv.Atoi(yearStr)
		if err != nil {
			log.Fatal().Err(err).Msg("wrong value for expiration year")
		}
		f.cardInfo.ExpirationYear = int32(year)
	})
}

func (f *CardInfoForm) Save(saveAction SaveAction) error {
	var err error
	switch saveAction {
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
