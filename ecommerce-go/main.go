package main

import (
	"encoding/json"
	"main/controllers"
	"main/database"
	"main/headers"
	"main/middleware"
	"net/http"
	"time"

	"github.com/rs/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	Description string
	Image       string
	Rating      int
	Price       int
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}

type ProductResponse struct {
	Product
	IsOnSale bool
}

func main() {
	database.Connect()
	database.SyncDatabase()

	router := http.NewServeMux()
	router.HandleFunc("/api/products", ProductHandler)
	router.HandleFunc("POST /signup", controllers.Signup)
	router.HandleFunc("OPTIONS /signup", controllers.Signup)
	router.HandleFunc("POST /login", controllers.Login)
	router.HandleFunc("OPTIONS /login", controllers.Login)
	router.HandleFunc("GET /validate", ValidateMiddleware)
	router.HandleFunc("POST /addproduct", AddProduct)
	router.HandleFunc("OPTIONS /addproduct", AddProduct)
	handler := cors.AllowAll().Handler(router)
	http.ListenAndServe(":8080", handler)

}

func ValidateMiddleware(w http.ResponseWriter, r *http.Request) {
	protected := middleware.RequireAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ProtectedEndpoint ve Validate fonksiyonlarını zincirleme
		controllers.Validate(w, r)
	}))
	protected.ServeHTTP(w, r)
}

func (p Product) IsOnSale() bool {
	return time.Since(p.CreatedAt).Hours() < 48
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	headers.SetHeaders(&w)
	dsn := "DBKEY"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Product{})
	var products []Product
	//db.Find(&products)
	db.Order("id desc, id").Find(&products)

	var response []ProductResponse
	for _, product := range products {
		response = append(response, ProductResponse{
			Product:  product,
			IsOnSale: product.IsOnSale(),
		})
	}

	json.NewEncoder(w).Encode(response)
	//json.NewEncoder(w).Encode(products)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	headers.SetHeaders(&w)
	dsn := "DBKEY"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createresult := db.Create(&product)
	if createresult.Error != nil {
		http.Error(w, "error ading product", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(product)
}
