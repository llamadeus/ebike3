package persistence

import (
	"context"
	"github.com/llamadeus/ebike3/packages/rentals/domain/model"
	"github.com/llamadeus/ebike3/packages/rentals/domain/port/out"
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

func (r *VehicleViewRepository) Get(id uint64) (*model.VehicleView, error) {
	var vehicle model.VehicleView

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&vehicle)
	if err != nil {
		return nil, err
	}

	return &vehicle, nil
}

func (r *VehicleViewRepository) Create(id uint64, type_ model.VehicleType) (*model.VehicleView, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, bson.M{
		"_id":  id,
		"type": type_,
	})
	if err != nil {
		return nil, err
	}

	return r.Get(id)
}

func (r *VehicleViewRepository) Delete(id uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})

	return err
}
