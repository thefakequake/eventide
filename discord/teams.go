package eventide

// https://discord.com/developers/docs/topics/teams#data-models-team-object
type Team struct {
	// A hash of the image of the team's icon
	Icon string `json:"icon"`

	// The unique ID of the team
	ID string `json:"id"`

	// The members of the team
	Members []*TeamMember `json:"members"`

	// The name of the team
	Name string `json:"name"`

	// The user ID of the current team owner
	OwnerUserID string `json:"owner_user_id"`
}

// https://discord.com/developers/docs/topics/teams#data-models-team-member-object
type TeamMember struct {
	// The user's membership state on the team
	MembershipState int `json:"membership_state"`

	// Will always be ["*"]
	Permissions []string `json:"permissions"`

	// The ID of the parent team of which they are a member
	TeamID string `json:"team_id"`

	// The avatar, discriminator, id, and username of the user
	User *User `json:"user"`
}
