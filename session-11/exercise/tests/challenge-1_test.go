package test

import (
	"exercise/helpers"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPria(t *testing.T) {
	type request struct {
		gender string
		age    int
	}
	person := []struct {
		req request
		res bool
		msg string
	}{
		{
			req: request{"pria", 16},
			res: false,
		},
		{
			req: request{"pria", 17},
			res: false,
		},
		{
			req: request{"pria", 50},
			res: true,
		},
		{
			req: request{"pria", 60},
			res: false,
		},
	}

	for k, people := range person {
		t.Run(fmt.Sprintf("test-%d", k+1), func(t *testing.T) {
			require.Equal(t, people.res, helpers.IsProductiveAge(people.req.gender, people.req.age))
		})
	}
}

func TestWanita(t *testing.T) {
	type request struct {
		gender string
		age    int
	}
	person := []struct {
		req request
		res bool
		msg string
	}{
		{
			req: request{"wanita", 16},
			res: false,
		},
		{
			req: request{"wanita", 21},
			res: true,
		},
		{
			req: request{"wanita", 50},
			res: true,
		},
		{
			req: request{"wanita", 60},
			res: false,
		},
	}

	for k, people := range person {
		t.Run(fmt.Sprintf("test-%d", k+1), func(t *testing.T) {
			require.Equal(t, people.res, helpers.IsProductiveAge(people.req.gender, people.req.age))
		})
	}
}
