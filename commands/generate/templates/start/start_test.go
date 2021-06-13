package start

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

const jsonString = `
{
	"project":{
		"name": "somename",
		"basepath": "github.com/somename"
	},
	"appcontexts": {},
	"configs": {},
	"models": [{
		"name": "user"
	}],
	"controllers": [],
	"routes": {}
}
`

func TestGenerateMain(t *testing.T) {
	configs := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonString), &configs)

	files, err := Generate("github.com/somename", configs)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(files))
	mainfile, found := files["main.go"]
	assert.True(t, found)
	assert.NotEqual(t, "", mainfile)

	expected := "package main\n\nimport (\n\t\"/config\"\n\t\"/appcontext\"\n\t\"/models\"\n\t\"/server\"\n\t\"github.com/urfave/cli/v2\"\n\t\"log\"\n\t\"os\"\n)\n\n\nfunc MigrateDB()  {\n\tconfig := config.InitConfig()\n\tappContext := context.InitAppContext(config)\n\tappContext.DB.AutoMigrate(\n\t    &models.User{},)\n}\n\n\nfunc InitServer()  {\n\tconfig := config.InitConfig()\n\tappContext := context.InitAppContext(config)\n\trouter := server.SetupRouter(appContext)\n\trouter.Run()\n}\n\n\nfunc main(){\n\tapp := &cli.App{\n\t    \n\t\tName:    \"somename\",\n\t\tVersion: \"\",\n\t\tCommands: []*cli.Command{\n            \n\t\t\t{\n\t\t\t\tName:    \"migrate\",\n\t\t\t\tAliases: []string{\"m\"},\n\t\t\t\tUsage:   \"migrate database\",\n\t\t\t\tAction:  func(c *cli.Context) error {\n\t\t\t\t\tMigrateDB()\n\t\t\t\t\treturn nil\n\t\t\t\t},\n\t\t\t},\n\t\t\t\n\t\t\t{\n\t\t\t\tName:    \"api\",\n\t\t\t\tAliases: []string{\"a\"},\n\t\t\t\tUsage:   \"start api server\",\n\t\t\t\tAction:  func(c *cli.Context) error {\n\t\t\t\t\tInitServer()\n\t\t\t\t\treturn nil\n\t\t\t\t},\n\t\t\t},\n\t\t},\n\t}\n\n\terr := app.Run(os.Args)\n\tif err != nil {\n\t\tlog.Fatal(err)\n\t}\n}"
	assert.Equal(t, expected, mainfile)
}
