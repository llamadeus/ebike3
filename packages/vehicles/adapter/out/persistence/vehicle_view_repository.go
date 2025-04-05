package persistence

import (
	"context"
	"github.com/llamadeus/ebike3/packages/vehicles/domain/model"
	"github.com/llamadeus/ebike3/packages/vehicles/domain/port/out"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type VehicleViewRepository struct {
	collection *mongo.Collection
}

var _ out.VehicleViewRepository = (*VehicleViewRepository)(nil)

func NewVehicleViewRepository(collection *mongo.Collection) *VehicleViewRepository {
	return &VehicleViewRepository{collection: collection}
}

func (r *VehicleViewRepository) GetVehicles() ([]*model.VehicleView, error) {
	var vehicles []*model.VehicleView

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var vehicle *model.VehicleView
		if err := cursor.Decode(&vehicle); err != nil {
			return nil, err
		}
		vehicles = append(vehicles, vehicle)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return vehicles, nil
}

func (r *VehicleViewRepository) GetAvailableVehicles() ([]*model.VehicleView, error) {
	var vehicles []*model.VehicleView

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{"activeRental": bson.M{"$exists": false}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var vehicle *model.VehicleView
		if err := cursor.Decode(&vehicle); err != nil {
			return nil, err
		}
		vehicles = append(vehicles, vehicle)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return vehicles, nil
}

func (r *VehicleViewRepository) GetVehicleByID(id uint64) (*model.VehicleView, error) {
	var vehicle model.VehicleView
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&vehicle)
	if err != nil {
		return nil, err
	}

	return &vehicle, nil
}

func (r *VehicleViewRepository) CreateVehicle(id uint64, type_ model.VehicleType, positionX float64, positionY float64, battery float64) (*model.VehicleView, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	vehicle := &model.VehicleView{
		ID:           id,
		Type:         type_,
		PositionX:    positionX,
		PositionY:    positionY,
		Battery:      battery,
		ActiveRental: nil,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	_, err := r.collection.InsertOne(ctx, vehicle)
	if err != nil {
		return nil, err
	}

	return vehicle, nil
}

func (r *VehicleViewRepository) UpdateVehicle(id uint64, positionX float64, positionY float64, battery float64) (*model.VehicleView, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	vehicle := &model.VehicleView{
		PositionX: positionX,
		PositionY: positionY,
		Battery:   battery,
		UpdatedAt: time.Now(),
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": vehicle}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return r.GetVehicleByID(id)
}

func (r *VehicleViewRepository) UpdateActiveRental(rentalID uint64, customerID uint64, vehicleID uint64, vehicleType string, start time.Time, cost int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	vehicle := &model.VehicleView{
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
	filter := bson.M{"_id": vehicleID}
	update := bson.M{"$set": vehicle}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *VehicleViewRepository) ResetActiveRental(id uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{
		"$unset": bson.M{"activeRental": 1},
	})

	return err
}

func (r *VehicleViewRepository) DeleteVehicle(id uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
