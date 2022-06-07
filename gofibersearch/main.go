package main
 
import localdb "carlafibersearch/database"
import "log"
import "github.com/gofiber/fiber/v2"
import "os" 

func main() {

     config := &localdb.Config{
        Host:     os.Getenv("DB_HOST"),
        Port:     os.Getenv("DB_PORT"),
        Password: os.Getenv("DB_PASS"),
        User:     os.Getenv("DB_USER"),
        SSLMode:  os.Getenv("DB_SSLMODE"),
        DBName:   os.Getenv("DB_NAME"),
    } 
  
    db, err := localdb.NewConnection(config)
    if err != nil {
        log.Fatal("could not load database")
    }

    err = localdb.MigrateFeaturedSearches(db)
    if err != nil {
        log.Fatal("could not migrate FeaturedSearches to db")
    }

    err = localdb.MigrateCommentsAnnotations(db)
    if err != nil {
        log.Fatal("could not migrate CommentsAnnotations to db")
    }
 

    r := &localdb.Repository{
        DB: db,
    }
    app := fiber.New()
    r.SetupRoutes(app)

    app.Listen(":8080")

   
}
 