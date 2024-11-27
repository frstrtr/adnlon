package main

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/tonkeeper/tongo"
	"github.com/tonkeeper/tongo/config"
	"github.com/tonkeeper/tongo/liteapi"
)

func Logger(level string) *zap.Logger {
	cfg := zap.NewProductionConfig()

	var lvl zapcore.Level
	if err := lvl.UnmarshalText([]byte(level)); err != nil {
		panic(err)
	}
	cfg.Level.SetLevel(lvl)

	lg, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return lg
}

func main() {
	log := Logger("INFO")

	liteServerCFG := config.LiteServer{
		Host: LiteServerHost,
		Key:  LiteServerKey,
	}

	tongoClient, err := liteapi.NewClient(liteapi.WithLiteServers([]config.LiteServer{liteServerCFG}))
	if err != nil {
		log.Fatal("NewClient error", zap.Error(err))
	} else {
		log.Info("NewClient success")
	}

	// Parse the account ID
	accountId := tongo.MustParseAccountID("0:E2D41ED396A9F1BA03839D63C5650FAFC6FCFB574FD03F2E67D6555B61A3ACD9")

	// Get the latest block
	oneSeqNo, err := tongoClient.GetSeqno(context.Background(), accountId)
	if err != nil {
		log.Fatal("Get seqno error", zap.Error(err))
	}

	// Get the account state
	state, err := tongoClient.GetAccountState(context.Background(), accountId)
	if err != nil {
		log.Fatal("Get account state error", zap.Error(err))
	}

	// Print the account status and balance
	fmt.Printf("Account status: %v\nBalance: %v\n", state.Account.Status(), state.Account.Account.Storage.Balance.Grams)
	// Print the latest block
	fmt.Printf("Latest block: %v\n", oneSeqNo)
}
