package main

import (
	"fmt"
	"os"
	"time"

	"github.com/casmelad/stori/cmd/codecs"
	"github.com/casmelad/stori/cmd/utilities"
	"github.com/casmelad/stori/persistence"
	"github.com/casmelad/stori/pkg/accounts"
	"github.com/go-co-op/gocron/v2"
	"github.com/go-kit/log"
)

func main() {
	//log
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// create a scheduler
	s, err := gocron.NewScheduler()
	if err != nil {
		fmt.Println(err)
	}

	// add a job to the scheduler
	_, err = s.NewJob(
		gocron.DurationJob(
			10*time.Second,
		),
		gocron.NewTask(
			func() {
				startAccountBalanceProccess(logger)
			},
		),
	)
	if err != nil {
		fmt.Println(err)
	}

	// start the scheduler
	s.Start()

	time.Sleep(50 * time.Second)

	// when you're done, shut it down
	err = s.Shutdown()
	if err != nil {
		fmt.Println(err)
	}
}

func startAccountBalanceProccess(log log.Logger) {
	txs := readTransactionsInfoFromFile(log)
	saveAndNotify(txs, log)
}

func readTransactionsInfoFromFile(log log.Logger) []accounts.AccountTransaction {

	fileReader := utilities.FileReader{}
	fileContent, err := fileReader.ReadTransationLinesFromCsvFile("https://raw.githubusercontent.com/casmelad/stori/main/txns.csv")
	if err != nil {
		fmt.Println(err)
	}

	txs, err := codecs.MapTransactionsFromText(fileContent)
	if err != nil {
		fmt.Println(err)
	}
	return txs
}

func saveAndNotify(txs []accounts.AccountTransaction, log log.Logger) {
	//save to Db
	dbConn, err := persistence.OpenDataBaseConnection()
	if err != nil {
		fmt.Println(err)
	}

	repository := persistence.NewPostgreRepository(dbConn)
	service := accounts.NewAccountService(repository)

	accountInfo, err := service.GetAccountInfo("100765987777")
	if err != nil {
		fmt.Println(err)
	}

	for index, tx := range txs {
		txs[index].AccountId = accountInfo.Id
		fmt.Println(tx.AccountId)
	}

	err = service.Save(txs)
	if err != nil {
		fmt.Println(err)
	}
	accountInfo.AccountTransactions = txs
	notifyAccountBalance(accountInfo, log)
}

func notifyAccountBalance(accountInfo accounts.Account, log log.Logger) {
	notificationData := codecs.MapAccountToAccountBalanceNotificationDto(accountInfo)
	emailService := utilities.NewAccountBalanceMailNotificationService()
	err := emailService.SendNotification(notificationData)
	if err != nil {
		fmt.Println(err)
	}
}
