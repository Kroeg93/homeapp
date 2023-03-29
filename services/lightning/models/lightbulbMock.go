package models

type LightsMock = map[string]LightMock

type LightMock struct {
	Name        string
	Type        string
	ProductName string
}
