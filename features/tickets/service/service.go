package service

import (
	"capstone-mikti/features/tickets"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

// Service

type TicketService struct {
	data tickets.TicketDataInterface
}

func New(d tickets.TicketDataInterface) *TicketService {
	return &TicketService{
		data: d,
	}
}

// Create
func (ts *TicketService) Create(new_data tickets.Ticket) (*tickets.Ticket, error) {
	// Call Data CheckEvent
	checkEvent, err := ts.data.CheckEvent(new_data.EventID)

	if err != nil {
		logrus.Error("Service : Error CheckEvent : ", err.Error())
		return nil, errors.New("ERROR CheckEvent")
	}

	if len(checkEvent) == 0 {
		logrus.Warn("Service : Warning CheckEvent")
		return nil, errors.New("WARNING Event Does Not Exists")
	}

	// Parse Ticket Date
	layout := "2006-01-02"

	parse_ticketDate, err := time.Parse(layout, new_data.TicketDate)

	if err != nil {
		logrus.Error("Service : Error Parse Ticket Date : ", err.Error())
		return nil, errors.New("ERROR Parse Ticket Date")
	}

	new_data.ParseTicketDate = parse_ticketDate

	// Call Data Create
	create, err := ts.data.Create(new_data)

	if err != nil {
		logrus.Error("Service : Create Error : ", err.Error())
		return nil, errors.New("ERROR Create")
	}

	return create, nil
}

// GetAll
func (ts *TicketService) GetAll() ([]tickets.TicketInfo, error) {
	// Call Data GetAll
	getAll, err := ts.data.GetAll()

	if err != nil {
		logrus.Error("Service : GetAll Error : ", err.Error())
		return nil, errors.New("ERROR GetAll")
	}

	return getAll, nil
}

// GetByID
func (ts *TicketService) GetByID(id int) ([]tickets.TicketInfo, error) {
	// Call Data GetByID
	getById, err := ts.data.GetByID(id)

	if err != nil {
		logrus.Error("Service : GetAll Error : ", err.Error())
		return nil, errors.New("ERROR GetAll")
	}

	return getById, nil
}

// Update
func (ts *TicketService) Update(id int, new_data tickets.Ticket) (bool, error) {
	// Call Data CheckEvent
	checkEvent, err := ts.data.CheckEvent(new_data.EventID)

	if err != nil {
		logrus.Error("Service : Error CheckEvent : ", err.Error())
		return false, errors.New("ERROR CheckEvent")
	}

	if len(checkEvent) == 0 {
		logrus.Warn("Service : Warning CheckEvent")
		return false, errors.New("WARNING Event Does Not Exists")
	}

	// Parse Ticket Date
	layout := "2006-01-02"

	parse_ticketDate, err := time.Parse(layout, new_data.TicketDate)

	if err != nil {
		logrus.Error("Service : Error Parse Ticket Date : ", err.Error())
		return false, errors.New("ERROR Parse Ticket Date")
	}

	new_data.ParseTicketDate = parse_ticketDate

	// Call Data Update
	update, err := ts.data.Update(id, new_data)

	if err != nil {
		logrus.Error("Service : Update Error : ", err.Error())
		return false, errors.New("ERROR Update")
	}

	return update, nil
}

// Delete
func (ts *TicketService) Delete(id int) (bool, error) {
	// Call Data Delete
	delete, err := ts.data.Delete(id)

	if err != nil {
		logrus.Error("Service : Delete Error : ", err.Error())
		return false, errors.New("ERROR Delete")
	}

	return delete, nil
}