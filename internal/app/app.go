package app

import (
	"embed"
	"fmt"
	"github.com/andReyM228/lib/database"

	"tx_service/internal/config"
	balancesHandler "tx_service/internal/handler/balances"
	"tx_service/internal/repository/balances"
	"tx_service/internal/repository/transactions"
	balancesService "tx_service/internal/service/balances"

	"github.com/andReyM228/lib/log"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type App struct {
	config           config.Config
	serviceName      string
	balancesRepo     balances.Repository
	transactionsRepo transactions.Repository
	balancesService  balancesService.Service
	balancesHandler  balancesHandler.Handler
	logger           log.Logger
	db               *sqlx.DB

	router *fiber.App
}

func New(name string) App {
	return App{
		serviceName: name,
	}
}

func (a *App) Run(fs embed.FS) {
	a.populateConfig()
	a.initLogger()
	a.initDatabase(fs)
	a.initRepos()
	a.initServices()
	a.initHandlers()
	a.initHTTP()
}

func (a *App) initHTTP() {
	a.router = fiber.New()

	a.router.Post("v1/tx-service/transfer", a.balancesHandler.Transfer)

	a.logger.Debug("fiber api started")
	_ = a.router.Listen(fmt.Sprintf(":%d", a.config.HTTP.Port))
}

func (a *App) initDatabase(fs embed.FS) {
	database.InitDatabase(a.logger, a.config.DB, fs)
}

func (a *App) initLogger() {
	a.logger = log.Init()
}

func (a *App) initRepos() {
	a.balancesRepo = balances.NewRepository(a.db, a.logger)
	a.transactionsRepo = transactions.NewRepository(a.db, a.logger)

	a.logger.Debug("repos created")
}

func (a *App) initServices() {
	a.balancesService = balancesService.NewService(a.balancesRepo, a.transactionsRepo)

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
