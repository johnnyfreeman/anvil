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
	StartService(serviceName string) string
	StopService(serviceName string) string
	EnableService(serviceName string) string
	RestartService(serviceName string) string
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

func (os DebianFamily) StartService(serviceName string) string {
	return fmt.Sprintf("systemctl start %s", serviceName)
}

func (os DebianFamily) StopService(serviceName string) string {
	return fmt.Sprintf("systemctl stop %s", serviceName)
}

func (os DebianFamily) EnableService(serviceName string) string {
	return fmt.Sprintf("systemctl enable %s", serviceName)
}

func (os DebianFamily) RestartService(serviceName string) string {
	return fmt.Sprintf("systemctl restart %s", serviceName)
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

func (os FedoraFamily) StartService(serviceName string) string {
	return fmt.Sprintf("systemctl start %s", serviceName)
}

func (os FedoraFamily) StopService(serviceName string) string {
	return fmt.Sprintf("systemctl stop %s", serviceName)
}

func (os FedoraFamily) EnableService(serviceName string) string {
	return fmt.Sprintf("systemctl enable %s", serviceName)
}

func (os FedoraFamily) RestartService(serviceName string) string {
	return fmt.Sprintf("systemctl restart %s", serviceName)
}

type Fedora struct{ FedoraFamily }
type RedHat struct{ FedoraFamily }
