package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Category struct {
	Name string
}

type ProductDesc struct {
	ID   int
	Name string
	Desc string
}

type ProductCategory struct {
	Product  int
	Category string
}

func Run() {
	router := gin.Default()

	router.LoadHTMLGlob("html/*.html")
	router.Static("/assets/", "resources")

	router.GET("/", handlerIndex)
	router.PUT("/category", handlerCreateCategory)
	router.PUT("/product", handlerCreateProduct)

	router.PUT("/product-category", handlerCreateProductCategory)

	router.GET("/category", handlerGetCategory)
	router.GET("/product", handlerGetProduct)

	router.Run("127.0.0.1:8080")
}

func handlerIndex(c *gin.Context) {

	c.HTML(200, "index", nil)
}

func handlerCreateCategory(c *gin.Context) {
	var category Category
	e := c.BindJSON(&category)
	if e != nil {
		fmt.Println("ERROR:", e)
		c.JSON(400, e.Error())
		return
	}

	_, e = connection.Exec(`INSERT INTO "Category"("Name") VALUES ($1)`, category.Name)
	if e != nil {
		fmt.Println("ERROR:", e)
		c.JSON(400, e.Error())
		return
	}
	c.JSON(200, nil)
}

func handlerCreateProduct(c *gin.Context) {
	type Product struct {
		Name string
		Desc string
	}
	var product Product
	e := c.BindJSON(&product)
	if e != nil {
		fmt.Println("ERROR:", e)
		c.JSON(400, e.Error())
		return
	}

	_, e = connection.Exec(`INSERT INTO "ProductDesc"("Name", "Description") VALUES ($1, $2)`, product.Name, product.Desc)
	if e != nil {
		fmt.Println("ERROR:", e)
		c.JSON(400, e.Error())
		return
	}
	c.JSON(200, nil)
}

func handlerGetProduct(c *gin.Context) {
	result, e := connection.Query(`
SELECT "ID", "Name", "Description"
FROM "ProductDesc"
ORDER BY "Name"`)
	if e != nil {
		fmt.Println("ERROR:", e)
		c.JSON(400, e.Error())
		return
	}

	defer result.Close()

	products := make([]ProductDesc, 0)
	product := ProductDesc{}

	for result.Next() {
		e = result.Scan(&product.ID, &product.Name, &product.Desc)
		if e != nil {
			fmt.Println("ERROR:", e)
			c.JSON(400, e.Error())
			return
		}

		products = append(products, product)
	}

	c.JSON(200, products)
}

func handlerGetCategory(c *gin.Context) {
	result, e := connection.Query(`
SELECT "Name"
FROM "Category"
ORDER BY "Name"`)
	if e != nil {
		fmt.Println("ERROR:", e)
		c.JSON(400, e.Error())
		return
	}

	defer result.Close()

	categories := make([]Category, 0)
	category := Category{}

	for result.Next() {
		e = result.Scan(&category.Name)
		if e != nil {
			fmt.Println("ERROR:", e)
			c.JSON(400, e.Error())
			return
		}

		categories = append(categories, category)
	}

	c.JSON(200, categories)
}

func handlerCreateProductCategory(c *gin.Context) {
	var pc ProductCategory
	e := c.BindJSON(&pc)
	if e != nil {
		fmt.Println("ERROR:", e)
		c.JSON(400, e.Error())
		return
	}

	_, e = connection.Exec(`INSERT INTO "ProductCategory" ("Product", "Category") VALUES ($1, $2)`, pc.Product, pc.Category)
	if e != nil {
		fmt.Println("ERROR:", e)
		c.JSON(400, e.Error())
		return
	}

	c.JSON(200, nil)
}
