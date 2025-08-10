package racetime_service

import (
	"fmt"
	"sync"

	"github.com/abadami/randomizer-statistics/domain"
	"github.com/jackc/pgx/v5"
)

type RacetimeRepository interface {
	GetRaceTitlesAndEntrantsByPage(pageNum int) domain.RaceByPageResponse
	GetRaceDetails(raceName string) domain.RaceDetail
}

type RaceRepository interface {
	GetRaceByName(queryArgs pgx.NamedArgs) (string, error)
	GetRacesByRaceEntrant(request domain.StatisticsRequest) ([]domain.RaceEntrantAndRaceRecord, error)
	InsertRaceDetails(details domain.RaceDetail) error
}

type TaskLogRepository interface {
	InsertTaskLog(success bool, racesFetched int) error
}

type RacetimeService struct {
	racetimeRepo RacetimeRepository
	raceRepo     RaceRepository
	taskLogRepo  TaskLogRepository
}

func NewService(ra RacetimeRepository, rac RaceRepository, task TaskLogRepository) *RacetimeService {
	return &RacetimeService{
		racetimeRepo: ra,
		raceRepo:     rac,
		taskLogRepo:  task,
	}
}

type GetPageWorkerParams struct {
	Id      int
	Jobs    <-chan int
	Results chan<- string
	Pagewg  *sync.WaitGroup
	Racewg  *sync.WaitGroup
}

type GetRaceWorkerParams struct {
	Id      int
	Jobs    <-chan string
	Results chan<- domain.RaceDetail
	Wg      *sync.WaitGroup
}

func (service *RacetimeService) GetPageWorker(params GetPageWorkerParams) {
	for j := range params.Jobs {
		response := service.racetimeRepo.GetRaceTitlesAndEntrantsByPage(j)
		for race := range response.Races {
			params.Racewg.Add(1)
			params.Results <- response.Races[race].Name
		}
		params.Pagewg.Done()
	}
}

func (service *RacetimeService) GetRaceWorker(params GetRaceWorkerParams) {
	for job := range params.Jobs {
		queryArgs := pgx.NamedArgs{
			"raceName": job,
		}

		_, queryError := service.raceRepo.GetRaceByName(queryArgs)

		//TODO: More specific error handling
		if queryError == nil {
			params.Wg.Done()
			continue
		}

		if queryError != pgx.ErrNoRows {
			fmt.Println("Weird Error", queryError)
			params.Wg.Done()
			continue
		}

		response := service.racetimeRepo.GetRaceDetails(job)

		insertRaceError := service.raceRepo.InsertRaceDetails(response)

		if insertRaceError != nil {
			params.Wg.Done()
			continue
		}

		params.Results <- response
		params.Wg.Done()
	}
}

func (service *RacetimeService) FetchRaceDetailsFromRacetime() {
	fmt.Println("Fetching race details from racetime...")

	racesResponse := service.racetimeRepo.GetRaceTitlesAndEntrantsByPage(1)

	jobs := make(chan int, racesResponse.NumPages)
	detailJobs := make(chan string, racesResponse.Count)
	results := make(chan domain.RaceDetail, racesResponse.Count)

	pagewg := new(sync.WaitGroup)
	racewg := new(sync.WaitGroup)

	for _, race := range racesResponse.Races {
		detailJobs <- race.Name
		racewg.Add(1)
	}

	pagewg.Add(racesResponse.NumPages - 1)

	for w := 0; w <= 10; w++ {
		go service.GetPageWorker(GetPageWorkerParams{
			Id:      w,
			Jobs:    jobs,
			Results: detailJobs,
			Pagewg:  pagewg,
			Racewg:  racewg,
		})
	}

	for page := 2; page <= racesResponse.NumPages; page++ {
		jobs <- page
	}
	close(jobs)

	for w := 0; w <= 1000; w++ {
		go service.GetRaceWorker(GetRaceWorkerParams{
			Id:      w,
			Jobs:    detailJobs,
			Results: results,
			Wg:      racewg,
		})
	}

	pagewg.Wait()

	//We know we have all of the detailJobs, so we can close here
	close(detailJobs)

	racewg.Wait()

	//We know we have all the results now, so close the channels
	close(results)

	taskLogError := service.taskLogRepo.InsertTaskLog(true, len(results))

	if taskLogError != nil {
		fmt.Println("Error inserting task log error. Oh no!")
	}

	fmt.Println("Finished fetching race data from racetime!")
}
