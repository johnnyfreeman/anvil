package main

import "testing"

func Test_UbuntuCreateUser(t *testing.T) {
	os := Ubuntu{}
	command := os.CreateUser("john")
	if command != "adduser john" {
		t.Fatalf("malformed command: %v", os)
	}
}

func Test_DebianCreateUser(t *testing.T) {
	os := Debian{}
	command := os.CreateUser("john")
	if command != "adduser john" {
		t.Fatalf("malformed command: %v", os)
	}
}

func Test_FedoraCreateUser(t *testing.T) {
	os := Fedora{}
	command := os.CreateUser("john")
	if command != "useradd john" {
		t.Fatalf("malformed command: %v", os)
	}
}

func Test_RedHatCreateUser(t *testing.T) {
	os := RedHat{}
	command := os.CreateUser("john")
	if command != "useradd john" {
		t.Fatalf("malformed command: %v", os)
	}
}
