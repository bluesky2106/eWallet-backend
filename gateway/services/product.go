package services

import (
	"context"

	errs "github.com/bluesky2106/eWallet-backend/errors"
	"github.com/bluesky2106/eWallet-backend/gateway/config"
	"github.com/bluesky2106/eWallet-backend/models"
	pb "github.com/bluesky2106/eWallet-backend/protobuf"
	"google.golang.org/grpc"
)

// IProductService : interface of product service
type IProductService interface {
	AddProductGroup(req *models.ProductGroup) (*pb.CreateProductGroupRes, error)
}

// ProductService : product service
type ProductService struct {
	IProductService

	conf *config.Config
}

// NewProductService : returns a pointer of ProductService
func NewProductService(conf *config.Config) *ProductService {
	return &ProductService{
		conf: conf,
	}
}

// AddProductGroup : create product group
func (prodSrv *ProductService) AddProductGroup(productGrp *models.ProductGroup) (*models.ProductGroup, error) {
	conn, err := grpc.Dial(prodSrv.conf.EntryCacheEndpoint, grpc.WithInsecure())
	if err != nil {
		return nil, errs.GRPCDialError(err)
	}
	defer conn.Close()

	c := pb.NewProductSrvClient(conn)
	res, err := c.CreateProductGroup(context.Background(), &pb.CreateProductGroupReq{
		Request: &pb.BaseReq{
			Action:     pb.Action_ACTION_CREATE,
			Message:    pb.Message_MESSAGE_CREATE_PRODUCT,
			ObjectType: pb.Object_OBJECT_PRODUCT_GROUP,
		},
		ProductGroup: &pb.ProductGroup{
			GId:         productGrp.GID,
			Name:        productGrp.Name,
			Description: productGrp.Description,
		},
	})

	if err != nil {
		return nil, errs.WithMessage(err, "c.CreateProductGroup")
	}

	return models.ToProductGroup(res.GetProductGroup()), nil
}
