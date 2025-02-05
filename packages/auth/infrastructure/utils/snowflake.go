package utils

import (
	"errors"
	"sync"
	"time"
)

const (
	ebikeEpoch = 1735689600000 // January 1, 2025 in milliseconds
	bitmask42  = (1 << 42) - 1 // Mask for 42 bits
)

// SnowflakeGenerator generates unique IDs based on the snowflake algorithm.
type SnowflakeGenerator struct {
	mu            sync.Mutex
	serviceID     uint64
	workerID      uint64
	sequence      uint64
	lastTimestamp uint64
	epoch         uint64
}

// NewSnowflakeGenerator creates a new SnowflakeGenerator with the given service and worker IDs.
func NewSnowflakeGenerator(serviceID, workerID int) (*SnowflakeGenerator, error) {
	if serviceID < 0 || serviceID > 15 {
		return nil, errors.New("serviceID must be between 0 and 15")
	}
	if workerID < 0 || workerID > 15 {
		return nil, errors.New("workerID must be between 0 and 15")
	}

	return &SnowflakeGenerator{
		serviceID:     uint64(serviceID),
		workerID:      uint64(workerID),
		epoch:         ebikeEpoch,
		lastTimestamp: 0,
		sequence:      0,
	}, nil
}

// Generate produces a new snowflake ID.
func (s *SnowflakeGenerator) Generate() (uint64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	currentTime := uint64(time.Now().UnixMilli())
	if currentTime < s.epoch {
		return 0, errors.New("current time is before the epoch")
	}

	if currentTime < s.lastTimestamp {
		return 0, errors.New("clock moved backwards")
	}

	if currentTime == s.lastTimestamp {
		s.sequence = (s.sequence + 1) & 0xF // Increment and mask to 4 bits
		if s.sequence == 0 {
			// Wait until next millisecond
			for currentTime <= s.lastTimestamp {
				time.Sleep(time.Millisecond / 10) // Prevent busy loop
				currentTime = uint64(time.Now().UnixMilli())
			}
		}
	} else {
		s.sequence = 0
	}

	s.lastTimestamp = currentTime

	// Compute the snowflake components
	timePart := (currentTime - s.epoch) & bitmask42
	timePart <<= 12 // Shift to occupy highest 42 bits

	servicePart := (s.serviceID & 0xF) << 8
	workerPart := (s.workerID & 0xF) << 4
	seqPart := s.sequence & 0xF

	snowflake := timePart | servicePart | workerPart | seqPart
	return snowflake, nil
}
