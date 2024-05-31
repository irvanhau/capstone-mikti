package events

import (
	"time"

	"github.com/labstack/echo/v4"
)

type AllEvent struct {
	ID            uint      `json:"id"`
	CategoryFK    int       `json:"category_id"`
	Title         string    `json:"event_title"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	StartingPrice int       `json:"starting_price"`
}

type Event struct {
	ID                   uint      `json:"id"`
	CategoryFK           int       `json:"category_id"`
	Title                string    `json:"event_title"`
	StartDate            string    `json:"start_date"`
	EndDate              string    `json:"end_date"`
	ParseStartDate       time.Time `json:"parse_start_date"`
	ParseEndDate         time.Time `json:"parse_end_date"`
	City                 string    `json:"city"`
	StartingPrice        int       `json:"starting_price"`
	Description          string    `json:"description"`
	Highlight            string    `json:"highlight"`
	ImportantInformation string    `json:"important_information"`
	Address              string    `json:"address"`
	Image                string    `json:"image_url"`
}

type UpdateEvent struct {
	ID                   uint      `json:"id"`
	CategoryFK           int       `json:"category_id"`
	Title                string    `json:"event_title"`
	StartDate            string    `json:"start_date"`
	EndDate              string    `json:"end_date"`
	ParseStartDate       time.Time `json:"parse_start_date"`
	ParseEndDate         time.Time `json:"parse_end_date"`
	City                 string    `json:"city"`
	StartingPrice        int       `json:"starting_price"`
	Description          string    `json:"description"`
	Highlight            string    `json:"highlight"`
	ImportantInformation string    `json:"important_information"`
	Address              string    `json:"address"`
	Image                string    `json:"image_url"`
	PublicID             string    `json:"public_id"`
}

type EventHandlerInterface interface {
	CreateEvent() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	GetDetail() echo.HandlerFunc
	UpdateEvent() echo.HandlerFunc
	DeleteEvent() echo.HandlerFunc
}

type EventServiceInterface interface {
	CreateEvent(newData Event) (*Event, error)
	GetAll() ([]AllEvent, error)
	GetDetail(id int) ([]Event, error)
	UpdateEvent(id int, newData UpdateEvent) (*UpdateEvent, error)
	DeleteEvent(id int) (bool, error)
}

type EventDataInterface interface {
	CreateEvent(newData Event) (*Event, error)
	GetByTitle(username string) ([]Event, error)
	GetAll() ([]AllEvent, error)
	GetDetail(id int) ([]Event, error)
	UpdateEvent(id int, newData UpdateEvent) (*UpdateEvent, error)
	DeleteEvent(id int) (bool, error)
}
