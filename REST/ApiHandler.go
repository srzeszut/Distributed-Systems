package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"sync"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/drivers", GetDriver)
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	return router
}

func GetDriver(c *gin.Context) {
	name := c.Query("name")
	driverChannel := make(chan Result, 3)
	flChannel := make(chan Result, 3)
	newsChannel := make(chan Result, 3)
	var wg sync.WaitGroup
	wg.Add(2)

	go func(name string) {
		log.Println("Getting driver data")
		defer wg.Done()

		driver, err := GetDriverFromApiByName(os.Getenv("F1_API_KEY"), os.Getenv("F1_DRIVERS_API_URL"), name)
		if err != nil {
			log.Println("Get Driver error: ", err)
			driverChannel <- Result{Err: err}
			flChannel <- Result{Err: err}
			//HandleErrors(err, c)
			return

		}

		fl, err := GetDriverFastestLapsFromApi(os.Getenv("F1_API_KEY"), os.Getenv("F1_RACES_API_URL"), int(driver.Id))
		if err != nil {
			log.Println("Get Fastest Laps error: ", err)
			driverChannel <- Result{Err: err}
			flChannel <- Result{Err: err}
			return

		}
		log.Println("sending data", Result{Data: driver}, Result{Data: fl})
		driverChannel <- Result{Data: driver}
		flChannel <- Result{Data: fl}
		log.Println("data sent")

		return

	}(name)

	go func(name string) {
		log.Println("Getting news data")
		defer wg.Done()
		news, err := GetDriverNewsFromApi(os.Getenv("NEWS_API_URL"), os.Getenv("NEWS_API_KEY"), name)
		if err != nil {
			log.Println("Get News error: ", err)
			newsChannel <- Result{Err: err}
			HandleErrors(err, c)
			return

		}
		//log.Println(news)
		log.Println("sending data", Result{Data: news})
		newsChannel <- Result{Data: news}
		log.Println("data sent")
		return
	}(name)

	var driver Driver
	var fl int
	var news []News
	wg.Wait()
	for i := 0; i < 2; i++ {
		select {
		case driverData := <-driverChannel:

			log.Println("driver data: ", driver)
			if driverData.Err != nil {
				HandleErrors(driverData.Err, c)
				return
			}
			driver = driverData.Data.(Driver)

		case fastestLaps := <-flChannel:
			if fastestLaps.Err != nil {
				HandleErrors(fastestLaps.Err, c)
				return
			}
			fl = fastestLaps.Data.(int)
		case newsData := <-newsChannel:
			log.Println("news data: ", newsData.Data.([]News))

			if newsData.Err != nil {
				HandleErrors(newsData.Err, c)
				return
			}
			news = newsData.Data.([]News)
		default:
			log.Println("No data available in any channel")
		}
	}

	log.Println("All goroutines finished")
	log.Println("Driver: ", driver)
	log.Println("news: ", news)
	close(driverChannel)
	close(flChannel)
	close(newsChannel)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"driver":      driver,
		"fastestLaps": fl,
		"news":        news,
		"driverFound": true,
	})

}

func HandleErrors(err error, c *gin.Context) {
	switch err.Error() {
	case "connection to service error":
		c.HTML(http.StatusServiceUnavailable, "index.html", gin.H{
			"driverFound": false})

	case "not found":
		c.HTML(http.StatusNotFound, "index.html", gin.H{
			"driverFound": false})

	case "json decode error":
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{
			"driverFound": false})

	default:
		log.Println("Unknown error: ", err)
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{
			"driverFound": false})

	}
}
