package registry

import "testing"

type testTool struct{}

func (t *testTool) Name() string {
	return "test"
}

func (t *testTool) Description() string {
	return "A test tool"
}

func (t *testTool) Execute(arguments map[string]any) (any, error) {
	return "tool executed", nil
}

func TestRegistryExecute(t *testing.T) {
	registry := New()

	registry.Register(&testTool{})

	result, err := registry.Execute("test", nil)
	if err != nil {
		t.Fatal(err)
	}

	if result != "tool executed" {
		t.Fatalf("unexpected result: %v", result)
	}
}