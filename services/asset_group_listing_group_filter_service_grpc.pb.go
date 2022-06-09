// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package services

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AssetGroupListingGroupFilterServiceClient is the client API for AssetGroupListingGroupFilterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssetGroupListingGroupFilterServiceClient interface {
	// Creates, updates or removes asset group listing group filters. Operation
	// statuses are returned.
	MutateAssetGroupListingGroupFilters(ctx context.Context, in *MutateAssetGroupListingGroupFiltersRequest, opts ...grpc.CallOption) (*MutateAssetGroupListingGroupFiltersResponse, error)
}

type assetGroupListingGroupFilterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAssetGroupListingGroupFilterServiceClient(cc grpc.ClientConnInterface) AssetGroupListingGroupFilterServiceClient {
	return &assetGroupListingGroupFilterServiceClient{cc}
}

func (c *assetGroupListingGroupFilterServiceClient) MutateAssetGroupListingGroupFilters(ctx context.Context, in *MutateAssetGroupListingGroupFiltersRequest, opts ...grpc.CallOption) (*MutateAssetGroupListingGroupFiltersResponse, error) {
	out := new(MutateAssetGroupListingGroupFiltersResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v11.services.AssetGroupListingGroupFilterService/MutateAssetGroupListingGroupFilters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssetGroupListingGroupFilterServiceServer is the server API for AssetGroupListingGroupFilterService service.
// All implementations must embed UnimplementedAssetGroupListingGroupFilterServiceServer
// for forward compatibility
type AssetGroupListingGroupFilterServiceServer interface {
	// Creates, updates or removes asset group listing group filters. Operation
	// statuses are returned.
	MutateAssetGroupListingGroupFilters(context.Context, *MutateAssetGroupListingGroupFiltersRequest) (*MutateAssetGroupListingGroupFiltersResponse, error)
	mustEmbedUnimplementedAssetGroupListingGroupFilterServiceServer()
}

// UnimplementedAssetGroupListingGroupFilterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAssetGroupListingGroupFilterServiceServer struct {
}

func (UnimplementedAssetGroupListingGroupFilterServiceServer) MutateAssetGroupListingGroupFilters(context.Context, *MutateAssetGroupListingGroupFiltersRequest) (*MutateAssetGroupListingGroupFiltersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAssetGroupListingGroupFilters not implemented")
}
func (UnimplementedAssetGroupListingGroupFilterServiceServer) mustEmbedUnimplementedAssetGroupListingGroupFilterServiceServer() {
}

// UnsafeAssetGroupListingGroupFilterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssetGroupListingGroupFilterServiceServer will
// result in compilation errors.
type UnsafeAssetGroupListingGroupFilterServiceServer interface {
	mustEmbedUnimplementedAssetGroupListingGroupFilterServiceServer()
}

func RegisterAssetGroupListingGroupFilterServiceServer(s grpc.ServiceRegistrar, srv AssetGroupListingGroupFilterServiceServer) {
	s.RegisterService(&AssetGroupListingGroupFilterService_ServiceDesc, srv)
}

func _AssetGroupListingGroupFilterService_MutateAssetGroupListingGroupFilters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAssetGroupListingGroupFiltersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetGroupListingGroupFilterServiceServer).MutateAssetGroupListingGroupFilters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v11.services.AssetGroupListingGroupFilterService/MutateAssetGroupListingGroupFilters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetGroupListingGroupFilterServiceServer).MutateAssetGroupListingGroupFilters(ctx, req.(*MutateAssetGroupListingGroupFiltersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AssetGroupListingGroupFilterService_ServiceDesc is the grpc.ServiceDesc for AssetGroupListingGroupFilterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AssetGroupListingGroupFilterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v11.services.AssetGroupListingGroupFilterService",
	HandlerType: (*AssetGroupListingGroupFilterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateAssetGroupListingGroupFilters",
			Handler:    _AssetGroupListingGroupFilterService_MutateAssetGroupListingGroupFilters_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v11/services/asset_group_listing_group_filter_service.proto",
}
