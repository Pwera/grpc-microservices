package api

import (
<<<<<<< HEAD
	"context"
=======
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
	"errors"
	"github.com/pwera/grpc-micros-order/internal/application/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockedPayment struct {
	mock.Mock
}

<<<<<<< HEAD
func (p *mockedPayment) Charge(ctx context.Context, order *domain.Order) error {
=======
func (p *mockedPayment) Charge(order *domain.Order) error {
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
	args := p.Called(order)
	return args.Error(0)
}

type mockedDb struct {
	mock.Mock
}

<<<<<<< HEAD
func (d *mockedDb) Save(ctx context.Context, order *domain.Order) error {
=======
func (d *mockedDb) Save(order *domain.Order) error {
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
	args := d.Called(order)
	return args.Error(0)
}

<<<<<<< HEAD
func (d *mockedDb) Get(ctx context.Context, id int32) (domain.Order, error) {
=======
func (d *mockedDb) Get(id int32) (domain.Order, error) {
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
	args := d.Called(id)
	return args.Get(0).(domain.Order), args.Error(1)
}

func TestPlaceOrder(t *testing.T) {
	payment := new(mockedPayment)
	db := new(mockedDb)
<<<<<<< HEAD
	payment.On("Charge", mock.Anything, mock.Anything).Return(nil)
	db.On("Save", mock.Anything, mock.Anything).Return(nil)

	application := NewApplication(db, payment)
	_, err := application.PlaceOrder(context.Background(), domain.Order{
=======
	payment.On("Charge", mock.Anything).Return(nil)
	db.On("Save", mock.Anything).Return(nil)

	application := NewApplication(db, payment)
	_, err := application.PlaceOrder(domain.Order{
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
		CustomerID: 123,
		OrderItems: []domain.OrderItem{
			{
				ProductCode: "camera",
				UnitPrice:   12.3,
				Quantity:    3,
			},
		},
		CreatedAt: 0,
	})
	assert.Nil(t, err)
}

func Test_Should_Return_Error_When_Db_Persistence_Fail(t *testing.T) {
	payment := new(mockedPayment)
	db := new(mockedDb)
<<<<<<< HEAD
	payment.On("Charge", mock.Anything, mock.Anything).Return(nil)
	db.On("Save", mock.Anything, mock.Anything).Return(errors.New("connection error"))

	ctx := context.Background()
	application := NewApplication(db, payment)
	_, err := application.PlaceOrder(ctx, domain.Order{
=======
	payment.On("Charge", mock.Anything).Return(nil)
	db.On("Save", mock.Anything).Return(errors.New("connection error"))

	application := NewApplication(db, payment)
	_, err := application.PlaceOrder(domain.Order{
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
		CustomerID: 123,
		OrderItems: []domain.OrderItem{
			{
				ProductCode: "phone",
				UnitPrice:   14.7,
				Quantity:    1,
			},
		},
		CreatedAt: 0,
	})
	assert.EqualError(t, err, "connection error")
}

//func Test_Should_Return_Error_When_Payment_Fail(t *testing.T) {
//	payment := new(mockedPayment)
//	db := new(mockedDb)
//	payment.On("Charge", mock.Anything).Return(errors.New("insufficient balance"))
//	db.On("Save", mock.Anything).Return(nil)
//
//	application := NewApplication(db, payment)
//	_, err := application.PlaceOrder(domain.Order{
//		CustomerID: 123,
//		OrderItems: []domain.OrderItem{
//			{
//				ProductCode: "bag",
//				UnitPrice:   2.5,
//				Quantity:    6,
//			},
//		},
//		CreatedAt: 0,
//	})
//	st, _ := status.FromError(err)
//	assert.Equal(t, st.Message(), "order creation failed")
//	assert.Equal(t, st.Details()[0].(*errdetails.BadRequest).FieldViolations[0].Description, "insufficient balance")
//	assert.Equal(t, st.Code(), codes.InvalidArgument)
//}
