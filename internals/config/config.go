package config

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	c "github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

type Config struct {
	Environment string
	Vpc         VPC
}

type VPC struct {
	CidrBlock string
}

func GetConfig(ctx *pulumi.Context) (config Config, err error) {
	var con Config
	cfg := c.New(ctx, "")
	con.Environment = cfg.Require("Environment")
	// fmt.Print(con.Environment)

	err = cfg.GetObject("vpc", &con.Vpc)
	if err != nil {

		fmt.Print("error getting vpc config")
		return con, nil
	}

	return con, nil
}
