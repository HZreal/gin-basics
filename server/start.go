package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)



func login(c *gin.Context) {
	c.String(http.StatusOK, "login success")
}
func register(c *gin.Context) {
	c.String(http.StatusOK, "register success")
}

func formData(c *gin.Context) {
	filename := c.PostForm("filename")
	file, _ := c.FormFile("file")
	fmt.Println(filename)
	fmt.Println(file)
	_ = c.SaveUploadedFile(file, file.Filename)

	c.JSON(http.StatusOK, gin.H{"code": 0})
}

// 取body原始数据进行处理
func parseRawBody(ctx *gin.Context) {
	// 获取原始body数据
	body, _ := ctx.GetRawData()
	fmt.Println("body ---------", body)

	// 解析原始body数据

	// 1. body为单个文件，可直接存储等操作
	// _ = ioutil.WriteFile("123.zip", body, 0644)

	// 2. 若为json，可解析到map或struct或interface{}
	// mapData := make(map[string]interface{})
	// _ = json.Unmarshal(body, &mapData)
	// fmt.Println(mapData)
	// structData := struct {
	//	name string
	//	age int
	// }{}
	// _ = json.Unmarshal(body, &structData)
	// fmt.Println(structData)
	// var interfaceData interface{}
	// _ = json.Unmarshal(body, &interfaceData)

	// 3. 以表单上传文件(multipart/form-data)   同样可以解析到map/struct/interface{}
	// mapData2 := make(map[string]interface{})
	// _ = json.Unmarshal(body, &mapData2)
	// fmt.Println(mapData2["filename"])
	// fmt.Println(mapData2["file"])        // file文件  []byte
	// if v, ok := mapData["file"].([]byte); ok {
	//	fmt.Println("断言为[]byte成功", v)
	//	_ = ioutil.WriteFile("123.zip", mapData["file"], 0644)
	// }
	// if v, ok := mapData["file"].(string); ok {
	//	fmt.Println("断言为string成功", v)
	//	_ = ioutil.WriteFile("123.zip", []byte(mapData["file"]), 0644)
	// }

	// structData2 := struct {
	//	filename string
	//	file []byte
	// }{}
	// _ = json.Unmarshal(body, &structData2)
	// fmt.Println(structData2)
	// _ = ioutil.WriteFile("123.zip", structData2.file, 0644)

	ctx.JSON(200, gin.H{"code": 0, "msg": "upload success", "data": nil})

}


func main() {
	// 1.创建路由
	var r *gin.Engine = gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	r.GET("/urlParams/string:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		fmt.Println(name, action)
		c.String(http.StatusOK, "%s is %s", name, action)
	})
	r.GET("/deviceDetail", func(c *gin.Context) {
		device_id := c.Query("device_id") // 参数不存在，默认值为空串
		// device_id := c.DefaultQuery("device_id", "defaultID")    //  必须设置默认值
		c.String(http.StatusOK, "device_id is %s", device_id)
	})
	api := r.Group("/api")
	{
		api.GET("/login", login)
		api.POST("/register", register)
		api.POST("/formData", formData)
		api.POST("/parseRawBody", parseRawBody)
	}

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run("0.0.0.0:8000")
}



