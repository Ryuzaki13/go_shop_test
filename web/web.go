package web

import (
	"fmt"
	"github.com/gin-contrib/static"
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
	router.Use(static.Serve("/images/", static.LocalFile("data/", false)))

	router.GET("/", handlerIndex)
	router.PUT("/category", handlerCreateCategory)
	router.PUT("/product", handlerCreateProduct)

	router.PUT("/product-category", handlerCreateProductCategory)

	router.GET("/category", handlerGetCategory)
	router.GET("/product", handlerGetProduct)

	_ = router.Run("127.0.0.1:8080")
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
	form, e := c.MultipartForm()
	if e != nil {
		fmt.Println("ERROR:", e)
		c.JSON(400, e.Error())
		return
	}

	var filename string

	files := form.File["File"]
	//for _, file := range files {
	if len(files) == 1 {
		e = c.SaveUploadedFile(files[0], "data/"+files[0].Filename)
		if e != nil {
			fmt.Println("ERROR:", e)
			c.JSON(400, e.Error())
			return
		}
		filename = "data/" + files[0].Filename
	}
	//}

	name := form.Value["Name"]
	desc := form.Value["Desc"]

	if len(name) == 0 || len(desc) == 0 {
		fmt.Println("ERROR:", e)
		c.JSON(400, e.Error())
		return
	}

	if len(filename) == 0 {
		_, e = connection.Exec(`INSERT INTO "ProductDesc"("Name", "Description", "Image") VALUES ($1, $2, $3)`, name[0], desc[0], nil)
	} else {
		_, e = connection.Exec(`INSERT INTO "ProductDesc"("Name", "Description", "Image") VALUES ($1, $2, $3)`, name[0], desc[0], filename)
	}
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
