package metadata

import "time"

type Metadata struct {
	createdAt     time.Time
	lastUpdatedAt time.Time
}

func CreateNewMetadataWithDetails(createdTime time.Time, lastUpdatedTime time.Time) Metadata {
	return Metadata{
		createdAt:     createdTime,
		lastUpdatedAt: lastUpdatedTime,
	}
}
func CreateNewMetadata() Metadata {
	return Metadata{
		createdAt:     time.Now(),
		lastUpdatedAt: time.Now(),
	}
}

func (m Metadata) GetCreatedAt() time.Time {
	return m.createdAt
}

func (m Metadata) GetLastUpdatedAt() time.Time {
	return m.lastUpdatedAt
}
