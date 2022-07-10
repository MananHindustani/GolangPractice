package	 main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sharmamanan796/onlineBank/Account"
	"github.com/sharmamanan796/onlineBank/Entity"
)

func main() {
	server := gin.Default()
	bank := server.Group("/api/v1/bank")
	{
		bank.POST("/account", func(context *gin.Context) {

			var account = &Entity.Account{}
			req := context.Request
			err := json.NewDecoder(req.Body).Decode(account)
			//		ctx.JSON(200, accountController.FindAll())
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				fmt.Println(err.Error())
				return
			}
			err = Account.CreateAccount(account)
			if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			context.JSON(http.StatusOK, gin.H{
				"status":  "account created",
				"message": account,
			})
		})
		bank.GET("/accounts", func(context *gin.Context) {
			//		ctx.JSON(200, accountController.FindAll())
			var account []Entity.Account
			account, err := Account.ListAccount(account)
			if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			context.JSON(http.StatusOK, gin.H{
				"status":  "account details",
				"message": account,
			})

		})
		bank.GET("/accounts/account/:id", func(context *gin.Context) {
			id, _ := strconv.Atoi(context.Param("id"))
			var account Entity.Account
			account, err, msg := Account.AccountById(id)
			if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			context.JSON(http.StatusOK, gin.H{
				"status":  "account details",
				"message": msg,
				"body":    account,
			})

		})
		bank.DELETE("/account/:id", func(context *gin.Context) {
			id, _ := strconv.Atoi(context.Param("id"))
			count, err := Account.DeleteAccount(id)
			if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			output := strconv.FormatInt(count, 10) + " row deleted"
			context.JSON(http.StatusOK, gin.H{
				"status":  "account delete",
				"message": output,
			})

		})
		bank.PUT("/account/:id", func(context *gin.Context) {
			var account = &Entity.Account{}
			req := context.Request
			err := json.NewDecoder(req.Body).Decode(account)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			id, _ := strconv.Atoi(context.Param("id"))
			count, err := Account.UpdateAccount(account, id)
			if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			output := strconv.FormatInt(count, 10) + " row update"
			context.JSON(http.StatusOK, gin.H{
				"status":  "account delete",
				"message": output,
			})
		})
	}

	server.GET("/ping", func(context *gin.Context) {
		//	context.JSON(200, "Pong")
		firstname := context.DefaultQuery("firstname", "Guest")
		lastname := context.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
		context.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	server.POST("/upload", func(context *gin.Context) {
		firstname := context.DefaultQuery("firstname", "Guest")
		email := context.Query("email") // shortcut for c.Request.URL.Query().Get("lastname")
		file, _ := context.FormFile("file")
		log.Println(file.Filename)
		// Upload the file to specific dst.
		context.SaveUploadedFile(file, "/home/billion")
		context.String(http.StatusOK, fmt.Sprintf("'%s' uploaded! %s %s", file.Filename, firstname, email))
	})
	server.Run(":7000")
}
