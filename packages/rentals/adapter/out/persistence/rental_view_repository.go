package persistence

import (
	"context"
	"errors"
	"github.com/guregu/null/v5"
	"github.com/llamadeus/ebike3/packages/rentals/domain/model"
	"github.com/llamadeus/ebike3/packages/rentals/domain/port/out"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type RentalViewRepository struct {
	collection *mongo.Collection
}

var _ out.RentalViewRepository = (*RentalViewRepository)(nil)

func NewRentalViewRepository(collection *mongo.Collection) *RentalViewRepository {
	return &RentalViewRepository{collection: collection}
}

func (r *RentalViewRepository) Get(id uint64) (*model.RentalView, error) {
	var rental model.RentalView

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&rental)
	if err != nil {
		return nil, err
	}

	return &rental, nil
}

func (r *RentalViewRepository) GetActiveRentalByCustomerID(customerID uint64) (*model.RentalView, error) {
	var rental model.RentalView

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := r.collection.FindOne(ctx, bson.M{"customerId": customerID, "end": bson.M{"$exists": false}}).Decode(&rental)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		return nil, err
	}

	return &rental, nil
}

func (r *RentalViewRepository) GetPastRentalsByCustomerID(customerID uint64) ([]*model.RentalView, error) {
	var rentals []*model.RentalView

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{"customerId": customerID, "end": bson.M{"$exists": true}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var rental *model.RentalView
		if err := cursor.Decode(&rental); err != nil {
			return nil, err
		}
		rentals = append(rentals, rental)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return rentals, nil
}

func (r *RentalViewRepository) Create(id uint64, customerID uint64, vehicleID uint64, vehicleType model.VehicleType, start time.Time) (*model.RentalView, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rental := &model.RentalView{
		ID:          id,
		CustomerID:  customerID,
		VehicleID:   vehicleID,
		VehicleType: vehicleType,
		Start:       start,
		End:         null.TimeFromPtr(nil),
		Cost:        0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_, err := r.collection.InsertOne(ctx, rental)
	if err != nil {
		return nil, err
	}

	return rental, nil
}

func (r *RentalViewRepository) Update(id uint64, end time.Time) (*model.RentalView, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	rental := &model.RentalView{
		End:       null.TimeFrom(end),
		UpdatedAt: time.Now(),
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": rental}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return r.Get(id)
}

func (r *RentalViewRepository) AddExpense(rentalID uint64, amount int32) (*model.RentalView, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": rentalID}, bson.M{
		"$inc": bson.M{"cost": amount},
	})
	if err != nil {
		return nil, err
	}

	return r.Get(rentalID)
}
