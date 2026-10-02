package applauncher

import (
	"fmt"
	"os/exec"
	"strings"

	"Jarvis/tools/registry"
)

type Tool struct{}

func NewTool() *Tool {
	return &Tool{}
}

func (t *Tool) Name() string {
	return "open_app"
}

func (t *Tool) Description() string {
	return "Open an allowed application on the Windows computer."
}

func (t *Tool) Parameters() []registry.Parameter {
	return []registry.Parameter{
		{
			Name:        "name",
			Type:        "string",
			Description: "Application to open, such as VS Code, Chrome, Spotify, or Notepad.",
			Required:    true,
		},
	}
}

func (t *Tool) Execute(arguments map[string]any) (any, error) {
	name, ok := arguments["name"].(string)
	if !ok || strings.TrimSpace(name) == "" {
		return nil, fmt.Errorf("application name is required")
	}

	apps := map[string]string{
		"vs code":  "code",
		"vscode":   "code",
		"chrome":   "chrome",
		"spotify":  "spotify",
		"notepad":  "notepad",
		"calculator": "calc",
	}

	key := strings.ToLower(strings.TrimSpace(name))

	command, ok := apps[key]
	if !ok {
		return nil, fmt.Errorf("application %q is not allowed", name)
	}

	if err := exec.Command("cmd", "/c", "start", "", command).Start(); err != nil {
		return nil, fmt.Errorf("failed to open %s: %w", name, err)
	}

	return fmt.Sprintf("%s opened successfully", name), nil
}

var _ registry.Tool = (*Tool)(nil)