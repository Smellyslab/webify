# webify
simple http webserver maker in golang for golang but you need basic understanding of the net/http module to actually use it, this is because you need to know how to make pages and handle them etc.. also requires some knowledge of golang...

// you do need to know how to make webpages in golang but that its pretty easy i will make a templat here for you!


// --------------------------HTTP PAGE TEMPLATE-------------------------- // 

 func testHandler(w http.ResponseWriter, r *http.Request) {

  
  err := fmt.FprintF(w, "Hello")
  
  if err != nil {
  
  fmt.Println(err)
  
  }
  
 }

// --------------------------HTTP PAGE TEMPLATE-------------------------- // 

// Im Not going to get into how to make pages and serve html files to the user as that would take to long to write and you must learn how to do things on your own

// to handle the page created above go to your main function in the file you defined the http page template in and write inside of the func main {} 
// http.HandleFunc("/<page>", testHandler) replace the <page> with the page you want it to serve on without the <> 
// like this http.HandleFunc("/test", testHandler)
// but you need to define the func testHandler() {} outside of the func main() {}

