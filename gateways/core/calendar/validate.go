/*
Copyright 2018 BlackRock, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package calendar

import (
	"context"
	"fmt"

	"github.com/argoproj/argo-events/common"
	"github.com/argoproj/argo-events/gateways"
	gwcommon "github.com/argoproj/argo-events/gateways/common"
	"github.com/argoproj/argo-events/pkg/apis/eventsources/v1alpha1"
	"github.com/ghodss/yaml"
)

// ValidateEventSource validates calendar event source
func (listener *EventSourceListener) ValidateEventSource(ctx context.Context, eventSource *gateways.EventSource) (*gateways.ValidEventSource, error) {
	var calendarEventSource *v1alpha1.CalendarEventSource
	if err := yaml.Unmarshal(eventSource.Value, &calendarEventSource); err != nil {
		listener.Logger.WithError(err).WithField(common.LabelEventSource, eventSource.Name).Errorln("failed to parse the event source")
		return &gateways.ValidEventSource{
			IsValid: false,
			Reason:  err.Error(),
		}, err
	}

	if err := validate(calendarEventSource); err != nil {
		listener.Logger.WithError(err).WithField(common.LabelEventSource, eventSource.Name).Errorln("failed to validate the event source")
		return &gateways.ValidEventSource{
			IsValid: false,
			Reason:  err.Error(),
		}, err
	}

	return &gateways.ValidEventSource{
		IsValid: true,
	}, nil
}

func validate(calendarEventSource *v1alpha1.CalendarEventSource) error {
	if calendarEventSource == nil {
		return gwcommon.ErrNilEventSource
	}
	if calendarEventSource.Schedule == "" && calendarEventSource.Interval == "" {
		return fmt.Errorf("must have either schedule or interval")
	}
	if _, err := resolveSchedule(calendarEventSource); err != nil {
		return err
	}
	return nil
}
