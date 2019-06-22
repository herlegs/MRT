package api

import (
	"fmt"

	"github.com/herlegs/MRT/pb"

	"github.com/herlegs/MRT/core"
)

func (s *Server) getCache(from, to string, mode core.Mode) (*pb.RouteResponse, bool) {
	key := getCacheKey(from, to, mode)
	result, exist := s.cache[key]
	return result, exist
}

func (s *Server) setCache(from, to string, mode core.Mode, resp *pb.RouteResponse) {
	key := getCacheKey(from, to, mode)
	s.cache[key] = resp
}

func getCacheKey(from, to string, mode core.Mode) string {
	key := "%v|%v|%v"
	return fmt.Sprintf(key, core.NormalizeName(from), core.NormalizeName(to), mode)
}
