package lemur

func SelectList(id string) {
	DefaultPusher.SelectList(id)
}

func (this *Pusher) SelectList(id string) {
	this.listId = id
}

func MemberInfo(listId string, emails []string) (err error) {
	return DefaultPusher.MemberInfo(listId, emails)
}

func (this *Pusher) MemberInfo(listId string, emails []string) (err error) {
	m := map[string]interface{}{}
	m["emails"] = emails
	return this.SendData(listId, "lists/member-info", m)
}

func Members(listId string, status string, opts map[string]interface{}) (err error) {
	return DefaultPusher.Members(listId, status, opts)
}

func (this *Pusher) Members(listId string, status string, opts map[string]interface{}) (err error) {
	m := map[string]interface{}{}
	m["status"] = status
	m["opts"] = opts
	return this.SendData(listId, "lists/members", m)
}

func MergeVarAdd(listId string, tag string, name string, opts map[string]interface{}) (err error) {
	return DefaultPusher.MergeVarAdd(listId, tag, name, opts)
}

func (this *Pusher) MergeVarAdd(listId string, tag string, name string, opts map[string]interface{}) (err error) {
	m := map[string]interface{}{}
	m["tag"] = tag
	m["name"] = name
	m["opts"] = opts
	return this.SendData(listId, "lists/merge-var-add", m)
}

func MergeVarDel(listId string, tag string) (err error) {
	return DefaultPusher.MergeVarDel(listId, tag)
}

func (this *Pusher) MergeVarDel(listId string, tag string) (err error) {
	m := map[string]interface{}{}
	m["tag"] = tag
	return this.SendData(listId, "lists/merge-var-del", m)
}

func ListUpdateMember(listId string, email string, mergeVars map[string]string) (err error) {
	return DefaultPusher.ListUpdateMember(listId, email, mergeVars)
}

func (this *Pusher) ListUpdateMember(listId string, email string, mergeVars map[string]string) (err error) {
	m := map[string]interface{}{}
	m["email"] = map[string]string{"email": email}
	m["merge_vars"] = mergeVars
	m["email_type"] = ""
	m["replace_interests"] = true
	return this.SendData(listId, "lists/update-member", m)
}

func ListSubscribe(listId string, email string, mergeVars map[string]string) (err error) {
	return DefaultPusher.ListSubscribe(listId, email, mergeVars)
}

func (this *Pusher) ListSubscribe(listId string, email string, mergeVars map[string]string) (err error) {
	return this.listSubscribe(
		listId,
		email,
		mergeVars,
		"text",
		false,
		false,
		true,
		false,
	)
}

func (this *Pusher) listSubscribe(listId string, email string, mergeVars map[string]string, emailType string, doubleOpt bool, updateExisting bool, replaceInterests bool, sendWelcome bool) (err error) {
	m := map[string]interface{}{}
	m["email"] = map[string]string{"email": email}
	m["merge_vars"] = mergeVars
	m["email_type"] = emailType
	m["double_optin"] = doubleOpt
	m["update_existing"] = updateExisting
	m["replace_interests"] = replaceInterests
	m["send_welcome"] = sendWelcome

	return this.SendData(listId, "lists/subscribe", m)
}

func ListUnsubscribe(listId string, email string) (err error) {
	return DefaultPusher.ListUnsubscribe(listId, email)
}

func (this *Pusher) ListUnsubscribe(listId string, email string) (err error) {
	return this.listUnsubscribe(listId, email, true, false, false)
}

func (this *Pusher) listUnsubscribe(listId string, email string, deleteMember bool, sendGoodbye bool, sendNotify bool) (err error) {
	m := map[string]interface{}{}
	m["email"] = map[string]string{"email": email}
	m["delete_member"] = deleteMember
	m["send_goodbye"] = sendGoodbye
	m["send_notify"] = sendNotify

	return this.SendData(listId, "lists/unsubscribe", m)
}
