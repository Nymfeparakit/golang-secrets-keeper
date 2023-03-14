package forms

import (
	"encoding/base64"
	"fmt"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/rivo/tview"
)

type BinaryInfoForm struct {
	instance *dto.BinaryInfo
	tview.Form
	addService            AddSecretService
	retrieveUpdateService UpdateRetrieveBinaryService
	FormWithSaveAction
	binFilePath string
}

func NewBinaryInfoForm(service AddSecretService, updateService UpdateRetrieveBinaryService) *BinaryInfoForm {
	bin := &dto.BinaryInfo{}
	form := tview.NewForm()
	return &BinaryInfoForm{instance: bin, Form: *form, addService: service, retrieveUpdateService: updateService}
}

func (f *BinaryInfoForm) AddInputs() {
	f.AddInputField("Name", f.instance.Name, 64, nil, func(name string) {
		f.instance.Name = name
	})
	f.AddInputField("Metadata", f.instance.Metadata, 128, nil, func(metadata string) {
		f.instance.Metadata = metadata
	})
	f.AddInputField("Binary file path", "", 32, nil, func(path string) {
		f.binFilePath = path
	})
}

func (f *BinaryInfoForm) Save() error {
	if f.binFilePath != "" {
		binData, err := readBinaryFile(f.binFilePath)
		if err != nil {
			return fmt.Errorf("can't read file with binary data: %v", err)
		}
		f.instance.Data = base64.StdEncoding.EncodeToString(binData)
	}
	var err error
	switch f.saveAction {
	case UPDATE:
		err = f.retrieveUpdateService.UpdateSecret(*f.instance)
	case CREATE:
		err = f.addService.AddBinaryInfo(f.instance)
	}
	return err
}

func (f *BinaryInfoForm) AddBtn(label string, selected func()) {
	f.AddButton(label, selected)
}

func (f *BinaryInfoForm) SetSecret(id string) error {
	bin, err := f.retrieveUpdateService.GetSecretByID(id)
	f.instance = &bin
	return err
}
