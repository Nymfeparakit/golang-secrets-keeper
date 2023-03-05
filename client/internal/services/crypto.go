package services

type KeysStorage interface {
	GetUserKey() ([]byte, error)
}

type CryptoService struct {
}

//func encryptItem(request interface{}) {
//	reqValue := reflect.ValueOf(request)
//	reqObjType := reqValue.Type()
//
//	for i := 0; i <= reqValue.NumField(); i++ {
//		field := reqObjType.Field(i)
//	}
//
//	binary.Write(a, binary.LittleEndian, myInt)
//}
