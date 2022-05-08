package processor

import (
	"bufio"
	"fmt"
	"ledgerCo-loans/controller"
	"ledgerCo-loans/entity"
	"ledgerCo-loans/utils"
	"os"
	"strings"
)

//FileProcessor struct
type FileProcessor struct {
	FilePath string
}

// FileProcessor method process
func (fileProcessor FileProcessor) Process() bool {
	var inputCmdLines = getInputsFromFile(fileProcessor.FilePath)
	for _, inputCmdData := range inputCmdLines {
		fields := strings.Split(inputCmdData, " ")
		if len(fields) > 0 {
			switch fields[0] {
			case "LOAN":
				processLoan(fields)
			case "PAYMENT":
				processPayment(fields)
			case "BALANCE":
				res, ok := processBalance(fields)
				if ok {
					fmt.Println(res)
				} else {
					fmt.Println("Error in processing balance")
				}
			}
		}
	}
	return true
}

var getInputsFromFile = func(filePath string) []string {
	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	return fileLines
}

func processLoan(fields []string) {
	if len(fields) == 6 { // simple check on number of fields in LOAN cmd
		loanCmdData := entity.LoanCmdData{
			BankName:        fields[1],
			BorrowerName:    fields[2],
			PrincipalAmount: utils.ParseFloat(fields[3]),
			TenureYears:     utils.ParseUint(fields[4]),
			InterestRate:    utils.ParseFloat(fields[5]),
		}
		ctlr := controller.LoanCtlr{}
		ctlr.Do(loanCmdData)
	}
}

func processPayment(fields []string) {
	if len(fields) == 5 { // simple check on number of fields in PAYMENT cmd
		payCmdData := entity.PaymentCmdData{
			BankName:     fields[1],
			BorrowerName: fields[2],
			AmountPaid:   utils.ParseFloat(fields[3]),
			EmiNumber:    utils.ParseUint(fields[4]),
		}
		ctlr := controller.PaymentCtlr{}
		ctlr.Do(payCmdData)
	}
}

func processBalance(fields []string) (entity.BalanceOutput, bool) {
	if len(fields) == 4 { // simple check on number of fields in BALANCE cmd
		balCmdData := entity.BalanceCmdData{
			BankName:     fields[1],
			BorrowerName: fields[2],
			EmiNumber:    utils.ParseUint(fields[3]),
		}
		ctlr := controller.BalanceCtlr{}
		return ctlr.Do(balCmdData)
	}
	return entity.BalanceOutput{}, false
}
