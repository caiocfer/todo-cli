package todo

var itemsList []TodoItem

func ListItems() []TodoItem {
	item1 := TodoItem{Id: 1, Name: "Item1", Completed: false}
	item2 := TodoItem{Id: 2, Name: "Item2", Completed: false}
	item3 := TodoItem{Id: 3, Name: "Item3", Completed: false}

	itemsList = append(itemsList, item1)
	itemsList = append(itemsList, item2)
	itemsList = append(itemsList, item3)

	return itemsList

}
