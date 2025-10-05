/*
TeaBrewSim バックエンドAPI
Ginを用いたフル機能サーバー
設定管理、ログ、データベース、複数APIエンドポイント対応
*/
package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-contrib/cors" // CORS対応
	"github.com/gin-gonic/gin"
)

// グローバル変数
var (
	config   *Config
	database *Database
	logger   *Logger
)

func main() {
	// 設定を読み込み
	var err error
	configPath := GetConfigPath()
	config, err = LoadConfig(configPath)
	if err != nil {
		panic("設定の読み込みに失敗: " + err.Error())
	}

	// ロガーを初期化
	logger = NewLogger(ParseLogLevel(config.Logging.Level), os.Stdout)
	logger.Info("TeaBrewSim サーバーを起動中...")

	// データベースを初期化
	dbPath := filepath.Join(filepath.Dir(configPath), "teabrewsim.db")
	database, err = NewDatabase(dbPath)
	if err != nil {
		logger.Fatal("データベースの初期化に失敗: " + err.Error())
	}
	defer database.Close()

	// Ginのルーターを作成
	r := gin.Default()
	
	// CORS設定
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600, // 12時間
	}))

	// ログミドルウェアを追加
	r.Use(LogMiddleware())

	// APIルートグループ
	api := r.Group("/api")
	{
		// ヘルスチェック
		api.GET("/health", HealthCheck)
		
		// お茶の種類一覧
		api.GET("/tea-types", GetTeaTypes)
		
		// 抽出シミュレーション
		api.POST("/simulate", func(c *gin.Context) {
			var req SimulateRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				logger.Error("無効なリクエスト", map[string]interface{}{
					"error": err.Error(),
				})
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
				return
			}

			// スコア計算
			a, s, aroma := calcScores(req)
			comment := makeComment(req)
			
			res := SimulateResponse{
				Astringency: a,
				Sweetness:   s,
				Aroma:       aroma,
				Comment:     comment,
			}

			logger.Info("抽出シミュレーション実行", map[string]interface{}{
				"tea_type":    req.TeaType,
				"temperature": req.Temp,
				"time":        req.Time,
			})

			c.JSON(http.StatusOK, res)
		})

		// 抽出履歴
		api.GET("/history", GetBrewingHistory)
		api.POST("/history", SaveBrewingResult)
		
		// 推奨設定
		api.GET("/recommendations/:type", GetRecommendations)
		
		// 統計情報
		api.GET("/statistics", GetStatistics)
	}

	// サーバー起動
	addr := config.Server.Host + ":" + config.Server.Port
	logger.Info("サーバーを起動", map[string]interface{}{
		"address": addr,
		"config":  configPath,
	})
	
	if err := r.Run(addr); err != nil {
		logger.Fatal("サーバーの起動に失敗: " + err.Error())
	}
}
