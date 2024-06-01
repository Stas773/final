package main

import (
	"context"
	"encoding/json"
	"final/billing"
	"final/email"
	"final/entities"

	"final/incident"
	"final/logger"
	"final/mms"
	"final/sms"
	"final/support"
	"final/usecase"
	"final/voice"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

var (
	SMSHandler      usecase.SMSWork      = &sms.SMSStruct{}
	MMSHandler      usecase.MMSWork      = &mms.MMSStract{}
	VoiceHandler    usecase.VoiceWork    = &voice.VoiceStruct{}
	EmailHandler    usecase.EmailWork    = &email.EmailStruct{}
	BillingHandler  usecase.BillingWork  = &billing.BillingStruct{}
	SupportHandler  usecase.SupportWork  = &support.SupportStract{}
	IncidenrHandler usecase.IncidentWork = &incident.IncidentStract{}
	SMSResult       [][]entities.SMSData
	MMSResult       [][]entities.MMSData
	VoiceResult     []entities.VoiceData
	EmailResult     map[string][][]entities.EmailData
	BillingResult   entities.BillingData
	SupportResult   []int
	IncidentResult  []entities.IncidentData
	ResultAll       []byte
	wg              sync.WaitGroup
)

func main() {
	logger.NewLogger()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	router := mux.NewRouter()
	router.HandleFunc("/", HandleConnection)
	server := &http.Server{
		Addr:    "localhost:8282",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Logger.Info("Server don't started", err)
			os.Exit(1)
		}
	}()
	logger.Logger.Info("Server started on port:", server.Addr)

	go func() {
		for {
			wg.Add(7)
			go func() {
				defer wg.Done()
				SMSResult = SMSReader()
			}()
			go func() {
				defer wg.Done()
				MMSResult = MMSReader()
			}()
			go func() {
				defer wg.Done()
				VoiceResult = VoiceReader()
			}()
			go func() {
				defer wg.Done()
				EmailResult = EmailReader()
			}()
			go func() {
				defer wg.Done()
				BillingResult = BillingReader()
			}()
			go func() {
				defer wg.Done()
				SupportResult = SupportReader()
			}()
			go func() {
				defer wg.Done()
				IncidentResult = IncidentReader()
			}()
			wg.Wait()
			ResultAll = ResultReader()

			time.Sleep(time.Second * 30)
		}
	}()

	<-done
	logger.Logger.Info("Stop signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Logger.Error("Shutdown error:", err)
	}
	logger.Logger.Info("Server stoped")
}

func SMSReader() [][]entities.SMSData {
	logger.Logger.Info("SMS file are contain:", SMSHandler.SMSReader())
	return SMSHandler.SMSReader()
}

func MMSReader() [][]entities.MMSData {
	data, err := MMSHandler.MMSReader()
	if err != nil {
		logger.Logger.Error(err)
	}
	if data == nil {
		logger.Logger.Warn("MMS data is empty:", data)
		return nil
	}
	logger.Logger.Info("MMS data are contain: ", data)
	return data
}

func EmailReader() map[string][][]entities.EmailData {
	logger.Logger.Info("Email file are contain:", EmailHandler.EmailReader())
	return EmailHandler.EmailReader()
}

func VoiceReader() []entities.VoiceData {
	logger.Logger.Info("Voice file are contain:", VoiceHandler.VoiceReader())
	return VoiceHandler.VoiceReader()
}

func BillingReader() entities.BillingData {
	logger.Logger.Info("Billing file are contain:", BillingHandler.BillingReader())
	return BillingHandler.BillingReader()
}

func SupportReader() []int {
	data, err := SupportHandler.SupportReader()
	if err != nil {
		logger.Logger.Error(err)
	}
	if data == nil {
		logger.Logger.Warn("Support data is empty:", data)
		return nil
	}
	logger.Logger.Info("Support data are contain:", data)
	return data
}

func IncidentReader() []entities.IncidentData {
	data, err := IncidenrHandler.IncidentReader()
	if err != nil {
		logger.Logger.Error(err)
	}
	if data == nil {
		logger.Logger.Warn("Incident data is empty:", data)
		return nil
	}
	logger.Logger.Info("Incident data are contain:", data)
	return data
}

func HandleConnection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Write(ResultAll)
}

func ResultReader() []byte {
	resultSet := entities.ResultSet{
		SMS:       SMSResult,
		MMS:       MMSResult,
		VoiceCall: VoiceResult,
		Email:     EmailResult,
		Billing:   BillingResult,
		Support:   SupportResult,
		Incidents: IncidentResult,
	}

	var result entities.Result
	if resultSet.SMS == nil || resultSet.MMS == nil || resultSet.VoiceCall == nil || resultSet.Email == nil || resultSet.Support == nil || resultSet.Incidents == nil {
		result.Status = false
		result.Error = "Error on collect data"
	} else {
		result.Status = true
		result.Data = resultSet
		result.Error = ""
	}

	jsonData, err := json.Marshal(&result)
	if err != nil {
		logger.Logger.Error("Error:", err)
		return nil
	}
	logger.Logger.Info("Result:", string(jsonData))
	return jsonData
}
