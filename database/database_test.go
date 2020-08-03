package database

import (
        "testing"
        "github.com/gofiber/fiber"
)

func TestProcessErr(t *testing.T) {
        type args struct {
                err error
        }
        tests := []struct {
                name string
                args args
        }{
                // TODO: Add test cases.
        }
        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        processErr(tt.args.err)
                })
        }
}

func TestGetFilm(t *testing.T) {
        type args struct {
                c *fiber.Ctx
        }
        tests := []struct {
                name string
                args args
        }{
                // TODO: Add test cases.
        }
        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        GetFilm(tt.args.c)
                })
        }
}

func TestGetRandomFilm(t *testing.T) {
        type args struct {
                c *fiber.Ctx
        }
        tests := []struct {
                name string
                args args
        }{
                // TODO: Add test cases.
        }
        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        GetRandomFilm(tt.args.c)
                })
        }
}

func TestGetPlayer(t *testing.T) {
        type args struct {
                c *fiber.Ctx
        }
        tests := []struct {
                name string
                args args
        }{
                // TODO: Add test cases.
        }
        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        GetPlayer(tt.args.c)
                })
        }
}

func TestGetRandomPlayer(t *testing.T) {
        type args struct {
                c *fiber.Ctx
        }
        tests := []struct {
                name string
                args args
        }{
                // TODO: Add test cases.
        }
        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        GetRandomPlayer(tt.args.c)
                })
        }
}

func TestDisconnect(t *testing.T) {
        tests := []struct {
                name string
        }{
                // TODO: Add test cases.
        }
        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        disconnect()
                })
        }
}

func TestClose(t *testing.T) {
        tests := []struct {
                name string
        }{
                // TODO: Add test cases.
        }
        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        Close()
                })
        }
}