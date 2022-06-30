package store

type StoreMock struct {
	WriteMock func(data interface{}) error
	ReadMock func(data interface{}) error
}

func (m StoreMock) Read(data interface{}) error { 
	return m.ReadMock(data)
}
func (m StoreMock) Write(data interface{}) error { 
	return m.WriteMock(data)
}




