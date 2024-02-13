package app

import (
	"context"
	"fmt"
	"github.com/andReyM228/one/chain_client"
	"tx_service/internal/domain"
	"tx_service/internal/handler"
	"tx_service/internal/repository"
	"tx_service/internal/repository/chain"
	"tx_service/internal/service"

	"tx_service/internal/config"
	balancesHandler "tx_service/internal/handler/balances"
	balancesService "tx_service/internal/service/balances"

	"github.com/andReyM228/lib/log"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	config          config.Config
	serviceName     string
	chainRepo       repository.Chain
	balancesService service.Balances
	balancesHandler handler.Balances
	logger          log.Logger
	chain           chain_client.Client
	router          *fiber.App
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
	a.initRepos()
	a.initServices()
	a.initHandlers()
	a.initHTTP()
}

func (a *App) initHTTP() {
	a.router = fiber.New()

	group := a.router.Group("/v1/tx-service")
	group.Post("/issue", a.balancesHandler.Issue)
	group.Post("/withdraw", a.balancesHandler.Withdraw)

	a.logger.Debug("fiber api started")
	_ = a.router.Listen(fmt.Sprintf(":%d", a.config.HTTP.Port))
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
	a.logger.Debug("handlers created")
}

func (a *App) populateConfig() {
	cfg, err := config.ParseConfig()
	if err != nil {
		a.logger.Fatalf("parse config error: ", err)
	}

	a.config = cfg
}
