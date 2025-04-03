package service

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/llamadeus/ebike3/packages/accounting/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/accounting/domain/events"
	"github.com/llamadeus/ebike3/packages/accounting/domain/model"
	"github.com/llamadeus/ebike3/packages/accounting/domain/port/in"
	"github.com/llamadeus/ebike3/packages/accounting/domain/port/out"
	"github.com/llamadeus/ebike3/packages/accounting/infrastructure/database"
	"github.com/llamadeus/ebike3/packages/accounting/infrastructure/micro"
	"time"
)

type AccountingService struct {
	kafka                        micro.Kafka
	transactor                   database.Transactor
	paymentRepository            out.PaymentRepository
	expenseRepository            out.ExpenseRepository
	preliminaryExpenseRepository out.PreliminaryExpenseRepository
}

var _ in.AccountingService = (*AccountingService)(nil)

func NewAccountingService(kafka micro.Kafka, transactor database.Transactor, paymentRepository out.PaymentRepository, expenseRepository out.ExpenseRepository, preliminaryExpenseRepository out.PreliminaryExpenseRepository) *AccountingService {
	return &AccountingService{
		kafka:                        kafka,
		transactor:                   transactor,
		paymentRepository:            paymentRepository,
		expenseRepository:            expenseRepository,
		preliminaryExpenseRepository: preliminaryExpenseRepository,
	}
}

func (s *AccountingService) GetAllPayments() ([]*model.Payment, error) {
	return s.paymentRepository.GetAll()
}

func (s *AccountingService) GetPaymentsForCustomer(customerID uint64) ([]*model.Payment, error) {
	return s.paymentRepository.GetByCustomerID(customerID)
}

func (s *AccountingService) CreatePayment(customerID uint64, amount int32) (*model.Payment, error) {
	payment, err := s.paymentRepository.Create(customerID, amount)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to create payment: %v", err))
	}

	event := micro.NewEvent(events.AccountingPaymentCreatedEventType, events.PaymentCreatedEvent{
		ID:         dto.IDToDTO(payment.ID),
		CustomerID: dto.IDToDTO(payment.CustomerID),
		Amount:     payment.Amount,
		Status:     dto.StatusToDTO(payment.Status),
	})
	err = s.kafka.Producer().Send(events.AccountingTopic, event.Payload.ID, event)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to send kafka event: %v", err))
	}

	return payment, nil
}

func (s *AccountingService) UpdatePayment(id uint64, status model.PaymentStatus) (*model.Payment, error) {
	payment, err := s.paymentRepository.Get(id)
	if err != nil {
		return nil, micro.NewNotFoundError(fmt.Sprintf("payment with id %d not found", id))
	}

	if payment.Status == model.PaymentStatusConfirmed {
		return nil, micro.NewBadRequestError("payment already confirmed")
	}

	updated, err := s.paymentRepository.Update(payment.ID, status)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to update payment: %v", err))
	}

	event := micro.NewEvent(events.AccountingPaymentUpdatedEventType, events.PaymentUpdatedEvent{
		ID:         dto.IDToDTO(updated.ID),
		CustomerID: dto.IDToDTO(updated.CustomerID),
		Amount:     updated.Amount,
		Status:     dto.StatusToDTO(updated.Status),
	})
	err = s.kafka.Producer().Send(events.AccountingTopic, event.Payload.ID, event)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to send kafka event: %v", err))
	}

	return updated, nil
}

func (s *AccountingService) DeleteCustomerPayment(id uint64, customerID uint64) (*model.Payment, error) {
	payment, err := s.paymentRepository.Get(id)
	if err != nil {
		return nil, micro.NewNotFoundError(fmt.Sprintf("payment with id %d not found", id))
	}

	if payment.CustomerID != customerID {
		return nil, micro.NewUnauthorizedError(fmt.Sprintf("customer %d not authorized to delete payment %d", customerID, id))
	}

	err = s.paymentRepository.Delete(id)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to delete payment: %v", err))
	}

	event := micro.NewEvent(events.AccountingPaymentDeletedEventType, events.PaymentDeletedEvent{
		ID:         dto.IDToDTO(payment.ID),
		CustomerID: dto.IDToDTO(payment.CustomerID),
		Amount:     payment.Amount,
		Status:     dto.StatusToDTO(payment.Status),
	})
	err = s.kafka.Producer().Send(events.AccountingTopic, event.Payload.ID, event)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to send kafka event: %v", err))
	}

	return payment, nil
}

