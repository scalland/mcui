package cmd

import (
	"fmt"
	"mcui/config"
	"mcui/handlers"
	"mcui/memcache"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var configPath string

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the Memcached UI server",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig(configPath)
		if err != nil {
			return err
		}

		address := fmt.Sprintf("%s:%d", cfg.Memcached.MCHost, cfg.Memcached.MCPort)
		memcache.Init(address)

		r := gin.Default()
		r.LoadHTMLGlob("templates/*")
		r.Static("/static", "./static")

		api := r.Group("/api")
		{
			api.POST("/set", handlers.SetKey)
			api.GET("/get/:key", handlers.GetKey)
			api.DELETE("/delete/:key", handlers.DeleteKey)
			api.GET("/stats", handlers.Stats)
		}

		r.GET("/", handlers.RenderIndex)
		r.POST("/set", handlers.HandleSetHTML)
		r.GET("/get", handlers.HandleGetHTML)
		r.POST("/delete", handlers.HandleDeleteHTML)
		r.GET("/stats", handlers.HandleStatsHTML)

		fmt.Printf("🚀 Starting server on http://%s:%d", cfg.App.Host, cfg.App.Port)
		return r.Run(fmt.Sprintf("%s:%d", cfg.App.Host, cfg.App.Port))
	},
}

func init() {
	ServeCmd.Flags().StringVarP(&configPath, "config", "c", "config.yaml", "Path to config file")
}
