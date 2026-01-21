package main

import (
	"context"
	"errors"
	"net/http"
	"time"
)

func (app *application) summary(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	const serviceCount = 3
	responses := make(chan Response, serviceCount)

	userService := UserService{}
	vectorMemory := VectorMemory{}
	permissionsService := PermissionsService{}

	go func() {
		responses <- userService.getUser(ctx)
	}()

	go func() {
		responses <- vectorMemory.getContext(ctx)
	}()

	go func() {
		responses <- permissionsService.checkAccess(ctx)
	}()

	responseMap := map[string]int{
		"UserService":       500,
		"VectorMemory":      500,
		"PermissionService": 500,
	}

received:
	for i := 0; i < serviceCount; i++ {
		select {
		case <-ctx.Done():
			break received
		case resp := <-responses:
			responseMap[resp.Service] = resp.Status
		}
	}

	if responseMap["UserService"] != 200 {
		app.serverError(w, r, errors.New("UserService took too long to answer or crashed"))
		return
	}

	if responseMap["PermissionService"] != 200 {
		app.serverError(w, r, errors.New("PermissionService took too long to answer or crashed"))
		return
	}

	err := app.writeJSON(w, 200,
		envelope{
			"user_service_status":       responseMap["UserService"],
			"vector_memory_status":      responseMap["VectorMemory"],
			"permission_service_status": responseMap["PermissionService"]})
	if err != nil {
		app.serverError(w, r, err)
	}
}
