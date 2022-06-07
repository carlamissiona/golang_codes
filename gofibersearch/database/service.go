package database

import "net/http"
import "github.com/go-playground/validator/v10"
import "github.com/gofiber/fiber/v2"
 
 
 
func (r *Repository) SetupRoutes(app *fiber.App) {
    api := app.Group("/api")
    api.Post("/create_feat", r.CreateSearches) 
    api.Get("/get_feat_searches/:id", r.GetFeaturedSearchesByID)
    api.Get("/get_all_feat_searches", r.GetAllFeaturedSearches) 
}

func (r *Repository) CreateSearches(context *fiber.Ctx) error {
    featd_search := FeaturedSearches{}

    err := context.BodyParser(&featd_search)
    if err != nil {
        context.Status(http.StatusUnprocessableEntity).JSON(
            &fiber.Map{"message": "request failed"})
        return err
    }
    validator := validator.New()
    err = validator.Struct(FeaturedSearches{})

    if err != nil {
        context.Status(http.StatusUnprocessableEntity).JSON(
            &fiber.Map{"message": err},
        )
        return err
    } 

    err = r.DB.Create(&featd_search).Error

    if err != nil {
        context.Status(http.StatusBadRequest).JSON(
            &fiber.Map{"message": "could not create FeaturedSearches"})
        return err
    }

    context.Status(http.StatusOK).JSON(&fiber.Map{
        "message": "Search has been successfully added",
    })
    return nil
}


func (r *Repository) GetFeaturedSearchesByID(context *fiber.Ctx) error {
    id := context.Params("id")
    featd_search := FeaturedSearches{}
    if id == "" {
        context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
            "message": "id cannot be empty",
        })
        return nil
    }

    err := r.DB.Where("id = ?", id).First(featd_search).Error
    if err != nil {
        context.Status(http.StatusBadRequest).JSON(
            &fiber.Map{"message": "could not get featured search"})
        return err
    }

    context.Status(http.StatusOK).JSON(&fiber.Map{
        "message": "Featured Search retrieved  successfully",
        "data":    featd_search,
    })
    return nil
}


func (r *Repository) GetAllFeaturedSearches(context *fiber.Ctx) error {
    featd_search := &[]FeaturedSearches{}

    err := r.DB.Find(featd_search).Error
    if err != nil {
        context.Status(http.StatusBadRequest).JSON(
            &fiber.Map{"message": "could not get books"})
        return err
    }

    context.Status(http.StatusOK).JSON(&fiber.Map{
        "message": "books gotten successfully",
        "data":    featd_search,
    })
    return nil

}