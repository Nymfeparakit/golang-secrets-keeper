package forms

import (
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/rivo/tview"
)

type TextInfoForm struct {
	textInfo *dto.TextInfo
	tview.Form
	itemService ItemsService
}

func NewTextInfoForm(service ItemsService) *TextInfoForm {
	textInfo := &dto.TextInfo{}
	form := tview.NewForm()
	return &TextInfoForm{textInfo: textInfo, Form: *form, itemService: service}
}

func (f *TextInfoForm) AddInputs() {
	f.AddInputField("Name", "", 64, nil, func(name string) {
		f.textInfo.Name = name
	})
	f.AddInputField("Text", "", 128, nil, func(text string) {
		f.textInfo.Text = text
	})
}

func (f *TextInfoForm) Save(saveAction SaveAction) error {
	var err error
	switch saveAction {
	case UPDATE:
		err = f.itemService.AddTextInfo(f.textInfo)
	case CREATE:
		err = f.itemService.AddTextInfo(f.textInfo)
	}
	return err
}

func (f *TextInfoForm) AddBtn(label string, selected func()) {
	f.AddButton(label, selected)
}
