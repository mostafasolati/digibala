package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Video struct {
	Name        string
	Num         int
	ContentType string
	Period      int
}

var videos []*Video

type VideoServices interface {
	Create(*Video) error
	Read(int) (*Video, error)
	ReadAll() ([]*Video, error)
	Update(*Video) error
	Delete(int) error
}

type videoServicesStruct struct{}

func init() {
	videos = make([]*Video, 0)
	for i := 0; i < 5; i++ {
		videos = append(videos, &Video{
			Name:        fmt.Sprint("video", i+1),
			Num:         i,
			ContentType: fmt.Sprint("content", i+1),
			Period:      (i + 1) * 60,
		})
	}
}

func (v *videoServicesStruct) Create(video *Video) error {
	videos = append(videos, video)
	fmt.Println("Video created")
	return nil
}

func (v *videoServicesStruct) Read(id int) (*Video, error) {
	return videos[id], nil
}

func (v *videoServicesStruct) ReadAll() ([]*Video, error) {
	return videos, nil
}

func (v *videoServicesStruct) Update(video *Video) error {
	for id, arrayVideo := range videos {
		if arrayVideo.Num == video.Num {
			videos[id] = video
			break
		}
	}
	fmt.Println("Video updated")
	return nil
}

func (v *videoServicesStruct) Delete(id int) error {
	for arrayID, _ := range videos {
		if arrayID == id {
			videos = append(videos[:id], videos[id+1:]...)
			break
		}
	}
	fmt.Println("Video deleted")
	return nil
}

func newVideoService() VideoServices {
	return &videoServicesStruct{}
}

func handleCreate(s VideoServices) echo.HandlerFunc {
	return func(c echo.Context) error {
		var video *Video
		if err := c.Bind(video); err != nil {
			return err
		}
		err := s.Create(video)
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		return c.JSON(http.StatusOK, video)
	}
}

func handleRead(s VideoServices) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		video, err := s.Read(id)
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		return c.JSON(http.StatusOK, video)
	}
}

func handleReadAll(s VideoServices) echo.HandlerFunc {
	return func(c echo.Context) error {
		reqVideos, err := s.ReadAll()
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		return c.JSON(http.StatusOK, reqVideos)
	}
}

func handleUpdate(s VideoServices) echo.HandlerFunc {
	return func(c echo.Context) error {
		var video *Video
		if err := c.Bind(video); err != nil {
			return err
		}
		err := s.Update(video)
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		return c.JSON(http.StatusOK, video)
	}
}

func handleDelete(s VideoServices) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		err := s.Delete(id)
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		return c.String(http.StatusOK, "User deleted")
	}
}

func VideoRoutes(server *echo.Echo) {
	videoService := newVideoService()
	server.POST("/video", handleCreate(videoService))
	server.PUT("/video", handleUpdate(videoService))
	server.DELETE("/video/:id", handleDelete(videoService))
	server.GET("/video/:id", handleRead(videoService))
	server.GET("/video", handleReadAll(videoService))
}
