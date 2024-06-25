package billing

import (
	"bufio"
	"final/entities"
	"final/logger"
	"math"
	"os"
	"strconv"
)

type BillingStruct struct {
}

func (bs *BillingStruct) BillingReader() entities.BillingData {
	var dataStructs entities.BillingData
	var sum uint8

	file, err := os.Open(fileName)
	if err != nil {
		logger.Logger.Panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineInt, err := strconv.Atoi(line)
		if err != nil {
			logger.Logger.Error("can't convert string to int:", err)
		}
		for i := 0; i < len(line); i++ {
			b := lineInt % 10
			switch i {
			case 0:
				if b > 0 {
					sum += uint8(math.Pow(2, float64(i)))
					dataStructs.CreateCustomer = true
				}
			case 1:
				if b > 0 {
					sum += uint8(math.Pow(2, float64(i)))
					dataStructs.Purchase = true
				}
			case 2:
				if b > 0 {
					sum += uint8(math.Pow(2, float64(i)))
					dataStructs.Payout = true
				}
			case 3:
				if b > 0 {
					sum += uint8(math.Pow(2, float64(i)))
					dataStructs.Recurring = true
				}
			case 4:
				if b > 0 {
					sum += uint8(math.Pow(2, float64(i)))
					dataStructs.FraudControl = true
				}
			case 5:
				if b > 0 {
					sum += uint8(math.Pow(2, float64(i)))
					dataStructs.CheckoutPage = true
				}
			}
			lineInt /= 10
		}
	}
	logger.Logger.Info("Billing sum:", sum)
	return dataStructs
}
