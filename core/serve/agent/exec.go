package agent

import (
	"log"

	"github.com/google/uuid"
)

type ExecPayload struct {
	Content          string
	Username         string
	CommandType      string
	WorkingDirectory string
	Timeout          uint
}

func (pod *SendPod) Exec(data *ExecPayload) (string, error) {

	v := &SocketData{
		TaskId:  uuid.NewString(),
		Method:  "Exec",
		Payload: data,
	}

	return v.TaskId, pod.Write(v)

}

func (pod *RespPod) Exec(rq *SocketData) {

	log.Println("Ping:resp:", rq.Payload)

}
