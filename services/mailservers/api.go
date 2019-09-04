package mailservers

import "context"

func NewAPI(db *Database) *API {
	return &API{db}
}

// API is class with methods available over RPC.
type API struct {
	db *Database
}

func (a *API) AddMailserver(ctx context.Context, m Mailserver) error {
	return a.db.Add(m)
}

func (a *API) GetMailservers(ctx context.Context) ([]Mailserver, error) {
	return a.db.Mailservers()
}

func (a *API) DeleteMailserver(ctx context.Context, id string) error {
	return a.db.Delete(id)
}

func (a *API) AddMailserverRequestGaps(ctx context.Context, gaps []MailserverRequestGap) error {
	return a.db.AddGaps(gaps)
}

func (a *API) GetMailserverRequestGaps(ctx context.Context, chatID string) ([]MailserverRequestGap, error) {
	return a.db.RequestGaps(chatID)
}

func (a *API) DeleteMailserverRequestGaps(ctx context.Context, ids []string) error {
	return a.db.DeleteGaps(ids)
}

func (a *API) DeleteMailserverRequestGapsByChatID(ctx context.Context, chatID string) error {
	return a.db.DeleteGapsByChatID(chatID)
}

func (a *API) AddMailserverTopic(ctx context.Context, topic MailserverTopic) error {
	return a.db.AddTopic(topic)
}

func (a *API) GetMailserverTopics(ctx context.Context) ([]MailserverTopic, error) {
	return a.db.Topics()
}

func (a *API) DeleteMailserverTopic(ctx context.Context, topic string) error {
	return a.db.DeleteTopic(topic)
}
