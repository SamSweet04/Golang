package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
}

type User struct {
	gorm.Model
	Username string
	Password string `gorm:"<-:false"`
}

var db *gorm.DB

func main() {
	var err error
	dsn := "user='' password='' dbname='' host=localhost port=5432 sslmode=disable TimeZone=UTC"
	db, err = gorm.Open("postgres", dsn)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Product{}, &User{})
	router := gin.Default()
	router.GET("/products", listProducts)
	router.GET("/products/:id", getProduct)
	router.POST("/productsCreate", createProduct)
	router.PUT("/productsUpdate/:id", updateProduct)
	router.DELETE("/productsDelete/:id", deleteProduct)
	router.GET("/search", searchProducts)
	router.GET("/sort", sortProducts)

	router.POST("/register", register)
	router.POST("/login", login)

	// Start the server
	router.Run(":8080")
}

// List all products
func listProducts(c *gin.Context) {
	products := []Product{}
	db.Find(&products)
	c.JSON(http.StatusOK, products)
}

// Get a specific product by ID
func getProduct(c *gin.Context) {
	productID := c.Param("id")
	product := Product{}
	db.First(&product, productID)
	if product.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// Create a new product
func createProduct(c *gin.Context) {
	var productData Product
	if err := c.ShouldBindJSON(&productData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&productData)
	c.JSON(http.StatusOK, productData)
}

// Update an existing product
func updateProduct(c *gin.Context) {
	productID := c.Param("id")
	existingProduct := Product{}
	db.First(&existingProduct, productID)
	if existingProduct.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	var updatedProductData Product
	if err := c.ShouldBindJSON(&updatedProductData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	existingProduct.Name = updatedProductData.Name
	existingProduct.Description = updatedProductData.Description
	existingProduct.Price = updatedProductData.Price
	db.Save(&existingProduct)
	c.JSON(http.StatusOK, existingProduct)
}
func deleteProduct(c *gin.Context) {
	productID := c.Param("id")
	existingProduct := Product{}
	db.First(&existingProduct, productID)
	if existingProduct.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	db.Delete(&existingProduct)
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}

// Register a new user
func register(c *gin.Context) {
	var userData User
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	userData.Password = string(hashedPassword)
	db.Create(&userData)
	c.JSON(http.StatusOK, userData)
}

// Authenticate a user and generate a JWT token
func login(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := User{}
	db.Where("username = ?", credentials.Username).First(&user)

	// If the user is not found or the password is incorrect, return a 401 error
	if user.ID == 0 || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)) != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	// Return the token in the response
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func searchProducts(c *gin.Context) {
	searchQuery := c.Query("q")
	products := []Product{}
	db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", searchQuery)).Find(&products)
	c.JSON(http.StatusOK, products)
}

func sortProducts(c *gin.Context) {
	products := []Product{}
	sortBy := c.Query("sort_by")
	sortOrder := c.Query("sort_order")

	query := db.Order(fmt.Sprintf("%s %s", sortBy, sortOrder))
	query.Find(&products)

	// Return the products in the response
	c.JSON(http.StatusOK, products)
}