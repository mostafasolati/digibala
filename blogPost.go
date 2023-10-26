package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	// "golang.org/x/tools/go/analysis/passes/nilfunc"
)

type blogPost struct{
	header 			string
	context			string
	authorId		int
	creatAt			time.Time
}

type blogPostService struct{}

type blogPostServiceInterface interface{
	Create(*blogPost) 		error
	Put(*blogPost)			error
	Delete(int) 			error
	Search(int) (*blogPost,	error)
	List() ([]*blogPost, 	error)
}

func newBlogPostService()blogPostServiceInterface{
	return &blogPostService{}
}

func (b *blogPostService) Create(blogpost *blogPost) error{
	fmt.Println("new post created! At", blogpost.creatAt)
	return nil
}

func (b *blogPostService) Put(blogpost *blogPost) error{
	fmt.Println("Post updated!", blogpost)
	return nil
}

func (b *blogPostService) Delete(id int) error{
	fmt.Println("Post deleted!", id)
	return nil
}

func (b *blogPostService) Search(id int) (*blogPost, error){
	fmt.Println("Post founded!", id)
	return nil, nil
}

func (b *blogPostService) List() ([]*blogPost, error){
	fmt.Println("Blog Posts Are:")
	return nil, nil
}

func createBlogPostHandler(service blogPostServiceInterface) echo.HandlerFunc{
	return func(c echo.Context) error {
		blogpost := new(blogPost)
		if err := c.Bind(blogpost); err != nil{
			return err
		}
		service.Create(blogpost)
		return c.JSON(http.StatusOK, service)
	}
}

func updateBlogPostHandler(service blogPostServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		blopost := new(blogPost)
		if err := c.Bind(blopost); err != nil {
			return err
		}
		if err := service.Put(blopost); err != nil {
			return c.JSON(http.StatusBadRequest, Message{Message: "Error on updating blog post"})
		}
		return c.JSON(http.StatusOK, blopost)
	}
}

func deleteBlogPostHandler(service blogPostServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := service.Delete(id); err != nil {
			return c.JSON(http.StatusBadRequest, Message{Message: "Error on deleting blog post"})
		}
		return c.JSON(http.StatusOK, id)
	}
}

func findBlogPostHandler(service blogPostServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		var blogpost *blogPost
		var err error
		blogpost, err = service.Search(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, Message{Message: "Blog post do not exist"})
		}
		return c.JSON(http.StatusOK, blogpost)
	}
}

func listBlogPostHandler(service blogPostServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		blogposts, err := service.List()
		if err != nil {
			return c.JSON(http.StatusNotFound, Message{Message: "Blog posts findig faced error!"})
		}
		return c.JSON(http.StatusOK, blogposts)
	}
}

func blogPostRoutes(server *echo.Echo) {
	blogpostService := newBlogPostService()
	server.POST("/blogpost", createBlogPostHandler(blogpostService))
	server.PUT("/blogpost", updateBlogPostHandler(blogpostService))
	server.DELETE("/blogpost", deleteBlogPostHandler(blogpostService))
	server.GET("/blogpost/:id", findBlogPostHandler(blogpostService))
	server.GET("/blogpost", listBlogPostHandler(blogpostService))
}
