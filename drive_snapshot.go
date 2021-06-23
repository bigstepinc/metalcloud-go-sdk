package metalcloud

import (
	"fmt"
)

//Snapshot A snapshot of a drive created at a specific time.
type Snapshot struct {
	DriveSnapshotID               int    `json:"drive_snapshot_id,omitempty"`
	DriveSnapshotLabel            string `json:"drive_snapshot_label,omitempty"`
	DriveID                       int    `json:"drive_id,omitempty"`
	DriveSnapshotCreatedTimestamp string `json:"drive_snapshot_created_timestamp,omitempty"`
}

//DriveSnapshotCreate creates a drive snapshot
func (c *Client) DriveSnapshotCreate(driveID int) (*Snapshot, error) {
	var createdObject Snapshot

	err := c.rpcClient.CallFor(
		&createdObject,
		"drive_snapshot_create",
		driveID,
	)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//DriveSnapshotDelete creates a drive snapshot
func (c *Client) DriveSnapshotDelete(driveSnapshotID int) error {
	if err := checkID(driveSnapshotID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call("drive_snapshot_delete", driveSnapshotID)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

//DriveSnapshotRollback rolls a Drive back to a specified DriveSnapshot. The specified snapshot is not destroyed and can be reused.
func (c *Client) DriveSnapshotRollback(driveSnapshotID int) error {
	if err := checkID(driveSnapshotID); err != nil {
		return err
	}

	resp, err := c.rpcClient.Call("drive_snapshot_rollback", driveSnapshotID)

	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}

//DriveSnapshotGet gets a drive snapshot
func (c *Client) DriveSnapshotGet(driveSnapshotID int) (*Snapshot, error) {
	var createdObject Snapshot

	err := c.rpcClient.CallFor(
		&createdObject,
		"drive_snapshot_get",
		driveSnapshotID,
	)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}

//DriveSnapshots retrieves a list of all the snapshot objects
func (c *Client) DriveSnapshots(driveID int) (*map[string]Snapshot, error) {
	var err error

	resp, err := c.rpcClient.Call(
		"drive_snapshots",
		driveID,
	)

	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}

	if err != nil {
		return nil, err
	}

	_, ok := resp.Result.([]interface{})
	if ok {
		var m = map[string]Snapshot{}
		return &m, nil
	}

	var createdObject map[string]Snapshot

	err = resp.GetObject(&createdObject)

	if err != nil {
		return nil, err
	}

	return &createdObject, nil
}
