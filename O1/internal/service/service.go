package service

import (
	"O1/internal/model"
	"sort"
)

type RouteService struct {
	routes *[]model.Route
}

func NewRouteService(routes []model.Route) *RouteService {
	return &RouteService{
		routes: &routes,
	}
}

func (rs *RouteService) SortByDistance() {
	sort.SliceStable(*rs.routes, func(i, j int) bool {
		return (*rs.routes)[i].Distance < (*rs.routes)[j].Distance
	})
}

func (rs *RouteService) FilterByAverageDistance(lessThen float32) []model.Route {
	var filteredRoutes []model.Route
	for _, route := range *rs.routes {
		if route.Distance < lessThen {
			filteredRoutes = append(filteredRoutes, route)
		}
	}
	return filteredRoutes
}

func (rs *RouteService) FilterByStart(start string) []model.Route {
	var filteredRoutes []model.Route
	for _, route := range *rs.routes {
		if route.Start == start {
			filteredRoutes = append(filteredRoutes, route)
		}
	}
	return filteredRoutes
}

func (rs *RouteService) MaxStopsCount() []model.Route {
	var maxStopsCount []model.Route
	for _, route := range *rs.routes {
		if len(maxStopsCount) > 0 && route.StopsCount == maxStopsCount[0].StopsCount {
			maxStopsCount = append(maxStopsCount, route)
		} else if len(maxStopsCount) > 0 && route.StopsCount > maxStopsCount[0].StopsCount {
			maxStopsCount = []model.Route{route}
		} else if len(maxStopsCount) == 0 {
			maxStopsCount = append(maxStopsCount, route)
		}
	}
	return maxStopsCount
}
