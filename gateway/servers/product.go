package servers

import (
	"context"

	errs "github.com/bluesky2106/eWallet-backend/errors"
	"github.com/bluesky2106/eWallet-backend/gateway/config"
	"github.com/bluesky2106/eWallet-backend/models"
	pb "github.com/bluesky2106/eWallet-backend/protobuf"
	"google.golang.org/grpc"
)

// IProductSrv : interface of product service
type IProductSrv interface {
	AddProductGroup(req *models.ProductGroup) (*pb.CreateProductGroupRes, error)
}

// ProductSrv : product service
type ProductSrv struct {
	IProductSrv

	conf *config.Config
}

// NewProductServer : returns a pointer of ProductSrv
func NewProductServer(conf *config.Config) *ProductSrv {
	return &ProductSrv{
		conf: conf,
	}
}

// AddProductGroup : create product group
func (prodSrv *ProductSrv) AddProductGroup(productGrp *models.ProductGroup) (*models.ProductGroup, error) {
	conn, err := grpc.Dial(prodSrv.conf.EntryCacheEndpoint, grpc.WithInsecure())
	if err != nil {
		return nil, errs.GRPCDialError(err)
	}
	defer conn.Close()

	c := pb.NewProductSvcClient(conn)
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