func (s *AccountingService) GetExpensesForCustomer(customerID uint64) ([]*model.Expense, error) {
	return s.expenseRepository.GetByCustomerID(customerID)
}

func (s *AccountingService) CreateExpense(customerID uint64, rentalID uint64, amount int32) (*model.Expense, error) {
	expense, err := s.expenseRepository.Create(customerID, rentalID, amount)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to create expense: %v", err))
	}

	event := micro.NewEvent(events.AccountingExpenseCreatedEventType, events.ExpenseCreatedEvent{
		ID:         dto.IDToDTO(expense.ID),
		CustomerID: dto.IDToDTO(expense.CustomerID),
		RentalID:   dto.IDToDTO(expense.RentalID),
		Amount:     expense.Amount,
	})
	err = s.kafka.Producer().Send(events.AccountingTopic, event.Payload.ID, event)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to send kafka event: %v", err))
	}

	return expense, nil
}

func (s *AccountingService) CreatePreliminaryExpense(inquiryID uint64, customerID uint64, amount int32) (*model.PreliminaryExpense, error) {
	expiresAt := time.Now().Add(time.Second * 10)
	preliminaryExpense, err := s.preliminaryExpenseRepository.Create(inquiryID, customerID, amount, expiresAt)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to create preliminary expense: %v", err))
	}

	event := micro.NewEvent(events.AccountingPreliminaryExpenseCreatedEventType, events.PreliminaryExpenseCreatedEvent{
		ID:         dto.IDToDTO(preliminaryExpense.ID),
		InquiryID:  dto.IDToDTO(preliminaryExpense.InquiryID),
		CustomerID: dto.IDToDTO(preliminaryExpense.CustomerID),
		Amount:     preliminaryExpense.Amount,
	})
	err = s.kafka.Producer().Send(events.AccountingTopic, event.Payload.ID, event)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to send kafka event: %v", err))
	}

	return preliminaryExpense, nil
}

func (s *AccountingService) FinalizePreliminaryExpense(id uint64, rentalID uint64) (*model.Expense, error) {
	preliminaryExpense, err := s.preliminaryExpenseRepository.Get(id)
	if err != nil {
		return nil, micro.NewNotFoundError(fmt.Sprintf("preliminary expense with id %d not found", id))
	}

	if time.Now().After(preliminaryExpense.ExpiresAt) {
		return nil, micro.NewBadRequestError("preliminary expense expired")
	}

	var expense *model.Expense

	err = database.RunInTx(s.transactor, func(tx *sqlx.Tx) (err error) {
		expense, err = s.expenseRepository.CreateWithTx(tx, preliminaryExpense.CustomerID, rentalID, preliminaryExpense.Amount)
		if err != nil {
			return err
		}

		return s.preliminaryExpenseRepository.DeleteWithTx(tx, preliminaryExpense.ID)
	})
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to finalize preliminary expense: %v", err))
	}

	return expense, nil
}

func (s *AccountingService) GetCreditBalanceForCustomer(customerID uint64) (int32, error) {
	payments, err := s.paymentRepository.GetByCustomerID(customerID)
	if err != nil {
		return 0, micro.NewInternalServerError(fmt.Sprintf("failed to get payments: %v", err))
	}

	expenses, err := s.expenseRepository.GetByCustomerID(customerID)
	if err != nil {
		return 0, micro.NewInternalServerError(fmt.Sprintf("failed to get expenses: %v", err))
	}

	preliminaryExpenses, err := s.preliminaryExpenseRepository.GetByCustomerID(customerID)
	if err != nil {
		return 0, micro.NewInternalServerError(fmt.Sprintf("failed to get preliminary expenses: %v", err))
	}

	creditBalance := int32(0)
	for _, payment := range payments {
		if payment.Status != model.PaymentStatusConfirmed {
			continue
		}

		creditBalance += payment.Amount
	}
	for _, expense := range expenses {
		creditBalance -= expense.Amount
	}
	for _, preliminaryExpense := range preliminaryExpenses {
		if time.Now().After(preliminaryExpense.ExpiresAt) {
			continue
		}

		creditBalance -= preliminaryExpense.Amount
	}

	return creditBalance, nil
}
