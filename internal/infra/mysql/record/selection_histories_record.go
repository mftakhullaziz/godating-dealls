package record

import "time"

type SelectionHistoryRecord struct {
	SelectionHistoryId  int64     `db:"selection_history_id"`
	AccountIdIdentifier int64     `db:"account_id_identifier"`
	AccountID           int64     `db:"account_id"`
	SelectionDate       time.Time `db:"selection_date"`
}

func (SelectionHistoryRecord) TableName() string {
	return "selection_histories"
}
