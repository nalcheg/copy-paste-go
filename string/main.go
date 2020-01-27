package main

import "log"

func main() {
	str := "-aaa-"
	withoutFirst := str[1:]
	withoutLast := str[:len(str)-1]

	log.Print(withoutFirst)
	log.Print(withoutLast)
}

type WhereCondition string

func (wc WhereCondition) AddCondition(condition string) WhereCondition {
	if len(wc) == 0 {
		wc += ` WHERE `
	} else {
		wc += ` AND `
	}
	wc += WhereCondition(condition)

	return wc
}

func (wc *WhereCondition) AddConditionPointer(condition string) {
	if len(*wc) == 0 {
		*wc += ` WHERE `
	} else {
		*wc += ` AND `
	}
	*wc += WhereCondition(condition)
}

func (wc WhereCondition) String() string {
	return string(wc)
}
