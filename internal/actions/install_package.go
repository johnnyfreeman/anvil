package actions

import (
	"context"

	"github.com/johnnyfreeman/anvil/internal/core"
)

type InstallPackageOpts struct {
	PackageName string
	Update      bool
}

type InstallPackageOptsFunc func(*InstallPackageOpts)

func WithUpdate() InstallPackageOptsFunc {
	return func(o *InstallPackageOpts) {
		o.Update = true
	}
}

type InstallPackage struct {
	InstallPackageOpts
}

func DefaultInstallPackageOpts() InstallPackageOpts {
	return InstallPackageOpts{
		PackageName: "",
		Update:      false,
	}
}

func NewInstallPackage(packageName string, opts ...InstallPackageOptsFunc) *InstallPackage {
	o := DefaultInstallPackageOpts()
	o.PackageName = packageName
	for _, fn := range opts {
		fn(&o)
	}
	return &InstallPackage{
		InstallPackageOpts: o,
	}
}

func (a InstallPackage) Handle(ctx context.Context, ex core.Executor, os core.OS, observer core.ActionObserver) error {
	return core.WithObserver(observer, func() error {
		if a.Update {
			_, err := ex.Execute(ctx, os.UpdatePackages(), observer)
			if err != nil {
				return err
			}
		}

		_, err := ex.Execute(ctx, os.InstallPackage(a.PackageName), observer)
		return err
	})
}

var _ core.Action = (*InstallPackage)(nil)