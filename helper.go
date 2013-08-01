package lemur

func Ping() (err error) {
	return DefaultPusher.Ping()
}

func (this *Pusher) Ping() (err error) {
	m := map[string]interface{}{}
	return this.SendData("helper/ping", m)
}

func AccountDetails(exclude []string) (err error) {
	return DefaultPusher.AccountDetails(exclude)
}

func (this *Pusher) AccountDetails(exclude []string) (err error) {
	m := map[string]interface{}{}
	m["exclude"] = exclude
	return this.SendData("helper/account-details", m)
}
