package forms

import (
	"context"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/rivo/tview"
)

// TextInfoForm - form to perform save operations with certain TextInfo.
type TextInfoForm struct {
	instance *dto.TextInfo
	tview.Form
	itemService           AddSecretService
	retrieveUpdateService UpdateRetrieveTextService
	FormWithSaveAction
}

// NewTextInfoForm creates TextInfoForm object.
func NewTextInfoForm(service AddSecretService, updateService UpdateRetrieveTextService) *TextInfoForm {
	textInfo := &dto.TextInfo{}
	form := tview.NewForm()
	return &TextInfoForm{instance: textInfo, Form: *form, itemService: service, retrieveUpdateService: updateService}
}

// AddInputs adds input fields to form.
func (f *TextInfoForm) AddInputs() {
	f.AddInputField("Name", f.instance.Name, 64, nil, func(name string) {
		f.instance.Name = name
	})
	f.AddInputField("Text", f.instance.Text, 128, nil, func(text string) {
		f.instance.Text = text
	})
	f.AddInputField("Metadata", f.instance.Metadata, 128, nil, func(metadata string) {
		f.instance.Metadata = metadata
	})
}

// Save performs operation with TextInfo when the save button is clicked.
func (f *TextInfoForm) Save(ctx context.Context) error {
	var err error
	switch f.saveAction {
	case UPDATE:
		err = f.retrieveUpdateService.UpdateSecret(ctx, *f.instance)
	case CREATE:
		err = f.itemService.AddTextInfo(ctx, f.instance)
	}
	return err
}

// AddBtn adds button with certain label and selected function.
func (f *TextInfoForm) AddBtn(label string, selected func()) {
	f.AddButton(label, selected)
}

// SetSecret associates specific TextInfo with form.
func (f *TextInfoForm) SetSecret(ctx context.Context, id string) error {
	txt, err := f.retrieveUpdateService.GetSecretByID(ctx, id)
	f.instance = &txt
	return err
}
