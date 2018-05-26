package robocli

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"robodb"
	"roboapi"
	log "github.com/sirupsen/logrus"
)

var listenURL string
var dbURL string

// 设置命令行参数
var rootCmd = &cobra.Command{
	Use:   "slncenter",
	Short: cmdIntroduce,
	Long:  cmdIntroduce,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := robodb.InitDB(dbURL)
		if err != nil {
			log.WithFields(log.Fields{
				"url": dbURL,
			}).Fatal("Cannot connect to database")
		}
		defer db.Close()
		roboapi.StartWebService(listenURL, db)
	},
}

// 初始化命令行
func InitCmd() {
	// print banner
	color.Cyan(cmdBanner)
	// register params
	rootCmd.PersistentFlags().StringVarP(&listenURL, "listen", "l", "127.0.0.1:9000", "listen url")
	rootCmd.PersistentFlags().StringVarP(&dbURL, "database", "d", "root:root@tcp(localhost:3306)/mysql?charset=utf-8", "database url")
	// execute
	rootCmd.Execute()
}
