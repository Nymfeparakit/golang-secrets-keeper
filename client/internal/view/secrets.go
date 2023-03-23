package view

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/Nymfeparakit/gophkeeper/client/internal/services"
	"github.com/Nymfeparakit/gophkeeper/client/internal/view/forms"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/rs/zerolog/log"
)

// ListAddSecretsService - service for adding and listing secrets.
type ListAddSecretsService interface {
	ListSecrets(ctx context.Context) (dto.SecretsList, error)
	AddCredentials(ctx context.Context, password *dto.LoginPassword) error
	AddTextInfo(ctx context.Context, text *dto.TextInfo) error
	AddCardInfo(ctx context.Context, card *dto.CardInfo) error
	AddBinaryInfo(ctx context.Context, bin *dto.BinaryInfo) error
}

// UpdateRetrievePasswordService - service for retrieving/updating certain LoginPassword instance.
type UpdateRetrievePasswordService interface {
	GetSecretByID(ctx context.Context, id string) (dto.LoginPassword, error)
	UpdateSecret(ctx context.Context, secret dto.LoginPassword) error
	DeleteSecret(ctx context.Context, id string) error
}

// UpdateRetrieveCardService - service for retrieving/updating certain CardInfo instance.
type UpdateRetrieveCardService interface {
	GetSecretByID(ctx context.Context, id string) (dto.CardInfo, error)
	UpdateSecret(ctx context.Context, secret dto.CardInfo) error
	DeleteSecret(ctx context.Context, id string) error
}

// UpdateRetrieveTextService - service for retrieving/updating certain TextInfo instance.
type UpdateRetrieveTextService interface {
	GetSecretByID(ctx context.Context, id string) (dto.TextInfo, error)
	UpdateSecret(ctx context.Context, secret dto.TextInfo) error
	DeleteSecret(ctx context.Context, id string) error
}

// UpdateRetrieveBinaryService - service for retrieving/updating certain BinaryInfo instance.
type UpdateRetrieveBinaryService interface {
	GetSecretByID(ctx context.Context, id string) (dto.BinaryInfo, error)
	UpdateSecret(ctx context.Context, secret dto.BinaryInfo) error
	DeleteSecret(ctx context.Context, id string) error
}

// NewFlexWithHint creates new flex with main view and text view with specified hint.
func NewFlexWithHint(mainView tview.Primitive, hint string) *tview.Flex {
	flex := tview.NewFlex().SetDirection(tview.FlexRow)
	flex.AddItem(mainView, 0, 6, true)
	hintView := tview.NewTextView().SetText(hint)
	flex.AddItem(hintView, 0, 1, false)
	return flex
}

// SecretsView - view for showing pages for creating, updating and listing secrets.
type SecretsView struct {
	PagesView
	secretsService     ListAddSecretsService
	pwdInstanceService UpdateRetrievePasswordService
	crdInstanceService UpdateRetrieveCardService
	txtInstanceService UpdateRetrieveTextService
	binInstanceService UpdateRetrieveBinaryService
	secrets            dto.SecretsList
}

// NewSecretsView creates new SecretsView object.
func NewSecretsView(
	secretsService ListAddSecretsService,
	pwdInstanceService UpdateRetrievePasswordService,
	crdInstanceService UpdateRetrieveCardService,
	txtInstanceService UpdateRetrieveTextService,
	binInstanceService UpdateRetrieveBinaryService,
) *SecretsView {
	pagesView := NewPagesView()
	return &SecretsView{
		PagesView:          *pagesView,
		secretsService:     secretsService,
		pwdInstanceService: pwdInstanceService,
		crdInstanceService: crdInstanceService,
		txtInstanceService: txtInstanceService,
		binInstanceService: binInstanceService,
	}
}

// AddSecretPage shows page to add new secret.
func (v *SecretsView) AddSecretPage(ctx context.Context, itemType dto.SecretType) {
	form := v.formFromSecretType(itemType)
	form, err := forms.FillSaveItemForm(ctx, form, forms.CREATE, "", v.processSaveSecretResult)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
	v.pages.AddPage("Add item", form, true, true)
	err = v.app.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
}

