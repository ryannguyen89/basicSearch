package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"basicSearch"
)

func main() {
	// Load data

	var (
		orgData    []*basicSearch.Organization
		userData   []*basicSearch.User
		ticketData []*basicSearch.Ticket
	)

	err := basicSearch.Load("data/organizations.json", &orgData)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = basicSearch.Load("data/users.json", &userData)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = basicSearch.Load("data/tickets.json", &ticketData)
	if err != nil {
		fmt.Println(err)
		return
	}

	orgIdx := basicSearch.IndexOrg(orgData)
	userIdx := basicSearch.IndexUser(userData)
	ticketIdx := basicSearch.IndexTicket(ticketData)

	var (
		indexing = basicSearch.Indexing{
			OIndex: orgIdx,
			UIndex: userIdx,
			TIndex: ticketIdx,
		}
	)

	basicSearch.UpdateRelatedEntities(indexing)

	for {
		fmt.Print("<Command> [help|quit] ➜ ")
		reader := bufio.NewReader(os.Stdin)
		cmd, err := reader.ReadString('\n')
		if err != nil {
			continue
		}
		cmd = strings.TrimSuffix(cmd, "\n")

		switch cmd {
		case "quit":
			fmt.Println("bye")
			return
		case "help":
			fmt.Println("describe")
			fmt.Println("\t\tdescribe-organizations\t\twill return search fields for organizations")
			fmt.Println("\t\tdescribe-users\t\t\twill return search fields for users")
			fmt.Println("\t\tdescribe-tickets\t\twill return search fields for tickets")
			fmt.Println("search")
			fmt.Println("\t\tsearch-organizations:tags=West\twill return any organizations who has West in their Tags in JSON")
			fmt.Println("\t\tsearch-users:alias=Miss Coffey\twill return any users whose alias is Miss Coffey in JSON")
			fmt.Println("\t\tsearch-tickets:status=pending\twill return any tickets with pending status in JSON")
		case "describe-organizations":
			for _, field := range basicSearch.SupportedOrgFields {
				fmt.Println(field)
			}
		case "describe-users":
			for _, field := range basicSearch.SupportedUserFields {
				fmt.Println(field)
			}
		case "describe-tickets":
			for _, field := range basicSearch.SupportedTicketFields {
				fmt.Println(field)
			}
		default:
			if strings.HasPrefix(cmd, "search-") {
				type funcSearch func(key, val string, indexing basicSearch.Indexing) *basicSearch.Result

				var (
					subCmd  string
					fSearch funcSearch
				)
				if strings.HasPrefix(cmd, "search-organizations:") {
					subCmd = strings.Replace(cmd, "search-organizations:", "", 1)
					fSearch = basicSearch.SearchOrg
				} else if strings.HasPrefix(cmd, "search-users:") {
					subCmd = strings.Replace(cmd, "search-users:", "", 1)
					fSearch = basicSearch.SearchUser
				} else if strings.HasPrefix(cmd, "search-tickets:") {
					subCmd = strings.Replace(cmd, "search-tickets:", "", 1)
					fSearch = basicSearch.SearchTicket
				} else {
					continue
				}

				kv := strings.Split(subCmd, "=")
				if len(kv) != 2 {
					continue
				}
				result := fSearch(kv[0], kv[1], indexing)
				if result == nil {
					continue
				}
				resultByte, err := json.MarshalIndent(result, "", "  ")
				if err != nil {
					fmt.Println("Prepare data:", err)
					continue
				}
				fmt.Println(string(resultByte))
				continue
			}

			fmt.Println("")
		}
	}
}
