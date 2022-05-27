package basicSearch

import (
	"fmt"
	"strconv"
)

const (
	idField               = "_id"
	urlField              = "url"
	externalIDField       = "external_id"
	nameField             = "name"
	aliasField            = "alias"
	orgDomainNamesField   = "domain_names"
	createdAtField        = "created_at"
	activeField           = "active"
	verifiedField         = "verified"
	sharedField           = "shared"
	localeField           = "locale"
	timeZoneField         = "timezone"
	lastLoginAtField      = "last_login_at"
	emailField            = "email"
	phoneField            = "phone"
	signatureField        = "signature"
	organizationIDField   = "organization_id"
	orgDetailField        = "details"
	orgSharedTicketsField = "shared_tickets"
	tagsField             = "tags"
	suspendedField        = "suspended"
	roleField             = "role"
	typeField             = "type"
	subjectField          = "subject"
	descriptionField      = "description"
	priorityField         = "priority"
	statusField           = "status"
	submitterIDField      = "submitter_id"
	assigneeIDField       = "assignee_id"
	hasIncidentsField     = "has_incidents"
	dueAtField            = "due_at"
	viaField              = "via"
)

var (
	SupportedOrgFields = []string{idField, urlField, externalIDField, nameField, orgDomainNamesField, createdAtField,
		orgDetailField, orgSharedTicketsField, tagsField}
	SupportedUserFields = []string{idField, urlField, externalIDField, nameField, aliasField, createdAtField,
		activeField, verifiedField, sharedField, localeField, timeZoneField, lastLoginAtField, emailField, phoneField,
		signatureField, organizationIDField, tagsField, suspendedField, roleField}
	SupportedTicketFields = []string{idField, urlField, externalIDField, createdAtField, typeField, subjectField,
		descriptionField, priorityField, statusField, submitterIDField, assigneeIDField, organizationIDField, tagsField,
		hasIncidentsField, dueAtField, viaField}
)

func IndexOrg(data []*Organization) *OrgIndex {
	var (
		i = OrgIndex{}
	)
	i.Indexes = make(map[string]map[string][]*Organization)
	for _, f := range SupportedOrgFields {
		i.Indexes[f] = make(map[string][]*Organization)
	}

	// Indexing
	for _, d := range data {
		d := d

		indexOrgField(i.Indexes[idField], d, d.ID)
		indexOrgField(i.Indexes[urlField], d, d.URL)
		indexOrgField(i.Indexes[externalIDField], d, d.ExternalID)
		indexOrgField(i.Indexes[nameField], d, d.Name)
		indexOrgField(i.Indexes[createdAtField], d, d.CreatedAt)
		indexOrgField(i.Indexes[orgDetailField], d, d.Details)
		indexOrgField(i.Indexes[orgSharedTicketsField], d, d.SharedTickets)
		for _, domainName := range d.DomainNames {
			indexOrgField(i.Indexes[orgDomainNamesField], d, domainName)
		}
		for _, tag := range d.Tags {
			indexOrgField(i.Indexes[tagsField], d, tag)
		}
	}

	return &i
}

func IndexUser(data []*User) *UserIndex {
	var (
		i UserIndex
	)

	i.Indexes = make(map[string]map[string][]*User)
	for _, f := range SupportedUserFields {
		i.Indexes[f] = make(map[string][]*User)
	}

	// Indexing
	for _, d := range data {
		d := d

		indexUserField(i.Indexes[idField], d, d.ID)
		indexUserField(i.Indexes[urlField], d, d.URL)
		indexUserField(i.Indexes[externalIDField], d, d.ExternalID)
		indexUserField(i.Indexes[nameField], d, d.Name)
		indexUserField(i.Indexes[aliasField], d, d.Alias)
		indexUserField(i.Indexes[createdAtField], d, d.CreatedAt)
		indexUserField(i.Indexes[activeField], d, d.Active)
		indexUserField(i.Indexes[verifiedField], d, d.Active)
		indexUserField(i.Indexes[sharedField], d, d.Shared)
		indexUserField(i.Indexes[localeField], d, d.Locale)
		indexUserField(i.Indexes[timeZoneField], d, d.Timezone)
		indexUserField(i.Indexes[lastLoginAtField], d, d.LastLoginAt)
		indexUserField(i.Indexes[emailField], d, d.Email)
		indexUserField(i.Indexes[phoneField], d, d.Phone)
		indexUserField(i.Indexes[signatureField], d, d.Signature)
		indexUserField(i.Indexes[organizationIDField], d, d.OrganizationID)
		indexUserField(i.Indexes[suspendedField], d, d.Suspended)
		indexUserField(i.Indexes[roleField], d, d.Role)
		for _, tag := range d.Tags {
			indexUserField(i.Indexes[tagsField], d, tag)
		}
	}

	return &i
}

