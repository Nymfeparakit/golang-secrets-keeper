package common

import (
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CredentialsToProto(pwd *dto.LoginPassword) *secrets.Password {
	dest := secrets.Password{
		Id:        pwd.ID,
		Name:      pwd.Name,
		Metadata:  pwd.Metadata,
		User:      pwd.User,
		Login:     pwd.Login,
		Password:  pwd.Password,
		UpdatedAt: timestamppb.New(pwd.UpdatedAt),
	}

	return &dest
}

func CardFromProto(crd *secrets.CardInfo) dto.CardInfo {
	itemDest := dto.BaseSecret{
		ID:        crd.Id,
		Name:      crd.Name,
		Metadata:  crd.Metadata,
		User:      crd.User,
		UpdatedAt: crd.UpdatedAt.AsTime(),
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
		Id:              crd.ID,
		Name:            crd.Name,
		Number:          crd.Number,
		ExpirationMonth: crd.ExpirationMonth,
		ExpirationYear:  crd.ExpirationYear,
		Cvv:             crd.CVV,
		Metadata:        crd.Metadata,
		User:            crd.User,
		UpdatedAt:       timestamppb.New(crd.UpdatedAt),
	}

	return &crdDest
}

func TextFromProto(txt *secrets.TextInfo) dto.TextInfo {
	itemDest := dto.BaseSecret{
		ID:        txt.Id,
		Name:      txt.Name,
		Metadata:  txt.Metadata,
		User:      txt.User,
		UpdatedAt: txt.UpdatedAt.AsTime(),
	}
	txtDest := dto.TextInfo{
		Text:       txt.Text,
		BaseSecret: itemDest,
	}

	return txtDest
}

func TextToProto(txt *dto.TextInfo) *secrets.TextInfo {
	txtDest := secrets.TextInfo{
		Id:        txt.ID,
		Name:      txt.Name,
		Text:      txt.Text,
		Metadata:  txt.Metadata,
		User:      txt.User,
		UpdatedAt: timestamppb.New(txt.UpdatedAt),
	}

	return &txtDest
}

func BinaryFromProto(bin *secrets.BinaryInfo) dto.BinaryInfo {
	itemDest := dto.BaseSecret{
		ID:        bin.Id,
		Name:      bin.Name,
		Metadata:  bin.Metadata,
		User:      bin.User,
		UpdatedAt: bin.UpdatedAt.AsTime(),
	}
	binDest := dto.BinaryInfo{
		Data:       bin.Data,
		BaseSecret: itemDest,
	}

	return binDest
}

func BinaryToProto(bin *dto.BinaryInfo) *secrets.BinaryInfo {
	dest := secrets.BinaryInfo{
		Id:        bin.ID,
		Name:      bin.Name,
		Data:      bin.Data,
		Metadata:  bin.Metadata,
		UpdatedAt: timestamppb.New(bin.UpdatedAt),
		User:      bin.User,
	}

	return &dest
}

func PasswordFromProto(pwd *secrets.Password) dto.LoginPassword {
	itemDest := dto.BaseSecret{
		ID:        pwd.Id,
		Name:      pwd.Name,
		Metadata:  pwd.Metadata,
		User:      pwd.User,
		UpdatedAt: pwd.UpdatedAt.AsTime(),
	}
	dest := dto.LoginPassword{
		Login:      pwd.Login,
		Password:   pwd.Password,
		BaseSecret: itemDest,
	}

	return dest
}