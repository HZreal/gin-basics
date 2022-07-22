package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {     // CheckOrigin防止跨站点的请求伪造
		return true
	},
}

func index(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", nil)

}

func ping(c *gin.Context) {
	// 升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close() // 返回前关闭

	for {
		// 读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		// 写入ws数据
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}


}

func start() {
	r := gin.Default()
	r.LoadHTMLGlob("./*.html")
	r.GET("index", index)
	r.GET("/ws", ping)

	r.Run("0.0.0.0:8000")
}


func handler(w http.ResponseWriter, r *http.Request) {

	c, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("cannot upgrade: %v\n", err)
		return
	}

	c.WriteMessage(1, []byte("hello"))


	// i := 0
	// for {
	// 	i++
	// 	c.WriteJSON(map[string]string{
	// 		"hello": "wensocket",
	// 		"msg_id": strconv.Itoa(i),
	// 	})
	// 	time.Sleep(2 * time.Second)
	// 	if i == 3 {
	// 		break
	// 	}
	// }



}

func start2()  {
	http.HandleFunc("/ws_connect", handler)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func main()  {
	start()
	// start2()
}

