package idea

type UserDetails struct {
	User User `json:"User"`
}

type Group struct {
	Description string `json:"description,omitempty"`
	Enabled     bool   `json:"enabled"`
	GroupName   string `json:"groupName"`
	ID          int    `json:"id"`
}

type Groups struct {
	Group []Group `json:"group"`
}

type User struct {
	Attribute Attribute `json:"attribute"`
	Enabled   bool      `json:"enabled"`
	Groups    Groups    `json:"groups"`
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
}

type Attribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (userDetails *UserDetails) IsAdmin() bool {
	return userDetails != nil && userDetails.User.Role == "ADMIN"
}

func (userDetails *UserDetails) IsAdminOfEntity(entity string) bool {
	if userDetails != nil {
		for _, group := range userDetails.User.Groups.Group {
			if group.GroupName == entity+" Admins" {
				return true
			}
		}
	}
	return false
}
