package entity

import "encoding/json"

type Product struct {
	id         string
	customerId string
	name       string
	price      int
	qty        int
}

func (c *Product) SetId(id string) {
	c.id = id
}

func (c *Product) GetId() string {
	return c.id
}

func (c *Product) SetCustomerId(customerId string) {
	c.customerId = customerId
}

func (c *Product) GetCustomerId() string {
	return c.customerId
}

func (c *Product) SetName(name string) {
	c.name = name
}
func (c *Product) GetName() string {
	return c.name
}

func (c *Product) SetPrice(price int) {
	c.price = price
}

func (c *Product) GetPrice() int {
	return c.price
}

func (c *Product) SetQty(qty int) {
	c.qty = qty
}

func (c *Product) GetQty() int {
	return c.qty
}

func (c *Product) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id         string `json:"id"`
		CustomerId string `json:"customer_id"`
		Name       string `json:"name"`
		Price      int    `json:"price"`
		Qty        int    `json:"qty"`
	}{
		Id:         c.id,
		CustomerId: c.customerId,
		Name:       c.name,
		Price:      c.price,
		Qty:        c.qty,
	})
}

func (c *Product) UnmarshalJSON(data []byte) error {
	alias := struct {
		Id         string `json:"id"`
		CustomerId string `json:"customer_id"`
		Name       string `json:"name"`
		Price      int    `json:"price"`
		Qty        int    `json:"qty"`
	}{}

	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	c.id = alias.Id
	c.customerId = alias.CustomerId
	c.name = alias.Name
	c.price = alias.Price
	c.qty = alias.Qty

	return nil
}
