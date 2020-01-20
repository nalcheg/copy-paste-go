package slices

type Element string
type Elements []Element

func (es Elements) Delete(key int) (Elements, error) {
	if len(es) <= 0 {
		return nil, nil
	}

	if key < len(es)-1 {
		copy(es[key:], es[key+1:])
	}
	es[len(es)-1] = ""

	return es[:len(es)-1], nil
}

func (es Elements) Insert(key int, str Element) Elements {
	if key > len(es) || key < 0 {
		key = len(es)
	}

	var emptyString Element = ""
	es = append(es, emptyString)
	copy(es[key+1:], es[key:])
	es[key] = str

	return es
}
