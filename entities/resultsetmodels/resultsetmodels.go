package resultsetmodels

import (
	"final/entities/billingmodels"
	"final/entities/emailmodels"
	"final/entities/incidentmodels"
	"final/entities/mmsmodels"
	"final/entities/smsmodels"
	"final/entities/voicemodels"
)

type ResultSet struct {
	SMS       [][]smsmodels.SMSData                `json:"sms"`
	MMS       [][]mmsmodels.MMSData                `json:"mms"`
	VoiceCall []voicemodels.VoiceData              `json:"voice_call"`
	Email     map[string][][]emailmodels.EmailData `json:"email"`
	Billing   billingmodels.BillingData            `json:"billing"`
	Support   []int                                `json:"support"`
	Incidents []incidentmodels.IncidentData        `json:"incident"`
}
