package rabbitMQ

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"smsServiceReport/internal/config"
	"strconv"
	"time"
)

var (
	layout = "2006-01-02 15:04:05"
)

type Send struct {
	Date     string `json:"date"`
	Mes      string `json:"mes"`
	Sequence int64  `json:"sequence"`
}

type Sent struct {
	Date          string `json:"send_date,omitempty"`
	Id            int    `json:"id,omitempty"`
	Sms_id        int    `json:"sms_id,omitempty"`
	Sms_text      string `json:"sms_text,omitempty"`
	Source_addr   string `json:"source_addr,omitempty"`
	Dest_addr     string `json:"dest_addr,omitempty"`
	Delivery_time string `json:"delivery_time,omitempty"`
	Sequence      int64  `json:"sequence,omitempty"`
}

type Receive struct {
	Date                   string `json:"date"`
	Message_id             string `json:"message_id"`
	User_message_reference int64  `json:"user_message_reference"`
	Sub                    string `json:"sub"`
	Dlvrd                  string `json:"dlvrd"`
	Submit_date            string `json:"submit_date"`
	Done_date              string `json:"done_date"`
	Stat                   string `json:"stat"`
	Err                    int64  `json:"err"`
	Text                   string `json:"text"`
	Source_addr            string `json:"source_addr"`
	Destination_addr       string `json:"destination_addr"`
}

type SentMesId struct {
	Message_id string `json:"message_id"`
	Sequence   int64  `json:"sequence"`
}

//type Rabbit struct {
//	rabbitConnection *amqp.Channel
//	rabbitError error
//}

func StartRabbitMQ() (*amqp.Channel, error) {
	con, err := amqp.Dial("amqp://smsvm:smspassword@192.168.143.208/")
	if err != nil {
		return nil, err
	}

	ch, err := con.Channel()

	if err != nil {
		return nil, err
	}

	//defer ch.Close()

	msgs_send_queue, err := ch.Consume(
		"send_queue", // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)

	msgs_link_queue, err := ch.Consume(
		"link_queue", // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)

	//msgs_receive_queue, err := ch.Consume(
	//	"receive_queue", // queue
	//	"",     // consumer
	//	true,   // auto-ack
	//	false,  // exclusive
	//	false,  // no-local
	//	false,  // no-wait
	//	nil,    // args
	//)

	//forev := make(chan bool)

	go func() {
		for d := range msgs_send_queue {

			log.Println(string(d.Body))
			var send Sent
			err := json.Unmarshal(d.Body, &send)
			if err != nil {
				log.Println(err)
			}

			date, err := time.Parse(layout, send.Date)

			if err != nil {
				log.Println(err)
			}

			dest_addr, err := strconv.ParseInt(send.Dest_addr, 10, 64)

			if err != nil {
				log.Println(dest_addr)
			}

			tx := config.Config.DB.MustBegin()

			tx.MustExec("INSERT INTO Send (date,id,sms_id,sms_text, source_addr,dest_addr,sequence ) VALUES ($1, $2, $3, $4, $5, $6, $7)", date, send.Id, send.Sms_id, send.Sms_text, send.Source_addr, dest_addr, send.Sequence)

			err = tx.Commit()
			if err != nil {
				log.Println(err)
			}

			log.Println("seeeeen", send)
		}
	}()

	go func() {
		for d := range msgs_link_queue {
			log.Println(string(d.Body))

			var sentMesId SentMesId
			err := json.Unmarshal(d.Body, &sentMesId)
			if err != nil {
				log.Println(err)
			}

			tx := config.Config.DB.MustBegin()

			tx.MustExec("INSERT INTO SentMesId (message_id, sequence ) VALUES ($1, $2)", sentMesId.Message_id, sentMesId.Sequence)

			err = tx.Commit()
			if err != nil {
				log.Println(err)
			}

			log.Println(sentMesId)
		}
	}()

	//go func() {
	//	for d:= range msgs_receive_queue{
	//		log.Println(string(d.Body))
	//		var receive Receive
	//		err := json.Unmarshal(d.Body, &receive)
	//		if (err != nil){
	//			log.Println(err)
	//		}
	//
	//		tx := config.Config.DB.MustBegin()
	//
	//		tx.MustExec("INSERT INTO Receive (date, message_id, user_message_reference, sub, dlvrd, submit_date, done_date, stat, err, text,source_addr,destination_addr) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)", rec.Date,rec.Message_id, rec.User_message_reference,rec.Sub,rec.Dlvrd, rec.Submit_date, rec.Done_date, rec.Stat, rec.Err, rec.Text, rec.Source_addr, rec.Destination_addr)
	//
	//
	//		err = tx.Commit()
	//		if err != nil{
	//			log.Println(err)
	//		}
	//
	//		//log.Println("receive",receive)
	//	}
	//}()

	//<-forev

	return ch, nil
}
