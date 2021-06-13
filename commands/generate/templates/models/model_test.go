package models

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

const jsonString = `
{
  "project": {
    "basepath": "github.com/gogen"
  },
  "models": [
	{
		"name": "user",
		"fields": [{
			"name": "username",
			"type": "string"
		},{
			"name": "email",
			"type": "string"
		}]
	}
  ]
}
`

func TestGenerateModels(t *testing.T) {
	var configs map[string]interface{}
	json.Unmarshal([]byte(jsonString), &configs)

	fileList, err := Generate("github.com/gogen", configs)
	assert.Nil(t, err)
	assert.NotNil(t, fileList)
	user,found := fileList["models/user.go"]
	assert.True(t,found)
	assert.NotNil(t, user)
	expected  := "package models\n\nimport  \"gorm.io/gorm\"\n\ntype User struct {\n\tgorm.Model\n\tID\tuint\t\t`gorm:\"primaryKey;autoIncrement:true\"`\n\t\n\t\n    Username string   `json:\"username\"`\n\t\n    Email string   `json:\"email\"`\n\t\n}\n\nfunc (m *User)Valid() bool {\n\treturn !(m.ID <= 0)\n}"
	assert.Equal(t, expected, user)
}
