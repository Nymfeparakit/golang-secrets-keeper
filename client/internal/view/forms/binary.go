package forms

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/rivo/tview"
	"time"
)

// BinaryInfoForm - form to perform save operations with certain BinaryInfo.
type BinaryInfoForm struct {
	instance *dto.BinaryInfo
	tview.Form
	addService            AddSecretService
	retrieveUpdateService UpdateRetrieveBinaryService
	FormWithSaveAction
	binFilePath string
}

// NewBinaryInfoForm creates BinaryInfoForm object.
func NewBinaryInfoForm(service AddSecretService, updateService UpdateRetrieveBinaryService) *BinaryInfoForm {
	bin := &dto.BinaryInfo{}
	form := tview.NewForm()
	return &BinaryInfoForm{instance: bin, Form: *form, addService: service, retrieveUpdateService: updateService}
}

// AddInputs adds input fields to form.
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

// Save performs operation with BinaryInfo when the save button is clicked.
func (f *BinaryInfoForm) Save(ctx context.Context) error {
	if f.binFilePath != "" {
		binData, err := readBinaryFile(f.binFilePath)
		if err != nil {
			return fmt.Errorf("can't read file with binary data: %v", err)
		}
		f.instance.Data = base64.StdEncoding.EncodeToString(binData)
	}
	f.instance.UpdatedAt = time.Now().UTC()
	var err error
	switch f.saveAction {
	case UPDATE:
		err = f.retrieveUpdateService.UpdateSecret(ctx, *f.instance)
	case CREATE:
		err = f.addService.AddBinaryInfo(ctx, f.instance)
	}
	return err
}

// AddBtn adds button with certain label and selected function.
func (f *BinaryInfoForm) AddBtn(label string, selected func()) {
	f.AddButton(label, selected)
}

// SetSecret associates specific BinaryInfo with form.
func (f *BinaryInfoForm) SetSecret(ctx context.Context, id string) error {
	bin, err := f.retrieveUpdateService.GetSecretByID(ctx, id)
	f.instance = &bin
	return err
}
