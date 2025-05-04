package core

import (
	"context"
	"fmt"
	"strings"
)

type OSInfo struct {
	ID       string
	IDLike   []string
	Version  string
	Pretty   string
	Detected OS
}

func DetectOS(ctx context.Context, executor Executor) (*OSInfo, error) {
	output, err := executor.Execute(ctx, "cat /etc/os-release", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to exec on target: %w", err)
	}

	info := OSInfo{}

	for line := range strings.SplitSeq(output, "\n") {
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := parts[0]
		value := strings.Trim(parts[1], `"`)

		switch key {
		case "ID":
			info.ID = value
		case "ID_LIKE":
			info.IDLike = strings.Fields(value)
		case "VERSION_ID":
			info.Version = value
		case "PRETTY_NAME":
			info.Pretty = value
		}
	}

	// Detection logic
	switch info.ID {
	case "ubuntu":
		info.Detected = Debian{}
	case "debian":
		info.Detected = Debian{}
	case "fedora":
		info.Detected = Fedora{}
	case "rhel", "redhat":
		info.Detected = RedHat{}
	default:
		for _, like := range info.IDLike {
			if like == "debian" {
				info.Detected = Debian{}
				break
			}
			if like == "rhel" || like == "fedora" {
				info.Detected = RedHat{}
				break
			}
		}
	}

	return &info, nil
}
