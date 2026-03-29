package handlers

import (
    "net/http"
    "recept-server/models"
    "recept-server/storage"
    
    "github.com/gin-gonic/gin"
)

type RecipeHandler struct {
    storage *storage.MemoryStorage
}

func NewRecipeHandler(storage *storage.MemoryStorage) *RecipeHandler {
    return &RecipeHandler{storage: storage}
}

// GetAllRecipes - GET /recipes
func (h *RecipeHandler) GetAllRecipes(c *gin.Context) {
    recipes, err := h.storage.GetAllRecipes()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, recipes)
}

// GetRecipeByID - GET /recipes/:id
func (h *RecipeHandler) GetRecipeByID(c *gin.Context) {
    id := c.Param("id")
    
    recipe, err := h.storage.GetRecipeByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, recipe)
}

// CreateRecipe - POST /recipes
func (h *RecipeHandler) CreateRecipe(c *gin.Context) {
    var recipe models.Recipe
    
    if err := c.ShouldBindJSON(&recipe); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    created, err := h.storage.CreateRecipe(recipe)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusCreated, gin.H{
        "message": "Рецепт успешно создан",
        "id":      created.ID,
    })
}

// UpdateRecipe - PUT /recipes/:id
func (h *RecipeHandler) UpdateRecipe(c *gin.Context) {
    id := c.Param("id")
    var recipe models.Recipe
    
    if err := c.ShouldBindJSON(&recipe); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    updated, err := h.storage.UpdateRecipe(id, recipe)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "message": "Рецепт успешно обновлен",
        "recipe":  updated,
    })
}

// DeleteRecipe - DELETE /recipes/:id
func (h *RecipeHandler) DeleteRecipe(c *gin.Context) {
    id := c.Param("id")
    
    err := h.storage.DeleteRecipe(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"message": "Рецепт успешно удален"})
}

// GetAllIngredients - GET /ingredients
func (h *RecipeHandler) GetAllIngredients(c *gin.Context) {
    ingredients, err := h.storage.GetAllIngredients()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, ingredients)
}