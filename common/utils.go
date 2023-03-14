package common

import (
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
)

func CardFromProto(crd *secrets.CardInfo) dto.CardInfo {
	itemDest := dto.BaseSecret{
		ID:       crd.Id,
		Name:     crd.Name,
		Metadata: crd.Metadata,
		User:     crd.User,
	}
	crdDest := dto.CardInfo{
		BaseSecret:      itemDest,
		Number:          crd.Number,
		CVV:             crd.Cvv,
		ExpirationMonth: crd.ExpirationMonth,
		ExpirationYear:  crd.ExpirationYear,
	}

	return crdDest
}

func CardToProto(crd *dto.CardInfo) *secrets.CardInfo {
	crdDest := secrets.CardInfo{
		Name:            crd.Name,
		Number:          crd.Number,
		ExpirationMonth: crd.ExpirationMonth,
		ExpirationYear:  crd.ExpirationYear,
		Cvv:             crd.CVV,
		Metadata:        crd.Metadata,
	}

	return &crdDest
}

func TextFromProto(txt *secrets.TextInfo) dto.TextInfo {
	itemDest := dto.BaseSecret{
		ID:       txt.Id,
		Name:     txt.Name,
		Metadata: txt.Metadata,
		User:     txt.User,
	}
	txtDest := dto.TextInfo{
		Text:       txt.Text,
		BaseSecret: itemDest,
	}

	return txtDest
}

func TextToProto(txt *dto.TextInfo) *secrets.TextInfo {
	txtDest := secrets.TextInfo{
		Name:     txt.Name,
		Text:     txt.Text,
		Metadata: txt.Metadata,
	}

	return &txtDest
}

func BinaryFromProto(bin *secrets.BinaryInfo) dto.BinaryInfo {
	itemDest := dto.BaseSecret{
		ID:       bin.Id,
		Name:     bin.Name,
		Metadata: bin.Metadata,
		User:     bin.User,
	}
	binDest := dto.BinaryInfo{
		Data:       bin.Data,
		BaseSecret: itemDest,
	}

	return binDest
}

func BinaryToProto(bin *dto.BinaryInfo) *secrets.BinaryInfo {
	dest := secrets.BinaryInfo{
		Name:     bin.Name,
		Data:     bin.Data,
		Metadata: bin.Metadata,
	}

	return &dest
}

func PasswordFromProto(pwd *secrets.Password) dto.LoginPassword {
	itemDest := dto.BaseSecret{
		ID:       pwd.Id,
		Name:     pwd.Name,
		Metadata: pwd.Metadata,
		User:     pwd.User,
	}
	dest := dto.LoginPassword{
		Login:      pwd.Login,
		Password:   pwd.Password,
		BaseSecret: itemDest,
	}

	return dest
}

func PasswordToProto(pwd *dto.LoginPassword) *secrets.Password {
	dest := secrets.Password{
		Name:     pwd.Name,
		Login:    pwd.Login,
		Password: pwd.Password,
		Metadata: pwd.Metadata,
	}

	return &dest
}
