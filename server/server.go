package main

import (
	"animeclub.com/calculatorpb"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type Server struct {
	calculatorpb.UnimplementedCalculatorServiceServer
}

func (s *Server) Calculator(req *calculatorpb.CalculatorRequest,stream calculatorpb.CalculatorService_CalculatorServer)error {
	number := req.GetNumber()
	var slice[]int =PrimeFactors(int(number))

	for i:=0;i<len(slice);i++ {
		res:=&calculatorpb.CalculatorResponse{Prime: int64(slice[i])}
		if err:=stream.Send(res); err!=nil{
			log.Fatalf("errir")
		}
		time.Sleep(time.Second)
	}
	return nil
}
func PrimeFactors(n int) (pfs []int) {
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:5555")
	if err != nil {
		log.Fatal("Failed to listen :%v", err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &Server{})
	log.Println("Server is running on port 5555 :")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

//Greet is an example of unary rpc call
