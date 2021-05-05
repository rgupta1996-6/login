package controllers

import (
	"fmt"
	"go-auth/database"
	"go-auth/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ShowAll(c *fiber.Ctx) error {
	var data map[string]string
	var count int64
	var DataToBeSent models.DataToBeSent

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	fmt.Println(data)

	accID := data["query"]
	fmt.Println("Account ID:", accID)

	if accID != "" {
		custDetails := []models.CustDetails{}
		database.DB.Find(&custDetails, accID)
		DataToBeSent.CustDetails = custDetails
		DataToBeSent.Count = 1
		fmt.Println("custDetails:", DataToBeSent.CustDetails)
		return c.JSON(DataToBeSent)
	} else {

		custDetails := []models.CustDetails{}
		rowsPerPage, _ := strconv.Atoi(data["rowsPerPage"])
		page, _ := strconv.Atoi(data["page"])
		database.DB.Find(&custDetails).Count(&count)
		database.DB.Limit(rowsPerPage).Offset(rowsPerPage * page).Order(data["orderBy"] + " " + data["order"]).Find(&custDetails)
		DataToBeSent.CustDetails = custDetails
		DataToBeSent.Count = count

		return c.JSON(DataToBeSent)
	}

}

func AddNewAccount(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	accID, _ := strconv.Atoi(data["accID"])
	contact, _ := strconv.Atoi(data["contact"])
	bal, _ := strconv.Atoi(data["bal"])
	custDetails := models.CustDetails{
		AccID:   int32(accID),
		AccType: data["accType"],
		BCode:   data["bCode"],
		Contact: int64(contact),
		Balance: int32(bal),
	}

	database.DB.Create(&custDetails)

	return nil
}

func CreditBalance(c *fiber.Ctx) error {

	var data map[string]string

	fmt.Println(data["accID"])
	fmt.Println(data["amount"])

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	accID, _ := strconv.Atoi(data["accID"])
	amount, _ := strconv.Atoi(data["amount"])

	fmt.Println("Hello")
	fmt.Println(accID, amount)

	custDetails := models.CustDetails{}

	database.DB.First(&custDetails, accID)
	newBal := custDetails.Balance + int32(amount)

	res := database.DB.Exec("UPDATE cust_details SET balance=? WHERE acc_id=?", newBal, custDetails.AccID)
	if res.Error != nil {
		return c.JSON(fiber.Map{
			"message": "ERROR",
		})
	} else {
		return c.JSON(fiber.Map{
			"message": "SUCCESS",
		})
	}

}

func DebitBalance(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	accID, _ := strconv.Atoi(data["accID"])
	amount, _ := strconv.Atoi(data["amount"])

	custDetails := models.CustDetails{}

	database.DB.First(&custDetails, accID)
	fmt.Println(custDetails)
	newBal := custDetails.Balance - int32(amount)

	res:= database.DB.Exec("UPDATE cust_details SET balance=? WHERE acc_id=?", newBal, custDetails.AccID)
	if res.Error != nil {
		return c.JSON(fiber.Map{
			"message": "ERROR",
		})
	} else {
		return c.JSON(fiber.Map{
			"message": "SUCCESS",
		})
	}


}

func DeleteAccount(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	accID, _ := strconv.Atoi(data["accID"])

	database.DB.Delete(&models.CustDetails{}, accID)
	return nil
}
