package newUrl

import (
	"fmt"
	"os"
)

func NewURL(key string) string {
	fmt.Print("Введите ключевое слово: ")
	_, err := fmt.Fscan(os.Stdin, &key)
	if err != nil {
		return "Ошибка"
	}

	url := fmt.Sprintf("https://www.fabrikant.ru/trades/procedure/search/?query=%s&type=0&org_type=org&currency=0&date_type=date_publication&ensure=all&paid_participation=0&okpd2_embedded=1&okdp_embedded=1&count_on_page=10&order_direction=1&type_hash=1561441166", key)

	return url
}