// UpdateSecretPage shows page to update exising secret.
func (v *SecretsView) UpdateSecretPage(ctx context.Context, itemType dto.SecretType, secretID string) {
	form := v.formFromSecretType(itemType)
	form, err := forms.FillSaveItemForm(ctx, form, forms.UPDATE, secretID, v.processSaveSecretResult)
	if err != nil {
		v.ResultPage(fmt.Sprintf("can't update secret: %v", err))
		return
	}
	v.pages.AddPage("Update item", form, true, true)
	err = v.app.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
}

// DeleteSecretPage deletes the secret and then shows to user result page.
func (v *SecretsView) DeleteSecretPage(ctx context.Context, itemType dto.SecretType, secretID string) {
	var err error
	switch itemType {
	case dto.PASSWORD:
		err = v.pwdInstanceService.DeleteSecret(ctx, secretID)
	case dto.TEXT:
		err = v.txtInstanceService.DeleteSecret(ctx, secretID)
	case dto.CARD:
		err = v.crdInstanceService.DeleteSecret(ctx, secretID)
	case dto.BINARY:
		err = v.binInstanceService.DeleteSecret(ctx, secretID)
	}
	if err != nil {
		v.ResultPage(fmt.Sprintf("Error on deleting secret: %v", err))
		return
	}
	v.ResultPage(fmt.Sprintf("Successfully deleted secret!"))
}

// ListSecretsPage shows page to list all user's secrets.
func (v *SecretsView) ListSecretsPage(ctx context.Context) {
	resultSecrets, err := v.secretsService.ListSecrets(ctx)
	if errors.Is(err, services.ErrTokenNotFound) {
		v.ResultPage("You are not authenticated - use 'login' command to set credentials")
		return
	}
	if err != nil {
		v.ResultPage(fmt.Sprintf("can not list secrets: %v", err.Error()))
		return
	}
	v.secrets = resultSecrets

	var pwdNames []string
	for _, pwd := range resultSecrets.Passwords {
		pwdNames = append(pwdNames, pwd.Name)
	}
	listPwdView := v.listSecretsView(pwdNames, v.detailedLoginPasswordView)
	v.pages.AddPage("List passwords", listPwdView, true, true)

	var txtNames []string
	for _, txt := range resultSecrets.Texts {
		txtNames = append(txtNames, txt.Name)
	}
	listTxtView := v.listSecretsView(txtNames, v.detailedTextInfoView)
	v.pages.AddPage("List texts", listTxtView, true, true)

	var cardNames []string
	for _, crd := range resultSecrets.Cards {
		cardNames = append(cardNames, crd.Name)
	}
	listCardView := v.listSecretsView(cardNames, v.detailedCardInfoView)
	v.pages.AddPage("List cards", listCardView, true, true)

	var binNames []string
	for _, bin := range resultSecrets.Bins {
		binNames = append(binNames, bin.Name)
	}
	listBinView := v.listSecretsView(binNames, v.detailedBinaryInfoView)
	v.pages.AddPage("List binaries", listBinView, true, true)

	buttonsList := tview.NewList()
	buttonsList.AddItem("Passwords", "", 0, func() {
		v.pages.SwitchToPage("List passwords")
	})
	buttonsList.AddItem("Texts", "", 0, func() {
		v.pages.SwitchToPage("List texts")
	})
	buttonsList.AddItem("Cards", "", 0, func() {
		v.pages.SwitchToPage("List cards")
	})
	buttonsList.AddItem("Binaries", "", 0, func() {
		v.pages.SwitchToPage("List binaries")
	})
	v.pages.AddPage("List secrets", buttonsList, true, true)

	err = v.app.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
}

func (v *SecretsView) processSaveSecretResult(saveAction forms.SaveAction, err error) {
	var resultMsg string
	switch saveAction {
	case forms.CREATE:
		resultMsg = "Successfully added new secret!"
	case forms.UPDATE:
		resultMsg = "Successfully updated secret!"
	}
	if err != nil {
		resultMsg = fmt.Sprintf("Error happened on saving secret: %v", err)
	}
	v.ResultPage(resultMsg)
}

