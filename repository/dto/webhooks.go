package dto

import (
	"time"

	"github.com/bitcoin-sv/pulse/notification"
)

// DbWebhook represent webhook saved in db.
type DbWebhook struct {
	Url               string    `db:"url"`
	TokenHeader       string    `db:"token_header"`
	Token             string    `db:"token"`
	CreatedAt         time.Time `db:"created_at"`
	LastEmitStatus    string    `db:"last_emit_status"`
	LastEmitTimestamp time.Time `db:"last_emit_timestamp"`
	ErrorsCount       int       `db:"errors_count"`
	Active            bool      `db:"is_active"`
}

// ToWebhook converts DbWebhook to Webhook.
func (dbt *DbWebhook) ToWebhook() *notification.Webhook {
	return &notification.Webhook{
		Url:         dbt.Url,
		TokenHeader: dbt.TokenHeader,
		Token:       dbt.Token,
		CreatedAt:   dbt.CreatedAt,
		ErrorsCount: dbt.ErrorsCount,
		Active:      dbt.Active,
	}
}

// ToDbWebhook converts Webhook to DbWebhook.
func ToDbWebhook(t *notification.Webhook) *DbWebhook {
	return &DbWebhook{
		Url:         t.Url,
		TokenHeader: t.TokenHeader,
		Token:       t.Token,
		CreatedAt:   t.CreatedAt,
		ErrorsCount: t.ErrorsCount,
		Active:      t.Active,
	}
}
