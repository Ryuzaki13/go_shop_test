function CreateCategory() {
    let name = document.querySelector("#CategoryName");

    let xhr = new XMLHttpRequest();
    xhr.open("PUT", "/category");
    xhr.send(JSON.stringify({
        name: name.value,
    }));
}

function CreateProduct() {
    let name = document.querySelector("#ProductName");
    let desc = document.querySelector("#ProductDesc");

    let xhr = new XMLHttpRequest();
    xhr.open("PUT", "/product");
    xhr.send(JSON.stringify({
        name: name.value,
        desc: desc.value,
    }));
}

function AddCategoryToProduct() {
    let product = document.querySelector("#SelectProduct");
    let category = document.querySelector("#SelectCategory");

    let xhr = new XMLHttpRequest();
    xhr.open("PUT", "/product-category");
    xhr.send(JSON.stringify({
        product: +product.value,
        category: category.value,
    }));
}

function LoadProducts() {
    let xhr = new XMLHttpRequest();
    xhr.open("GET", "/product");
    xhr.onload = buildProducts;
    xhr.send();
}

function buildProducts(event) {
    let products = JSON.parse(event.target.response);
    let productElement = document.querySelector("#SelectProduct");

    for (let i = 0; i < products.length; i++) {
        productElement.options.add(new Option(
            products[i].Name,
            products[i].ID,
        ));
    }
}

function LoadCategories() {
    let xhr = new XMLHttpRequest();
    xhr.open("GET", "/category");
    xhr.onload = buildCategories;
    xhr.send();
}

function buildCategories(event) {
    let products = JSON.parse(event.target.response);
    let productElement = document.querySelector("#SelectCategory");

    for (let i = 0; i < products.length; i++) {
        productElement.options.add(new Option(
            products[i].Name,
            products[i].Name,
        ));
    }
}

LoadProducts();
LoadCategories();