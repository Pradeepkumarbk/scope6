// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cloudfront

import (
	"github.com/aws/aws-sdk-go/private/waiter"
)

func (c *CloudFront) WaitUntilDistributionDeployed(input *GetDistributionInput) error {
	waiterCfg := waiter.Config{
		Operation:   "GetDistribution",
		Delay:       60,
		MaxAttempts: 25,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "path",
				Argument: "Status",
				Expected: "Deployed",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *CloudFront) WaitUntilInvalidationCompleted(input *GetInvalidationInput) error {
	waiterCfg := waiter.Config{
		Operation:   "GetInvalidation",
		Delay:       20,
		MaxAttempts: 30,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "path",
				Argument: "Status",
				Expected: "Completed",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *CloudFront) WaitUntilStreamingDistributionDeployed(input *GetStreamingDistributionInput) error {
	waiterCfg := waiter.Config{
		Operation:   "GetStreamingDistribution",
		Delay:       60,
		MaxAttempts: 25,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "path",
				Argument: "Status",
				Expected: "Deployed",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}
