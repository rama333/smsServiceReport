package rabbitMQ

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"smsServiceReport/internal/config"
	"strconv"
	"strings"
	"time"
)

var (
	layout       = "2006-01-02 15:04:05"
	WaitTimeList = config.Config.RABBITMQWAITTIME
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
	Date                   string `json:"send_date"`
	Message_id             string `json:"message_id"`
	User_message_reference int64  `json:"user_message_reference"`
	Sub                    string `json:"sub"`
	Dlvrd                  string `json:"dlvrd"`
	Submit_date            string `json:"submit_date"`
	Done_date              string `json:"done_date"`
	Stat                   string `json:"stat"`
	Err                    int    `json:"err"`
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

	msgs_receive_queue, err := ch.Consume(
		"receive_queue", // queue
		"",              // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)

	//forev := make(chan bool)

	sendList := []Sent{}

	go func() {
		timeNow := time.Now()
		for d := range msgs_send_queue {

			timeNowTemp := time.Now()
			//log.Println("timeNowTemp", timeNowTemp)
			//log.Println(string(d.Body))

			tes := string(d.Body)

			var send Sent
			err := json.Unmarshal([]byte(strings.ReplaceAll(tes, "\n", "")), &send)
			if err != nil {

				log.Fatal("err send", string(d.Body), err)
			}

			sendList = append(sendList, send)

			if len(sendList) >= 1000 || (timeNowTemp.Sub(timeNow)).Seconds() >= WaitTimeList && len(sendList) > 0 {
				fmt.Println("-------------")
				fmt.Println("len list", len(sendList))
				fmt.Println("timeNow 1", timeNow)
				fmt.Println("time now temp 2", timeNowTemp)

				fmt.Println("duration", (timeNowTemp.Sub(timeNow)).Seconds())

				fmt.Println("-------------")

				tx, _ := config.Config.DB.Begin()
				stmp, _ := tx.Prepare("INSERT INTO Send (date,id,sms_id,sms_text, source_addr,dest_addr,sequence ) VALUES ($1, $2, $3, $4, $5, $6, $7)")

				for _, rec := range sendList {
					dest_addr, err := strconv.ParseInt(rec.Dest_addr, 10, 64)

					if err != nil {
						log.Println(dest_addr)
					}
					_, err = stmp.Exec(rec.Date, rec.Id, rec.Sms_id, rec.Sms_text, rec.Source_addr, dest_addr, rec.Sequence)
					if err != nil {
						log.Fatal("err", err)
					}
				}

				if err := tx.Commit(); err != nil {
					log.Fatal("err", err)
				}

				sendList = sendList[:0]

				timeNow = time.Now()
			}

		}
	}()

	sentMesIdList := []SentMesId{}

	go func() {
		timeNow := time.Now()
		for d := range msgs_link_queue {

			timeNowTemp := time.Now()
			//log.Println(string(d.Body))

			var sentMesId SentMesId
			err := json.Unmarshal(d.Body, &sentMesId)
			if err != nil {
				log.Fatal("err", err)
			}

			sentMesIdList = append(sentMesIdList, sentMesId)

			if len(sentMesIdList) >= 1000 || (timeNowTemp.Sub(timeNow)).Seconds() >= WaitTimeList && len(sentMesIdList) > 0 {
				fmt.Println("-------------")
				fmt.Println("len list", len(sentMesIdList))
				fmt.Println("timeNow 1", timeNow)
				fmt.Println("time now temp 2", timeNowTemp)

				fmt.Println("duration", (timeNowTemp.Sub(timeNow)).Seconds())

				fmt.Println("-------------")

				tx, _ := config.Config.DB.Begin()
				stmp, _ := tx.Prepare("INSERT INTO SentMesId (message_id, sequence ) VALUES ($1, $2)")

				for _, rec := range sentMesIdList {
					_, err := stmp.Exec(rec.Message_id, rec.Sequence)
					if err != nil {
						log.Fatal("err", err)
					}
				}

				if err := tx.Commit(); err != nil {
					log.Fatal("err", err)
				}

				sentMesIdList = sentMesIdList[0:]
				timeNow = time.Now()
			}
		}
	}()

	receiveList := []Receive{}

	go func() {
		timeNow := time.Now()
		for d := range msgs_receive_queue {

			timeNowTemp := time.Now()
			//k := d.Body
			var receive Receive
			err := json.Unmarshal(d.Body, &receive)
			if err != nil {
				log.Fatal("err receive", err)
			}

			receiveList = append(receiveList, receive)

			if len(receiveList) >= 1000 || (timeNowTemp.Sub(timeNow)).Seconds() >= WaitTimeList && len(receiveList) > 0 {
				fmt.Println("-------------")
				fmt.Println("len list", len(receiveList))
				fmt.Println("timeNow 1", timeNow)
				fmt.Println("time now temp 2", timeNowTemp)

				fmt.Println("duration", (timeNowTemp.Sub(timeNow)).Seconds())

				log.Println()

				fmt.Println("-------------")

				tx, _ := config.Config.DB.Begin()
				stmp, _ := tx.Prepare("INSERT INTO Receive (date, message_id, user_message_reference, sub, dlvrd, submit_date, done_date, stat, err, text,source_addr,destination_addr) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)")

				for _, rec := range receiveList {

					date, err := time.Parse(layout, rec.Date)

					if err != nil {
						log.Fatal("err", err)
					}

					submit_date, err := time.Parse(layout, rec.Date)

					if err != nil {
						log.Fatal("err", err)
					}

					done_date, err := time.Parse(layout, rec.Date)

					if err != nil {
						log.Fatal("err", err)
					}

					dest_addr, err := strconv.ParseInt(rec.Destination_addr, 10, 64)

					if err != nil {
						log.Fatal("err", err)
					}

					errReceive := strconv.Itoa(rec.Err)

					_, err = stmp.Exec(date, rec.Message_id, rec.User_message_reference, rec.Sub, rec.Dlvrd, submit_date, done_date, rec.Stat, errReceive, rec.Text, rec.Source_addr, dest_addr)
					if err != nil {
						log.Fatal("err stp", err)
					}

				}

				if err := tx.Commit(); err != nil {
					log.Fatal("err", err)
				}

				timeNow = time.Now()

				receiveList = receiveList[0:]
			}

			//log.Println("receive",receive)
		}
	}()

	//<-forev

	return ch, nil
}
