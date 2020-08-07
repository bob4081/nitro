package nitro

import (
	"strings"

	"github.com/craftcms/nitro/internal/validate"
)

func MountDir(name, source, target string) (*Action, error) {
	if err := validate.MachineName(name); err != nil {
		return nil, err
	}
	if err := validate.Path(source); err != nil {
		return nil, err
	}

	// make sure target has a forward slash
	if !strings.HasPrefix(target, "/") {
		target = "/" + target
	}

	if strings.HasSuffix(source, "/") {
		source = strings.TrimRight(source, "/")
	}

	return &Action{
		Type:       "mount",
		UseSyscall: false,
		Args:       []string{"mount", source, name + ":" + target},
	}, nil
}
