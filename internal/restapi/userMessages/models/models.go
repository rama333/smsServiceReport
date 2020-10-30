package models

import "time"

type Messages struct {
	Date        time.Time `json:"date" db:"Send.date"`
	Submit_date time.Time `json:"submit_date" db:"submit_date"`
	Done_date   time.Time `json:"done_date" db:"done_date"`
	Dest_addr   string    `json:"dest_addr" db:"Receive.destination_addr"`
	Id          int       `json:"id" db:"id"`
	Sms_text    string    `json:"sms_text,omitempty" db:"sms_text"`
	Source_addr string    `json:"source_addr" db:"Send.source_addr"`
	Message_id  string    `json:"message_id" db:"SentMesId.message_id"`
	Stat        string    `json:"stat" db:"stat"`
}

type DurationDate struct {
	Dest_adr      string `json:"dest_addr"`
	StartDuration string `json:"since_date"`
	EndDuration   string `json:"on_date"`
}
