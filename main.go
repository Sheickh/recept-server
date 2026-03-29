package main

import (
    "recept-server/handlers"
    "recept-server/storage"
    
    "github.com/gin-gonic/gin"
)

func main() {
    // Создаю хранилище
    store := storage.NewMemoryStorage()
    
    // Создаю хендлеры
    recipeHandler := handlers.NewRecipeHandler(store)
    
    // Создаю роутер
    router := gin.Default()
    
    // Маршруты для рецептов
    router.GET("/recipes", recipeHandler.GetAllRecipes)
    router.GET("/recipes/:id", recipeHandler.GetRecipeByID)
    router.POST("/recipes", recipeHandler.CreateRecipe)
    router.PUT("/recipes/:id", recipeHandler.UpdateRecipe)
    router.DELETE("/recipes/:id", recipeHandler.DeleteRecipe)
    
    // Маршрут для ингредиентов
    router.GET("/ingredients", recipeHandler.GetAllIngredients)
    
    // Запускаю сервер
    router.Run(":8080")
}