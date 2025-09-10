package Models

type BitmapModel struct {
	Ext         string  `json:"ext"`
	Url         string  `json:"url"`
	Hash        string  `json:"hash"`
	Mime        string  `json:"mime"`
	Name        string  `json:"name"`
	Size        float64 `json:"size"`
	Width       int     `json:"width"`
	Height      int     `json:"height"`
	SizeInBytes int     `json:"sizeInBytes"`
}

type ImageModel struct {
	Name            string `json:"name"`
	AlternativeText string `json:"alternativeText"`
	Caption         string `json:"caption"`
	Width           int    `json:"width"`
	Height          int    `json:"height"`
	Formats         struct {
		Large     BitmapModel `json:"large"`
		Small     BitmapModel `json:"small"`
		Medium    BitmapModel `json:"medium"`
		Thumbnail BitmapModel `json:"thumbnail"`
	} `json:"formats"`
	Hash             string      `json:"hash"`
	Ext              string      `json:"ext"`
	Mime             string      `json:"mime"`
	Size             float64     `json:"size"`
	Url              string      `json:"url"`
	PreviewUrl       string      `json:"previewUrl"`
	Provider         string      `json:"provider"`
	ProviderMetadata interface{} `json:"provider_metadata"`
}
