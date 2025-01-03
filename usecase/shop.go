package usecase

import (
	"fmt"
	"github.com/go-to/egp-backend/usecase/input"
	"github.com/go-to/egp-backend/usecase/output"
	"github.com/go-to/egp-protobuf/pb"
)

func GetShops(input *input.ShopsInput) (output.ShopsOutput, error) {
	fmt.Println(&input)

	var shops []*pb.Shop
	for i := 1; i <= 10; i++ {
		shops = append(shops, &pb.Shop{
			ID: int64(i),
		})
	}

	return output.ShopsOutput{
		ShopsResponse: pb.ShopsResponse{
			Shops: shops,
		},
	}, nil
}
