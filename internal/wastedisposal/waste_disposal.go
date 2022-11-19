package wastedisposal

import "fmt"

type WasteDisposal struct {
	Id       int64
	Name     string
	City     string
	Location Location
}

func (w WasteDisposal) String() string {
	return fmt.Sprintf("ID: %d\n NAME: %s\n CITY: %s\n", w.Id, w.Name, w.City)
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (l Location) String() string {
	return fmt.Sprintf("Latitude: %f\n Longitude: %f\n", l.Latitude, l.Longitude)
}

// Define contract for retriving external informations
type WasteDisposalListRetriever interface {
	List() []WasteDisposal
}

type FindNearestRequest struct {
	WantedLoc Location `json:"location"`
}

type FindNearestResponse struct {
	Id       int64    `json:"id"`
	Name     string   `json:"name"`
	City     string   `json:"city"`
	Location Location `json:"location"`
}

type Service struct {
	retriever WasteDisposalListRetriever
}

func NewService(retriever WasteDisposalListRetriever) *Service {
	return &Service{
		retriever: retriever,
	}
}

func (s *Service) FindNearest(request FindNearestRequest) FindNearestResponse {
	centers := s.retriever.List()

	for _, c := range centers {
		fmt.Println(c)
	}

	candidate := centers[0]

	return FindNearestResponse{
		Id:       candidate.Id,
		Name:     candidate.Name,
		City:     candidate.City,
		Location: candidate.Location,
	}
}
