package types

import (
	"time"
)

type Campaign struct {
	Info                           map[string]interface{} `json:"_info,omitempty"`
	ActualCost                     float64                `json:"actualCost,omitempty"`
	ActualGrossMargin              float64                `json:"actualGrossMargin,omitempty"`
	ActualROI                      float64                `json:"actualROI,omitempty"`
	ActualRevenue                  float64                `json:"actualRevenue,omitempty"`
	BudgetCost                     float64                `json:"budgetCost,omitempty"`
	BudgetGrossMargin              float64                `json:"budgetGrossMargin,omitempty"`
	BudgetROI                      float64                `json:"budgetROI,omitempty"`
	BudgetRevenue                  float64                `json:"budgetRevenue,omitempty"`
	DefaultGroup                   Group                  `json:"defaultGroup,omitempty"`
	EmailsSent                     int                    `json:"emailsSent,omitempty"`
	EndDate                        time.Time              `json:"endDate,omitempty"`
	ID                             int                    `json:"id,omitempty"`
	Impressions                    int                    `json:"impressions,omitempty"`
	Inactive                       bool                   `json:"inactive,omitempty"`
	InactiveDaysAfterEnd           int                    `json:"inactiveDaysAfterEnd,omitempty"`
	LocationId                     int                    `json:"locationId,omitempty"`
	MarketingManagerDefaultTrackId int                    `json:"marketingManagerDefaultTrackId,omitempty"`
	Member                         Member                 `json:"member,omitempty"`
	Name                           string                 `json:"name"`
	Notes                          string                 `json:"notes,omitempty"`
	OpportunityDefaultTrackId      int                    `json:"opportunityDefaultTrackId,omitempty"`
	StartDate                      time.Time              `json:"startDate"`
	Status                         CampaignStatus         `json:"status,omitempty"`
	SubType                        CampaignSubType        `json:"subType"`
	Type                           CampaignType           `json:"type"`
}

type CampaignSubTypeCampaignSubType struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name"`
	Type CampaignType           `json:"type"`
}

type CampaignAudit struct {
	CampaignId       int    `json:"campaignId,omitempty"`
	CreatedBy        string `json:"createdBy,omitempty"`
	DateCreated      string `json:"dateCreated,omitempty"`
	DocumentsCreated int    `json:"documentsCreated,omitempty"`
	EmailSubject     string `json:"emailSubject,omitempty"`
	EmailsSent       int    `json:"emailsSent,omitempty"`
	EmailsUnsent     int    `json:"emailsUnsent,omitempty"`
	Group            Group  `json:"group,omitempty"`
	ID               int    `json:"id,omitempty"`
}

type CampaignStatus struct {
	Info         map[string]interface{} `json:"_info,omitempty"`
	DefaultFlag  bool                   `json:"defaultFlag,omitempty"`
	ID           int                    `json:"id,omitempty"`
	InactiveFlag bool                   `json:"inactiveFlag,omitempty"`
	Name         string                 `json:"name"`
}

type CampaignType struct {
	Info        map[string]interface{} `json:"_info,omitempty"`
	DefaultFlag bool                   `json:"defaultFlag,omitempty"`
	ID          int                    `json:"id,omitempty"`
	Name        string                 `json:"name"`
}

type CampaignTypeInfo struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type EmailOpened struct {
	CampaignId int       `json:"campaignId,omitempty"`
	ContactId  int       `json:"contactId,omitempty"`
	DateOpened time.Time `json:"dateOpened,omitempty"`
	ID         int       `json:"id,omitempty"`
}

type FormSubmitted struct {
	CampaignId    int       `json:"campaignId,omitempty"`
	ContactId     int       `json:"contactId,omitempty"`
	DateSubmitted time.Time `json:"dateSubmitted,omitempty"`
	ID            int       `json:"id,omitempty"`
	PageSubType   string    `json:"pageSubType,omitempty"`
	PageType      string    `json:"pageType,omitempty"`
	QueryString   string    `json:"queryString,omitempty"`
	Status        string    `json:"status,omitempty"`
	Topic         string    `json:"topic,omitempty"`
	URL           string    `json:"url"`
	Version       string    `json:"version,omitempty"`
}

type Group struct {
	Info              map[string]interface{} `json:"_info,omitempty"`
	ID                int                    `json:"id,omitempty"`
	InactiveFlag      bool                   `json:"inactiveFlag,omitempty"`
	Name              string                 `json:"name"`
	PublicDescription string                 `json:"publicDescription,omitempty"`
	PublicFlag        bool                   `json:"publicFlag,omitempty"`
}

type GroupInfo struct {
	Info map[string]interface{} `json:"_info,omitempty"`
	ID   int                    `json:"id,omitempty"`
	Name string                 `json:"name,omitempty"`
}

type LinkClicked struct {
	CampaignId  int       `json:"campaignId,omitempty"`
	ContactId   int       `json:"contactId,omitempty"`
	DateClicked time.Time `json:"dateClicked,omitempty"`
	ID          int       `json:"id,omitempty"`
	QueryString string    `json:"queryString,omitempty"`
	URL         string    `json:"url"`
}

type MarketingCompany struct {
	Info               map[string]interface{} `json:"_info,omitempty"`
	AllContactsFlag    bool                   `json:"allContactsFlag,omitempty"`
	DefaultContactFlag bool                   `json:"defaultContactFlag,omitempty"`
	GroupId            int                    `json:"groupId,omitempty"`
	ID                 int                    `json:"id,omitempty"`
	UnsubscribeFlag    bool                   `json:"unsubscribeFlag,omitempty"`
}

type MarketingContact struct {
	Info            map[string]interface{} `json:"_info,omitempty"`
	GroupId         int                    `json:"groupId,omitempty"`
	ID              int                    `json:"id,omitempty"`
	Note            string                 `json:"note,omitempty"`
	UnsubscribeFlag bool                   `json:"unsubscribeFlag,omitempty"`
}

type TypeSubTypeCampaignSubType struct {
	Info   map[string]interface{} `json:"_info,omitempty"`
	ID     int                    `json:"id,omitempty"`
	Name   string                 `json:"name,omitempty"`
	TypeId int                    `json:"typeId,omitempty"`
}
