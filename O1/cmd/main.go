package main

import (
	"O1/internal/scanner"
	"O1/internal/service"
	"fmt"
)

func main() {
	routes := scanner.Scan("./data/route.json")
	fmt.Println(routes)
	routeService := service.NewRouteService(routes)
	routeService.SortByDistance()
	fmt.Println(routes)
	distance := routeService.FilterByAverageDistance(100)
	fmt.Println(distance)
	byStart := routeService.FilterByStart("A")
	fmt.Println(byStart)
	maxStopsCount := routeService.MaxStopsCount()
	fmt.Println(maxStopsCount)
}
