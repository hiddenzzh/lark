// Code generated by lark_sdk_gen. DO NOT EDIT.
/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Application_Sample_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		cli.Mock().MockGetAppAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.Application

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CheckUserIsInApplicationPaidScope(ctx, &lark.CheckUserIsInApplicationPaidScopeReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplication(ctx, &lark.GetApplicationReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationAppAdminUserList(ctx, &lark.GetApplicationAppAdminUserListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationAppList(ctx, &lark.GetApplicationAppListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationAppVisibility(ctx, &lark.GetApplicationAppVisibilityReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationFeedbackList(ctx, &lark.GetApplicationFeedbackListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationOrder(ctx, &lark.GetApplicationOrderReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationOrderList(ctx, &lark.GetApplicationOrderListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationUnderAuditList(ctx, &lark.GetApplicationUnderAuditListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationUsageOverview(ctx, &lark.GetApplicationUsageOverviewReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationUsageTrend(ctx, &lark.GetApplicationUsageTrendReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationUserAdminScope(ctx, &lark.GetApplicationUserAdminScopeReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationUserVisibleApp(ctx, &lark.GetApplicationUserVisibleAppReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationVersion(ctx, &lark.GetApplicationVersionReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.IsApplicationUserAdmin(ctx, &lark.IsApplicationUserAdminReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateApplication(ctx, &lark.UpdateApplicationReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateApplicationAppVisibility(ctx, &lark.UpdateApplicationAppVisibilityReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateApplicationFeedback(ctx, &lark.UpdateApplicationFeedbackReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateApplicationVersion(ctx, &lark.UpdateApplicationVersionReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

	})

	t.Run("request mock failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.Application

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApplicationCheckUserIsInApplicationPaidScope(func(ctx context.Context, request *lark.CheckUserIsInApplicationPaidScopeReq, options ...lark.MethodOptionFunc) (*lark.CheckUserIsInApplicationPaidScopeResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApplicationCheckUserIsInApplicationPaidScope()

			_, _, err := moduleCli.CheckUserIsInApplicationPaidScope(ctx, &lark.CheckUserIsInApplicationPaidScopeReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApplicationGetApplication(func(ctx context.Context, request *lark.GetApplicationReq, options ...lark.MethodOptionFunc) (*lark.GetApplicationResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApplicationGetApplication()

			_, _, err := moduleCli.GetApplication(ctx, &lark.GetApplicationReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApplicationGetApplicationAppAdminUserList(func(ctx context.Context, request *lark.GetApplicationAppAdminUserListReq, options ...lark.MethodOptionFunc) (*lark.GetApplicationAppAdminUserListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApplicationGetApplicationAppAdminUserList()

			_, _, err := moduleCli.GetApplicationAppAdminUserList(ctx, &lark.GetApplicationAppAdminUserListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApplicationGetApplicationAppList(func(ctx context.Context, request *lark.GetApplicationAppListReq, options ...lark.MethodOptionFunc) (*lark.GetApplicationAppListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApplicationGetApplicationAppList()

			_, _, err := moduleCli.GetApplicationAppList(ctx, &lark.GetApplicationAppListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApplicationGetApplicationAppVisibility(func(ctx context.Context, request *lark.GetApplicationAppVisibilityReq, options ...lark.MethodOptionFunc) (*lark.GetApplicationAppVisibilityResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApplicationGetApplicationAppVisibility()

			_, _, err := moduleCli.GetApplicationAppVisibility(ctx, &lark.GetApplicationAppVisibilityReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApplicationGetApplicationFeedbackList(func(ctx context.Context, request *lark.GetApplicationFeedbackListReq, options ...lark.MethodOptionFunc) (*lark.GetApplicationFeedbackListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApplicationGetApplicationFeedbackList()

			_, _, err := moduleCli.GetApplicationFeedbackList(ctx, &lark.GetApplicationFeedbackListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApplicationGetApplicationOrder(func(ctx context.Context, request *lark.GetApplicationOrderReq, options ...lark.MethodOptionFunc) (*lark.GetApplicationOrderResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApplicationGetApplicationOrder()

			_, _, err := moduleCli.GetApplicationOrder(ctx, &lark.GetApplicationOrderReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApplicationGetApplicationOrderList(func(ctx context.Context, request *lark.GetApplicationOrderListReq, options ...lark.MethodOptionFunc) (*lark.GetApplicationOrderListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApplicationGetApplicationOrderList()

			_, _, err := moduleCli.GetApplicationOrderList(ctx, &lark.GetApplicationOrderListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApplicationGetApplicationUnderAuditList(func(ctx context.Context, request *lark.GetApplicationUnderAuditListReq, options ...lark.MethodOptionFunc) (*lark.GetApplicationUnderAuditListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApplicationGetApplicationUnderAuditList()

			_, _, err := moduleCli.GetApplicationUnderAuditList(ctx, &lark.GetApplicationUnderAuditListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApplicationGetApplicationUsageOverview(func(ctx context.Context, request *lark.GetApplicationUsageOverviewReq, options ...lark.MethodOptionFunc) (*lark.GetApplicationUsageOverviewResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApplicationGetApplicationUsageOverview()

			_, _, err := moduleCli.GetApplicationUsageOverview(ctx, &lark.GetApplicationUsageOverviewReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApplicationGetApplicationUsageTrend(func(ctx context.Context, request *lark.GetApplicationUsageTrendReq, options ...lark.MethodOptionFunc) (*lark.GetApplicationUsageTrendResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApplicationGetApplicationUsageTrend()

			_, _, err := moduleCli.GetApplicationUsageTrend(ctx, &lark.GetApplicationUsageTrendReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApplicationGetApplicationUserAdminScope(func(ctx context.Context, request *lark.GetApplicationUserAdminScopeReq, options ...lark.MethodOptionFunc) (*lark.GetApplicationUserAdminScopeResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApplicationGetApplicationUserAdminScope()

			_, _, err := moduleCli.GetApplicationUserAdminScope(ctx, &lark.GetApplicationUserAdminScopeReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApplicationGetApplicationUserVisibleApp(func(ctx context.Context, request *lark.GetApplicationUserVisibleAppReq, options ...lark.MethodOptionFunc) (*lark.GetApplicationUserVisibleAppResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApplicationGetApplicationUserVisibleApp()

			_, _, err := moduleCli.GetApplicationUserVisibleApp(ctx, &lark.GetApplicationUserVisibleAppReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApplicationGetApplicationVersion(func(ctx context.Context, request *lark.GetApplicationVersionReq, options ...lark.MethodOptionFunc) (*lark.GetApplicationVersionResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApplicationGetApplicationVersion()

			_, _, err := moduleCli.GetApplicationVersion(ctx, &lark.GetApplicationVersionReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApplicationIsApplicationUserAdmin(func(ctx context.Context, request *lark.IsApplicationUserAdminReq, options ...lark.MethodOptionFunc) (*lark.IsApplicationUserAdminResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApplicationIsApplicationUserAdmin()

			_, _, err := moduleCli.IsApplicationUserAdmin(ctx, &lark.IsApplicationUserAdminReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApplicationUpdateApplication(func(ctx context.Context, request *lark.UpdateApplicationReq, options ...lark.MethodOptionFunc) (*lark.UpdateApplicationResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApplicationUpdateApplication()

			_, _, err := moduleCli.UpdateApplication(ctx, &lark.UpdateApplicationReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApplicationUpdateApplicationAppVisibility(func(ctx context.Context, request *lark.UpdateApplicationAppVisibilityReq, options ...lark.MethodOptionFunc) (*lark.UpdateApplicationAppVisibilityResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApplicationUpdateApplicationAppVisibility()

			_, _, err := moduleCli.UpdateApplicationAppVisibility(ctx, &lark.UpdateApplicationAppVisibilityReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApplicationUpdateApplicationFeedback(func(ctx context.Context, request *lark.UpdateApplicationFeedbackReq, options ...lark.MethodOptionFunc) (*lark.UpdateApplicationFeedbackResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApplicationUpdateApplicationFeedback()

			_, _, err := moduleCli.UpdateApplicationFeedback(ctx, &lark.UpdateApplicationFeedbackReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockApplicationUpdateApplicationVersion(func(ctx context.Context, request *lark.UpdateApplicationVersionReq, options ...lark.MethodOptionFunc) (*lark.UpdateApplicationVersionResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApplicationUpdateApplicationVersion()

			_, _, err := moduleCli.UpdateApplicationVersion(ctx, &lark.UpdateApplicationVersionReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

	})

	t.Run("response is failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.Application

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CheckUserIsInApplicationPaidScope(ctx, &lark.CheckUserIsInApplicationPaidScopeReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplication(ctx, &lark.GetApplicationReq{
				AppID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationAppAdminUserList(ctx, &lark.GetApplicationAppAdminUserListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationAppList(ctx, &lark.GetApplicationAppListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationAppVisibility(ctx, &lark.GetApplicationAppVisibilityReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationFeedbackList(ctx, &lark.GetApplicationFeedbackListReq{
				AppID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationOrder(ctx, &lark.GetApplicationOrderReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationOrderList(ctx, &lark.GetApplicationOrderListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationUnderAuditList(ctx, &lark.GetApplicationUnderAuditListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationUsageOverview(ctx, &lark.GetApplicationUsageOverviewReq{
				AppID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationUsageTrend(ctx, &lark.GetApplicationUsageTrendReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationUserAdminScope(ctx, &lark.GetApplicationUserAdminScopeReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationUserVisibleApp(ctx, &lark.GetApplicationUserVisibleAppReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationVersion(ctx, &lark.GetApplicationVersionReq{
				AppID:     "x",
				VersionID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.IsApplicationUserAdmin(ctx, &lark.IsApplicationUserAdminReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateApplication(ctx, &lark.UpdateApplicationReq{
				AppID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateApplicationAppVisibility(ctx, &lark.UpdateApplicationAppVisibilityReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateApplicationFeedback(ctx, &lark.UpdateApplicationFeedbackReq{
				AppID:      "x",
				FeedbackID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateApplicationVersion(ctx, &lark.UpdateApplicationVersionReq{
				AppID:     "x",
				VersionID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

	})

	t.Run("fake request is failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.Application
		cli.Mock().MockRawRequest(func(ctx context.Context, req *lark.RawRequestReq, resp interface{}) (response *lark.Response, err error) {
			return nil, fmt.Errorf("fake raw request")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CheckUserIsInApplicationPaidScope(ctx, &lark.CheckUserIsInApplicationPaidScopeReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplication(ctx, &lark.GetApplicationReq{
				AppID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationAppAdminUserList(ctx, &lark.GetApplicationAppAdminUserListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationAppList(ctx, &lark.GetApplicationAppListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationAppVisibility(ctx, &lark.GetApplicationAppVisibilityReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationFeedbackList(ctx, &lark.GetApplicationFeedbackListReq{
				AppID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationOrder(ctx, &lark.GetApplicationOrderReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationOrderList(ctx, &lark.GetApplicationOrderListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationUnderAuditList(ctx, &lark.GetApplicationUnderAuditListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationUsageOverview(ctx, &lark.GetApplicationUsageOverviewReq{
				AppID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationUsageTrend(ctx, &lark.GetApplicationUsageTrendReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationUserAdminScope(ctx, &lark.GetApplicationUserAdminScopeReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationUserVisibleApp(ctx, &lark.GetApplicationUserVisibleAppReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetApplicationVersion(ctx, &lark.GetApplicationVersionReq{
				AppID:     "x",
				VersionID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.IsApplicationUserAdmin(ctx, &lark.IsApplicationUserAdminReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateApplication(ctx, &lark.UpdateApplicationReq{
				AppID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateApplicationAppVisibility(ctx, &lark.UpdateApplicationAppVisibilityReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateApplicationFeedback(ctx, &lark.UpdateApplicationFeedbackReq{
				AppID:      "x",
				FeedbackID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateApplicationVersion(ctx, &lark.UpdateApplicationVersionReq{
				AppID:     "x",
				VersionID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

	})
}
