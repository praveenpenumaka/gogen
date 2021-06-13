package router

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
  "router":{
    "cruds": [{
      "name": "snip"
    }]
  }
}
`

func TestGenerateRouter(t *testing.T) {
	var configs map[string]interface{}
	json.Unmarshal([]byte(jsonString), &configs)

	path, generated, err := GenerateRouter("github.com/gogen", configs["router"].(map[string]interface{}))
	assert.Nil(t, err)
	expected := "package server\n\nimport (\n\t\"github.com/gogen/context\"\n\t\"github.com/gogen/controllers\"\n\t\"github.com/gin-gonic/gin\"\n)\n\nfunc SetupRouter(context context.AppContext) *gin.Engine {\n\trouter := gin.New()\n\trouter.Use(gin.Logger())\n\trouter.Use(gin.Recovery())\n    \n        router =SnipCrudSetup(router,context)\n        router =UserCrudSetup(router,context)\n\n\treturn router\n}"
	assert.Equal(t, expected, generated)
	assert.Equal(t, "server/router.go", path)
}

func TestGenerateCruds(t *testing.T) {
	var configs map[string]interface{}
	json.Unmarshal([]byte(jsonString), &configs)

	appContextConfig := configs["router"].(map[string]interface{})
	sublist := appContextConfig["cruds"].([]interface{})
	subContextConfig := sublist[0]
	path, generated, err := GenerateCruds("github.com/gogen", subContextConfig.(map[string]interface{}))
	assert.Nil(t, err)
	expected := "package server\n\nimport (\n\t\"github.com/gogen/context\"\n\t\"github.com/gogen/controllers\"\n\t\"github.com/gin-gonic/gin\"\n)\n\n\nfunc SnipCrudSetup(router *gin.Engine,ctx context.AppContext) *gin.Engine{\n\tcontroller := controllers.SnipController{DB:ctx.DB}\n\tcrud := router.Group(\"/snips\")\n\t{\n\t\tcrud.GET(\"/\",controller.ListSnips)\n\t\tcrud.POST(\"/new\",controller.CreateSnip)\n\t\tcrud.POST(\"/bulk-new\",controller.CreateBulkSnips)\n\t\tcrud.PUT(\"/:id/update\",controller.ListSnips)\n\t\tcrud.GET(\"/:id\",controller.GetSnip)\n\t\tcrud.DELETE(\"/:id/delete\",controller.DeleteSnip)\n\t}\n\treturn router\n}"
	assert.Equal(t, expected, generated)
	assert.Equal(t, "server/snip.go", path)
}

func TestGetAll(t *testing.T) {
	var configs map[string]interface{}
	json.Unmarshal([]byte(jsonString), &configs)

	resources, err := Generate("github.com/gogen", configs)
	assert.Nil(t, err)
	assert.NotNil(t, resources)
	_, found1 := resources["server/router.go"]
	_, found2 := resources["server/snip.go"]
	assert.True(t, found1)
	assert.True(t, found2)
	//assert.Equal(t, "appcontext/config.go", "")
}
