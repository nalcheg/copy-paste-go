package slices

type Element struct {
	value string
}

type Elements []*Element

func (es Elements) Delete(key int) ([]*Element, error) {
	if len(es) <= 0 {
		return nil, nil
	}

	if key < len(es)-1 {
		copy(es[key:], es[key+1:])
	}
	es[len(es)-1] = nil

	return es[:len(es)-1], nil
}

func (es Elements) Insert(key int, order *Element) []*Element {
	if key > len(es) || key < 0 {
		key = len(es)
	}

	es = append(es, &Element{})
	copy(es[key+1:], es[key:])
	es[key] = order

	return es
}
