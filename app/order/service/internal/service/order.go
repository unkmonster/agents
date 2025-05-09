package service

import (
	"context"

	pb "agents/api/order/service/v1"
	"agents/app/order/service/internal/biz"
)

type OrderService struct {
	pb.UnimplementedOrderServer
	order *biz.OrderUseCase
}

func NewOrderService(order *biz.OrderUseCase) *OrderService {
	return &OrderService{
		order: order,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderReply, error) {
	order, err := s.order.Create(ctx, &biz.Order{
		Id:          req.Id,
		PaymentType: req.PaymentType,
		Name:        req.Name,
		Amount:      req.Amount,
		Domain:      req.Domain,
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateOrderReply{
		Id:          order.Id,
		PaymentType: order.PaymentType,
		Name:        order.Name,
		Amount:      order.Amount,
		Domain:      order.Domain,
	}, nil
}
func (s *OrderService) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderReply, error) {
	order, err := s.order.Get(ctx, req.Id)

	if err != nil {
		return nil, err
	}

	return &pb.GetOrderReply{
		Id:          order.Id,
		PaymentType: order.PaymentType,
		Name:        order.Name,
		Amount:      order.Amount,
		Domain:      order.Domain,
	}, nil
}
func (s *OrderService) ListOrder(ctx context.Context, req *pb.ListOrderRequest) (*pb.ListOrderReply, error) {
	orders, err := s.order.List(ctx)
	if err != nil {
		return nil, err
	}

	reply := pb.ListOrderReply{}
	for _, order := range orders {
		reply.Orders = append(reply.Orders, &pb.ListOrderReply_Order{
			Id:          order.Id,
			PaymentType: order.PaymentType,
			Name:        order.Name,
			Amount:      order.Amount,
			Domain:      order.Domain,
		})
	}
	return &reply, nil
}
func (s *OrderService) ListOrderByUser(ctx context.Context, req *pb.ListOrderByUserRequest) (*pb.ListOrderReply, error) {
	orders, err := s.order.ListByUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	reply := pb.ListOrderReply{}
	for _, order := range orders {
		reply.Orders = append(reply.Orders, &pb.ListOrderReply_Order{
			Id:          order.Id,
			PaymentType: order.PaymentType,
			Name:        order.Name,
			Amount:      order.Amount,
			Domain:      order.Domain,
		})
	}
	return &reply, nil
}
func (s *OrderService) ListOrderByDomain(ctx context.Context, req *pb.ListOrderByDomainRequest) (*pb.ListOrderReply, error) {
	orders, err := s.order.ListByUser(ctx, req.Domain)
	if err != nil {
		return nil, err
	}

	reply := pb.ListOrderReply{}
	for _, order := range orders {
		reply.Orders = append(reply.Orders, &pb.ListOrderReply_Order{
			Id:          order.Id,
			PaymentType: order.PaymentType,
			Name:        order.Name,
			Amount:      order.Amount,
			Domain:      order.Domain,
		})
	}
	return &reply, nil
}
