// Code generated by protoc-gen-mrpc.

package message

import "errors"

// ArithService Defining Computational Digital Services
type ArithService struct{}

// Add addition
func (this *ArithService) Add(args *ArithRequest, reply *ArithResponse) error {
	// define your service ...
	reply.C = args.A + args.B
	return nil
}

// Sub subtraction
func (this *ArithService) Sub(args *ArithRequest, reply *ArithResponse) error {
	// define your service ...
	reply.C = args.A - args.B
	return nil
}

// Mul multiplication
func (this *ArithService) Mul(args *ArithRequest, reply *ArithResponse) error {
	// define your service ...
	reply.C = args.A * args.B
	return nil
}

// Div division
func (this *ArithService) Div(args *ArithRequest, reply *ArithResponse) error {
	// define your service ...
	if args.B == 0 {
		return errors.New("div is zero")
	}
	reply.C = args.A / args.B
	return nil
}
