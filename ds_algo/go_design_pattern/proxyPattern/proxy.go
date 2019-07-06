package proxyPattern

type UserFinder interface {
	FindUser(id int32) (User, error)
}

type User struct {
	ID int32
}

type UserList []User

type UserListProxy struct {
	SomeDatabase              UserList
	StackCache                UserList
	StackCapacity             int
	DidDidLastSearchUsedCache bool
}
