package models

import "time"

type Messages struct {
	Date        time.Time `json:"date"`
	Submit_date time.Time `json:"submit_date"`
	Done_date   time.Time `json:"done_date"`
	Dest_addr   string    `json:"dest_addr"`
	Id          int       `json:"id"`
	Sms_text    string    `json:"sms_text,omitempty"`
	Source_addr string    `json:"source_addr"`
	Message_id  string    `json:"message_id"`
	Stat        string    `json:"stat"`
}

type DurationDate struct {
	StartDuration string `json:"since_date"`
	EndDuration   string `json:"on_date"`
}
