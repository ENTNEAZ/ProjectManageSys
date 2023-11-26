package dataStruct

import "encoding/json"

type Worker struct {
	WorkerId       int
	WorkerName     string
	WorkerGender   string
	WorkerBirth    string
	WorkerJoinTime string
	WorkerJob      string
}

func (w Worker) ToJson() ([]byte, error) {
	return json.Marshal(w)
}

/*
create table worker
(
    worker_id        int auto_increment
        primary key,
    worker_name      varchar(255) not null,
    worker_gender    varchar(255) not null,
    worker_birth     date         not null,
    worker_join_time date         not null,
    worker_job       varchar(255) not null
);
*/
