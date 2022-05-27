package basicSearch

type (
	Organization struct {
		ID            int32    `json:"_id"`
		URL           string   `json:"url"`
		ExternalID    string   `json:"external_id"`
		Name          string   `json:"name"`
		DomainNames   []string `json:"domain_names"`
		CreatedAt     string   `json:"created_at"`
		Details       string   `json:"details"`
		SharedTickets bool     `json:"shared_tickets"`
		Tags          []string `json:"tags"`
		Tickets       []string `json:"tickets,omitempty"`
		Users         []string `json:"users"`
	}

	User struct {
		ID               int32    `json:"_id"`
		URL              string   `json:"url"`
		ExternalID       string   `json:"external_id"`
		Name             string   `json:"name"`
		Alias            string   `json:"alias"`
		CreatedAt        string   `json:"created_at"`
		Active           bool     `json:"active"`
		Verified         bool     `json:"verified"`
		Shared           bool     `json:"shared"`
		Locale           string   `json:"locale"`
		Timezone         string   `json:"timezone"`
		LastLoginAt      string   `json:"last_login_at"`
		Email            string   `json:"email"`
		Phone            string   `json:"phone"`
		Signature        string   `json:"signature"`
		OrganizationID   int32    `json:"organization_id"`
		Tags             []string `json:"tags"`
		Suspended        bool     `json:"suspended"`
		Role             string   `json:"role"`
		OrganizationName string   `json:"organization_name,omitempty"`
		SubmittedTickets []string `json:"submitted_tickets,omitempty"`
		AssignedTickets  []string `json:"assigned_tickets,omitempty"`
	}

	Ticket struct {
		ID               string   `json:"_id"`
		URL              string   `json:"url"`
		ExternalID       string   `json:"external_id"`
		CreatedAt        string   `json:"created_at"`
		Type             string   `json:"type"`
		Subject          string   `json:"subject"`
		Description      string   `json:"description"`
		Priority         string   `json:"priority"`
		Status           string   `json:"status"`
		SubmitterID      int32    `json:"submitter_id"`
		AssigneeID       int32    `json:"assignee_id"`
		OrganizationID   int32    `json:"organization_id"`
		Tags             []string `json:"tags"`
		HasIncidents     bool     `json:"has_incidents"`
		DueAt            string   `json:"due_at"`
		Via              string   `json:"via"`
		SubmitterName    string   `json:"submitter_name,omitempty"`
		AssigneeName     string   `json:"assignee_name,omitempty"`
		OrganizationName string   `json:"organization_name,omitempty"`
	}

	OrgIndex struct {
		Indexes map[string]map[string][]*Organization
	}

	UserIndex struct {
		Indexes map[string]map[string][]*User
	}

	TicketIndex struct {
		Indexes map[string]map[string][]*Ticket
	}

	Indexing struct {
		OIndex *OrgIndex
		UIndex *UserIndex
		TIndex *TicketIndex
	}

	Result struct {
		Hits int32       `json:"hits"`
		Data interface{} `json:"data"`
	}
)
