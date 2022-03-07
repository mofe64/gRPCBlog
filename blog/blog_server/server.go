package blog_server

import (
	"grpcBlog/blog/blog_pb"
)

type Server struct {
	blog_pb.UnimplementedBlogServiceServer
}
