package example2

type User struct {
	Name string
	ID   int

	paswd string
}

type auser struct {
	Name string
	ID   int
}

type Manager struct {
	auser
	Title string
}