func IndexTicket(data []*Ticket) *TicketIndex {
	var (
		i TicketIndex
	)

	i.Indexes = make(map[string]map[string][]*Ticket)
	for _, f := range SupportedTicketFields {
		i.Indexes[f] = make(map[string][]*Ticket)
	}

	// Indexing
	for _, d := range data {
		d := d

		indexTicketField(i.Indexes[idField], d, d.ID)
		indexTicketField(i.Indexes[urlField], d, d.URL)
		indexTicketField(i.Indexes[externalIDField], d, d.ExternalID)
		indexTicketField(i.Indexes[createdAtField], d, d.CreatedAt)
		indexTicketField(i.Indexes[typeField], d, d.Type)
		indexTicketField(i.Indexes[subjectField], d, d.Subject)
		indexTicketField(i.Indexes[descriptionField], d, d.Description)
		indexTicketField(i.Indexes[priorityField], d, d.Priority)
		indexTicketField(i.Indexes[statusField], d, d.Status)
		indexTicketField(i.Indexes[submitterIDField], d, d.SubmitterID)
		indexTicketField(i.Indexes[assigneeIDField], d, d.AssigneeID)
		indexTicketField(i.Indexes[organizationIDField], d, d.OrganizationID)
		indexTicketField(i.Indexes[hasIncidentsField], d, d.HasIncidents)
		indexTicketField(i.Indexes[dueAtField], d, d.DueAt)
		indexTicketField(i.Indexes[viaField], d, d.Via)

		for _, tag := range d.Tags {
			indexTicketField(i.Indexes[tagsField], d, tag)
		}
	}

	return &i
}

func UpdateRelatedEntities(indexing Indexing) {
	// Update org
	go func() {
		for k, v := range indexing.OIndex.Indexes[idField] {
			tickets := indexing.TIndex.Indexes[organizationIDField]
			for _, ticket := range tickets[k] {
				v[0].Tickets = append(v[0].Tickets, ticket.Subject)
			}

			users := indexing.UIndex.Indexes[organizationIDField]
			for _, user := range users[k] {
				v[0].Users = append(v[0].Users, user.Name)
			}
		}
	}()

	// Update user
	go func() {
		for k, v := range indexing.UIndex.Indexes[idField] {
			orgIDIdx := indexing.OIndex.Indexes[idField]
			orgID := fmt.Sprintf("%v", v[0].OrganizationID)
			if org, ok := orgIDIdx[orgID]; ok {
				v[0].OrganizationName = org[0].Name
			}

			submittedTickets := indexing.TIndex.Indexes[submitterIDField]
			for _, t := range submittedTickets[k] {
				v[0].SubmittedTickets = append(v[0].SubmittedTickets, t.Subject)
			}

			assignedTickets := indexing.TIndex.Indexes[assigneeIDField]
			for _, t := range assignedTickets[k] {
				v[0].AssignedTickets = append(v[0].AssignedTickets, t.Subject)
			}
		}
	}()

	// update ticket
	go func() {
		for _, v := range indexing.TIndex.Indexes[idField] {
			v[0].OrganizationName = getOrgName(strconv.Itoa(int(v[0].OrganizationID)), indexing)
			v[0].SubmitterName = getUserName(strconv.Itoa(int(v[0].SubmitterID)), indexing)
			v[0].AssigneeName = getUserName(strconv.Itoa(int(v[0].AssigneeID)), indexing)
		}
	}()
}

func getOrgName(id string, indexing Indexing) string {
	orgIDIdx := indexing.OIndex.Indexes[idField]
	if org, ok := orgIDIdx[id]; ok {
		return org[0].Name
	}
	return ""
}

func getUserName(id string, indexing Indexing) string {
	userIDIdx := indexing.UIndex.Indexes[idField]
	if user, ok := userIDIdx[id]; ok {
		return user[0].Name
	}
	return ""
}

func indexOrgField(idx map[string][]*Organization, data *Organization, val interface{}) {
	valStr := fmt.Sprintf("%v", val)
	idx[valStr] = append(idx[valStr], data)
}

func indexUserField(idx map[string][]*User, data *User, val interface{}) {
	valStr := fmt.Sprintf("%v", val)
	idx[valStr] = append(idx[valStr], data)
}

func indexTicketField(idx map[string][]*Ticket, data *Ticket, val interface{}) {
	valStr := fmt.Sprintf("%v", val)
	idx[valStr] = append(idx[valStr], data)
}
