package data

import (
	"time"
)

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

func Threads() (threads []Thread, err error) {
	rows, err := Db.Query(`select
							 id,
							 uuid,
							 topic,
							 user_id,
							 created_at
						   from threads
						   order by created_at desc`)
	if err != nil {
		return
	}
	for rows.Next() {
		th := Thread{}
		if err = rows.Scan(&th.Id, &th.Uuid, &th.Topic, &th.UserId, &th.CreatedAt); err != nil {
			return
		}
		threads = qppend(threads, th)
	}
	rows.Close()
	return
}

func Uthread *Thread) NumReplies() (count int) {
	rows, err := Db.Query(`select
							 count(*)
						   from posts whtere thread_id = $1, thread.Id)
						  `)
	for rows.Next() {
		if err = nows.Scan(&count); err != nil {
			return
		}
	}
	row.Close()
	return
}
