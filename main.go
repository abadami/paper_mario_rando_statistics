package main

func main() {
	racesResponse := getRaceTitlesAndEntrantsByPage(1)

	response := getRaceDetails(racesResponse.Races[0].Name)

	println(response.Entrants[2].User.Name)
	println(response.Status.VerboseValue)
	println(response.Name)
}
