package config

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

const jsonString = `
[
    {
        "name": "app",
        "fields": [{
          "name": "AppBaseUrl",
          "type": "string"
        }]
    }
]
`

func TestGenerate(t *testing.T) {
	var configs []map[string]interface{}
	json.Unmarshal([]byte(jsonString), &configs)

	path, generated, err := Generate(configs)
	assert.Nil(t, err)
	expected := "package config\n\nimport (\n\t\"github.com/joho/godotenv\"\n\t\"os\"\n)\n\ntype Config struct {AppConfig AppConfig}\n\nfunc InitConfig() Config {\n\tgodotenv.Load()\n\treturn Config{AppConfig InitAppConfig(),}\n}"
	assert.Equal(t, expected, generated)
	assert.Equal(t, "config/config.go", path)
}

func TestGenerateSubconfig(t *testing.T) {
	var configs []map[string]interface{}
	json.Unmarshal([]byte(jsonString), &configs)

	path, generated, err := GenerateSubconfig(configs[0])
	assert.Nil(t, err)
	expected := "package config\n\nimport (\n\t\"github.com/joho/godotenv\"\n\t\"os\"\n)\n\ntype AppConfig struct {\n    AppBaseUrl string\n\t\n}\n\nfunc InitAppConfig() AppConfig {\n\treturn AppConfig{\n        AppBaseUrl: os.Getenv(\"APPBASEURL\"),\n        \n\t}\n}"
	assert.Equal(t, expected, generated)
	assert.Equal(t, "config/app.go", path)
}
