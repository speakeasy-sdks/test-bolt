// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/test-bolt/pkg/utils"
	"time"
)

type EventType string

const (
	EventTypeGroup EventType = "group"
	EventTypeList  EventType = "list"
)

type Event struct {
	EventGroup *EventGroup
	EventList  *EventList

	Type EventType
}

func CreateEventGroup(group EventGroup) Event {
	typ := EventTypeGroup
	typStr := EventGroupTag(typ)
	group.DotTag = typStr

	return Event{
		EventGroup: &group,
		Type:       typ,
	}
}

func CreateEventList(list EventList) Event {
	typ := EventTypeList
	typStr := EventListTag(typ)
	list.DotTag = typStr

	return Event{
		EventList: &list,
		Type:      typ,
	}
}

func (u *Event) UnmarshalJSON(data []byte) error {

	type discriminator struct {
		DotTag string
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.DotTag {
	case "group":
		eventGroup := new(EventGroup)
		if err := utils.UnmarshalJSON(data, &eventGroup, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.EventGroup = eventGroup
		u.Type = EventTypeGroup
		return nil
	case "list":
		eventList := new(EventList)
		if err := utils.UnmarshalJSON(data, &eventList, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.EventList = eventList
		u.Type = EventTypeList
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Event) MarshalJSON() ([]byte, error) {
	if u.EventGroup != nil {
		return utils.MarshalJSON(u.EventGroup, "", true)
	}

	if u.EventList != nil {
		return utils.MarshalJSON(u.EventList, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type Webhook struct {
	// The time at which the webhook was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Event     Event      `json:"event"`
	// The webhook's ID
	ID *string `json:"id,omitempty"`
	// The webhook's URL
	URL string `json:"url"`
}

func (w Webhook) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *Webhook) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Webhook) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Webhook) GetEvent() Event {
	if o == nil {
		return Event{}
	}
	return o.Event
}

func (o *Webhook) GetEventGroup() *EventGroup {
	return o.GetEvent().EventGroup
}

func (o *Webhook) GetEventList() *EventList {
	return o.GetEvent().EventList
}

func (o *Webhook) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Webhook) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

type WebhookInput struct {
	Event Event `json:"event"`
	// The webhook's URL
	URL string `json:"url"`
}

func (o *WebhookInput) GetEvent() Event {
	if o == nil {
		return Event{}
	}
	return o.Event
}

func (o *WebhookInput) GetEventGroup() *EventGroup {
	return o.GetEvent().EventGroup
}

func (o *WebhookInput) GetEventList() *EventList {
	return o.GetEvent().EventList
}

func (o *WebhookInput) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}
