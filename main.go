package main

import (
	"log"

	"demo/src/products/application"
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

	// Configuración de CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"}, // Frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // Necesario para autenticación con cookies o tokens
	}))

	// Manejo de preflight requests explícitamente
	router.OPTIONS("/*any", func(c *gin.Context) {
		log.Println("Preflight request:", c.Request.Method, c.Request.URL.Path) // Log para ver si se ejecuta
		c.Header("Access-Control-Allow-Origin", "http://localhost:4200")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Status(200) // Código 200 para la solicitud OPTIONS
	})
	mysql := infraestructure.NewMySQL()
	defer mysql.Close()

	productRepo := repositories.NewProductRepository(mysql.DB)
	createProduct := application.NewCreateProduct(*productRepo)
	getProducts := application.NewGetProducts(*productRepo)
	getProductById := application.NewGetProductById(*productRepo)
	updateProduct := application.NewUpdateProduct(*productRepo)
	deleteProduct := application.NewDeleteProduct(*productRepo)

	createProductController := infraestructure.NewCreateProductController(createProduct)
	getProductsController := infraestructure.NewGetProductsController(getProducts)
	getProductByIdController := infraestructure.NewGetProductByIdController(getProductById)
	updateProductController := infraestructure.NewUpdateProductController(updateProduct)
	deleteProductController := infraestructure.NewDeleteProductController(deleteProduct)

	clientRepo := clients_repositories.NewClientRepository(mysql.DB)
	createClient := clients_application.NewCreateClient(clientRepo)
	getClients := clients_application.NewGetClient(clientRepo)
	updateClient := clients_application.NewUpdateClient(clientRepo)
	deleteClient := clients_application.NewDeleteClient(clientRepo)

	createClientController := clients_infraestructure.NewCreateClientController(createClient)
	getClientsController := clients_infraestructure.NewGetClientsController(getClients)
	updateClientController := clients_infraestructure.NewUpdateClientController(updateClient)
	deleteClientController := clients_infraestructure.NewDeleteClientController(deleteClient)

	// Configurar el enrutador de Gin


	// Configuración de rutas para productos
	productRoutes := infraestructure.NewProductRoutes(
		createProductController,
		getProductsController,
		updateProductController,
		deleteProductController,
		getProductByIdController,
	)
	productRoutes.SetupRoutes(router)

	// Configuración de rutas para clientes
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
