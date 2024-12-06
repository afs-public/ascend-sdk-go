// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"time"

	"github.com/afs-public/ascend-sdk-go/internal/utils"
)

// File - The details of file containing the snapshot data. This contains the download uri and uri expiry time.
type File struct {
	// The signed download uri for the file. This allows the client to download the file.
	DownloadURI *string `json:"download_uri,omitempty"`
	// The timestamp at which the download uri expires in UTC. This is set to 1 hour after the request time.
	URIExpiryTime *time.Time `json:"uri_expiry_time,omitempty"`
}

func (f File) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(f, "", false)
}

func (f *File) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &f, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *File) GetDownloadURI() *string {
	if o == nil {
		return nil
	}
	return o.DownloadURI
}

func (o *File) GetURIExpiryTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.URIExpiryTime
}

// SnapshotProcessDate - The process date of the snapshot. This date corresponds to the underlying data within the snapshot.
type SnapshotProcessDate struct {
	// Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.
	Day *int `json:"day,omitempty"`
	// Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.
	Month *int `json:"month,omitempty"`
	// Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.
	Year *int `json:"year,omitempty"`
}

func (o *SnapshotProcessDate) GetDay() *int {
	if o == nil {
		return nil
	}
	return o.Day
}

func (o *SnapshotProcessDate) GetMonth() *int {
	if o == nil {
		return nil
	}
	return o.Month
}

func (o *SnapshotProcessDate) GetYear() *int {
	if o == nil {
		return nil
	}
	return o.Year
}

// Snapshot - A snapshot of the data at a specific point in time. This contains the snapshot details and the download_uri of the file containing the snapshot data.
type Snapshot struct {
	// The timestamp at which the snapshot was created at in UTC.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The details of file containing the snapshot data. This contains the download uri and uri expiry time.
	File *File `json:"file,omitempty"`
	// The name of the resource.
	Name *string `json:"name,omitempty"`
	// The process date of the snapshot. This date corresponds to the underlying data within the snapshot.
	ProcessDate *SnapshotProcessDate `json:"process_date,omitempty"`
	// The unique identifier of the snapshot file.
	SnapshotID *string `json:"snapshot_id,omitempty"`
	// The type of the snapshot.
	SnapshotType *string `json:"snapshot_type,omitempty"`
	// The version of the snapshot.
	Version *string `json:"version,omitempty"`
}

func (s Snapshot) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *Snapshot) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Snapshot) GetCreateTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreateTime
}

func (o *Snapshot) GetFile() *File {
	if o == nil {
		return nil
	}
	return o.File
}

func (o *Snapshot) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Snapshot) GetProcessDate() *SnapshotProcessDate {
	if o == nil {
		return nil
	}
	return o.ProcessDate
}

func (o *Snapshot) GetSnapshotID() *string {
	if o == nil {
		return nil
	}
	return o.SnapshotID
}

func (o *Snapshot) GetSnapshotType() *string {
	if o == nil {
		return nil
	}
	return o.SnapshotType
}

func (o *Snapshot) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}