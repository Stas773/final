package main

import (
	"context"
	"encoding/json"
	"final/cmd/models"
	"final/controller"
	"final/entities"
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
	SMSResult      [][]entities.SMSData
	MMSResult      [][]entities.MMSData
	VoiceResult    []entities.VoiceData
	EmailResult    map[string][][]entities.EmailData
	BillingResult  entities.BillingData
	SupportResult  []int
	IncidentResult []entities.IncidentData
	ResultAll      []byte
	wg             sync.WaitGroup
	Logger         *logrus.Logger
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
				SMSResult = controller.SMSReader(Logger)
			}()
			go func() {
				defer wg.Done()
				MMSResult = controller.MMSReader(Logger)
			}()
			go func() {
				defer wg.Done()
				VoiceResult = controller.VoiceReader(Logger)
			}()
			go func() {
				defer wg.Done()
				EmailResult = controller.EmailReader(Logger)
			}()
			go func() {
				defer wg.Done()
				BillingResult = controller.BillingReader(Logger)
			}()
			go func() {
				defer wg.Done()
				SupportResult = controller.SupportReader(Logger)
			}()
			go func() {
				defer wg.Done()
				IncidentResult = controller.IncidentReader(Logger)
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
