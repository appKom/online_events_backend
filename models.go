package main

type User struct {
	ID            string          `json:"id"`
	Grade         string          `json:"grade"`
	Email         string          `json:"email"`
	Image         *string         `json:"imageSrc,omitempty"`
	InterestGroup []InterestGroup `json:"interestGroup"`
}

type InterestGroup struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Members     []User `json:"members"`
}
