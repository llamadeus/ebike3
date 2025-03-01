package persistence

import (
	"context"
	"github.com/llamadeus/ebike3/packages/stations/domain/model"
	"github.com/llamadeus/ebike3/packages/stations/domain/port/out"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type StationViewRepository struct {
	collection *mongo.Collection
}

var _ out.StationViewRepository = (*StationViewRepository)(nil)

func NewStationViewRepository(collection *mongo.Collection) *StationViewRepository {
	return &StationViewRepository{collection: collection}
}

func (r *StationViewRepository) GetStations() ([]*model.StationView, error) {
	var stations []*model.StationView

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var station *model.StationView
		if err := cursor.Decode(&station); err != nil {
			return nil, err
		}
		stations = append(stations, station)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return stations, nil
}

func (r *StationViewRepository) GetStationByID(id uint64) (*model.StationView, error) {
	var station model.StationView
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&station)
	if err != nil {
		return nil, err
	}

	return &station, nil
}

func (r *StationViewRepository) CreateStation(id uint64, name string, positionX float64, positionY float64) (*model.StationView, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	station := &model.StationView{
		ID:        id,
		Name:      name,
		PositionX: positionX,
		PositionY: positionY,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := r.collection.InsertOne(ctx, station)
	if err != nil {
		return nil, err
	}

	return station, nil
}

func (r *StationViewRepository) UpdateStation(id uint64, name string, positionX float64, positionY float64) (*model.StationView, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	station := &model.StationView{
		Name:      name,
		PositionX: positionX,
		PositionY: positionY,
		UpdatedAt: time.Now(),
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": station}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return r.GetStationByID(id)
}

func (r *StationViewRepository) DeleteStation(id uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
