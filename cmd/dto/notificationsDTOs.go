package dto

type AccountBalanceNotificationDto struct {
	To                    string
	Balance               float64
	AvgCreditAmmount      float64
	AvgDebitAmmount       float64
	JanuaryTransactions   int
	FebruaryTransactions  int
	MarchTransactions     int
	AprilTransactions     int
	MayTransactions       int
	JuneTransactions      int
	JulyTransactions      int
	AugustTransactions    int
	SeptemberTransactions int
	OctoberTransactions   int
	NovemberTransactions  int
	DecemberTransactions  int
}
