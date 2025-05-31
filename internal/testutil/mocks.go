package testutil

// MockOS is a test implementation of core.OS
type MockOS struct{}

func (o *MockOS) CreateUser(username string) string {
	return "create-user " + username
}

func (o *MockOS) CheckUser(username string) string {
	return "check-user " + username
}

func (o *MockOS) GroupUser(username string, group string) string {
	return "group-user " + username + " " + group
}

func (o *MockOS) InstallPackage(packageName string) string {
	return "install " + packageName
}

func (o *MockOS) RemovePackage(packageName string) string {
	return "remove " + packageName
}

func (o *MockOS) UpdatePackages() string {
	return "update-packages"
}

func (o *MockOS) StartService(serviceName string) string {
	return "start " + serviceName
}

func (o *MockOS) StopService(serviceName string) string {
	return "stop " + serviceName
}

func (o *MockOS) EnableService(serviceName string) string {
	return "enable " + serviceName
}

func (o *MockOS) RestartService(serviceName string) string {
	return "restart " + serviceName
}

// MockObserver is a test implementation of core.ActionObserver
type MockObserver struct {
	StartCalled       bool
	EndCalled         bool
	OutputCalled      bool
	ActionStartCalled bool
	ActionEndCalled   bool
	Commands          []string
	Outputs           []string
}

func (o *MockObserver) OnExecutionStart(command string) error {
	o.StartCalled = true
	o.Commands = append(o.Commands, command)
	return nil
}

func (o *MockObserver) OnExecutionEnd() error {
	o.EndCalled = true
	return nil
}

func (o *MockObserver) OnExecutionOutput(output string) error {
	o.OutputCalled = true
	o.Outputs = append(o.Outputs, output)
	return nil
}

func (o *MockObserver) OnActionStart() error {
	o.ActionStartCalled = true
	return nil
}

func (o *MockObserver) OnActionEnd() error {
	o.ActionEndCalled = true
	return nil
}