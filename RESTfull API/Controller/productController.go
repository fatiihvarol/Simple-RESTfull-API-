package Controller

import (
	"encoding/json"
	"first_project/Model"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

type ProductController struct {
	DB *gorm.DB
}

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{
		DB: db,
	}
}

func (pc *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	// JSON verisini çözümle (request body'den)
	var newProduct Model.Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newProduct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Yeni bir ürün oluştur
	createdProduct := Model.NewProduct(
		newProduct.Name,     // Ürün Adı
		newProduct.Type,     // Ürün Tipi
		newProduct.Category, // Ürün Kategorisi
		newProduct.Price,    // Ürün Fiyatı
	)

	// Ürünü veritabanına kaydet
	pc.DB.Create(&createdProduct)

	// Ürünü JSON formatına çevir
	productJSON, err := json.Marshal(createdProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// JSON yanıtını yaz
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(productJSON)
}

func (pc *ProductController) GetProduct(w http.ResponseWriter, r *http.Request) {
	// URL'den ürün kimliğini al
	parts := strings.Split(r.URL.Path, "/")
	productIDStr := parts[len(parts)-1]

	// Ürün kimliğini uint64'e dönüştür
	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Veritabanından ürünü al
	var retrievedProduct Model.Product
	result := pc.DB.First(&retrievedProduct, productID)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	// Ürünü JSON formatına çevir
	productJSON, err := json.Marshal(retrievedProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// JSON yanıtını yaz
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(productJSON)
}

func (pc *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// URL'den ürün kimliğini al
	parts := strings.Split(r.URL.Path, "/")
	productIDStr := parts[len(parts)-1]

	// Ürün kimliğini uint64'e dönüştür
	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Veritabanından ürünü sil
	result := pc.DB.Delete(&Model.Product{}, productID)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	// Başarılı yanıtı dön
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

// Diğer HTTP işlemleri için fonksiyonlar ekleyebilirsiniz...
