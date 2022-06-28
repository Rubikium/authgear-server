package model

import (
	"fmt"
	"time"
)

type SubscriptionCheckoutStatus string

const (
	SubscriptionCheckoutStatusOpen       SubscriptionCheckoutStatus = "open"
	SubscriptionCheckoutStatusCompleted  SubscriptionCheckoutStatus = "completed"
	SubscriptionCheckoutStatusSubscribed SubscriptionCheckoutStatus = "subscribed"
)

type Subscription struct {
	ID                   string
	AppID                string
	StripeSubscriptionID string
	StripeCustomerID     string
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

type SubscriptionCheckout struct {
	ID                      string
	AppID                   string
	StripeCheckoutSessionID string
	StripeCustomerID        *string
	Status                  SubscriptionCheckoutStatus
	CreatedAt               time.Time
	UpdatedAt               time.Time
	ExpireAt                time.Time
}

type PriceType string

const (
	PriceTypeFixed PriceType = "fixed"
	PriceTypeUsage PriceType = "usage"
)

func (t PriceType) Valid() error {
	switch t {
	case PriceTypeFixed:
		return nil
	case PriceTypeUsage:
		return nil
	}
	return fmt.Errorf("unknown price_type: %#v", t)
}

type UsageType string

const (
	UsageTypeNone UsageType = ""
	UsageTypeSMS  UsageType = "sms"
	UsageTypeMAU  UsageType = "mau"
)

func (t UsageType) Valid() error {
	switch t {
	case UsageTypeNone:
		return nil
	case UsageTypeSMS:
		return nil
	case UsageTypeMAU:
		return nil
	}
	return fmt.Errorf("unknown usage_type: %#v", t)
}

type SMSRegion string

const (
	SMSRegionNone         SMSRegion = ""
	SMSRegionNorthAmerica SMSRegion = "north-america"
	SMSRegionOtherRegions SMSRegion = "other-regions"
)

func (r SMSRegion) Valid() error {
	switch r {
	case SMSRegionNone:
		return nil
	case SMSRegionNorthAmerica:
		return nil
	case SMSRegionOtherRegions:
		return nil
	}
	return fmt.Errorf("unknown sms_region: %#v", r)
}

type SubscriptionUsage struct {
	NextBillingDate time.Time                `json:"nextBillingDate"`
	Items           []*SubscriptionUsageItem `json:"items,omitempty"`
}

type SubscriptionUsageItem struct {
	Type        PriceType `json:"type"`
	UsageType   UsageType `json:"usageType"`
	SMSRegion   SMSRegion `json:"smsRegion"`
	Quantity    int       `json:"quantity"`
	Currency    *string   `json:"currency"`
	UnitAmount  *int      `json:"unitAmount"`
	TotalAmount *int      `json:"totalAmount"`
}