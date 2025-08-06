package handler

import (
	"net/http"
)

type ExpenseHandler struct {
}

func (eh *ExpenseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// m := r.Method
	// var res string
	// dbConfig := config.ExpenseTrackerDBConfig{DbName: "expensetracker", DbHost: "localhost", DbPort: "5440", DbUser: "user", DbPassword: "admin"}
	// repo := repository.NewExpenseRepository(dbConfig)
	// service := service.NewExpenseService(repo, dbConfig)

	// switch m {
	// case "GET":
	// 	fmt.Println(r.URL.Path)
	// 	expense, err := service.GetExpense(2)
	// 	expenseResponse := dto.BuildExpenseResponse(expense)
	// 	jsonData, err := json.Marshal(expenseResponse)
	// 	if err != nil {
	// 		fmt.Println("Error marshalling JSON:", err)
	// 		return
	// 	}
	// 	data, err := w.Write(jsonData)
	// 	if err != nil {
	// 		fmt.Println("Error marshalling JSON:", err)
	// 		return
	// 	}
	// 	fmt.Println("Write response :: ", data)
	// case "POST":
	// 	res = "this is a s default response from POST method."
	// case "PUT":
	// 	res = "this a default response from PUT method."
	// case "DELETE":
	// 	res = "this a default response from DELETE method."
	// default:
	// 	res = "this a default response from NO DETECTED method."
	// }
	// io.WriteString(w, res)
}
