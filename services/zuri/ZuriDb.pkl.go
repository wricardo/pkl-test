// Code generated from Pkl module `Zuri`. DO NOT EDIT.
package zuri

import "bitbucket.org/zetaactions/pkltest/lib"

type ZuriDb interface {
	lib.ConnectionDetails
}

var _ ZuriDb = (*ZuriDbImpl)(nil)

type ZuriDbImpl struct {
	Username string `pkl:"username"`

	Password string `pkl:"password"`

	Host string `pkl:"host"`

	Port int `pkl:"port"`

	DbName string `pkl:"dbName"`
}

func (rcv *ZuriDbImpl) GetUsername() string {
	return rcv.Username
}

func (rcv *ZuriDbImpl) GetPassword() string {
	return rcv.Password
}

func (rcv *ZuriDbImpl) GetHost() string {
	return rcv.Host
}

func (rcv *ZuriDbImpl) GetPort() int {
	return rcv.Port
}

func (rcv *ZuriDbImpl) GetDbName() string {
	return rcv.DbName
}
