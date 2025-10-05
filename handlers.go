/*
TeaBrewSim API ハンドラー
各種エンドポイントの処理を定義
*/
package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// HealthCheck ヘルスチェックエンドポイント
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "healthy",
		"timestamp": time.Now().Unix(),
		"service":   "teabrewsim",
		"version":   "1.0.0",
	})
}

// GetTeaTypes お茶の種類一覧を取得
func GetTeaTypes(c *gin.Context) {
	teaTypes := []gin.H{
		{"id": "green", "name": "緑茶", "description": "日本茶の代表的なお茶"},
		{"id": "black", "name": "紅茶", "description": "発酵させたお茶"},
		{"id": "oolong", "name": "烏龍茶", "description": "半発酵のお茶"},
		{"id": "white", "name": "白茶", "description": "微発酵のお茶"},
		{"id": "pu-erh", "name": "プーアル茶", "description": "後発酵のお茶"},
		{"id": "herbal", "name": "ハーブティー", "description": "ハーブを原料としたお茶"},
	}
	c.JSON(http.StatusOK, gin.H{"teaTypes": teaTypes})
}

// GetBrewingHistory 抽出履歴を取得（簡易版）
func GetBrewingHistory(c *gin.Context) {
	// 実際の実装ではデータベースから取得
	history := []gin.H{
		{
			"id":          1,
			"teaType":     "green",
			"temperature": 80.0,
			"time":        120,
			"timestamp":   time.Now().Add(-1 * time.Hour).Unix(),
			"rating":      4.5,
		},
		{
			"id":          2,
			"teaType":     "black",
			"temperature": 95.0,
			"time":        180,
			"timestamp":   time.Now().Add(-2 * time.Hour).Unix(),
			"rating":      4.0,
		},
	}
	c.JSON(http.StatusOK, gin.H{"history": history})
}

// SaveBrewingResult 抽出結果を保存
func SaveBrewingResult(c *gin.Context) {
	var req struct {
		TeaType     string  `json:"teaType" binding:"required"`
		Temperature float64 `json:"temperature" binding:"required"`
		Time        int     `json:"time" binding:"required"`
		Rating      float64 `json:"rating"`
		Comment     string  `json:"comment"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// 実際の実装ではデータベースに保存
	result := gin.H{
		"id":          time.Now().Unix(),
		"teaType":     req.TeaType,
		"temperature": req.Temperature,
		"time":        req.Time,
		"rating":      req.Rating,
		"comment":     req.Comment,
		"timestamp":   time.Now().Unix(),
	}

	c.JSON(http.StatusCreated, gin.H{"result": result})
}

// GetRecommendations お茶の種類に基づく推奨設定を取得
func GetRecommendations(c *gin.Context) {
	teaType := c.Param("type")
	
	recommendations := map[string]gin.H{
		"green": {
			"temperature": 80.0,
			"time":        120,
			"tips":        []string{"高温すぎると苦味が強くなります", "1煎目は短めに抽出"},
		},
		"black": {
			"temperature": 95.0,
			"time":        180,
			"tips":        []string{"しっかりと蒸らすことが重要", "ミルクを入れる場合は濃いめに"},
		},
		"oolong": {
			"temperature": 90.0,
			"time":        150,
			"tips":        []string{"香りを楽しむため温度に注意", "何煎も楽しめるお茶"},
		},
	}

	if rec, exists := recommendations[teaType]; exists {
		c.JSON(http.StatusOK, gin.H{"recommendation": rec})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "tea type not found"})
	}
}

// GetStatistics 統計情報を取得
func GetStatistics(c *gin.Context) {
	// 実際の実装ではデータベースから集計
	stats := gin.H{
		"totalBrews":     42,
		"averageRating": 4.2,
		"favoriteTea":   "green",
		"totalTime":      "12h 30m",
		"thisWeek": gin.H{
			"brews":  8,
			"rating": 4.3,
		},
	}
	c.JSON(http.StatusOK, gin.H{"statistics": stats})
}
