package Models

type ImageModel struct {
	Name            string      `json:"name"`
	AlternativeText interface{} `json:"alternativeText"`
	Caption         interface{} `json:"caption"`
	Width           int         `json:"width"`
	Height          int         `json:"height"`
	Formats         struct {
		Large struct {
			Ext         string      `json:"ext"`
			Url         string      `json:"url"`
			Hash        string      `json:"hash"`
			Mime        string      `json:"mime"`
			Name        string      `json:"name"`
			Path        interface{} `json:"path"`
			Size        float64     `json:"size"`
			Width       int         `json:"width"`
			Height      int         `json:"height"`
			SizeInBytes int         `json:"sizeInBytes"`
		} `json:"large"`
		Small struct {
			Ext         string      `json:"ext"`
			Url         string      `json:"url"`
			Hash        string      `json:"hash"`
			Mime        string      `json:"mime"`
			Name        string      `json:"name"`
			Path        interface{} `json:"path"`
			Size        float64     `json:"size"`
			Width       int         `json:"width"`
			Height      int         `json:"height"`
			SizeInBytes int         `json:"sizeInBytes"`
		} `json:"small"`
		Medium struct {
			Ext         string      `json:"ext"`
			Url         string      `json:"url"`
			Hash        string      `json:"hash"`
			Mime        string      `json:"mime"`
			Name        string      `json:"name"`
			Path        interface{} `json:"path"`
			Size        float64     `json:"size"`
			Width       int         `json:"width"`
			Height      int         `json:"height"`
			SizeInBytes int         `json:"sizeInBytes"`
		} `json:"medium"`
		Thumbnail struct {
			Ext         string      `json:"ext"`
			Url         string      `json:"url"`
			Hash        string      `json:"hash"`
			Mime        string      `json:"mime"`
			Name        string      `json:"name"`
			Path        interface{} `json:"path"`
			Size        float64     `json:"size"`
			Width       int         `json:"width"`
			Height      int         `json:"height"`
			SizeInBytes int         `json:"sizeInBytes"`
		} `json:"thumbnail"`
	} `json:"formats"`
	Hash             string      `json:"hash"`
	Ext              string      `json:"ext"`
	Mime             string      `json:"mime"`
	Size             float64     `json:"size"`
	Url              string      `json:"url"`
	PreviewUrl       interface{} `json:"previewUrl"`
	Provider         string      `json:"provider"`
	ProviderMetadata interface{} `json:"provider_metadata"`
}
