package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSApiGatewayDomainName AWS CloudFormation Resource (AWS::ApiGateway::DomainName)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html
type AWSApiGatewayDomainName struct {

	// CertificateArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html#cfn-apigateway-domainname-certificatearn
	CertificateArn *Value `json:"CertificateArn,omitempty"`

	// DomainName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html#cfn-apigateway-domainname-domainname
	DomainName *Value `json:"DomainName,omitempty"`

	// EndpointConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html#cfn-apigateway-domainname-endpointconfiguration
	EndpointConfiguration *AWSApiGatewayDomainName_EndpointConfiguration `json:"EndpointConfiguration,omitempty"`

	// RegionalCertificateArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html#cfn-apigateway-domainname-regionalcertificatearn
	RegionalCertificateArn *Value `json:"RegionalCertificateArn,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayDomainName) AWSCloudFormationType() string {
	return "AWS::ApiGateway::DomainName"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSApiGatewayDomainName) MarshalJSON() ([]byte, error) {
	type Properties AWSApiGatewayDomainName
	return json.Marshal(&struct {
		Type       string
		Properties Properties
	}{
		Type:       r.AWSCloudFormationType(),
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSApiGatewayDomainName) UnmarshalJSON(b []byte) error {
	type Properties AWSApiGatewayDomainName
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSApiGatewayDomainName(*res.Properties)
	}

	return nil
}

// GetAllAWSApiGatewayDomainNameResources retrieves all AWSApiGatewayDomainName items from an AWS CloudFormation template
func (t *Template) GetAllAWSApiGatewayDomainNameResources() map[string]AWSApiGatewayDomainName {
	results := map[string]AWSApiGatewayDomainName{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSApiGatewayDomainName:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::ApiGateway::DomainName" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSApiGatewayDomainName{}
						if err := result.UnmarshalJSON(b); err == nil {
							results[name] = *result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSApiGatewayDomainNameWithName retrieves all AWSApiGatewayDomainName items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayDomainNameWithName(name string) (AWSApiGatewayDomainName, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSApiGatewayDomainName:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::ApiGateway::DomainName" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSApiGatewayDomainName{}
						if err := result.UnmarshalJSON(b); err == nil {
							return *result, nil
						}
					}
				}
			}
		}
	}
	return AWSApiGatewayDomainName{}, errors.New("resource not found")
}
