package main

import (
	"fmt"
)
//Сортироват по приоритету создат приоритет
//создат тип ЛистНаме
//Аддресс на предадуший элемент- prev
//ддресс на следующий элемент - next
//value///

//Тип List
//размер ListNode
//firstNode - первий элемент
//lastNode - последний элемент зависит от метода пушбек
//написат метод ремув которий заканчивает очерид
type Student struct {
	name string
	score int64
}

type ListNode struct {
	next *ListNode
	//previs *ListNode
	name string
	purchase int64
	//prioritet  int64
}

func main() {
	var firsNode = ListNode{
		next:     nil,
		name:     "Jonibek Mahmudov",
		purchase: 100,
	}
	var secondNode = ListNode{
		next:     nil,
		name:     "Nekruz Rahimov",
		purchase: 3,
	}
	firsNode.next = &secondNode
	current := firsNode
	//fmt.Println(current)
	for current.next != nil {
		current = *current.next
	//	fmt.Println(current)
	}
	//arrayFirst := []int64{4, 5, 4}



	array := []Student{
		{
			name:  "Jonibek",
			score: 100,
		},
		{
			name: "Nekruz",
			score: 25,
		},
		{
			name:  "Sayfulo",
			score: 15,
		},
	}
	fmt.Println(array)
	for index := 0; index < len(array); index++{
		for jndex := 1; jndex < len(array); jndex++ {
			if array[index - 1].score > array[jndex].score{
				helper := array[jndex - 1]
				array[jndex - 1] = array[jndex]
				array[jndex] = helper
			}
		}
	}
	fmt.Println(array)
}

//создать метод для ТипЛис  имя метода пушбек