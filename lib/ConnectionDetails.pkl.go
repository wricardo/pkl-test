// Code generated from Pkl module `Postgres`. DO NOT EDIT.
package lib

type ConnectionDetails interface {
	GetUsername() string

	GetPassword() string

	GetHost() string

	GetPort() int

	GetDbName() string
}
