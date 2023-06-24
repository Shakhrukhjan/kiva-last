package Loan

import (
	t "Kiva-api/autorization"
	logged "Kiva-api/logs"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
func SendPostIndividualLoanDraft(c *gin.Context) {
	var loan IndividualLoan

	var entreps Entreps
	url := "https://partner-api.k1.kiva.org/v3/partner/63/loan_draft"
	method := "POST"
	activityID, err := strconv.Atoi(c.Query("activity_id"))
	if err != nil {
		logged.ErrL.Printf("Преобразование типа  %v activityID./n", err)
		return
	}
	clientWaiverSigned, err := strconv.ParseBool(c.Query("client_waiver_signed"))
	if err != nil {
		logged.ErrL.Printf("Преобразование типа  %v ClientWaiverSigned./n", err)
		return
	}
	descriptionLanguageID, err := strconv.Atoi(c.Query("description_language_id"))
	if err != nil {
		logged.ErrL.Printf("Преобразование типа DescriptionLanguageID %v ", err)
		return
	}
	themeTypeID, err := strconv.Atoi(c.Query("theme_type_id"))
	if err != nil {
		logged.ErrL.Printf("Преобразование типа ThemeTypeID %v ", err)
	}
	var currency string = c.Query("currency")

	var description string = c.Query("description")
	var disburseTime string = c.Query("disburse_time")
	var loanuse string = c.Query("loanuse")
	var location string = c.Query("location")
	schedule_count, err := strconv.Atoi(c.Query("schedule_count"))

	if err != nil {
		logged.ErrL.Println("Преобразование типа schedule_count ", err.Error())
	}
	err = c.Bind(&entreps)
	if err != nil {
		logged.ErrL.Printf("Error in entreps %v ", err.Error())
		return
	}
	loan.SetEntreps(entreps)
	schedules := make([]Schedule, schedule_count)
	for i := 0; i < schedule_count; i++ {

		schedules[i].Date = c.Query(fmt.Sprintf("date_%d", i+1))
		schedules[i].Interest, err = strconv.ParseFloat(c.Query(fmt.Sprintf("interest_%d", i+1)), 8)
		if err != nil {
			logged.ErrL.Println("Error in transformation Interest", err.Error())
		}
		schedules[i].Principal, err = strconv.ParseFloat(c.Query(fmt.Sprintf("principal_%d", i+1)), 8)
		if err != nil {
			logged.ErrL.Println("Error in transformation Principal", err.Error())
		}
	}
	loan.SetSchedule(schedules)
	loan.SetActivityID(activityID)
	loan.SetClient_waiver_signed(clientWaiverSigned)
	loan.SetCurrency(currency)
	loan.SetDescription(description)
	loan.SetDescriptionLanguageID(descriptionLanguageID)
	loan.SetDisburseTime(disburseTime)
	loan.SetLoanuse(loanuse)
	loan.SetLocation(location)
	loan.SetThemeTypeID(themeTypeID)
	log.Println("===============================================================================================================================================================================================================")
	log.Println("Структура loanDraft ----> ", loan)
	log.Println("===============================================================================================================================================================================================================")


	if loan.Entreps[0].LastName == "" {
		loan.Entreps[0].LastName = " "
	}

	bytesRepresentation, err := json.Marshal(loan)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(bytesRepresentation))

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", t.AccessToken()))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return
	}
	logged.InfoL.Println("Контракт успешно отправлен на сайт!")
	contract := fmt.Sprintf("%v", string(body))
	c.JSON(http.StatusOK, gin.H{
		"message":  "Контракт успешно отправлен на сайт!",
		"response": contract,
	})
	fmt.Println(string(body))
}
