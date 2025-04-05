package persistence

import (
	"context"
	"github.com/guregu/null/v5"
	"github.com/llamadeus/ebike3/packages/customers/domain/model"
	"github.com/llamadeus/ebike3/packages/customers/domain/port/out"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type CustomerViewRepository struct {
	collection *mongo.Collection
}

var _ out.CustomerViewRepository = (*CustomerViewRepository)(nil)

func NewCustomerViewRepository(collection *mongo.Collection) *CustomerViewRepository {
	return &CustomerViewRepository{collection: collection}
}

func (r *CustomerViewRepository) GetCustomers() ([]*model.CustomerView, error) {
	var customers []*model.CustomerView

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var customer *model.CustomerView
		if err := cursor.Decode(&customer); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}

func (r *CustomerViewRepository) GetCustomerByID(id uint64) (*model.CustomerView, error) {
	var customer model.CustomerView
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&customer)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (r *CustomerViewRepository) CreateCustomer(id uint64, name string, positionX float64, positionY float64, creditBalance int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	customer := &model.CustomerView{
		ID:            id,
		Name:          name,
		PositionX:     positionX,
		PositionY:     positionY,
		CreditBalance: creditBalance,
		ActiveRental:  nil,
		LastLogin:     null.TimeFromPtr(nil),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	_, err := r.collection.InsertOne(ctx, customer)

	return err
}

func (r *CustomerViewRepository) UpdateCustomerViewPosition(id uint64, positionX float64, positionY float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	customer := &model.CustomerView{
		PositionX: positionX,
		PositionY: positionY,
		UpdatedAt: time.Now(),
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": customer}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *CustomerViewRepository) UpdateCustomerViewCreditBalance(id uint64, amount int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{
		"$inc": bson.M{"creditBalance": amount},
		"$set": bson.M{"updatedAt": time.Now()},
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *CustomerViewRepository) UpdateCustomerViewLastLogin(id uint64, lastLogin time.Time) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	customer := &model.CustomerView{
		LastLogin: null.TimeFromPtr(&lastLogin),
		UpdatedAt: time.Now(),
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": customer}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *CustomerViewRepository) UpdateCustomerViewActiveRental(customerID uint64, rentalID uint64, vehicleID uint64, vehicleType string, start time.Time, cost int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	customer := &model.CustomerView{
		ActiveRental: &model.RentalView{
			ID:          rentalID,
			CustomerID:  customerID,
			VehicleID:   vehicleID,
			VehicleType: vehicleType,
			Start:       start,
			Cost:        cost,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		UpdatedAt: time.Now(),
	}
	filter := bson.M{"_id": customerID}
	update := bson.M{"$set": customer}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *CustomerViewRepository) ResetCustomerViewActiveRental(id uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{
		"$unset": bson.M{"activeRental": 1},
	})

	return err
}
