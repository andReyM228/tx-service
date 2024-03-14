package app

import (
	"context"
	"fmt"
	"github.com/andReyM228/lib/bus"
	"github.com/andReyM228/lib/rabbit"
	"github.com/andReyM228/one/chain_client"
	"tx_service/internal/delivery"
	"tx_service/internal/delivery/broker/transfers"
	balancesHandler "tx_service/internal/delivery/http/balances"
	"tx_service/internal/domain"
	"tx_service/internal/repositories"
	"tx_service/internal/repositories/chain"
	"tx_service/internal/services"

	"tx_service/internal/config"
	balancesService "tx_service/internal/services/balances"

	"github.com/andReyM228/lib/log"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	config          config.Config
	serviceName     string
	chainRepo       repositories.Chain
	balancesService services.Balances
	balancesHandler delivery.Balances
	transfersBroker delivery.TransfersBroker
	logger          log.Logger
	chain           chain_client.Client
	router          *fiber.App
	rabbit          rabbit.Rabbit
}

func New(name string) App {
	return App{
		serviceName: name,
	}
}

func (a *App) Run(ctx context.Context) {
	a.populateConfig()
	a.initLogger()
	a.initChainClient(ctx)
	a.initRabbit()
	a.initRepos()
	a.initServices()
	a.initHandlers()
	a.listenRabbit()
	a.initHTTP()
}

func (a *App) initHTTP() {
	a.router = fiber.New()

	group := a.router.Group("/v1/tx-services")
	group.Post("/issue", a.balancesHandler.Issue)
	group.Post("/withdraw", a.balancesHandler.Withdraw)

	a.logger.Debug("fiber api started")
	_ = a.router.Listen(fmt.Sprintf(":%d", a.config.HTTP.Port))
}

func (a *App) listenRabbit() {
	err := a.rabbit.Consume(bus.SubjectTxServiceIssue, a.transfersBroker.BrokerIssue)
	if err != nil {
		return
	}

	err = a.rabbit.Consume(bus.SubjectTxServiceWithdraw, a.transfersBroker.BrokerWithdraw)
	if err != nil {
		return
	}
}

func (a *App) initChainClient(ctx context.Context) {
	a.chain = chain_client.NewClient(a.config.Chain)

	err := a.chain.AddAccount(ctx, domain.SignerAccount, a.config.Extra.Mnemonic)
	if err != nil {
		a.logger.Fatal(err.Error())
	}
}

func (a *App) initLogger() {
	a.logger = log.Init()
}

func (a *App) initRepos() {
	a.chainRepo = chain.NewRepository(a.chain, a.logger)

	a.logger.Debug("repos created")
}

func (a *App) initServices() {
	a.balancesService = balancesService.NewService(a.chainRepo, a.logger)

	a.logger.Debug("repos created")
}

func (a *App) initHandlers() {
	a.balancesHandler = balancesHandler.NewHandler(a.balancesService)

	a.transfersBroker = transfers.NewHandler(a.rabbit, a.logger, a.balancesService)

	a.logger.Debug("handlers created")
}

func (a *App) populateConfig() {
	cfg, err := config.ParseConfig()
	if err != nil {
		a.logger.Fatalf("parse config error: ", err)
	}

	a.config = cfg
}

func (a *App) initRabbit() {
	var err error
	a.rabbit, err = rabbit.NewRabbitMQ(a.config.Rabbit.Url, a.logger)
	if err != nil {
		a.logger.Fatal(err.Error())
	}
}
