package entities

const (
	RootEndpoint		= "/"
	UsersEndpoint		= "/users"
)

var Routes = []Route {
	{Name: "RootEndpoint", Path: RootEndpoint},
	{Name: "UsersEndpoint", Path: UsersEndpoint},
}

type Route struct {
	Name	string
	Path	string
}
