package main

import (
	"fmt"
	"strconv"

	st "../Stack"
)

//判斷一個字符是不是運算符(* + - /) 42 43 45 47
func IsOperator(val int) bool {

	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

//運算方法
func Cal(num1, num2 int, oper int) (res int) {
	// num1 = num1 - 48
	// num2 = num2 - 48
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("運算符錯誤")
	}
	return res
}

//判斷優先級 (*, / 1) (+, - 0)
func operPriority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}
	return res
}

func main() {
	numStack := st.NewStack()
	operStack := st.NewStack()

	exp := "300+30*6-4-6"

	//定義一個index幫助掃描
	index := 0
	//為了配合運算
	num1 := 0
	num2 := 0
	oper := 0
	result := 0
	keepNum := ""
	for {
		//增加一個邏輯
		//處理多位數問題
		ch := exp[index : index+1] //從exp取出charactor
		temp := int([]byte(ch)[0])

		if IsOperator(temp) { //如果是符號
			if operStack.Top == -1 { //如果是空
				operStack.Push(temp)
			} else {
				//判斷Stack內部的符號優先級
				if operPriority(operStack.Data[operStack.Top]) > operPriority(temp) {
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					oper, _ = operStack.Pop()
					result = Cal(num1, num2, oper)
					numStack.Push(result)
					operStack.Push(temp)
				} else {
					operStack.Push(temp)
				}
			}

		} else { //如果是數字
			//處理多位數的思路
			//先定義一個變數keepnum string做拼接
			keepNum += ch
			//每次要向index的後面字符測試一下 看看是不是運算符
			if index == len(exp)-1 { //如果已經到表達是最後 直接將keepNum做拼接
				val, _ := strconv.ParseInt(keepNum, 10, 64)
				numStack.Push(int(val))
			} else {
				//向後探一位
				if IsOperator(int([]byte(exp[index+1 : index+2])[0])) {
					val, _ := strconv.ParseInt(keepNum, 10, 64)
					numStack.Push(int(val))
					keepNum = ""
				}
			}

		}
		//判斷是否繼續掃描
		if index == len(exp)-1 {
			break
		} else {
			index++
		}
	}
	for {
		if operStack.Top == -1 {
			break
		}
		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = operStack.Pop()
		result = Cal(num1, num2, oper)
		numStack.Push(result)
	}
	//如果算法沒有問題 表達是也是正確 結果就是numStack最後的數
	res, _ := numStack.Pop()
	fmt.Println("表達式結果", exp, "=", res)
}
