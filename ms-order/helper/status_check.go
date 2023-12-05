package helper

import (
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func StatusCheck(os, us string) error {
	us = strings.ToUpper(us)

	if us == "CANCEL" {
		switch os {
		case "CANCEL":
			return status.Error(codes.InvalidArgument, "order already canceled")
		case "PAID":
			return status.Error(codes.InvalidArgument, "order already paid")
		case "SUCCESS":
			return status.Error(codes.InvalidArgument, "order already success")
		}
	}

	if os == "PENDING" && us == "SUCCESS" {
		return status.Error(codes.InvalidArgument, "order not paid yet")
	}

	if os == "PAID" && us == "PENDING" {
		return status.Error(codes.InvalidArgument, "order already paid")
	}

	if os == "CANCEL" {
		return status.Error(codes.InvalidArgument, "order already canceled")
	}

	if os == "SUCCESS" {
		return status.Error(codes.InvalidArgument, "order already success")
	}

	return nil
}
