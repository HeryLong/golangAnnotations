// Generated automatically: do not edit manually

package example

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/satori/go.uuid"
)

const GamblerCreatedEventName = "GamblerCreated"

func (s *GamblerCreated) Wrap(uid string) (*Envelope, error) {
	envelope := new(Envelope)
	envelope.Uuid = uuid.NewV1().String()
	envelope.SequenceNumber = 0 // Set later by event-store
	envelope.Timestamp = time.Now()
	envelope.AggregateName = GamblerAggregateName
	envelope.AggregateUid = uid
	envelope.EventTypeName = GamblerCreatedEventName
	blob, err := json.Marshal(s)
	if err != nil {
		log.Printf("Error marshalling GamblerCreated payload %+v", err)
		return nil, err
	}
	envelope.EventData = string(blob)

	return envelope, nil
}

func IsGamblerCreated(envelope *Envelope) bool {
	return envelope.EventTypeName == "GamblerCreated"
}

func GetIfIsGamblerCreated(envelop *Envelope) (*GamblerCreated, bool) {
	if IsGamblerCreated(envelop) == false {
		return nil, false
	}
	event, err := UnWrapGamblerCreated(envelop)
	if err != nil {
		return nil, false
	}
	return event, true
}

func UnWrapGamblerCreated(envelop *Envelope) (*GamblerCreated, error) {
	if IsGamblerCreated(envelop) == false {
		return nil, fmt.Errorf("Not a GamblerCreated")
	}
	var event GamblerCreated
	err := json.Unmarshal([]byte(envelop.EventData), &event)
	if err != nil {
		log.Printf("Error unmarshalling GamblerCreated payload %+v", err)
		return nil, err
	}

	return &event, nil
}
