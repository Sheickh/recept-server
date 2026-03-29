package storage

import (
    "errors"
    "recept-server/models"
    "strconv"
    "sync"
    "time"
)

// MemoryStorage - хранилище в памяти
type MemoryStorage struct {
    mu       sync.RWMutex
    recipes  map[string]models.Recipe
    nextID   int
}

// NewMemoryStorage - создает новое хранилище
func NewMemoryStorage() *MemoryStorage {
    return &MemoryStorage{
        recipes: make(map[string]models.Recipe),
        nextID:  1,
    }
}

// GetAllRecipes - возвращает все рецепты
func (s *MemoryStorage) GetAllRecipes() ([]models.Recipe, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    
    recipes := make([]models.Recipe, 0, len(s.recipes))
    for _, recipe := range s.recipes {
        recipes = append(recipes, recipe)
    }
    return recipes, nil
}

// GetRecipeByID - возвращает рецепт по ID
func (s *MemoryStorage) GetRecipeByID(id string) (*models.Recipe, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    
    recipe, exists := s.recipes[id]
    if !exists {
        return nil, errors.New("рецепт не найден")
    }
    return &recipe, nil
}

// CreateRecipe - создает новый рецепт
func (s *MemoryStorage) CreateRecipe(recipe models.Recipe) (models.Recipe, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    
    // Генерируем ID только для рецепта
    recipe.ID = strconv.Itoa(s.nextID)
    s.nextID++
    
    // НЕ генерируем ID для ингредиентов - убираем этот код
    // НЕ генерируем ID для шагов - убираем этот код
    
    recipe.CreatedAt = time.Now().Format(time.RFC3339)
    s.recipes[recipe.ID] = recipe
    
    return recipe, nil
}

// UpdateRecipe - обновляет существующий рецепт
func (s *MemoryStorage) UpdateRecipe(id string, recipe models.Recipe) (*models.Recipe, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    
    _, exists := s.recipes[id]
    if !exists {
        return nil, errors.New("рецепт не найден")
    }
    
    recipe.ID = id
    s.recipes[id] = recipe
    return &recipe, nil
}

// DeleteRecipe - удаляет рецепт
func (s *MemoryStorage) DeleteRecipe(id string) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    
    _, exists := s.recipes[id]
    if !exists {
        return errors.New("рецепт не найден")
    }
    
    delete(s.recipes, id)
    return nil
}

// GetAllIngredients - возвращает все ингредиенты
func (s *MemoryStorage) GetAllIngredients() ([]models.Ingredient, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    
    ingredients := make([]models.Ingredient, 0)
    
    for _, recipe := range s.recipes {
        ingredients = append(ingredients, recipe.Ingredients...)
    }
    return ingredients, nil
}