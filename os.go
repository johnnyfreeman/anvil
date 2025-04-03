package main

import (
	"fmt"
)

type OS interface {
	CreateUser(username string) string
	CheckUser(username string) string
	GroupUser(username string, group string) string
}

type DebianFamily struct{}

func (os DebianFamily) CreateUser(username string) string {
	return fmt.Sprintf("adduser %s", username)
}

func (os DebianFamily) CheckUser(username string) string {
	return fmt.Sprintf("id -u %s", username)
}

func (os DebianFamily) GroupUser(username string, group string) string {
	return fmt.Sprintf("usermod -aG %s %s", group, username)
}

type Ubuntu struct{ DebianFamily }
type Debian struct{ DebianFamily }

type FedoraFamily struct{}

func (f FedoraFamily) CreateUser(username string) string {
	return fmt.Sprintf("useradd %s", username)
}

func (os FedoraFamily) CheckUser(username string) string {
	return fmt.Sprintf("id -u %s", username)
}

func (os FedoraFamily) GroupUser(username string, group string) string {
	return fmt.Sprintf("usermod -aG %s %s", group, username)
}

type Fedora struct{ FedoraFamily }
type RedHat struct{ FedoraFamily }
