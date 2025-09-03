/*
TeaBrewSim バックエンドAPI
Ginを用いたシンプルなサーバー
/api/simulate POSTエンドポイント雛形
*/
package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors" // CORS対応
	"github.com/gin-gonic/gin"
)

func main() {
	// Ginのデフォルトルーターを作成
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// /api/simulate POSTエンドポイント
	r.POST("/api/simulate", func(c *gin.Context) {
		/*
		   リクエストをパースし、スコア・コメントを計算して返す
		*/
		var req SimulateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}
		a, s, aroma := calcScores(req)
		comment := makeComment(req)
		res := SimulateResponse{
			Astringency: a,
			Sweetness:   s,
			Aroma:       aroma,
			Comment:     comment,
		}
		c.JSON(http.StatusOK, res)
	})

	// ポート8080で起動
	r.Run(":8080")
}
