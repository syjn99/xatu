package v1

import (
	"context"
	"errors"

	"github.com/ethpandaops/xatu/pkg/proto/xatu"
	"github.com/sirupsen/logrus"
)

const (
	EventsContributionAndProofType = "BEACON_API_ETH_V1_EVENTS_CONTRIBUTION_AND_PROOF"
)

type EventsContributionAndProof struct {
	log       logrus.FieldLogger
	event     *xatu.DecoratedEvent
	networkID uint64
}

func NewEventsContributionAndProof(log logrus.FieldLogger, event *xatu.DecoratedEvent, networkID uint64) *EventsContributionAndProof {
	return &EventsContributionAndProof{
		log:       log.WithField("event", EventsContributionAndProofType),
		event:     event,
		networkID: networkID,
	}
}

func (b *EventsContributionAndProof) Type() string {
	return EventsContributionAndProofType
}

func (b *EventsContributionAndProof) Validate(ctx context.Context) error {
	_, ok := b.event.GetData().(*xatu.DecoratedEvent_EthV1EventsContributionAndProof)
	if !ok {
		return errors.New("failed to cast event data")
	}

	return nil
}

func (b *EventsContributionAndProof) Filter(ctx context.Context) bool {
	networkID := b.event.GetMeta().GetClient().GetEthereum().GetNetwork().GetId()

	return networkID != b.networkID
}
