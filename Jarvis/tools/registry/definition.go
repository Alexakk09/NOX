package registry

type Parameter struct {
	Name        string
	Type        string
	Description string
	Required    bool
}

type Definition struct {
	Name        string
	Description string
	Parameters  []Parameter
}

func (r *Registry) Definitions() []Definition {
	definitions := make([]Definition, 0, len(r.tools))

	for _, tool := range r.tools {
		definition := Definition{
			Name:        tool.Name(),
			Description: tool.Description(),
		}

		if parameterized, ok := tool.(interface {
			Parameters() []Parameter
		}); ok {
			definition.Parameters = parameterized.Parameters()
		}

		definitions = append(definitions, definition)
	}

	return definitions
}