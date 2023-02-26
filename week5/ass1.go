package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Item represents an item in the system
type Item struct {
	ID    int
	Name  string
	Price float64
	Rate  float64
}

// User represents a user in the system
type User struct {
	ID       int
	Username string
	Password string
}
// System represents the main system
type System struct {
	Users []*User
	Items []*Item
}

// NewSystem creates a new instance of the System
func NewSystem() *System {
	return &System{}
}

// LoadUsers loads the users from the users.txt file
func (s *System) LoadUsers() error {
	file, err := os.Open("users.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			continue
		}
		id, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}
		user := &User{
			ID:       id,
			Username: parts[1],
			Password: parts[2],
		}
		s.Users = append(s.Users, user)
	}
	return scanner.Err()
}

// SaveUsers saves the users to the users.txt file
func (s *System) SaveUsers() error {
	file, err := os.Create("users.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, user := range s.Users {
		line := fmt.Sprintf("%d,%s,%s\n", user.ID, user.Username, user.Password)
		_, err := writer.WriteString(line)
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}
// LoadItems loads the items from the items.txt file
func (s *System) LoadItems() error {
	file, err := os.Open("items.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 4 {
			continue
		}
		id, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}
		price, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			continue
		}
		rate, err := strconv.ParseFloat(parts[3], 64)
		if err != nil {
			continue
		}
		item := &Item{
			ID:    id,
			Name:  parts[1],
			Price: price,
			Rate:  rate,
		}
		s.Items = append(s.Items, item)
	}
	return scanner.Err()
}

// SaveItems saves the items to the items.txt file
func (s *System) SaveItems() error {
	file, err := os.Create("items.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, item := range s.Items {
		line := fmt.Sprintf("%d,%s,%.2f,%.2f\n", item.ID, item.Name, item.Price, item.Rate)
		_, err := writer.WriteString(line)
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

// SearchItems searches for items based on their name
func (s *System) SearchItems(name string) []*Item {
	result := []*Item{}
	for _, item := range s.Items {
		if strings.Contains(item.Name, name) {
			result = append(result, item)
		}
	}
	return result
}
// FilterItems filters items based on their price and rating
func (s *System) FilterItems(price float64, rate float64) []*Item {
	result := []*Item{}
	for _, item := range s.Items {
		if item.Price <= price && item.Rate >= rate {
			result = append(result, item)
		}
	}
	return result
}
// GiveRating gives a rating to an item and saves it
func (s *System) GiveRating(id int, rate float64) error {
	for _, item := range s.Items {
		if item.ID == id {
			item.Rate =  rate
			break
		}
	}
	return s.SaveItems()
}