package types

import (
	"time"
)

type Calendar struct {
	Info interface{} `json:"_info,omitempty"`
	FridayEndTime string `json:"fridayEndTime,omitempty"`
	FridayStartTime string `json:"fridayStartTime,omitempty"`
	HolidayList HolidayListReference `json:"holidayList,omitempty"`
	ID int `json:"id,omitempty"`
	MondayEndTime string `json:"mondayEndTime,omitempty"`
	MondayStartTime string `json:"mondayStartTime,omitempty"`
	Name string `json:"name"`
	SaturdayEndTime string `json:"saturdayEndTime,omitempty"`
	SaturdayStartTime string `json:"saturdayStartTime,omitempty"`
	SundayEndTime string `json:"sundayEndTime,omitempty"`
	SundayStartTime string `json:"sundayStartTime,omitempty"`
	ThursdayEndTime string `json:"thursdayEndTime,omitempty"`
	ThursdayStartTime string `json:"thursdayStartTime,omitempty"`
	TuesdayEndTime string `json:"tuesdayEndTime,omitempty"`
	TuesdayStartTime string `json:"tuesdayStartTime,omitempty"`
	WednesdayEndTime string `json:"wednesdayEndTime,omitempty"`
	WednesdayStartTime string `json:"wednesdayStartTime,omitempty"`
}

type CalendarInfo struct {
	Info interface{} `json:"_info,omitempty"`
	FridayEndTime string `json:"fridayEndTime,omitempty"`
	FridayStartTime string `json:"fridayStartTime,omitempty"`
	HolidayList HolidayListReference `json:"holidayList,omitempty"`
	ID int `json:"id,omitempty"`
	MondayEndTime string `json:"mondayEndTime,omitempty"`
	MondayStartTime string `json:"mondayStartTime,omitempty"`
	Name string `json:"name,omitempty"`
	SaturdayEndTime string `json:"saturdayEndTime,omitempty"`
	SaturdayStartTime string `json:"saturdayStartTime,omitempty"`
	SundayEndTime string `json:"sundayEndTime,omitempty"`
	SundayStartTime string `json:"sundayStartTime,omitempty"`
	ThursdayEndTime string `json:"thursdayEndTime,omitempty"`
	ThursdayStartTime string `json:"thursdayStartTime,omitempty"`
	TuesdayEndTime string `json:"tuesdayEndTime,omitempty"`
	TuesdayStartTime string `json:"tuesdayStartTime,omitempty"`
	WednesdayEndTime string `json:"wednesdayEndTime,omitempty"`
	WednesdayStartTime string `json:"wednesdayStartTime,omitempty"`
}

type Holiday struct {
	Info interface{} `json:"_info,omitempty"`
	AllDayFlag bool `json:"allDayFlag,omitempty"`
	Date string `json:"date"`
	HolidayList HolidayListReference `json:"holidayList,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name"`
	TimeEnd string `json:"timeEnd,omitempty"`
	TimeStart string `json:"timeStart,omitempty"`
}

type HolidayInfo struct {
	Info interface{} `json:"_info,omitempty"`
	AllDayFlag bool `json:"allDayFlag,omitempty"`
	Date string `json:"date,omitempty"`
	HolidayList HolidayListReference `json:"holidayList,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	TimeEnd string `json:"timeEnd,omitempty"`
	TimeStart string `json:"timeStart,omitempty"`
}

type HolidayList struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name"`
}

