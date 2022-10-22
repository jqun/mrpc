package message

import "errors"

type ArithService struct {
}

func (s *ArithService) Add(args *ArithRequest, reply *ArithResponse) error {
	reply.C = args.A + args.B
	return nil
}

func (s *ArithService) Sub(args *ArithRequest, reply *ArithResponse) error {
	reply.C = args.A - args.B
	return nil
}

func (s *ArithService) Mul(args *ArithRequest, reply *ArithResponse) error {
	reply.C = args.A * args.B
	return nil
}

func (s *ArithService) Div(args *ArithRequest, reply *ArithResponse) error {
	if args.B == 0 {
		return errors.New("div is zero")
	}

	reply.C = args.A / args.B
	return nil
}