func (v *SecretsView) listSecretsView(
	secretsNames []string,
	getDetailedSecretView func(i int) *tview.Flex,
) *tview.Flex {
	listView := tview.NewList()
	for _, name := range secretsNames {
		listView.AddItem(name, "", 0, nil)
	}

	flex := tview.NewFlex()
	flex.AddItem(listView, 0, 1, true)
	listView.SetSelectedFunc(func(i int, s string, s2 string, r rune) {
		flex.Clear()
		flex.AddItem(listView, 0, 1, true)
		flex.AddItem(getDetailedSecretView(i), 0, 4, false)
	})
	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			v.pages.SwitchToPage("List secrets")
			return nil
		}
		return event
	})
	hintFlex := NewFlexWithHint(flex, "press ENTER to choose secret; press ESC to exit")

	return hintFlex
}

func (v *SecretsView) detailedLoginPasswordView(i int) *tview.Flex {
	pwd := v.secrets.Passwords[i]
	flex := tview.NewFlex().SetDirection(tview.FlexRow)
	flex.AddItem(tview.NewTextView().SetText(pwd.Name).SetLabel("Name:"), 0, 1, false)
	flex.AddItem(tview.NewTextView().SetText(pwd.Login).SetLabel("Login:"), 0, 1, false)
	flex.AddItem(tview.NewTextView().SetText(pwd.Password).SetLabel("Password:"), 0, 1, false)
	flex.AddItem(tview.NewTextView().SetText(pwd.Metadata).SetLabel("Metadata:"), 0, 1, false)

	return flex
}

func (v *SecretsView) detailedTextInfoView(i int) *tview.Flex {
	text := v.secrets.Texts[i]
	flex := tview.NewFlex().SetDirection(tview.FlexRow)
	flex.AddItem(tview.NewTextView().SetText(text.Name).SetLabel("Name:"), 0, 1, false)
	flex.AddItem(tview.NewTextView().SetText(text.Text).SetLabel("Text:"), 0, 1, false)
	flex.AddItem(tview.NewTextView().SetText(text.Metadata).SetLabel("Metadata:"), 0, 1, false)

	return flex
}

func (v *SecretsView) detailedCardInfoView(i int) *tview.Flex {
	card := v.secrets.Cards[i]
	flex := tview.NewFlex().SetDirection(tview.FlexRow)
	flex.AddItem(tview.NewTextView().SetText(card.Name).SetLabel("Name:"), 0, 1, false)
	flex.AddItem(tview.NewTextView().SetText(card.Number).SetLabel("Number:"), 0, 1, false)
	flex.AddItem(tview.NewTextView().SetText(card.CVV).SetLabel("CVV:"), 0, 1, false)
	flex.AddItem(tview.NewTextView().SetText(card.ExpirationMonth).SetLabel("Expiration month:"), 0, 1, false)
	flex.AddItem(tview.NewTextView().SetText(card.ExpirationYear).SetLabel("Expiration year:"), 0, 1, false)
	flex.AddItem(tview.NewTextView().SetText(card.Metadata).SetLabel("Metadata:"), 0, 1, false)

	return flex
}

func (v *SecretsView) detailedBinaryInfoView(i int) *tview.Flex {
	bin := v.secrets.Bins[i]
	flex := tview.NewFlex().SetDirection(tview.FlexRow)
	flex.AddItem(tview.NewTextView().SetText(bin.Name).SetLabel("Name:"), 0, 1, false)
	data, err := base64.StdEncoding.DecodeString(bin.Data)
	if err != nil {
		log.Fatal().Err(err).Msg("can't decode binary data")
	}
	flex.AddItem(tview.NewTextView().SetText(string(data)).SetLabel("Data:"), 0, 1, false)
	flex.AddItem(tview.NewTextView().SetText(bin.Metadata).SetLabel("Metadata:"), 0, 1, false)

	return flex
}

func (v *SecretsView) formFromSecretType(itemType dto.SecretType) forms.SaveItemForm {
	var form forms.SaveItemForm
	switch itemType {
	case dto.PASSWORD:
		form = forms.NewLoginPasswordForm(v.secretsService, v.pwdInstanceService)
	case dto.TEXT:
		form = forms.NewTextInfoForm(v.secretsService, v.txtInstanceService)
	case dto.CARD:
		form = forms.NewCardInfoForm(v.secretsService, v.crdInstanceService)
	case dto.BINARY:
		form = forms.NewBinaryInfoForm(v.secretsService, v.binInstanceService)
	}
	return form
}
