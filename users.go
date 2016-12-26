package twitch

// UserType - describes a user on twitch
type UserType struct {
	Type        string            `json:"type"`
	Bio         string            `json:"bio"`
	Logo        string            `json:"logo"`
	DisplayName string            `json:"display_name"`
	CreatedAt   string            `json:"created_at"`
	UpdatedAt   string            `json:"updated_at"`
	ID          int               `json:"_id"`
	Name        string            `json:"name"`
	Links       map[string]string `json:"_links"`
}
