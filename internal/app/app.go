package app

import (
	"fmt"
	"log"
	"user_service/internal/config"
	"user_service/internal/repository/balances"
	"user_service/internal/repository/transactions"
	balances2 "user_service/internal/service/balances"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type App struct {
	config           config.Config
	serviceName      string
	balancesRepo     balances.Repository
	transactionsRepo transactions.Repository
	balancesService  balances2.Service
	logger           *logrus.Logger
	db               *sqlx.DB

	router *fiber.App
}

func New(name string) App {
	return App{
		serviceName: name,
	}
}

func (a *App) Run() {
	a.populateConfig()
	a.initLogger()
	a.initDatabase()
	a.initHTTP()
	a.initRepos()
	a.initServices()
	a.initHandlers()
}

func (a *App) initHTTP() {
	a.router = fiber.New()

	//роуты

	a.logger.Debug("fiber api started")
	_ = a.router.Listen(fmt.Sprintf(":%d", a.config.HTTP.Port))
}

func (a *App) initDatabase() {
	a.logger.Debug("opening database connection")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		a.config.DB.Host, a.config.DB.Port, a.config.DB.User, a.config.DB.Password, a.config.DB.DBname)

	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	a.db = db
	a.logger.Debug("database connected")
}

func (a *App) initLogger() {
	a.logger = logrus.New()
	a.logger.SetLevel(logrus.DebugLevel)
}

func (a *App) initRepos() {
	a.balancesRepo = balances.NewRepository(a.db, a.logger)
	a.transactionsRepo = transactions.NewRepository(a.db, a.logger)

	a.logger.Debug("repos created")
}

func (a *App) initServices() {
	a.balancesService = balances2.NewService(a.balancesRepo, a.transactionsRepo)

	a.logger.Debug("repos created")
}

func (a *App) initHandlers() {
	//пример: a.userHandler = users_handler.NewHandler(a.userRepo)
	a.logger.Debug("handlers created")
}

func (a *App) populateConfig() {
	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal()
	}

	a.config = cfg
}
