package main
 
import (
	"encoding/json"
	"fmt"
	_ "github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
  "net/url"
	pc "codetunadev/paymentcode" 

	_ "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/charge"
)
  
func main() {
  
type NasaItem struct {
	Date           string `json:"date"`
	Explanation    string `json:"explanation"`
	ServiceVersion string `json:"service_version"`
	Title          string `json:"title"`
	Hdurl          string `json:"hdurl"`
	Url            string `json:"url"`
	MediaType      string `json:"media_type"`
  Copyright      string `json:"copyright"`
}

  var ImagesParsed []NasaItem 
  var NasaImgsParsed []NasaItem 
  location := url.URL{Path: "/error-page-500",}
  
	r := gin.Default() 
  r.StaticFS("/assets", http.Dir("assets")) 
  r.StaticFS("/images", http.Dir("images")) 
	// r.Use(static.Serve("/assets", static.LocalFile("assets", false)))
  r.LoadHTMLGlob("templates/*") 
	// r.GET("/app", func(c *gin.Context) {
	// 	http.ServeFile(c.Writer, c.Request, "index.html") 
	// })

	r.GET("/error-page-500", func(c *gin.Context) { 
		
       
      	c.HTML(http.StatusOK, "error.html", gin.H{
          "Title": "Daily Nasa Images Error Page", 
        })
      
  
  })

 r.GET("/checkout", func(c *gin.Context) { 
		
       
      	c.HTML(http.StatusOK, "error.html", gin.H{
          "Title": "Daily Nasa Images Error Page", 
        })
      
  
  })

  
	r.GET("/", func(c *gin.Context) { 
		 

		resp, err := http.Get("https://api.nasa.gov/planetary/apod?api_key=oeKTapuXFk5txKkL7eLJec0QXHAFKPgYlOteMDZz&start_date=2020-01-01&end_date=2020-01-20")  

     
		// resp, err := http.Get("https://api.nasa.gov/planetary/apod?api_key=DEMO_KEY&start_date=2020-01-01&end_date=2020-01-02") 
		if err != nil {
			fmt.Println("No response from request")

      c.Redirect(http.StatusFound, location.RequestURI())
      
        
		}
 

		body, err := ioutil.ReadAll(resp.Body) // response body is []byte
    if err != nil {
			fmt.Println("No response from request")  
      fmt.Println(err) 

      c.Redirect(http.StatusFound, location.RequestURI())
       
		}     
    
    err = json.Unmarshal(body, &ImagesParsed) 
    // fmt.Println(err)  
    // fmt.Println("====ImagesParsed Json request ImagesParsed==========") 
    // fmt.Println(len(ImagesParsed)) 
    // fmt.Println(ImagesParsed)
    // fmt.Println("Json request ImagesParsed==========ERRORCopyrightCopyright[index][index][index][index]") 
    var result map[string]interface{}

    // Unmarshal or Decode the JSON to the interface.
    err = json.Unmarshal(body, &result)
    if err != nil {
			fmt.Println("Unmarshal Error from request")  
      fmt.Println(err) 

      // c.Redirect(http.StatusFound, location.RequestURI())
      
 		}   
    // fmt.Println(result)  
    // fmt.Println(err)  
    // fmt.Println("==== string(body)string(body)string(body)Json request resultresultresultresult ImagesParsed") 
    // // fmt.Println(string(body)) 
    // fmt.Println("Json requestresultresultresultresult ImagesParsed==========") 
	 
    // fmt.Println("Json request ImagesParsed") 
    // fmt.Println(ImagesParsed) 
    // fmt.Println("Json request string body") 
    // // fmt.Println(string(body)) 
    // fmt.Println("Json request  body") 
    // // fmt.Println(body) 
    // fmt.Println("====Json request=====")  
    
		// resp, err = http.Get("https://api.nasa.gov/planetary/apod?api_key=oeKTapuXFk5txKkL7eLJec0QXHAFKPgYlOteMDZz&start_date=2020-01-01&end_date=2020-01-02")  

     
		// resp, err := http.Get("https://api.nasa.gov/planetary/apod?api_key=DEMO_KEY&start_date=2020-01-01&end_date=2020-01-02") 
	 
    //resp, err = http.Get("https://api.nasa.gov/planetary/apod?api_key=DEMO_KEY&start_date=2020-01-01&end_date=2020-01-02") 

    
    err=nil 
		resp, err = http.Get("https://api.nasa.gov/planetary/apod?api_key=oeKTapuXFk5txKkL7eLJec0QXHAFKPgYlOteMDZz&start_date=2021-01-01&end_date=2021-01-15")  
 
  	if err != nil { 
  			fmt.Println("No response from request")
        NasaImgsParsed = ImagesParsed
      	c.HTML(http.StatusOK, "index.html", gin.H{
    			"title": "Daily Nasa Images",
          "json" : ImagesParsed,
    			"NasaImgs":  NasaImgsParsed, 
    		}) 
    } else{ 
      		body, err = ioutil.ReadAll(resp.Body) // response body is []byte 
          if err != nil {
      			fmt.Println("Error on parsing with ioutil ")    
            fmt.Println(err) 
            // c.Redirect(http.StatusFound, location.RequestURI())
      		}     
           
          err = json.Unmarshal(body, &NasaImgsParsed) 
          fmt.Println(err) 
           if err != nil {
      			fmt.Println("Error on parsing with ioutil ")    
            fmt.Println(err) 
            // c.Redirect(http.StatusFound, location.RequestURI())
      		}    
          fmt.Println("====ImagesParsed Json request ImagesParsed==========") 
          fmt.Println(len(NasaImgsParsed)) 
          fmt.Println(NasaImgsParsed) 
          
          defer resp.Body.Close()
      		c.HTML(http.StatusOK, "index.html", gin.H{
      			"title": "Daily Nasa Images",
            "json" : ImagesParsed,
      			"NasaImgs":  NasaImgsParsed, 
      		}) 

    }
	})
  
  pc.Pay()
  
  r.Run()

}
