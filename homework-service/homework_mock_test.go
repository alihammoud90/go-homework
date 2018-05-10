package main_test

import (
	"reflect"
	"testing"
	//"time"

	"github.com/golang/mock/gomock"
	//"golang.org/x/net/context"
	pb "go-homework/homework-service/proto"
	mock "go-homework/homework-service/mock_homework"
	"golang.org/x/net/context"
)

func TestGetSum(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHomeWorkClient := mock.NewMockHomeworkServiceClient(ctrl)
	req := &pb.Numbers{Nb1: 4, Nb2: 5}
	mockHomeWorkClient.EXPECT().GetSum(context.Background(), req).Return(&pb.Sum{Sum: 9}, nil).Times(1)
	testGetSum(t, mockHomeWorkClient)
}


func testGetSum(t *testing.T, client pb.HomeworkServiceClient) {
	r, err := client.GetSum(context.Background(), &pb.Numbers{Nb1: 4, Nb2: 5})
	if err != nil || r.Sum != 9 {
		t.Errorf("Mocking Sum Failed")
	}
	t.Log("Reply : ", r.Sum)
}

func TestOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHomeWorkClient := mock.NewMockHomeworkServiceClient(ctrl)
	var array []int64
	array = append(array, 1, 5, 2, 8, 3)

	var correctArray []int64
	correctArray = append(correctArray, 1, 2, 3, 5, 8)
	req := &pb.NumbersToOrder{Array: array}
	mockHomeWorkClient.EXPECT().OrderNumbers(context.Background(), req).Return(&pb.OrderedNumbers{OrderedArray: correctArray}, nil).Times(1)
	testOrder(t, mockHomeWorkClient)
}


func testOrder(t *testing.T, client pb.HomeworkServiceClient) {
	var array []int64
	array = append(array, 1, 5, 2, 8, 3)

	var correctArray []int64
	correctArray = append(correctArray, 1, 2, 3, 5, 8)

	r, err := client.OrderNumbers(context.Background(), &pb.NumbersToOrder{Array: array})
	if err != nil || (!reflect.DeepEqual(r.OrderedArray, correctArray)) {
		t.Errorf("Mocking Order Failed")
	}
	t.Log("Reply : ", r.OrderedArray)
}