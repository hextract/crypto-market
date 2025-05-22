package binance

import "fmt"

type Fill struct {
	Price           string
	Qty             string
	Commission      string
	CommissionAsset string
}

func ParseFills(response map[string]interface{}) []Fill {
	fillsRaw, ok := response["fills"].([]interface{})
	if !ok {
		fmt.Println("Не удалось прочитать массив fills")
		return nil
	}

	var fills []Fill

	for _, f := range fillsRaw {
		fillMap, ok := f.(map[string]interface{})
		if !ok {
			continue
		}

		price := fillMap["price"].(string)
		qty := fillMap["qty"].(string)
		commission := fillMap["commission"].(string)
		commissionAsset := fillMap["commissionAsset"].(string)

		fills = append(fills, Fill{
			Price:           price,
			Qty:             qty,
			Commission:      commission,
			CommissionAsset: commissionAsset,
		})
	}

	//for _, f := range fills {
	//	fmt.Printf("price: %s qty: %s commission: %s commissionAsset: %s\n", f.Price, f.Qty, f.Commission, f.CommissionAsset)
	//}
	return fills
}
