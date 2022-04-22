package hotelbeds

type AccommodationType struct {
	Code                 string  `json:"code"`
	TypeDescription      string  `json:"typeDescription"`
	TypeMultiDescription Content `json:"typeMultiDescription"`
}

type Content struct {
	Content      string `json:"content"`
	LanguageCode string `json:"languageCode"`
}

type Hotel struct {
	AccommodationType AccommodationType `json:"accommodationType"`
}
