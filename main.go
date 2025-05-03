package main

import (
	"fmt"
	"math"
)

type Point struct {
	lat float64
	lon float64
}

func main() {
	point1 := Point{lat: 50.06638888888889, lon: -5.714722222222222}
	point2 := Point{lat: 58.64388888888889, lon: -3.0700000000000003}

	haversineDistance := haversine(point1, point2)
	fmt.Printf("Distance: %.2f km\n", haversineDistance)
}

// This is a simple program to calculate the distance between two points on the earth's surface
// using the Haversine formula.
// The Haversine formula is an equation that gives the distance between two points on the surface of a sphere
// given their longitudes and latitudes.
// The formula is:
// a = sin²(Δφ/2) + cos φ₁ ⋅ cos φ₂ ⋅ sin²(Δλ/2)
// c = 2 ⋅ atan2(√a, √(1−a))
// d = R ⋅ c
// where φ is latitude, λ is longitude, R is the earth’s radius (mean radius = 6,371 km)
// The program takes two points as input and calculates the distance between them in kilometers.
func haversine(point1, point2 Point) float64 {
	// convert degrees to radians
	point1.lat = point1.lat * (3.141592653589793 / 180)
	point1.lon = point1.lon * (3.141592653589793 / 180)
	point2.lat = point2.lat * (3.141592653589793 / 180)
	point2.lon = point2.lon * (3.141592653589793 / 180)

	// haversine formula
	dlat := point2.lat - point1.lat
	dlon := point2.lon - point1.lon
	a := (math.Sin(dlat/2) * math.Sin(dlat/2)) + (math.Cos(point1.lat) * math.Cos(point2.lat) * math.Sin(dlon/2) * math.Sin(dlon/2))
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	r := float64(6371) // Radius of earth in kilometers
	return c * r
}
