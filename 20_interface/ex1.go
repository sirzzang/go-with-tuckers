package main

type OurDB struct {
	Name string
}

type DB interface {
	GetData() string
	WriteData(data string) string
	Close() error
}
