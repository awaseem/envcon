package main

type aesCrypMock struct {
	err        error
	key        []byte
	salt       []byte
	encryptRtn string
	decryptRtn string
}

func (a *aesCrypMock) keyGen(pass []byte) (key []byte, salt []byte, err error) {
	return a.key, a.salt, a.err
}
func (a *aesCrypMock) keyGenWithSalt(pass []byte, salt []byte) (key []byte) {
	return a.key
}
func (a *aesCrypMock) encrypt(key, text []byte) (string, error) {
	return a.encryptRtn, a.err
}
func (a *aesCrypMock) decrypt(key []byte, cryptoText string) (string, error) {
	return a.decryptRtn, a.err
}
