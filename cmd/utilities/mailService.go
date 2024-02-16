package utilities

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/caarlos0/env"
	"github.com/casmelad/stori/cmd/dto"
)

type emailConfig struct {
	Password string `env:"SMTP_PASSWORD"`
	From     string `env:"SMTP_USEREMAIL"`
	Host     string `env:"SMTP_HOST"`
	Port     string `env:"SMTP_PORT"`
	To       string
	Subject  string
	Body     string
}

func NewAccountBalanceMailNotificationService() AccountBalanceMailNotificationService {
	cfg := emailConfig{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
	return AccountBalanceMailNotificationService{
		emailConfig: &cfg,
	}
}

type AccountBalanceMailNotificationService struct {
	emailConfig *emailConfig
}

func (ac AccountBalanceMailNotificationService) SendNotification(accountBalanceInfo dto.AccountBalanceNotificationDto) error {
	htmlString := fmt.Sprintf(`<html><head><title>Title of the document</title></head><body><table style="float:center"><thead><tr><th>
								<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/b/b0/Stori_Logo_2023.svg/512px-Stori_Logo_2023.svg.png" alt="imagen izzi" style="width:100px"></th>
								</tr><tr><th>Detalle de cuenta</th></tr></thead><tbody>{{montly_movements}}<tr><td style="text-align:center">Average Debit Ammount : %f</td></tr>
								<tr><td style="text-align:center">Average Debit Ammount :%f</td></tr><tr><td style="text-align:center">Total Balance :%f</td></tr></tbody></table></body></html>`,
		accountBalanceInfo.AvgDebitAmmount, accountBalanceInfo.AvgCreditAmmount, accountBalanceInfo.Balance)

	monthlyHTMLDetail := getHTMLTransactionsByMonth(accountBalanceInfo)
	htmlString = strings.Replace(htmlString, "{{montly_movements}}", monthlyHTMLDetail, 1)

	ac.emailConfig.To = accountBalanceInfo.To
	ac.emailConfig.Body = htmlString
	err := sendEmail(ac.emailConfig)
	if err != nil {
		return err
	}

	return nil
}

func getHTMLTransactionsByMonth(accData dto.AccountBalanceNotificationDto) string {

	htmlString := ""
	htmlString += getHTMLTextFromTransactionsByMonth(accData.JanuaryTransactions, "January")
	htmlString += getHTMLTextFromTransactionsByMonth(accData.FebruaryTransactions, "February")
	htmlString += getHTMLTextFromTransactionsByMonth(accData.MarchTransactions, "March")
	htmlString += getHTMLTextFromTransactionsByMonth(accData.AprilTransactions, "April")
	htmlString += getHTMLTextFromTransactionsByMonth(accData.MayTransactions, "May")
	htmlString += getHTMLTextFromTransactionsByMonth(accData.JuneTransactions, "June")
	htmlString += getHTMLTextFromTransactionsByMonth(accData.JulyTransactions, "July")
	htmlString += getHTMLTextFromTransactionsByMonth(accData.AugustTransactions, "August")
	htmlString += getHTMLTextFromTransactionsByMonth(accData.SeptemberTransactions, "September")
	htmlString += getHTMLTextFromTransactionsByMonth(accData.OctoberTransactions, "October")
	htmlString += getHTMLTextFromTransactionsByMonth(accData.NovemberTransactions, "November")
	htmlString += getHTMLTextFromTransactionsByMonth(accData.DecemberTransactions, "December")

	return htmlString
}

func getHTMLTextFromTransactionsByMonth(txQty int, monthName string) string {
	if txQty > 0 {
		return fmt.Sprintf("<tr style=\"text-align:center\"><td>Number of transactions in %s: %d</td></tr>", monthName, txQty)
	}
	return ""
}

func sendEmail(emailData *emailConfig) error {
	to := []string{emailData.To}
	address := emailData.Host + ":" + emailData.Port

	subject := "Subject: Revisa tu estado de cuentasa par favar\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := []byte(subject + mime + emailData.Body)

	auth := smtp.PlainAuth("", emailData.From, emailData.Password, emailData.Host)
	err := smtp.SendMail(address, auth, emailData.From, to, message)
	if err != nil {
		return err
	}
	return nil
}
