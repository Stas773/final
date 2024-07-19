package main

import (
	"context"
	"encoding/json"
	"final/cmd/models"
	"final/entities"
	"final/usecase"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path"
	"runtime"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var (
	SMSHandler      usecase.SMSWork      = &usecase.BuilderStruct{}
	MMSHandler      usecase.MMSWork      = &usecase.BuilderStruct{}
	VoiceHandler    usecase.VoiceWork    = &usecase.BuilderStruct{}
	EmailHandler    usecase.EmailWork    = &usecase.BuilderStruct{}
	BillingHandler  usecase.BillingWork  = &usecase.BuilderStruct{}
	SupportHandler  usecase.SupportWork  = &usecase.BuilderStruct{}
	IncidenrHandler usecase.IncidentWork = &usecase.BuilderStruct{}
	SMSResult       [][]entities.SMSData
	MMSResult       [][]entities.MMSData
	VoiceResult     []entities.VoiceData
	EmailResult     map[string][][]entities.EmailData
	BillingResult   entities.BillingData
	SupportResult   []int
	IncidentResult  []entities.IncidentData
	ResultAll       []byte
	wg              sync.WaitGroup
	Logger          *logrus.Logger
)

func main() {
	Logger := NewLogger()
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
			Logger.Info("Server don't started", err)
			os.Exit(1)
		}
	}()
	Logger.Info("Server started on port:", server.Addr)

	go func() {
		for {
			wg.Add(7)
			go func() {
				defer wg.Done()
				SMSResult = SMSReader(Logger)
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
	Logger.Info("Stop signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		Logger.Error("Shutdown error:", err)
	}
	Logger.Info("Server stoped")
}

func NewLogger() *logrus.Logger {
	Logger = logrus.New()
	Logger.SetReportCaller(true)
	Logger.SetFormatter(&logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
		ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
	return Logger
}

func SMSReader(l *logrus.Logger) [][]entities.SMSData {
	l.Info("SMS file are contain:", SMSHandler.SMSReader(l))
	return SMSHandler.SMSReader(l)
}

func MMSReader() [][]entities.MMSData {
	data, err := MMSHandler.MMSReader()
	if err != nil {
		Logger.Error(err)
	}
	if data == nil {
		Logger.Warn("MMS data is empty:", data)
		return nil
	}
	Logger.Info("MMS data are contain:", data)
	return data
}

func EmailReader() map[string][][]entities.EmailData {
	Logger.Info("Email file are contain:", EmailHandler.EmailReader())
	return EmailHandler.EmailReader()
}

func VoiceReader() []entities.VoiceData {
	Logger.Info("Voice file are contain:", VoiceHandler.VoiceReader())
	return VoiceHandler.VoiceReader()
}

func BillingReader() entities.BillingData {
	Logger.Info("Billing file are contain:", BillingHandler.BillingReader())
	return BillingHandler.BillingReader()
}

func SupportReader() []int {
	data, err := SupportHandler.SupportReader()
	if err != nil {
		Logger.Error(err)
	}
	if data == nil {
		Logger.Warn("Support data is empty:", data)
		return nil
	}
	Logger.Info("Support data are contain:", data)
	return data
}

func IncidentReader() []entities.IncidentData {
	data, err := IncidenrHandler.IncidentReader()
	if err != nil {
		Logger.Error(err)
	}
	if data == nil {
		Logger.Warn("Incident data is empty:", data)
		return nil
	}
	Logger.Info("Incident data are contain:", data)
	return data
}

func HandleConnection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Write(ResultAll)
}

func ResultReader() []byte {
	resultSet := models.ResultSet{
		SMS:       SMSResult,
		MMS:       MMSResult,
		VoiceCall: VoiceResult,
		Email:     EmailResult,
		Billing:   BillingResult,
		Support:   SupportResult,
		Incidents: IncidentResult,
	}

	var result models.Result
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
		Logger.Error("Error:", err)
		return nil
	}
	Logger.Info("Result:", string(jsonData))
	return jsonData
}
