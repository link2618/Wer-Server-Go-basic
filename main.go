package main

func main() {

	server := NewServer(":4000")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("POST", "/create", PostRequest)
	server.Handle("POST", "/create/user", UserPostRequest)
	server.Handle("POST", "/api", server.AddMiddleware(HandlerHome, CheckAuth(), Logging()))
	server.Listen()

}
