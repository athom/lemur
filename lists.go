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

func ListUpdateMember(email string, mergeVars map[string]string) (err error) {
	return DefaultPusher.ListUpdateMember(email, mergeVars)
}

func (this *Pusher) ListUpdateMember(email string, mergeVars map[string]string) (err error) {
	m := map[string]interface{}{}
	m["email"] = map[string]string{"email": email}
	m["merge_vars"] = mergeVars
	m["email_type"] = ""
	m["replace_interests"] = true
	return this.SendData("lists/update-member", m)
}

func ListSubscribe(email string, mergeVars map[string]string) (err error) {
	return DefaultPusher.ListSubscribe(email, mergeVars)
}

func (this *Pusher) ListSubscribe(email string, mergeVars map[string]string) (err error) {
	return this.listSubscribe(
		email,
		mergeVars,
		"text",
		false,
		false,
		true,
		false,
	)
}

func (this *Pusher) listSubscribe(email string, mergeVars map[string]string, emailType string, doubleOpt bool, updateExisting bool, replaceInterests bool, sendWelcome bool) (err error) {
	m := map[string]interface{}{}
	m["email"] = map[string]string{"email": email}
	m["merge_vars"] = mergeVars
	m["email_type"] = emailType
	m["double_optin"] = doubleOpt
	m["update_existing"] = updateExisting
	m["replace_interests"] = replaceInterests
	m["send_welcome"] = sendWelcome

	return this.SendData("lists/subscribe", m)
}

func ListUnsubscribe(email string) (err error) {
	return DefaultPusher.ListUnsubscribe(email)
}

func (this *Pusher) ListUnsubscribe(email string) (err error) {
	return this.listUnsubscribe(email, true, false, false)
}

func (this *Pusher) listUnsubscribe(email string, deleteMember bool, sendGoodbye bool, sendNotify bool) (err error) {
	m := map[string]interface{}{}
	m["email"] = map[string]string{"email": email}
	m["delete_member"] = deleteMember
	m["send_goodbye"] = sendGoodbye
	m["send_notify"] = sendNotify

	return this.SendData("lists/unsubscribe", m)
}
