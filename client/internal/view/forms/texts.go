package forms

import (
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/rivo/tview"
)

type TextInfoForm struct {
	instance *dto.TextInfo
	tview.Form
	itemService           AddSecretService
	retrieveUpdateService UpdateRetrieveTextService
	FormWithSaveAction
}

func NewTextInfoForm(service AddSecretService, updateService UpdateRetrieveTextService) *TextInfoForm {
	textInfo := &dto.TextInfo{}
	form := tview.NewForm()
	return &TextInfoForm{instance: textInfo, Form: *form, itemService: service, retrieveUpdateService: updateService}
}

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

func (f *TextInfoForm) Save() error {
	var err error
	switch f.saveAction {
	case UPDATE:
		err = f.retrieveUpdateService.UpdateSecret(*f.instance)
	case CREATE:
		err = f.itemService.AddTextInfo(f.instance)
	}
	return err
}

func (f *TextInfoForm) AddBtn(label string, selected func()) {
	f.AddButton(label, selected)
}

func (f *TextInfoForm) SetSecret(id string) error {
	txt, err := f.retrieveUpdateService.GetSecretByID(id)
	f.instance = &txt
	return err
}
