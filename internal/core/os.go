package core

import (
	"fmt"
)

type OS interface {
	CreateUser(username string) string
	CheckUser(username string) string
	GroupUser(username string, group string) string
	InstallPackage(packageName string) string
	RemovePackage(packageName string) string
	UpdatePackages() string
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

func (os DebianFamily) InstallPackage(packageName string) string {
	return fmt.Sprintf("apt-get install -y %s", packageName)
}

func (os DebianFamily) RemovePackage(packageName string) string {
	return fmt.Sprintf("apt-get remove -y %s", packageName)
}

func (os DebianFamily) UpdatePackages() string {
	return "apt-get update && apt-get upgrade -y"
}

type Ubuntu struct{ DebianFamily }
type Debian struct{ DebianFamily }

type FedoraFamily struct{}

func (os FedoraFamily) CreateUser(username string) string {
	return fmt.Sprintf("useradd %s", username)
}

func (os FedoraFamily) CheckUser(username string) string {
	return fmt.Sprintf("id -u %s", username)
}

func (os FedoraFamily) GroupUser(username string, group string) string {
	return fmt.Sprintf("usermod -aG %s %s", group, username)
}

func (os FedoraFamily) InstallPackage(packageName string) string {
	return fmt.Sprintf("dnf install -y %s", packageName)
}

func (os FedoraFamily) RemovePackage(packageName string) string {
	return fmt.Sprintf("dnf remove -y %s", packageName)
}

func (os FedoraFamily) UpdatePackages() string {
	return "dnf update -y"
}

type Fedora struct{ FedoraFamily }
type RedHat struct{ FedoraFamily }
