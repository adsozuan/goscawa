package ademe

import "adnotanumber.com/goscawa/internal/wastedisposal"

type FakeAdemeRetriever struct {
}

type AdemeWasteCenter struct {
	Id       int64  `json:"C_SERVICE"`
	Name     string `json:"N_SERVICE"`
	City     string `json:"L_VILLE_SITE"`
	GeoPoint string `json:"_geopoint"`
}

func (ar *FakeAdemeRetriever) List() []wastedisposal.WasteDisposal {
	fake_list := [3]AdemeWasteCenter{
		{Id: 1, Name: "Chambé clean", City: "Chambéry", GeoPoint: "49.04705842759184,3.3984973997281323"},
		{Id: 2, Name: "Montbé clean", City: "Montbéliard", GeoPoint: "49.14705842759184,2.1984973997281323"},
		{Id: 3, Name: "Melun clean", City: "Melun", GeoPoint: "49.24705842759184,1.1984973997281323"},
	}

	outputs := [3]wastedisposal.WasteDisposal{}

	for i, wc := range fake_list {

		outputs[i] = wastedisposal.WasteDisposal{
			Id:   wc.Id,
			Name: wc.Name,
			City: wc.City,
		}

	}

	return outputs[:]

}
