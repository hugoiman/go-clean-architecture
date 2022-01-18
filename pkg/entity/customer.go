package entity

import "encoding/json"

type Customer struct {
	id       string
	username string
	name     string
}

func (c *Customer) SetId(id string) {
	c.id = id
}

func (c *Customer) GetId() string {
	return c.id
}

func (c *Customer) SetUsername(username string) {
	c.username = username
}

func (c *Customer) GetUsername() string {
	return c.username
}

func (c *Customer) SetName(name string) {
	c.name = name
}
func (c *Customer) GetName() string {
	return c.name
}

func (c *Customer) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id       string `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
	}{
		Id:       c.id,
		Name:     c.name,
		Username: c.username,
	})
}

func (c *Customer) UnmarshalJSON(data []byte) error {
	alias := struct {
		Id       string `json:"id"`
		Username string `json:"username"`
		Name     string `json:"name"`
	}{}

	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	c.id = alias.Id
	c.username = alias.Username
	c.name = alias.Name

	return nil
}
