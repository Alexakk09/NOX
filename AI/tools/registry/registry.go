package registry

import "fmt"

type Tool interface {
	Name() string
	Description() string
	Execute(arguments map[string]any) (any, error)
}

type Registry struct {
	tools map[string]Tool
}

func New() *Registry {
	return &Registry{
		tools: make(map[string]Tool),
	}
}

func (r *Registry) Register(tool Tool) {
	r.tools[tool.Name()] = tool
}

func (r *Registry) Get(name string) (Tool, bool) {
	tool, ok := r.tools[name]
	return tool, ok
}

func (r *Registry) Execute(name string, arguments map[string]any) (any, error) {
	tool, ok := r.tools[name]
	if !ok {
		return nil, fmt.Errorf("tool %q not found", name)
	}

	return tool.Execute(arguments)
}
func (r *Registry) Tools() []Tool {
	tools := make([]Tool, 0, len(r.tools))

	for _, tool := range r.tools {
		tools = append(tools, tool)
	}

	return tools
}