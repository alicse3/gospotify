package models

// GetBrowseCategoriesRequest represents the get browse categories request information.
type GetBrowseCategoriesRequest struct {
	Locale string
	Limit  int
	Offset int
}

// GetBrowseCategoryRequest represents the get browse category request information.
type GetBrowseCategoryRequest struct {
	CategoryId string // Required: The Spotify category ID for the category.
	Locale     string
}

// Categories represents the categories information retrieved from the Spotify API.
type Categories struct {
	Categories struct {
		Href     string     `json:"href"`
		Limit    int        `json:"limit"`
		Next     string     `json:"next"`
		Offset   int        `json:"offset"`
		Previous string     `json:"previous"`
		Total    int        `json:"total"`
		Items    []Category `json:"items"`
	} `json:"categories"`
}

// Category represents the category's information retrieved from the Spotify API.
type Category struct {
	Href  string `json:"href"`
	Icons []struct {
		Url    string `json:"url"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
	} `json:"icons"`
	Id   string `json:"id"`
	Name string `json:"name"`
}
