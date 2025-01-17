package django_test

import (
	gojinja2 "github.com/jmolinski/gosible-templates"
	"testing"

	"github.com/jmolinski/gosible-templates/ext/django"
	tu "github.com/jmolinski/gosible-templates/testutils"
)

func Env(root string) *gojinja2.Environment {
	env := tu.TestEnv(root)
	env.Filters.Update(django.Filters)
	env.Statements.Update(django.Statements)
	return env
}

func TestDjangoTemplates(t *testing.T) {
	root := "./testData"
	env := Env(root)
	tu.GlobTemplateTests(t, root, env)
}

func TestDjangoFilters(t *testing.T) {
	root := "./testData/filters"
	env := Env(root)
	tu.GlobTemplateTests(t, root, env)
}

func TestDjangoStatements(t *testing.T) {
	root := "./testData/statements"
	env := Env(root)
	tu.GlobTemplateTests(t, root, env)
}
