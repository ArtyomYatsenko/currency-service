package repository

// Имитация базы данных

type MyBase map[string]string

func NewMyBase() MyBase {
	return make(MyBase)
}
