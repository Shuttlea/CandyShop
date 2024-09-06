package handlers

import "errors"

func calculateCost(order Order) (int, error) {
	price, err := checkPrice(order.CandyType)
	if err != nil {
		return 0, err
	}
	if order.CandyCount <= 0 {
		err = errors.New("candy count must be positive")
		return 0, err
	}
	return price * order.CandyCount, nil
}

func checkPrice(t string) (int, error) {
	switch t {
	case "CE":
		return 10, nil
	case "AA":
		return 15, nil
	case "NT":
		return 17, nil
	case "DE":
		return 21, nil
	case "YR":
		return 23, nil
	}
	err := errors.New("wrong candy type name")
	return 0, err
}
