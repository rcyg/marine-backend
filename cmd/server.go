package cmd

import (
	"context"
	"errors"
	"fmt"
	"marine-backend/cmd/flags"
	"marine-backend/internal/conf"
	"marine-backend/pkg/utils"
	"marine-backend/server"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ServerCmd represents the server command
var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the server at the specified address",
	Long: `Start the server at the specified address
the address is defined in config file`,
	Run: func(cmd *cobra.Command, args []string) {
		Init()
		if conf.Conf.DelayedStart != 0 {
			utils.Log.Infof("delayed start for %d seconds", conf.Conf.DelayedStart)
			time.Sleep(time.Duration(conf.Conf.DelayedStart) * time.Second)
		}
		if !flags.Debug && !flags.Dev {
			gin.SetMode(gin.ReleaseMode)
		}
		r := gin.New()
		r.Use(gin.LoggerWithWriter(log.StandardLogger().Out), gin.RecoveryWithWriter(log.StandardLogger().Out))
		server.Init(r)
		var httpSrv *http.Server
		if conf.Conf.HttpPort != -1 {
			httpBase := fmt.Sprintf("%s:%d", conf.Conf.Address, conf.Conf.HttpPort)
			utils.Log.Infof("start HTTP server @ %s", httpBase)
			httpSrv = &http.Server{Addr: httpBase, Handler: r}
			go func() {
				err := httpSrv.ListenAndServe()
				if err != nil && !errors.Is(err, http.ErrServerClosed) {
					utils.Log.Fatalf("failed to start http: %s", err.Error())
				}
			}()
		}
		// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 1 second.
		quit := make(chan os.Signal, 1)
		// kill (no param) default send syscanll.SIGTERM
		// kill -2 is syscall.SIGINT
		// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		utils.Log.Println("Shutdown server...")
		Release()
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		var wg sync.WaitGroup
		if conf.Conf.HttpPort != -1 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if err := httpSrv.Shutdown(ctx); err != nil {
					utils.Log.Fatal("HTTP server shutdown err: ", err)
				}
			}()
		}
		wg.Wait()
		utils.Log.Println("Server exit")
	},
}

func init() {
	ServerCmd.PersistentFlags().StringVar(&flags.DataDir, "data", "data", "data folder")
	ServerCmd.PersistentFlags().BoolVar(&flags.Debug, "debug", false, "start with debug mode")
	ServerCmd.PersistentFlags().BoolVar(&flags.Dev, "dev", false, "start with dev mode")
	ServerCmd.PersistentFlags().BoolVar(&flags.LogStd, "log-std", false, "Force to log to std")
}

func Execute() {
	if err := ServerCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
