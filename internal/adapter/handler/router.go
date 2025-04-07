package handler

import (
	"database/sql"
	"go-forum/internal/adapter/handler/middleware"
	"go-forum/internal/adapter/storage"
	"log"
	"log/slog"
	"net/http"
)

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", middleware.Middleware(http.NotFoundHandler().ServeHTTP))

	return mux
}

type APIServer struct {
	address string
	mux     *http.ServeMux
	db      *sql.DB
	storage *storage.StorageInitializer
	logger  *slog.Logger
}

func NewAPIServer(address string, db *sql.DB, storage *storage.StorageInitializer, logger *slog.Logger) *APIServer {
	return &APIServer{
		address: address,
		mux:     Routes(),
		db:      db,
		storage: storage,
		logger:  logger,
	}
}

func (s *APIServer) Run() {
	// Logging http server initialization
	s.logger.Info("API server listening on " + s.address)

	// #######################
	// Repository Layer
	// #######################
	// inventoryRepository := repository.NewInventoryRepository(s.db)
	// menuRepository := repository.NewMenuRepository(s.db)
	// orderRepository := repository.NewOrderRepository(s.db)
	// transactionRepository := repository.NewTransactionRepository(s.db)

	// #######################
	// Business Layer
	// #######################
	// inventoryService := service.NewInventoryService(inventoryRepository)
	// menuService := service.NewMenuService(menuRepository)
	// orderService := service.NewOrderService(orderRepository)
	// transactionService := service.NewTransactionService(transactionRepository)

	// #######################
	// Presentation Layer
	// #######################
	// inventoryHandler := NewInventoryHandler(inventoryService, s.logger)
	// menuHandler := NewMenuHandler(menuService, s.logger)
	// orderHandler := NewOrderHandler(orderService, s.logger)
	// transactionHandler := NewTransactionHandler(transactionService, s.logger)

	// #######################
	// Registering Endpoints
	// #######################
	// inventoryHandler.RegisterEndpoints(s.mux)
	// menuHandler.RegisterEndpoints(s.mux)
	// orderHandler.RegisterEndpoints(s.mux)
	// transactionHandler.RegisterEndpoints(s.mux)

	// #######################
	// Repository Layer
	// #######################
	//repositoryLayer := repository.NewRepository(s.db, s.logger)

	// #######################
	// Business Layer
	// #######################
	//serviceLayer := service.NewService(repositoryLayer, s.logger)

	// #######################
	// Presentation Layer
	// #######################
	//httpLayer := handlers.NewHandler(serviceLayer, s.logger)

	s.logger.Info("API server listening on " + s.address)
	log.Fatal(http.ListenAndServe(s.address, s.mux))
}
