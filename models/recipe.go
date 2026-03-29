package models

// Recipe - основная структура рецепта
type Recipe struct {
    ID          string       `json:"id"`                      // Только у рецепта есть ID
    Title       string       `json:"title" binding:"required"`
    Description string       `json:"description"`
    PrepTime    int          `json:"prepTime" binding:"required,min=1"`
    Difficulty  string       `json:"difficulty" binding:"required,oneof=easy medium hard"`
    Ingredients []Ingredient `json:"ingredients" binding:"required,min=1"`
    Steps       []Step       `json:"steps" binding:"required,min=1"`
    CreatedAt   string       `json:"createdAt"`
}

// Ingredient - структура ингредиента (без ID)
type Ingredient struct {
    Name     string `json:"name" binding:"required"`      // Название ингредиента
    Quantity string `json:"quantity" binding:"required"`  // Количество
}

// Step - структура шага приготовления (без ID)
type Step struct {
	OrderNumber int    `json:"orderNumber" binding:"required,min=1"` // Порядковый номер
    Instruction string `json:"instruction" binding:"required"` // Текст инструкции
}