type HolidayListInfo struct {
	Info interface{} `json:"_info,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type PortalCalendar struct {
	Info interface{} `json:"_info,omitempty"`
	Adjust1End string `json:"adjust1End,omitempty"`
	Adjust1Hours float64 `json:"adjust1Hours,omitempty"`
	Adjust1Start string `json:"adjust1Start,omitempty"`
	Adjust2End string `json:"adjust2End,omitempty"`
	Adjust2Hours float64 `json:"adjust2Hours,omitempty"`
	Adjust2Start string `json:"adjust2Start,omitempty"`
	Adjust3End string `json:"adjust3End,omitempty"`
	Adjust3Hours float64 `json:"adjust3Hours,omitempty"`
	Adjust3Start string `json:"adjust3Start,omitempty"`
	ID int `json:"id,omitempty"`
	WeekStart string `json:"weekStart,omitempty"`
}

type ScheduleColor struct {
	Info interface{} `json:"_info,omitempty"`
	Color string `json:"color"`
	EndPercent int `json:"endPercent,omitempty"`
	ID int `json:"id,omitempty"`
	StartPercent int `json:"startPercent,omitempty"`
}

type ScheduleDetail struct {
	Info interface{} `json:"_info,omitempty"`
	DateEnd string `json:"dateEnd,omitempty"`
	DateStart string `json:"dateStart,omitempty"`
	ID int `json:"id,omitempty"`
	Member MemberReference `json:"member,omitempty"`
	ScheduleEntry ScheduleEntryReference `json:"scheduleEntry,omitempty"`
}

type ScheduleEntry struct {
	Info interface{} `json:"_info,omitempty"`
	AcknowledgedDate time.Time `json:"acknowledgedDate,omitempty"`
	AcknowledgedFlag bool `json:"acknowledgedFlag,omitempty"`
	AddMemberToProjectFlag bool `json:"addMemberToProjectFlag,omitempty"`
	AllowScheduleConflictsFlag bool `json:"allowScheduleConflictsFlag,omitempty"`
	CloseDate time.Time `json:"closeDate,omitempty"`
	DateEnd time.Time `json:"dateEnd,omitempty"`
	DateStart time.Time `json:"dateStart,omitempty"`
	DoneFlag bool `json:"doneFlag,omitempty"`
	EndTimeSet bool `json:"endTimeSet,omitempty"`
	Hours float64 `json:"hours,omitempty"`
	ID int `json:"id,omitempty"`
	MeetingFlag bool `json:"meetingFlag,omitempty"`
	Member MemberReference `json:"member,omitempty"`
	MobileGuid string `json:"mobileGuid,omitempty"`
	Name string `json:"name,omitempty"`
	NotificationResponse string `json:"notificationResponse,omitempty"`
	NotificationSent bool `json:"notificationSent,omitempty"`
	NotifyResource bool `json:"notifyResource,omitempty"`
	ObjectId int `json:"objectId,omitempty"`
	OwnerFlag bool `json:"ownerFlag,omitempty"`
	ProjectRoleId int `json:"projectRoleId,omitempty"`
	Reminder ReminderReference `json:"reminder,omitempty"`
	Span ScheduleSpanReference `json:"span,omitempty"`
	StartTimeSet bool `json:"startTimeSet,omitempty"`
	Status ScheduleStatusReference `json:"status,omitempty"`
	Type ScheduleTypeReference `json:"type"`
	Where ServiceLocationReference `json:"where,omitempty"`
}

type ScheduleEntryDetail struct {
	Info interface{} `json:"_info,omitempty"`
	DateEnd string `json:"dateEnd,omitempty"`
	DateStart string `json:"dateStart,omitempty"`
	HoursScheduled float64 `json:"hoursScheduled,omitempty"`
	ID int `json:"id,omitempty"`
	ScheduleEntry ScheduleEntryReference `json:"scheduleEntry,omitempty"`
}

type ScheduleReminderTime struct {
	Info interface{} `json:"_info,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Time int `json:"time,omitempty"`
}

type ScheduleStatus struct {
	Info interface{} `json:"_info,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name"`
	ShowAsTentativeFlag bool `json:"showAsTentativeFlag,omitempty"`
}

type ScheduleStatusInfo struct {
	Info interface{} `json:"_info,omitempty"`
	DefaultFlag bool `json:"defaultFlag,omitempty"`
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	ShowAsTentativeFlag bool `json:"showAsTentativeFlag,omitempty"`
}

type ScheduleType struct {
	Info interface{} `json:"_info,omitempty"`
	ChargeCode ChargeCodeReference `json:"chargeCode,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier"`
	Name string `json:"name"`
	SystemFlag bool `json:"systemFlag,omitempty"`
	Where ServiceLocationReference `json:"where,omitempty"`
}

type ScheduleTypeInfo struct {
	Info interface{} `json:"_info,omitempty"`
	ChargeCode ChargeCodeReference `json:"chargeCode,omitempty"`
	ID int `json:"id,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Name string `json:"name,omitempty"`
	SystemFlag bool `json:"systemFlag,omitempty"`
	Where ServiceLocationReference `json:"where,omitempty"`
}

