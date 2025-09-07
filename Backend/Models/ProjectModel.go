package Models

type ProjectModel struct {
	Title           string       `json:"Title"`
	Locale          string       `json:"locale"`
	Description     string       `json:"Description"`
	TaskRealized    string       `json:"TaskRealized"`
	Technologies    []string     `json:"Technologies"`
	HasPrizes       bool         `json:"HasPrizes"`
	Slug            string       `json:"Slug"`
	IsMadeInCompany bool         `json:"IsMadeInCompany"`
	Thumbnail       ImageModel   `json:"Thumbnail"`
	Medias          []ImageModel `json:"Medias"`
	Prizes          []PrizeModel `json:"prizes"`
	Company         CompanyModel `json:"company"`
}
