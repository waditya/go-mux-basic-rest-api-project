# go-mux-basic-rest-api-project

## Project Information 

### Product Table 

Product will have attributes - Id , name, quantity and price. The products will be stored in database 

API endpoints planned for the project 

1. GET  /products  -- Get a list of all the products
2. GET  /product/id -- gets the information about the respective product
3. POST /product    -- creates a new productbased on the given information form the user and saves it to the database
4. PUT /product/id  -- updates the respective product with the given information from the user  
5. DELETE /product/id -- deletes the respective product

## Packages to be Installed 

1. Third-party Router - gorilla/mux router
    * Package gorilla/mux implements a request router and dispatcher for matching incoming requests to their 
    respective handlers. 
    * The name mux stands for "HTTP request multiplexer"
    * Like the standard http.ServeMux,
        * mux.Router matches incoming requests against a list of registered routes and 
        calls a handler for the route that matches the URL or other conditions

    * Main features of gotilla/mux router 

        1. Support for method-based routing (GET,POST,PATCH, etc)
        2. Supports variable in URL paths (GET/{1})
        3. Supports regex route patterns (/product/{[a-z-]+})
    * Installation command - 
        go get github.com/gorilla/mux

