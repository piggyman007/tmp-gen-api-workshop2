package db

import "workshop2/model"

// MockDB implements TransferDBInterface for testing

type MockDB struct {
	Users     map[interface{}]model.User
	Transfers []model.Transfer
}

func NewMockDB() *MockDB {
	return &MockDB{
		Users:     make(map[interface{}]model.User),
		Transfers: []model.Transfer{},
	}
}

func (m *MockDB) FindUserByCode(code interface{}) (model.User, error) {
	user, ok := m.Users[code]
	if !ok {
		return model.User{}, ErrNotFound
	}
	return user, nil
}

func (m *MockDB) FindUserByID(id int) (model.User, error) {
	for _, user := range m.Users {
		if user.ID == id {
			return user, nil
		}
	}
	return model.User{}, ErrNotFound
}

func (m *MockDB) CreateTransfer(transfer *model.Transfer) error {
	m.Transfers = append(m.Transfers, *transfer)
	return nil
}

func (m *MockDB) FindTransfersByUser(userID int, limit int) []model.Transfer {
	var result []model.Transfer
	for _, t := range m.Transfers {
		if t.SenderID == userID || t.ReceiverID == userID {
			result = append(result, t)
		}
	}
	if len(result) > limit {
		return result[:limit]
	}
	return result
}

var ErrNotFound = &NotFoundError{}

// NotFoundError is a custom error for not found

type NotFoundError struct{}

func (e *NotFoundError) Error() string { return "not found" }
