package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"

	"github.com/go-swagger/go-swagger/examples/generated/models"
)

/*GetPetByIDOK successful operation

swagger:response getPetByIdOK
*/
type GetPetByIDOK struct {

	// In: body
	Payload *models.Pet `json:"body,omitempty"`
}

// NewGetPetByIDOK creates GetPetByIDOK with default headers values
func NewGetPetByIDOK() GetPetByIDOK {
	return GetPetByIDOK{}
}

// WriteResponse to the client
func (o *GetPetByIDOK) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetPetByIDBadRequest Invalid ID supplied

swagger:response getPetByIdBadRequest
*/
type GetPetByIDBadRequest struct {
}

// NewGetPetByIDBadRequest creates GetPetByIDBadRequest with default headers values
func NewGetPetByIDBadRequest() GetPetByIDBadRequest {
	return GetPetByIDBadRequest{}
}

// WriteResponse to the client
func (o *GetPetByIDBadRequest) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(400)
}

/*GetPetByIDNotFound Pet not found

swagger:response getPetByIdNotFound
*/
type GetPetByIDNotFound struct {
}

// NewGetPetByIDNotFound creates GetPetByIDNotFound with default headers values
func NewGetPetByIDNotFound() GetPetByIDNotFound {
	return GetPetByIDNotFound{}
}

// WriteResponse to the client
func (o *GetPetByIDNotFound) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(404)
}
