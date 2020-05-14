package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	pb "github.com/bubble-hook/lightnet-assignment/calculator_proto"
	"github.com/bubble-hook/lightnet-assignment/shared"
	"google.golang.org/grpc"
)

// Calculate Calculate
func Calculate(w http.ResponseWriter, r *http.Request) {

	request := shared.CalculateRequest{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		shared.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	internalServiceEnpoint := os.Getenv("SERVICE_ENDPOINT")

	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(internalServiceEnpoint, opts...)
	if err != nil {
		shared.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	client := pb.NewCalculateClient(conn)

	services := map[string]func(ctx context.Context, in *pb.CalculateRequest, opts ...grpc.CallOption) (*pb.CalculateResponnse, error){
		"/calculator.sum": client.Sum,
		"/calculator.sub": client.Sub,
		"/calculator.mul": client.Mul,
		"/calculator.div": client.Div,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if calculateFunc, ok := services[r.RequestURI]; ok {

		resp, err := calculateFunc(ctx, &pb.CalculateRequest{
			A: request.A,
			B: request.B,
		})
		if err != nil {
			shared.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		shared.RespondJSON(w, http.StatusOK, shared.CalculateResponse{
			CalculateRequest: shared.CalculateRequest{
				A: resp.GetA(),
				B: resp.GetB(),
			},
			Result: resp.GetResult(),
		})
	} else {
		shared.RespondError(w, http.StatusBadRequest, "Not Avalible Service")
	}

}
