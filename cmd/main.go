package main

import (
	"context"
	app "github.com/cloudslit/newca/internal"
	"github.com/cloudslit/newca/pkg/logger"
	"github.com/urfave/cli/v2"
	"os"
)

var VERSION = "0.0.1"

func main() {
	ctx := logger.NewTagContext(context.Background(), "__main__")

	app := cli.NewApp()
	app.Name = "ca"
	app.Version = VERSION
	app.Usage = "CA PKI"
	app.Commands = []*cli.Command{
		newTlsCmd(ctx),
	}
	err := app.Run(os.Args)
	if err != nil {
		logger.WithContext(ctx).Errorf(err.Error())
	}
}

func newTlsCmd(ctx context.Context) *cli.Command {
	return &cli.Command{
		Name:  "tls",
		Usage: "Run Tls server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "conf",
				Aliases:  []string{"c"},
				Usage:    "App configuration file(.json,.yaml,.toml)",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			return app.Run(ctx,
				app.SetConfigFile(c.String("conf")),
				app.SetVersion(VERSION))
		},
	}
}

//
//// InitTlsServer Initialize TLS service
//func InitTlsServer(ctx context.Context, handler *mux.Router) func() {
//	addr := core.Is.Config.HTTP.CaListen
//	tlsCfg := &tls.Config{
//		GetCertificate: func(info *tls.ClientHelloInfo) (*tls.Certificate, error) {
//			return keymanager.GetKeeper().GetCachedTLSKeyPair()
//		},
//		InsecureSkipVerify: true,
//		ClientAuth:         tls.NoClientCert,
//	}
//	srv := &http.Server{
//		Addr:         addr,
//		TLSConfig:    tlsCfg,
//		Handler:      handler,
//		ReadTimeout:  5 * time.Second,
//		WriteTimeout: 10 * time.Second,
//		IdleTimeout:  15 * time.Second,
//	}
//
//	go func() {
//		logger.Infof("TLS server is running at %s.", addr)
//		err := srv.ListenAndServeTLS("", "")
//		if err != nil && err != http.ErrServerClosed {
//			panic(err)
//		}
//	}()
//
//	return func() {
//		ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(30))
//		defer cancel()
//
//		srv.SetKeepAlivesEnabled(false)
//		if err := srv.Shutdown(ctx); err != nil {
//			logger.Errorf(err.Error())
//		}
//	}
//}
//
//func RunTls(ctx context.Context) error {
//	state := 1
//	sc := make(chan os.Signal, 1)
//	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
//
//	app, err := singleca.Server()
//	if err != nil {
//		return err
//	}
//	cleanFunc := InitTlsServer(ctx, app)
//
//EXIT:
//	for {
//		sig := <-sc
//		logger.Infof("Received signal[%s]", sig.String())
//		switch sig {
//		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
//			state = 0
//			break EXIT
//		case syscall.SIGHUP:
//		default:
//			break EXIT
//		}
//	}
//
//	cleanFunc()
//	logger.Infof("TLS service exit")
//	time.Sleep(time.Second)
//	os.Exit(state)
//	return nil
//}
