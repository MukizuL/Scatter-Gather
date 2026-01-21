package main

import (
	"context"
	"math/rand/v2"
	"time"
)

type Response struct {
	Service string `json:"service"`
	Status  int    `json:"status"`
}

type UserService struct{}

func (service *UserService) getUser(ctx context.Context) Response {
	select {
	case <-ctx.Done():
		return Response{Service: "UserService", Status: 500}
	case <-time.After(10 * time.Millisecond):
		if rand.Float32() > 0.3 {
			return Response{Service: "UserService", Status: 200}
		}

		return Response{Service: "UserService", Status: 500}
	}
}

type VectorMemory struct{}

func (service *VectorMemory) getContext(ctx context.Context) Response {
	select {
	case <-ctx.Done():
		return Response{Service: "VectorMemory", Status: 500}
	case <-time.After(time.Duration(rand.IntN(2900)+100) * time.Millisecond):
		if rand.Float32() > 0.3 {
			return Response{Service: "VectorMemory", Status: 200}
		}
		return Response{Service: "VectorMemory", Status: 500}
	}
}

type PermissionsService struct{}

func (service *PermissionsService) checkAccess(ctx context.Context) Response {
	select {
	case <-ctx.Done():
		return Response{Service: "PermissionService", Status: 500}
	case <-time.After(50 * time.Millisecond):
		if rand.Float32() > 0.3 {
			return Response{Service: "PermissionService", Status: 200}
		}

		return Response{Service: "PermissionService", Status: 500}
	}
}
