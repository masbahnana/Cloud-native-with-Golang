// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package wafregionaliface provides an interface to enable mocking the AWS WAF Regional service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package wafregionaliface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/waf"
	"github.com/aws/aws-sdk-go/service/wafregional"
)

// WAFRegionalAPI provides an interface to enable mocking the
// wafregional.WAFRegional service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS WAF Regional.
//    func myFunc(svc wafregionaliface.WAFRegionalAPI) bool {
//        // Make svc.AssociateWebACL request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := wafregional.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockWAFRegionalClient struct {
//        wafregionaliface.WAFRegionalAPI
//    }
//    func (m *mockWAFRegionalClient) AssociateWebACL(input *wafregional.AssociateWebACLInput) (*wafregional.AssociateWebACLOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockWAFRegionalClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type WAFRegionalAPI interface {
	AssociateWebACL(*wafregional.AssociateWebACLInput) (*wafregional.AssociateWebACLOutput, error)
	AssociateWebACLWithContext(aws.Context, *wafregional.AssociateWebACLInput, ...request.Option) (*wafregional.AssociateWebACLOutput, error)
	AssociateWebACLRequest(*wafregional.AssociateWebACLInput) (*request.Request, *wafregional.AssociateWebACLOutput)

	CreateByteMatchSet(*waf.CreateByteMatchSetInput) (*waf.CreateByteMatchSetOutput, error)
	CreateByteMatchSetWithContext(aws.Context, *waf.CreateByteMatchSetInput, ...request.Option) (*waf.CreateByteMatchSetOutput, error)
	CreateByteMatchSetRequest(*waf.CreateByteMatchSetInput) (*request.Request, *waf.CreateByteMatchSetOutput)

	CreateIPSet(*waf.CreateIPSetInput) (*waf.CreateIPSetOutput, error)
	CreateIPSetWithContext(aws.Context, *waf.CreateIPSetInput, ...request.Option) (*waf.CreateIPSetOutput, error)
	CreateIPSetRequest(*waf.CreateIPSetInput) (*request.Request, *waf.CreateIPSetOutput)

	CreateRule(*waf.CreateRuleInput) (*waf.CreateRuleOutput, error)
	CreateRuleWithContext(aws.Context, *waf.CreateRuleInput, ...request.Option) (*waf.CreateRuleOutput, error)
	CreateRuleRequest(*waf.CreateRuleInput) (*request.Request, *waf.CreateRuleOutput)

	CreateSizeConstraintSet(*waf.CreateSizeConstraintSetInput) (*waf.CreateSizeConstraintSetOutput, error)
	CreateSizeConstraintSetWithContext(aws.Context, *waf.CreateSizeConstraintSetInput, ...request.Option) (*waf.CreateSizeConstraintSetOutput, error)
	CreateSizeConstraintSetRequest(*waf.CreateSizeConstraintSetInput) (*request.Request, *waf.CreateSizeConstraintSetOutput)

	CreateSqlInjectionMatchSet(*waf.CreateSqlInjectionMatchSetInput) (*waf.CreateSqlInjectionMatchSetOutput, error)
	CreateSqlInjectionMatchSetWithContext(aws.Context, *waf.CreateSqlInjectionMatchSetInput, ...request.Option) (*waf.CreateSqlInjectionMatchSetOutput, error)
	CreateSqlInjectionMatchSetRequest(*waf.CreateSqlInjectionMatchSetInput) (*request.Request, *waf.CreateSqlInjectionMatchSetOutput)

	CreateWebACL(*waf.CreateWebACLInput) (*waf.CreateWebACLOutput, error)
	CreateWebACLWithContext(aws.Context, *waf.CreateWebACLInput, ...request.Option) (*waf.CreateWebACLOutput, error)
	CreateWebACLRequest(*waf.CreateWebACLInput) (*request.Request, *waf.CreateWebACLOutput)

	CreateXssMatchSet(*waf.CreateXssMatchSetInput) (*waf.CreateXssMatchSetOutput, error)
	CreateXssMatchSetWithContext(aws.Context, *waf.CreateXssMatchSetInput, ...request.Option) (*waf.CreateXssMatchSetOutput, error)
	CreateXssMatchSetRequest(*waf.CreateXssMatchSetInput) (*request.Request, *waf.CreateXssMatchSetOutput)

	DeleteByteMatchSet(*waf.DeleteByteMatchSetInput) (*waf.DeleteByteMatchSetOutput, error)
	DeleteByteMatchSetWithContext(aws.Context, *waf.DeleteByteMatchSetInput, ...request.Option) (*waf.DeleteByteMatchSetOutput, error)
	DeleteByteMatchSetRequest(*waf.DeleteByteMatchSetInput) (*request.Request, *waf.DeleteByteMatchSetOutput)

	DeleteIPSet(*waf.DeleteIPSetInput) (*waf.DeleteIPSetOutput, error)
	DeleteIPSetWithContext(aws.Context, *waf.DeleteIPSetInput, ...request.Option) (*waf.DeleteIPSetOutput, error)
	DeleteIPSetRequest(*waf.DeleteIPSetInput) (*request.Request, *waf.DeleteIPSetOutput)

	DeleteRule(*waf.DeleteRuleInput) (*waf.DeleteRuleOutput, error)
	DeleteRuleWithContext(aws.Context, *waf.DeleteRuleInput, ...request.Option) (*waf.DeleteRuleOutput, error)
	DeleteRuleRequest(*waf.DeleteRuleInput) (*request.Request, *waf.DeleteRuleOutput)

	DeleteSizeConstraintSet(*waf.DeleteSizeConstraintSetInput) (*waf.DeleteSizeConstraintSetOutput, error)
	DeleteSizeConstraintSetWithContext(aws.Context, *waf.DeleteSizeConstraintSetInput, ...request.Option) (*waf.DeleteSizeConstraintSetOutput, error)
	DeleteSizeConstraintSetRequest(*waf.DeleteSizeConstraintSetInput) (*request.Request, *waf.DeleteSizeConstraintSetOutput)

	DeleteSqlInjectionMatchSet(*waf.DeleteSqlInjectionMatchSetInput) (*waf.DeleteSqlInjectionMatchSetOutput, error)
	DeleteSqlInjectionMatchSetWithContext(aws.Context, *waf.DeleteSqlInjectionMatchSetInput, ...request.Option) (*waf.DeleteSqlInjectionMatchSetOutput, error)
	DeleteSqlInjectionMatchSetRequest(*waf.DeleteSqlInjectionMatchSetInput) (*request.Request, *waf.DeleteSqlInjectionMatchSetOutput)

	DeleteWebACL(*waf.DeleteWebACLInput) (*waf.DeleteWebACLOutput, error)
	DeleteWebACLWithContext(aws.Context, *waf.DeleteWebACLInput, ...request.Option) (*waf.DeleteWebACLOutput, error)
	DeleteWebACLRequest(*waf.DeleteWebACLInput) (*request.Request, *waf.DeleteWebACLOutput)

	DeleteXssMatchSet(*waf.DeleteXssMatchSetInput) (*waf.DeleteXssMatchSetOutput, error)
	DeleteXssMatchSetWithContext(aws.Context, *waf.DeleteXssMatchSetInput, ...request.Option) (*waf.DeleteXssMatchSetOutput, error)
	DeleteXssMatchSetRequest(*waf.DeleteXssMatchSetInput) (*request.Request, *waf.DeleteXssMatchSetOutput)

	DisassociateWebACL(*wafregional.DisassociateWebACLInput) (*wafregional.DisassociateWebACLOutput, error)
	DisassociateWebACLWithContext(aws.Context, *wafregional.DisassociateWebACLInput, ...request.Option) (*wafregional.DisassociateWebACLOutput, error)
	DisassociateWebACLRequest(*wafregional.DisassociateWebACLInput) (*request.Request, *wafregional.DisassociateWebACLOutput)

	GetByteMatchSet(*waf.GetByteMatchSetInput) (*waf.GetByteMatchSetOutput, error)
	GetByteMatchSetWithContext(aws.Context, *waf.GetByteMatchSetInput, ...request.Option) (*waf.GetByteMatchSetOutput, error)
	GetByteMatchSetRequest(*waf.GetByteMatchSetInput) (*request.Request, *waf.GetByteMatchSetOutput)

	GetChangeToken(*waf.GetChangeTokenInput) (*waf.GetChangeTokenOutput, error)
	GetChangeTokenWithContext(aws.Context, *waf.GetChangeTokenInput, ...request.Option) (*waf.GetChangeTokenOutput, error)
	GetChangeTokenRequest(*waf.GetChangeTokenInput) (*request.Request, *waf.GetChangeTokenOutput)

	GetChangeTokenStatus(*waf.GetChangeTokenStatusInput) (*waf.GetChangeTokenStatusOutput, error)
	GetChangeTokenStatusWithContext(aws.Context, *waf.GetChangeTokenStatusInput, ...request.Option) (*waf.GetChangeTokenStatusOutput, error)
	GetChangeTokenStatusRequest(*waf.GetChangeTokenStatusInput) (*request.Request, *waf.GetChangeTokenStatusOutput)

	GetIPSet(*waf.GetIPSetInput) (*waf.GetIPSetOutput, error)
	GetIPSetWithContext(aws.Context, *waf.GetIPSetInput, ...request.Option) (*waf.GetIPSetOutput, error)
	GetIPSetRequest(*waf.GetIPSetInput) (*request.Request, *waf.GetIPSetOutput)

	GetRule(*waf.GetRuleInput) (*waf.GetRuleOutput, error)
	GetRuleWithContext(aws.Context, *waf.GetRuleInput, ...request.Option) (*waf.GetRuleOutput, error)
	GetRuleRequest(*waf.GetRuleInput) (*request.Request, *waf.GetRuleOutput)

	GetSampledRequests(*waf.GetSampledRequestsInput) (*waf.GetSampledRequestsOutput, error)
	GetSampledRequestsWithContext(aws.Context, *waf.GetSampledRequestsInput, ...request.Option) (*waf.GetSampledRequestsOutput, error)
	GetSampledRequestsRequest(*waf.GetSampledRequestsInput) (*request.Request, *waf.GetSampledRequestsOutput)

	GetSizeConstraintSet(*waf.GetSizeConstraintSetInput) (*waf.GetSizeConstraintSetOutput, error)
	GetSizeConstraintSetWithContext(aws.Context, *waf.GetSizeConstraintSetInput, ...request.Option) (*waf.GetSizeConstraintSetOutput, error)
	GetSizeConstraintSetRequest(*waf.GetSizeConstraintSetInput) (*request.Request, *waf.GetSizeConstraintSetOutput)

	GetSqlInjectionMatchSet(*waf.GetSqlInjectionMatchSetInput) (*waf.GetSqlInjectionMatchSetOutput, error)
	GetSqlInjectionMatchSetWithContext(aws.Context, *waf.GetSqlInjectionMatchSetInput, ...request.Option) (*waf.GetSqlInjectionMatchSetOutput, error)
	GetSqlInjectionMatchSetRequest(*waf.GetSqlInjectionMatchSetInput) (*request.Request, *waf.GetSqlInjectionMatchSetOutput)

	GetWebACL(*waf.GetWebACLInput) (*waf.GetWebACLOutput, error)
	GetWebACLWithContext(aws.Context, *waf.GetWebACLInput, ...request.Option) (*waf.GetWebACLOutput, error)
	GetWebACLRequest(*waf.GetWebACLInput) (*request.Request, *waf.GetWebACLOutput)

	GetWebACLForResource(*wafregional.GetWebACLForResourceInput) (*wafregional.GetWebACLForResourceOutput, error)
	GetWebACLForResourceWithContext(aws.Context, *wafregional.GetWebACLForResourceInput, ...request.Option) (*wafregional.GetWebACLForResourceOutput, error)
	GetWebACLForResourceRequest(*wafregional.GetWebACLForResourceInput) (*request.Request, *wafregional.GetWebACLForResourceOutput)

	GetXssMatchSet(*waf.GetXssMatchSetInput) (*waf.GetXssMatchSetOutput, error)
	GetXssMatchSetWithContext(aws.Context, *waf.GetXssMatchSetInput, ...request.Option) (*waf.GetXssMatchSetOutput, error)
	GetXssMatchSetRequest(*waf.GetXssMatchSetInput) (*request.Request, *waf.GetXssMatchSetOutput)

	ListByteMatchSets(*waf.ListByteMatchSetsInput) (*waf.ListByteMatchSetsOutput, error)
	ListByteMatchSetsWithContext(aws.Context, *waf.ListByteMatchSetsInput, ...request.Option) (*waf.ListByteMatchSetsOutput, error)
	ListByteMatchSetsRequest(*waf.ListByteMatchSetsInput) (*request.Request, *waf.ListByteMatchSetsOutput)

	ListIPSets(*waf.ListIPSetsInput) (*waf.ListIPSetsOutput, error)
	ListIPSetsWithContext(aws.Context, *waf.ListIPSetsInput, ...request.Option) (*waf.ListIPSetsOutput, error)
	ListIPSetsRequest(*waf.ListIPSetsInput) (*request.Request, *waf.ListIPSetsOutput)

	ListResourcesForWebACL(*wafregional.ListResourcesForWebACLInput) (*wafregional.ListResourcesForWebACLOutput, error)
	ListResourcesForWebACLWithContext(aws.Context, *wafregional.ListResourcesForWebACLInput, ...request.Option) (*wafregional.ListResourcesForWebACLOutput, error)
	ListResourcesForWebACLRequest(*wafregional.ListResourcesForWebACLInput) (*request.Request, *wafregional.ListResourcesForWebACLOutput)

	ListRules(*waf.ListRulesInput) (*waf.ListRulesOutput, error)
	ListRulesWithContext(aws.Context, *waf.ListRulesInput, ...request.Option) (*waf.ListRulesOutput, error)
	ListRulesRequest(*waf.ListRulesInput) (*request.Request, *waf.ListRulesOutput)

	ListSizeConstraintSets(*waf.ListSizeConstraintSetsInput) (*waf.ListSizeConstraintSetsOutput, error)
	ListSizeConstraintSetsWithContext(aws.Context, *waf.ListSizeConstraintSetsInput, ...request.Option) (*waf.ListSizeConstraintSetsOutput, error)
	ListSizeConstraintSetsRequest(*waf.ListSizeConstraintSetsInput) (*request.Request, *waf.ListSizeConstraintSetsOutput)

	ListSqlInjectionMatchSets(*waf.ListSqlInjectionMatchSetsInput) (*waf.ListSqlInjectionMatchSetsOutput, error)
	ListSqlInjectionMatchSetsWithContext(aws.Context, *waf.ListSqlInjectionMatchSetsInput, ...request.Option) (*waf.ListSqlInjectionMatchSetsOutput, error)
	ListSqlInjectionMatchSetsRequest(*waf.ListSqlInjectionMatchSetsInput) (*request.Request, *waf.ListSqlInjectionMatchSetsOutput)

	ListWebACLs(*waf.ListWebACLsInput) (*waf.ListWebACLsOutput, error)
	ListWebACLsWithContext(aws.Context, *waf.ListWebACLsInput, ...request.Option) (*waf.ListWebACLsOutput, error)
	ListWebACLsRequest(*waf.ListWebACLsInput) (*request.Request, *waf.ListWebACLsOutput)

	ListXssMatchSets(*waf.ListXssMatchSetsInput) (*waf.ListXssMatchSetsOutput, error)
	ListXssMatchSetsWithContext(aws.Context, *waf.ListXssMatchSetsInput, ...request.Option) (*waf.ListXssMatchSetsOutput, error)
	ListXssMatchSetsRequest(*waf.ListXssMatchSetsInput) (*request.Request, *waf.ListXssMatchSetsOutput)

	UpdateByteMatchSet(*waf.UpdateByteMatchSetInput) (*waf.UpdateByteMatchSetOutput, error)
	UpdateByteMatchSetWithContext(aws.Context, *waf.UpdateByteMatchSetInput, ...request.Option) (*waf.UpdateByteMatchSetOutput, error)
	UpdateByteMatchSetRequest(*waf.UpdateByteMatchSetInput) (*request.Request, *waf.UpdateByteMatchSetOutput)

	UpdateIPSet(*waf.UpdateIPSetInput) (*waf.UpdateIPSetOutput, error)
	UpdateIPSetWithContext(aws.Context, *waf.UpdateIPSetInput, ...request.Option) (*waf.UpdateIPSetOutput, error)
	UpdateIPSetRequest(*waf.UpdateIPSetInput) (*request.Request, *waf.UpdateIPSetOutput)

	UpdateRule(*waf.UpdateRuleInput) (*waf.UpdateRuleOutput, error)
	UpdateRuleWithContext(aws.Context, *waf.UpdateRuleInput, ...request.Option) (*waf.UpdateRuleOutput, error)
	UpdateRuleRequest(*waf.UpdateRuleInput) (*request.Request, *waf.UpdateRuleOutput)

	UpdateSizeConstraintSet(*waf.UpdateSizeConstraintSetInput) (*waf.UpdateSizeConstraintSetOutput, error)
	UpdateSizeConstraintSetWithContext(aws.Context, *waf.UpdateSizeConstraintSetInput, ...request.Option) (*waf.UpdateSizeConstraintSetOutput, error)
	UpdateSizeConstraintSetRequest(*waf.UpdateSizeConstraintSetInput) (*request.Request, *waf.UpdateSizeConstraintSetOutput)

	UpdateSqlInjectionMatchSet(*waf.UpdateSqlInjectionMatchSetInput) (*waf.UpdateSqlInjectionMatchSetOutput, error)
	UpdateSqlInjectionMatchSetWithContext(aws.Context, *waf.UpdateSqlInjectionMatchSetInput, ...request.Option) (*waf.UpdateSqlInjectionMatchSetOutput, error)
	UpdateSqlInjectionMatchSetRequest(*waf.UpdateSqlInjectionMatchSetInput) (*request.Request, *waf.UpdateSqlInjectionMatchSetOutput)

	UpdateWebACL(*waf.UpdateWebACLInput) (*waf.UpdateWebACLOutput, error)
	UpdateWebACLWithContext(aws.Context, *waf.UpdateWebACLInput, ...request.Option) (*waf.UpdateWebACLOutput, error)
	UpdateWebACLRequest(*waf.UpdateWebACLInput) (*request.Request, *waf.UpdateWebACLOutput)

	UpdateXssMatchSet(*waf.UpdateXssMatchSetInput) (*waf.UpdateXssMatchSetOutput, error)
	UpdateXssMatchSetWithContext(aws.Context, *waf.UpdateXssMatchSetInput, ...request.Option) (*waf.UpdateXssMatchSetOutput, error)
	UpdateXssMatchSetRequest(*waf.UpdateXssMatchSetInput) (*request.Request, *waf.UpdateXssMatchSetOutput)
}

var _ WAFRegionalAPI = (*wafregional.WAFRegional)(nil)
