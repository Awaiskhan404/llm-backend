/*
Package Name: bootstrap
File Name: routes.go
Abstract: The wrapper for setting up all the routes.

Author: Alejandro Modroño <alex@sureservice.es>
Created: 07/08/2023
Last Updated: 07/24/2023

# MIT License

# Copyright 2023 Alejandro Modroño Vara

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package bootstrap

import (
	"llm-backend/pkg/auth"
	"llm-backend/pkg/users"
)

// ======== TYPES ========

// Route interface
type Route interface {
	Setup()
}

// Routes contains multiple routes
type Routes []Route

// ======== PUBLIC METHODS ========

// GetRoutes provides all the routes
func GetRoutes(
	userRoutes users.UsersRoutes,
	authRoutes auth.AuthRoutes,
) Routes {
	return Routes{
		userRoutes,
		authRoutes,
	}
}

// Sets up all the routes
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
