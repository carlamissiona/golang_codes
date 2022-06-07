package database

import "gorm.io/gorm"

 
type Repository struct { 
    DB *gorm.DB
}
  
  
type FeaturedSearches struct {
	ID         uint    `gorm:"primary key;autoIncrement" json:"id"`
	Author     *string `json:"author"`
	Title      *string `json:"title"`
	Source     *string `json:"source"`
	Url        *string `json:"url"`
	Query      *string `json:"query"`
	Highlights *string `json:"highlight"`
	Page       *string `json:"page"`
}

   
     
type CommentsAnnotations struct {
	ID             uint    `gorm:"primary key;autoIncrement" json:"id"`
	UserId       *string `json:"user_id"`
	FeatSearchId  *string `json:"feat_search_id"`
	Comment       *string `json:"comment"`  
}
 

 
// Migration To PG 
func MigrateFeaturedSearches(db *gorm.DB) error {
	err := db.AutoMigrate(&FeaturedSearches{})
	return err
} 
func MigrateCommentsAnnotations(db *gorm.DB) error {
	err := db.AutoMigrate(&CommentsAnnotations{})
	return err
} 

// Serializers

type FeaturedSearch struct { 
	Author     *string `json:"author" validate:"required"`
	Title      *string `json:"title"validate:"required"`
	Source     *string `json:"source"validate:"required"`
	Url        *string `json:"url" validate:"required"`
	Query      *string `json:"query" validate:"required"`
	Highlights *string `json:"highlight" validate:"required"`
	Page       *string `json:"page" validate:"required"` 
}
  
type CommentAnnotation struct {
	ID             uint    `gorm:"primary key;autoIncrement" json:"id"`
	UserId       *string `json:"user_id" validate:"required`
	FeatSearchId  *string `json:"feat_search_id" validate:"required`
	Comment       *string `json:"comment validate:"required"`  
}