/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"QUICK-Template/api"
	"QUICK-Template/app"
	"QUICK-Template/cache/redis"
	"QUICK-Template/config"
	"QUICK-Template/logger"
	"QUICK-Template/models"
	"QUICK-Template/module/encoder"
	"QUICK-Template/module/hash"
	"QUICK-Template/persistence/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the quicktmp API server",
	Run: func(cmd *cobra.Command, args []string) {
		// init dependencies

		config := config.New()
		logger := logger.New()
		pswdHasher := hash.NewArgon2(hash.DefaultArgon2Config)

		// sql storage
		sqlDB, err := sql.Connect(config.SQL)
		cobra.CheckErr(err)
		storage := sql.NewStorage(sqlDB)
		defer storage.Close(cmd.Context())

		// redis cache
		redisCli, err := redis.Connect(config.Redis)
		cobra.CheckErr(err)
		walletCacher := redis.NewCacher[models.Wallet](redisCli, "wallet", 30*time.Minute, encoder.MsgPack)
		sessionCacher := redis.NewCacher[models.Session](redisCli, "session", 15*time.Minute, encoder.MsgPack)
		defer redisCli.Close()

		app := app.New(logger, storage, pswdHasher, walletCacher, sessionCacher)
		api := api.New(config.API, logger, app)

		// run server
		logger.WithField("port", config.API).Info("http server is starting...")

		if err := api.Serve(); err != nil {
			logger.Fatal("failed to gracefully shutdown http server")
		}

		logger.Info("http server was shutdown gracefully")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
