package routes

import (
	run "Kiva-api/models/Loan"

	"github.com/gin-gonic/gin"
)

func Listen() {
	router := gin.Default()
	router.POST("/requestIndividualLoan", run.SendPostIndividualLoanDraft)
	router.Run(":8181")
}
