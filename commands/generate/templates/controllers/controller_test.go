package controllers

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
  "controllers": [
	{
		"name": "user"
	}
  ]
}
`

func TestGenerateControllers(t *testing.T) {
	var configs map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &configs)

	fileList, err := GenerateControllers("github.com/gogen", configs)
	assert.Nil(t, err)
	assert.NotNil(t, fileList)
	user, found := fileList["controllers/user.go"]
	assert.True(t, found)
	assert.NotNil(t, user)
	assert.NotEqual(t, "", user)
	expected  := "package controllers\n\nimport (\n\t\"github.com/gogen/models\"\n\t\"github.com/gin-gonic/gin\"\n\t\"gorm.io/gorm\"\n\t\"net/http\"\n\t\"strconv\"\n)\n\ntype UserController struct {\n\tDB *gorm.DB\n}\n\nfunc (base *UserController) CreateUser(c *gin.Context) {\n\tobj := new(models.User)\n\tif err := c.ShouldBindJSON(&obj); err != nil {\n\t\tc.JSON(http.StatusBadRequest, gin.H{\"error\": \"invalid user values\"})\n\t\treturn\n\t}\n\tbase.DB.FirstOrCreate(&obj)\n\tc.JSON(200, obj)\n}\n\nfunc (base *UserController) CreateBulkUsers(c *gin.Context) {\n\tobj := new([]models.User)\n\tif err := c.ShouldBindJSON(&obj); err != nil {\n\t\tc.JSON(http.StatusBadRequest, gin.H{\"error\": \"invalid users values\"})\n\t\treturn\n\t}\n\tbase.DB.Create(obj)\n\tc.JSON(200, obj)\n}\n\nfunc (base *UserController) UpdateUser(c *gin.Context) {\n\tid, exists := c.Params.Get(\"id\")\n\tif !exists {\n\t\tc.JSON(http.StatusBadRequest, gin.H{\"error\": \"id is missing\"})\n\t\treturn\n\t}\n\tuid, err := strconv.ParseUint(id, 10, 32)\n\tif err != nil {\n\t\tc.JSON(http.StatusBadRequest, gin.H{\"error\": \"invalid id provided\"})\n\t\treturn\n\t}\n\tobj := new(models.User)\n\tif err := c.ShouldBindJSON(&obj); err != nil {\n\t\treturn\n\t}\n\ttx := base.DB.Model(&models.User{ID: uint(uid)}).Updates(obj)\n\tif tx.Error != nil {\n\t\tc.JSON(http.StatusBadRequest, gin.H{\"error\": tx.Error})\n\t\treturn\n\t}\n\tc.JSON(200, obj)\n}\n\nfunc (base *UserController) Get<no value>(c *gin.Context) {\n\tid, exists := c.Params.Get(\"id\")\n\tif !exists {\n\t\tc.JSON(http.StatusBadRequest, gin.H{\"error\": \"id is missing\"})\n\t\treturn\n\t}\n\tuid, err := strconv.ParseUint(id, 10, 64)\n\tif err != nil {\n\t\tc.JSON(http.StatusBadRequest, gin.H{\"error\": \"value of id is invalid\"})\n\t\treturn\n\t}\n\tquery := &models.User{ID: uint(uid)}\n\tobj := new(models.User)\n\tbase.DB.Where(&query).Find(&obj)\n\tc.JSON(200, obj)\n}\n\nfunc (base *UserController) ListUsers(c *gin.Context) {\n\tquery := new(models.User)\n\tobj := new([]models.User)\n\tbase.DB.Where(&query).Find(&obj)\n\tc.JSON(200, obj)\n}\n\nfunc (base *UserController) DeleteUser(c *gin.Context) {\n\tid, exists := c.Params.Get(\"id\")\n\tif !exists {\n\t\tc.JSON(http.StatusBadRequest, gin.H{\"error\": \"id is missing\"})\n\t\treturn\n\t}\n\tuid, err := strconv.ParseUint(id, 10, 64)\n\tif err != nil {\n\t\tc.JSON(http.StatusBadRequest, gin.H{\"error\": \"value of id is invalid\"})\n\t\treturn\n\t}\n\tquery := &models.User{ID: uint(uid)} // TODO: Soft delete user\n\tobj := new([]models.User)\n\tbase.DB.Where(&query).Delete(&obj)\n\tc.JSON(200, obj)\n}"
	assert.Equal(t, expected, user)
}
