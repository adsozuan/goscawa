package goscawa

import "fmt"

type WasteDisposal struct {
	id       int64
	name     string
	city     string
	location Location
}

func (w WasteDisposal) String() string {
	return fmt.Sprintf("ID: %d\n NAME: %s\n CITY: %s\n", w.id, w.name, w.city)
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (l Location) String() string {
	return fmt.Sprintf("Latitude: %f\n Longitude: %f\n", l.Latitude, l.Longitude)
}

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
		Id:       candidate.id,
		Name:     candidate.name,
		City:     candidate.city,
		Location: candidate.location,
	}
}
