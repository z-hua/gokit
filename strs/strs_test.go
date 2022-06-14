package strs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCamelCase(t *testing.T) {
	assert.Equal(t, "myName", CamelCase("MyName"))
	assert.Equal(t, "myName", CamelCase("myName"))
	assert.Equal(t, "myName", CamelCase("my_name"))
	assert.Equal(t, "myName", CamelCase("my_Name"))
	assert.Equal(t, "myName", CamelCase("My_Name"))
}

func TestPascalCase(t *testing.T) {
	assert.Equal(t, "MyName", PascalCase("MyName"))
	assert.Equal(t, "MyName", PascalCase("myName"))
	assert.Equal(t, "MyName", PascalCase("my_name"))
	assert.Equal(t, "MyName", PascalCase("my_Name"))
	assert.Equal(t, "MyName", PascalCase("My_Name"))
}

func TestSnakeCase(t *testing.T) {
	assert.Equal(t, "my_name", SnakeCase("MyName"))
	assert.Equal(t, "my_name", SnakeCase("myName"))
	assert.Equal(t, "my_name", SnakeCase("my_name"))
	assert.Equal(t, "my_name", SnakeCase("my_Name"))
	assert.Equal(t, "my_name", SnakeCase("My_Name"))
}
