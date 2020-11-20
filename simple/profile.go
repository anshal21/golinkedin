package simple

// Profile is simplified Profile object
type Profile struct {
	Username          string `json:"username"`
	ProfilePicture    *Image `json:"profilePicture"`
	BackgroundPicture *Image `json:"BackgroundPicture"`
	FirstName         string `json:"firstName,omitempty"`
	LastName          string `json:"lastName,omitempty"`
	Headline          string `json:"headline,omitempty"`
	LocationName      string `json:"geoLocationName,omitempty"`
	About             string `json:"aboout,omitempty"`
}
