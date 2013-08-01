package lemur

func SelectList(id string) {
	DefaultPusher.SelectList(id)
}

func (this *Pusher) SelectList(id string) {
	this.listId = id
}

func MemberInfo(emails []string) (err error) {
	return DefaultPusher.MemberInfo(emails)
}

func (this *Pusher) MemberInfo(emails []string) (err error) {
	m := map[string]interface{}{}
	m["emails"] = emails
	return this.SendData("lists/member-info", m)
}

func Members(status string, opts map[string]interface{}) (err error) {
	return DefaultPusher.Members(status, opts)
}

func (this *Pusher) Members(status string, opts map[string]interface{}) (err error) {
	m := map[string]interface{}{}
	m["status"] = status
	m["opts"] = opts
	return this.SendData("lists/members", m)
}

func MergeVarAdd(tag string, name string, opts map[string]interface{}) (err error) {
	return DefaultPusher.MergeVarAdd(tag, name, opts)
}

func (this *Pusher) MergeVarAdd(tag string, name string, opts map[string]interface{}) (err error) {
	m := map[string]interface{}{}
	m["tag"] = tag
	m["name"] = name
	m["opts"] = opts
	return this.SendData("lists/merge-var-add", m)
}

func MergeVarDel(tag string) (err error) {
	return DefaultPusher.MergeVarDel(tag)
}

func (this *Pusher) MergeVarDel(tag string) (err error) {
	m := map[string]interface{}{}
	m["tag"] = tag
	return this.SendData("lists/merge-var-del", m)
}

func ListUpdateMember(email map[string]string, mergeVars map[string]string, emailType string, replace bool) (err error) {
	return DefaultPusher.ListUpdateMember(email, mergeVars, emailType, replace)
}

func (this *Pusher) ListUpdateMember(email map[string]string, mergeVars map[string]string, emailType string, replace bool) (err error) {
	m := map[string]interface{}{}
	m["email"] = email
	m["merge_vars"] = mergeVars
	m["email_type"] = emailType
	m["replace_interests"] = replace
	return this.SendData("lists/update-member", m)
}
