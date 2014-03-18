package lemur

func Ping(listId string) (err error) {
	return DefaultPusher.Ping(listId)
}

func (this *Pusher) Ping(listId string) (err error) {
	m := map[string]interface{}{}
	return this.SendData(listId, "helper/ping", m)
}

func AccountDetails(listId string, exclude []string) (err error) {
	return DefaultPusher.AccountDetails(listId, exclude)
}

func (this *Pusher) AccountDetails(listId string, exclude []string) (err error) {
	m := map[string]interface{}{}
	m["exclude"] = exclude
	return this.SendData(listId, "helper/account-details", m)
}
