package tree_sitter_c_sharp_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_c_sharp "github.com/BrokkAi/tree-sitter-c-sharp/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_c_sharp.Language())
	if language == nil {
		t.Errorf("Error loading CSharp grammar")
	}
}
