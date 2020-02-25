package netbox

import (
	"strconv"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/netbox-community/go-netbox/netbox/client/dcim"
	"github.com/netbox-community/go-netbox/netbox/models"
	"github.com/stretchr/testify/assert"
)

// These tests assume a netbox running locally, with the netbox test fixtures loaded.
// This is achievable with the following netbox management commands, after install:
//    python3 netbox/manage.py loaddata dcim extras ipam
//    python3 netbox/manage.py runserver 0.0.0.0:8000 --insecure

func TestRetrieveDeviceList(t *testing.T) {
	c := NewNetboxWithAPIKey("localhost:8000", "0123456789abcdef0123456789abcdef01234567")

	list, err := c.Dcim.DcimDevicesList(nil, nil)

	assert.Nil(t, err)
	assert.EqualValues(t, 7, *list.Payload.Count)
}

func TestSubdeviceRole(t *testing.T) {
	c := NewNetboxWithAPIKey("localhost:8000", "0123456789abcdef0123456789abcdef01234567")

	role := "parent"
	manufacturerID := int64(1)
	model := "Test model"
	slug := "test-slug"
	newDeviceType := &models.WritableDeviceType{
		SubdeviceRole: role,
		Comments:      "Test device type",
		Manufacturer:  &manufacturerID,
		Model:         &model,
		Slug:          &slug,
	}
	err := newDeviceType.Validate(strfmt.Default)
	assert.Nil(t, err)

	createRequest := dcim.NewDcimDeviceTypesCreateParams().WithData(newDeviceType)
	createResponse, err := c.Dcim.DcimDeviceTypesCreate(createRequest, nil)
	assert.Nil(t, err)

	//newID := strconv.FormatInt(createResponse.Payload.ID)
	newID := int64(createResponse.Payload.ID)
	assert.NotEqual(t, 0, newID)
	IDIn := strconv.FormatInt(newID, 10)

	retrieveResponse, err := c.Dcim.DcimDeviceTypesList(dcim.NewDcimDeviceTypesListParams().WithIDIn(&IDIn), nil)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, *retrieveResponse.Payload.Count)
	assert.EqualValues(t, "Test device type", retrieveResponse.Payload.Results[0].Comments)

	deleteRequest := dcim.NewDcimDeviceTypesDeleteParams().WithID(newID)
	_, err = c.Dcim.DcimDeviceTypesDelete(deleteRequest, nil)
	assert.Nil(t, err)
}
