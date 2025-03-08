package main

import (
	"log"

	"demo/src/products/application"
	"demo/src/products/application/messaging"
	"demo/src/products/infraestructure"
	"demo/src/products/infraestructure/repositories"

	clients_application "demo/src/Clients/applications"
	clients_infraestructure "demo/src/Clients/infraestructure"
	clients_repositories "demo/src/Clients/infraestructure/repositories"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Middleware para depuración (Verifica si las solicitudes llegan correctamente)
	router.Use(func(c *gin.Context) {
		log.Println("[Middleware] Solicitud recibida:", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})

	// Configuración de CORS (Permitir Angular y Vite)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200", "http://localhost:5173"}, // Angular y Vite
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Conexión a la base de datos
	mysql := infraestructure.NewMySQL()
	defer mysql.Close()

	rabbitMQ := infraestructure.NewRabbitMQ()
	defer rabbitMQ.Close()
	// Repositorios y servicios para productos
	productRepo := repositories.NewProductRepository(mysql.DB)
	publishProductCreated := messaging.NewPublishProductCreated(rabbitMQ)
	createProduct := application.NewCreateProduct(*productRepo, publishProductCreated)
	getProducts := application.NewGetProducts(*productRepo)
	getProductById := application.NewGetProductById(*productRepo)
	updateProduct := application.NewUpdateProduct(*productRepo)
	deleteProduct := application.NewDeleteProduct(*productRepo)

	// Controladores de productos
	createProductController := infraestructure.NewCreateProductController(createProduct)
	getProductsController := infraestructure.NewGetProductsController(getProducts)
	getProductByIdController := infraestructure.NewGetProductByIdController(getProductById)
	updateProductController := infraestructure.NewUpdateProductController(updateProduct)
	deleteProductController := infraestructure.NewDeleteProductController(deleteProduct)

	// Repositorios y servicios para clientes
	clientRepo := clients_repositories.NewClientRepository(mysql.DB)
	createClient := clients_application.NewCreateClient(clientRepo)
	getClients := clients_application.NewGetClient(clientRepo)
	updateClient := clients_application.NewUpdateClient(clientRepo)
	deleteClient := clients_application.NewDeleteClient(clientRepo)

	// Controladores de clientes
	createClientController := clients_infraestructure.NewCreateClientController(createClient)
	getClientsController := clients_infraestructure.NewGetClientsController(getClients)
	updateClientController := clients_infraestructure.NewUpdateClientController(updateClient)
	deleteClientController := clients_infraestructure.NewDeleteClientController(deleteClient)

	// Rutas de productos
	productRoutes := infraestructure.NewProductRoutes(
		createProductController,
		getProductsController,
		updateProductController,
		deleteProductController,
		getProductByIdController,
	)
	productRoutes.SetupRoutes(router)

	// Rutas de clientes
	clientsRoutes := clients_infraestructure.NewClientRoutes(
		createClientController,
		getClientsController,
		updateClientController,
		deleteClientController,
	)
	clientsRoutes.SetupRoutes(router)

	// Iniciar el servidor
	log.Println("[Main] Servidor corriendo en http://localhost:8080")
	router.Run(":8080")
}
