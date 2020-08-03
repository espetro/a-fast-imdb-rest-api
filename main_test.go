package main

import (
        "testing"
        "github.com/gofiber/fiber"
)

func Test_setupRoutes(t *testing.T) {
        type args struct {
                app *fiber.App
        }
        tests := []struct {
                name string
                args args
        }{
                // TODO: Add test cases.
        }
        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        setupRoutes(tt.args.app)
                })
        }
}

func Test_main(t *testing.T) {
        tests := []struct {
                name string
        }{
                // TODO: Add test cases.
        }
        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        main()
                })
        }
}