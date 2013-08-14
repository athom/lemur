package main

import (
	"fmt"
	"github.com/athom/lemur"
	"os"
)

func main() {
	apiKey := os.Getenv("MAILCHIMP_API_KEY")
	listId := os.Getenv("MAILCHIMP_LIST_ID")

	apiKey = "4bad4d3a5985f2e1a5a1c6660f0a86bf-us4"
	listId = "0a5644dbaf"

	lemur.SetApiKey(apiKey)
	//err := lemur.Ping()
	//err := lemur.AccountDetails(nil)
	lemur.SelectList(listId)
	//err := lemur.MemberInfo( []string{"a@theplant.jp"})
	//err := lemur.Members( "", nil)
	//err := lemur.MergeVarAdd( "ORGNAME", "OrganizationName", nil)
	//err := lemur.MergeVarDel( "ORGNAME")

	//err := lemur.ListUpdateMember(
	//"frankyue1019@gmail.com",
	//map[string]string{
	//"FNAME": "Famm ",
	//"LNAME": "Ge",
	//},
	//)

	err := lemur.ListSubscribe(
		"yeerkunth+1313@gmail.com",
		map[string]string{
			"FNAME": "Yeer",
			"LNAME": "Kunnnnn",
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	err = lemur.ListUpdateMember(
		"yeerkunth+1313@gmail.com",
		map[string]string{
			"new-email": "yeerkunth+1344@gmail.com",
		},
	)

	//err := lemur.ListUnsubscribe(
	//"yeerkunth+1309@gmail.com",
	//)

	return
}
