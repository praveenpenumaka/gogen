package makefile

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

const jsonString = `
{
  "project": {
	"name": "gogen",
    "basepath": "github.com/gogen"
  },
  "appcontext": {
    "subcontexts": [{
      "name": "config",
      "type": "config.Config"
    },{
      "name": "DB",
      "type": "*gorm.DB"
    }, {
      "name": "redis",
      "type": "*redis.Client"
    }]

  }
}
`

func TestGenerateMakefile(t *testing.T) {
	var configs map[string]interface{}
	json.Unmarshal([]byte(jsonString), &configs)

	all, err := GenerateMakefile("github.com/gogen", configs)
	assert.Nil(t, err)
	if generated,found := all["Makefile"];found {
		assert.True(t, found)
		expected  := "migrate:\n    go run start.go migrate\n\nclean:\n\trm out/gogen\n\nbuild:\n\tgo build -o out/gogen\n\ninstall:\n\tgo install\n\nrun-api:\n    go run start.go api"
		assert.Equal(t, expected, generated)
	}else{
		assert.Fail(t,"make file not found")
	}
}
