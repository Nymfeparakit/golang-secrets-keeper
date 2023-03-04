package view

import (
	"errors"
	"fmt"
	"github.com/Nymfeparakit/gophkeeper/client/internal/services"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/rivo/tview"
	"github.com/rs/zerolog/log"
	"strconv"
)

type ItemsService interface {
	AddPassword(password *dto.LoginPassword) error
	AddTextInfo(text *dto.TextInfo) error
	AddCardInfo(card *dto.CardInfo) error
	ListItems() (dto.ItemsList, error)
}

type ItemsView struct {
	PagesView
	itemsService ItemsService
	items        []interface{}
}

func NewItemsView(itemsService ItemsService) *ItemsView {
	pagesView := NewPagesView()
	return &ItemsView{PagesView: *pagesView, itemsService: itemsService}
}

func (v *ItemsView) AddPasswordPage() {
	form := v.loginPasswordForm()
	v.pages.AddPage("Add login-password", form, true, true)
	err := v.app.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
}

func (v *ItemsView) AddTextInfoPage() {
	form := v.textInfoForm()
	v.pages.AddPage("Add text info", form, true, true)
	err := v.app.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
}

func (v *ItemsView) AddCardInfoPage() {
	form := v.cardInfoForm()
	v.pages.AddPage("Add card info", form, true, true)
	err := v.app.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
}

func (v *ItemsView) ListItemsPage() {
	listView := tview.NewList().ShowSecondaryText(false)
	v.pages.AddPage("List items", listView, true, true)

	resultItems, err := v.itemsService.ListItems()
	if errors.Is(err, services.ErrTokenNotFound) {
		v.ResultPage("You are not authenticated - use 'login' command to set credentials")
		return
	}
	if err != nil {
		v.ResultPage(fmt.Sprintf("can not list items: %v", err.Error()))
		return
	}

	for _, pwd := range resultItems.Passwords {
		listView.AddItem(pwd.Name, "", 0, nil)
		v.items = append(v.items, pwd)
	}
	for _, txt := range resultItems.Texts {
		listView.AddItem(txt.Name, "", 0, nil)
		v.items = append(v.items, txt)
	}

	err = v.app.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
}

// todo: make basic add item form
func (v *ItemsView) loginPasswordForm() *tview.Form {
	var loginPwd dto.LoginPassword
	form := tview.NewForm()
	// todo: make base name input field?
	form.AddInputField("Name", "", 64, nil, func(name string) {
		loginPwd.Name = name
	})
	form.AddInputField("Login", "", 32, nil, func(login string) {
		loginPwd.Login = login
	})
	form.AddInputField("Password", "", 64, nil, func(pwd string) {
		loginPwd.Password = pwd
	})
	form.AddButton("Save", func() {
		err := v.itemsService.AddPassword(&loginPwd)
		resultMsg := "Successfully added new item!"
		if err != nil {
			resultMsg = fmt.Sprintf("Error happened on adding new item: %v", err)
		}
		v.ResultPage(resultMsg)
	})

	return form
}

func (v *ItemsView) textInfoForm() *tview.Form {
	var textInfo dto.TextInfo
	form := tview.NewForm()
	// todo: make base name input field?
	form.AddInputField("Name", "", 64, nil, func(name string) {
		textInfo.Name = name
	})
	form.AddInputField("Text", "", 128, nil, func(text string) {
		textInfo.Text = text
	})
	form.AddButton("Save", func() {
		err := v.itemsService.AddTextInfo(&textInfo)
		resultMsg := "Successfully added new item!"
		if err != nil {
			resultMsg = fmt.Sprintf("Error happened on adding new item: %v", err)
		}
		v.ResultPage(resultMsg)
	})

	return form
}

func (v *ItemsView) cardInfoForm() *tview.Form {
	var cardInfo dto.CardInfo
	form := tview.NewForm()
	form.AddInputField("Name", "", 64, nil, func(name string) {
		cardInfo.Name = name
	})
	form.AddInputField("Card number", "", 16, nil, func(number string) {
		cardInfo.CardNumber = number
	})
	// todo: how to validate value correctly?
	form.AddInputField("CVV", "", 3, nil, func(cvv string) {
		cvvNum, err := strconv.Atoi(cvv)
		if err != nil {
			v.ResultPage("Wrong value for cvv code")
		}
		cardInfo.CVV = int32(cvvNum)
	})
	form.AddInputField("Expiration Month", "", 2, nil, func(monthStr string) {
		month, err := strconv.Atoi(monthStr)
		if err != nil {
			v.ResultPage("Wrong value for expiration month")
		}
		cardInfo.ExpirationMonth = int32(month)
	})
	form.AddInputField("Expiration Year", "", 4, nil, func(yearStr string) {
		year, err := strconv.Atoi(yearStr)
		if err != nil {
			v.ResultPage("Wrong value for expiration month")
		}
		cardInfo.ExpirationYear = int32(year)
	})
	form.AddButton("Save", func() {
		err := v.itemsService.AddCardInfo(&cardInfo)
		resultMsg := "Successfully added new item!"
		if err != nil {
			resultMsg = fmt.Sprintf("Error happened on adding new item: %v", err)
		}
		v.ResultPage(resultMsg)
	})

	return form
}
