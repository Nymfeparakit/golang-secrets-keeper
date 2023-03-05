package view

import (
	"errors"
	"fmt"
	"github.com/Nymfeparakit/gophkeeper/client/internal/services"
	"github.com/Nymfeparakit/gophkeeper/client/internal/view/forms"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/rs/zerolog/log"
)

type ItemsService interface {
	ListItems() (dto.ItemsList, error)
	AddPassword(password *dto.LoginPassword) error
	AddTextInfo(text *dto.TextInfo) error
	AddCardInfo(card *dto.CardInfo) error
}

type FlexWithHint struct {
	tview.Flex
}

func NewFlexWithHint(mainView tview.Primitive, hint string) *FlexWithHint {
	flex := tview.NewFlex()
	flex.AddItem(mainView, 0, 6, true)
	hintView := tview.NewTextView().SetText(hint)
	flex.AddItem(hintView, 0, 1, false)
	return &FlexWithHint{Flex: *flex}
}

type ItemsView struct {
	PagesView
	itemsService ItemsService
	items        dto.ItemsList
}

func NewItemsView(itemsService ItemsService) *ItemsView {
	pagesView := NewPagesView()
	return &ItemsView{PagesView: *pagesView, itemsService: itemsService}
}

func (v *ItemsView) AddItemPage(itemType dto.ItemType) {
	var form forms.SaveItemForm
	switch itemType {
	case dto.PASSWORD:
		form = forms.NewLoginPasswordForm(v.itemsService)
	case dto.TEXT:
		form = forms.NewTextInfoForm(v.itemsService)
	case dto.CARD:
		form = forms.NewCardInfoForm(v.itemsService)
	}
	forms.FillSaveItemForm(form, forms.CREATE, v.processSaveItemResult)
	v.pages.AddPage("Add item", form, true, true)
	err := v.app.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
}

func (v *ItemsView) ListItemsPage() {
	resultItems, err := v.itemsService.ListItems()
	if errors.Is(err, services.ErrTokenNotFound) {
		v.ResultPage("You are not authenticated - use 'login' command to set credentials")
		return
	}
	if err != nil {
		v.ResultPage(fmt.Sprintf("can not list items: %v", err.Error()))
		return
	}
	v.items = resultItems

	var pwdNames []string
	for _, pwd := range resultItems.Passwords {
		pwdNames = append(pwdNames, pwd.Name)
	}
	listPwdView := v.listItemsView(pwdNames, v.detailedLoginPasswordView)
	v.pages.AddPage("List passwords", listPwdView, true, true)

	var txtNames []string
	for _, txt := range resultItems.Texts {
		txtNames = append(txtNames, txt.Name)
	}
	listTxtView := v.listItemsView(txtNames, v.detailedLoginPasswordView)
	v.pages.AddPage("List texts", listTxtView, true, true)

	buttonsList := tview.NewList()
	// todo: return to list items page on backspace
	buttonsList.AddItem("Passwords", "", 0, func() {
		v.pages.SwitchToPage("List passwords")
	})
	buttonsList.AddItem("Texts", "", 0, func() {
		v.pages.SwitchToPage("List passwords")
	})
	v.pages.AddPage("List items", buttonsList, true, true)

	err = v.app.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
}

func (v *ItemsView) processSaveItemResult(saveAction forms.SaveAction, err error) {
	var resultMsg string
	switch saveAction {
	case forms.CREATE:
		resultMsg = "Successfully added new item!"
	case forms.UPDATE:
		resultMsg = "Successfully updated item!"
	}
	if err != nil {
		resultMsg = fmt.Sprintf("Error happened on saving item: %v", err)
	}
	v.ResultPage(resultMsg)
}

func (v *ItemsView) listItemsView(
	itemsNames []string,
	getDetailedItemView func(i int) *tview.Flex,
) *tview.Flex {
	listView := tview.NewList()
	for _, name := range itemsNames {
		listView.AddItem(name, "", 0, nil)
	}

	flex := tview.NewFlex()
	flex.AddItem(listView, 0, 1, true)
	listView.SetSelectedFunc(func(i int, s string, s2 string, r rune) {
		flex.Clear()
		flex.AddItem(listView, 0, 1, true)
		flex.AddItem(getDetailedItemView(i), 0, 4, false)
	})
	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.pages.SwitchToPage("List items")
			return nil
		}
		return event
	})

	return flex
}

func (v *ItemsView) detailedLoginPasswordView(i int) *tview.Flex {
	pwd := v.items.Passwords[i]
	flex := tview.NewFlex().SetDirection(tview.FlexRow)
	flex.AddItem(tview.NewTextView().SetText(pwd.Name).SetLabel("Name"), 0, 1, false)
	flex.AddItem(tview.NewTextView().SetText(pwd.Login).SetLabel("Login"), 0, 1, false)
	// todo: don't show password
	flex.AddItem(tview.NewTextView().SetText(pwd.Password).SetLabel("Password"), 0, 1, false)
	flex.AddItem(tview.NewTextView().SetText(pwd.Metadata).SetLabel("Metadata"), 0, 1, false)

	return flex
}

func (v *ItemsView) detailedTextInfoView(i int) *tview.Flex {
	text := v.items.Texts[i]
	flex := tview.NewFlex()
	flex.AddItem(tview.NewTextView().SetText(text.Name).SetLabel("Name"), 0, 1, false)
	flex.AddItem(tview.NewTextView().SetText(text.Text).SetLabel("Text"), 0, 1, false)
	flex.AddItem(tview.NewTextView().SetText(text.Metadata).SetLabel("Metadata"), 0, 1, false)

	return flex
}
