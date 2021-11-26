package models

type LoremLocalData interface {
	Words() []string
	Suplemental() []string
}

type Lorem struct{}
