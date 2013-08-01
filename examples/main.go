package main

import (
	"fmt"
	"github.com/athom/lemur"
	"os"
)

func main() {
	apiKey := os.Getenv("MAILCHIMP_API_KEY")
	listId := os.Getenv("MAILCHIMP_LIST_ID")

	lemur.SetApiKey(apiKey)
	//err := lemur.Ping()
	//err := lemur.AccountDetails(nil)
	lemur.SelectList(listId)
	//err := lemur.MemberInfo( []string{"a@theplant.jp"})
	//err := lemur.Members( "", nil)
	//err := lemur.MergeVarAdd( "ORGNAME", "OrganizationName", nil)
	//err := lemur.MergeVarDel( "ORGNAME")
	//err := lemur.ListUpdateMember( lemur.NewMemberInput("a@theplant.jp", "", ""), []string{"First Name"}, "", true)
	//err := lemur.ListUpdateMember(
	//"athom@126.com",
	//map[string]string{
	//"FNAME": "xxm",
	//"LNAME": "y",
	//},
	//"text",
	//true,
	//)
	err := lemur.ListAddMember(0, "athom@126.com")
	if err != nil {
		fmt.Println(err)
	}
	return
}
