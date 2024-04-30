package aws

import (
	"context"
	"html/template"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/iam"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

func (c *AWSClient) ReconcileIAMUser(username string) (*iam.User, error) {
	svc := iam.New(c.sess)

	if user, err := svc.GetUser(&iam.GetUserInput{
		UserName: &username,
	}); err == nil {
		return user.User, nil // User exists.
	} else if aerr, ok := err.(awserr.Error); ok {
		if aerr.Code() != iam.ErrCodeNoSuchEntityException {
			return nil, err
		}
	} else {
		return nil, err
	}

	// User does not exist. Create user.
	if user, err := svc.CreateUser(&iam.CreateUserInput{
		UserName: &username,
	}); err == nil {
		return user.User, nil
	} else {
		return nil, err
	}
}

func (c *AWSClient) ReconcileIAMRole(ctx context.Context, roleName, namespace string) (*iam.Role, error) {
	log := log.FromContext(ctx)

	svc := iam.New(c.sess)

	if role, err := svc.GetRole(&iam.GetRoleInput{
		RoleName: &roleName,
	}); err == nil {
		return role.Role, nil // Role exists.
	} else if aerr, ok := err.(awserr.Error); ok {
		if aerr.Code() == iam.ErrCodeNoSuchEntityException {
			// Role does not exist. Continue.
			log.Info("Role does not exist", "role", roleName, "namespace", namespace)
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}

	// Create role.
	trustPolicy, err := os.ReadFile("../templates/aws/policies/trust-policy.json")
	if err != nil {
		return nil, err
	}
	tmpl, err := template.New("trust-policy").Parse(string(trustPolicy))
	if err != nil {
		return nil, err
	}
	assumeRolePolicyDocument := new(strings.Builder)
	if err := tmpl.Execute(assumeRolePolicyDocument, map[string]any{
		"accountID": c.config.AccountID,
		"oidc": map[string]any{
			"provider": c.config.OIDC.Provider,
		},
		"namespace":      namespace,
		"serviceAccount": "workspace-controller",
	}); err != nil {
		return nil, err
	}
	if role, err := svc.CreateRole(&iam.CreateRoleInput{
		RoleName:                 &roleName,
		Path:                     aws.String("/"),
		AssumeRolePolicyDocument: aws.String(assumeRolePolicyDocument.String()),
	}); err == nil {
		log.Info("Role created", "role", roleName, "namespace", namespace)
		return role.Role, nil
	} else {
		return nil, err
	}
}

func (c *AWSClient) DeleteIAMRole(ctx context.Context, roleName string) error {
	log := log.FromContext(ctx)

	svc := iam.New(c.sess)

	if _, err := svc.DeleteRole(&iam.DeleteRoleInput{
		RoleName: &roleName,
	}); err == nil {
		log.Info("Role deleted", "role", roleName)
		return nil
	} else if aerr, ok := err.(awserr.Error); ok {
		if aerr.Code() == iam.ErrCodeNoSuchEntityException {
			// Role does not exist. Continue.
			log.Info("Role does not exist", "role", roleName)
		} else {
			return err
		}
	} else {
		return err
	}
	return nil
}

func (c *AWSClient) ReconcileIAMRolePolicy(ctx context.Context, policyName string, role *iam.Role) (*string, error) {
	log := log.FromContext(ctx)

	svc := iam.New(c.sess)

	if policy, err := svc.GetRolePolicy(&iam.GetRolePolicyInput{
		PolicyName: &policyName,
		RoleName:   role.RoleName,
	}); err == nil {
		return policy.PolicyDocument, nil // Policy exists.
	} else if aerr, ok := err.(awserr.Error); ok {
		if aerr.Code() == iam.ErrCodeNoSuchEntityException {
			// Policy does not exist. Continue.
			log.Info("Policy does not exist", "policy", policyName, "role", role.RoleName)
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}

	// Create policy.
	policyDoumentTemplate, err := os.ReadFile("../templates/aws/policies/efs-role-policy.json")
	if err != nil {
		return nil, err
	}
	tmpl, err := template.New("efs-role-policy").Parse(string(policyDoumentTemplate))
	if err != nil {
		return nil, err
	}
	rolePolicyDocument := new(strings.Builder)
	if err = tmpl.Execute(rolePolicyDocument, map[string]any{
		"accountID": c.config.AccountID,
		"efsID":     c.config.Storage.EFSID,
	}); err != nil {
		return nil, err
	}
	if policy, err := svc.PutRolePolicy(&iam.PutRolePolicyInput{
		PolicyDocument: aws.String(rolePolicyDocument.String()),
		PolicyName:     &policyName,
		RoleName:       role.RoleName,
	}); err == nil {
		log.Info("Policy created", "policy", policyName, "role", role.RoleName)
		p := policy.String()
		return &p, nil
	} else {
		return nil, err
	}
}

func (c *AWSClient) DeleteIAMRolePolicy(ctx context.Context, policyName string) error {
	log := log.FromContext(ctx)

	svc := iam.New(c.sess)

	if _, err := svc.DeleteRolePolicy(&iam.DeleteRolePolicyInput{
		PolicyName: &policyName,
		RoleName:   &policyName,
	}); err == nil {
		log.Info("Policy deleted", "policy", policyName)
		return nil
	} else if aerr, ok := err.(awserr.Error); ok {
		if aerr.Code() == iam.ErrCodeNoSuchEntityException {
			// Policy does not exist. Continue.
			log.Info("Policy does not exist", "policy", policyName)
		} else {
			return err
		}
	} else {
		return err
	}
	return nil
}