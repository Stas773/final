package main

import (
	"context"
	"encoding/json"
	"final/cmd/models"
	"final/controller"
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
	SMSHandler      usecase.SMSWork      = &controller.Handler{}
	MMSHandler      usecase.MMSWork      = &controller.Handler{}
	VoiceHandler    usecase.VoiceWork    = &controller.Handler{}
	EmailHandler    usecase.EmailWork    = &controller.Handler{}
	BillingHandler  usecase.BillingWork  = &controller.Handler{}
	SupportHandler  usecase.SupportWork  = &controller.Handler{}
	IncidenrHandler usecase.IncidentWork = &controller.Handler{}
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
			Logger.Info("Server didn't start", err)
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
				MMSResult = MMSReader(Logger)
			}()
			go func() {
				defer wg.Done()
				VoiceResult = VoiceReader(Logger)
			}()
			go func() {
				defer wg.Done()
				EmailResult = EmailReader(Logger)
			}()
			go func() {
				defer wg.Done()
				BillingResult = BillingReader(Logger)
			}()
			go func() {
				defer wg.Done()
				SupportResult = SupportReader(Logger)
			}()
			go func() {
				defer wg.Done()
				IncidentResult = IncidentReader(Logger)
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
	data := SMSHandler.SMSReader(l)
	return data
}

func MMSReader(l *logrus.Logger) [][]entities.MMSData {
	data, err := MMSHandler.MMSReader(l)
	if err != nil {
		l.Error(err)
		return nil
	}
	if data == nil {
		l.Warn("MMS data is empty:", data)
		return nil
	}
	return data
}

func EmailReader(l *logrus.Logger) map[string][][]entities.EmailData {
	data := EmailHandler.EmailReader(l)
	return data
}

func VoiceReader(l *logrus.Logger) []entities.VoiceData {
	data := VoiceHandler.VoiceReader(l)
	return data
}

func BillingReader(l *logrus.Logger) entities.BillingData {
	data := BillingHandler.BillingReader(l)
	return data
}

func SupportReader(l *logrus.Logger) []int {
	data, err := SupportHandler.SupportReader(l)
	if err != nil {
		l.Error(err)
		return nil
	}
	if data == nil {
		l.Warn("Support data is empty:", data)
		return nil
	}
	return data
}

func IncidentReader(l *logrus.Logger) []entities.IncidentData {
	data, err := IncidenrHandler.IncidentReader(l)
	if err != nil {
		l.Error(err)
		return nil
	}
	if data == nil {
		l.Warn("Incident data is empty:", data)
		return nil
	}
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
