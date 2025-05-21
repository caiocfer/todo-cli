package todo

var ItemsList []*TodoItem

func ListItems() {
	item1 := &TodoItem{Id: 1, Name: "Item1", Completed: false}
	item2 := &TodoItem{Id: 2, Name: "Item2", Completed: false}
	item3 := &TodoItem{Id: 3, Name: "Item3", Completed: false}
	ItemsList = append(ItemsList, item1, item2, item3)

}
