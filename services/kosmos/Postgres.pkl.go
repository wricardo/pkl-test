// Code generated from Pkl module `Kosmos`. DO NOT EDIT.
package kosmos

import "bitbucket.org/zetaactions/pkltest/lib"

type Postgres interface {
	lib.ConnectionDetails
}

var _ Postgres = (*PostgresImpl)(nil)

type PostgresImpl struct {
	Host string `pkl:"host"`

	DbName string `pkl:"dbName"`

	Username string `pkl:"username"`

	Password string `pkl:"password"`

	Port int `pkl:"port"`
}

func (rcv *PostgresImpl) GetHost() string {
	return rcv.Host
}

func (rcv *PostgresImpl) GetDbName() string {
	return rcv.DbName
}

func (rcv *PostgresImpl) GetUsername() string {
	return rcv.Username
}

func (rcv *PostgresImpl) GetPassword() string {
	return rcv.Password
}

func (rcv *PostgresImpl) GetPort() int {
	return rcv.Port
}
