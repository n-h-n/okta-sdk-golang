/*
 * Copyright 2018 - Present Okta, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package integration

import (
	"context"
	"testing"

	"github.com/n-h-n/okta-sdk-golang/v2/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListRoleSubscriptions(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	subscriptions, _, err := client.Subscription.ListRoleSubscriptions(ctx, "SUPER_ADMIN")
	require.NoError(t, err, "Getting subscriptions for a role should not error")

	assert.True(t, len(subscriptions) > 0, "There should be subscriptions")
}

func TestGetRoleSubscriptionByNotificationType(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	subscription, _, err := client.Subscription.GetRoleSubscriptionByNotificationType(ctx, "SUPER_ADMIN", "OKTA_ANNOUNCEMENT")
	require.NoError(t, err, "Getting subscription for a role by a notification type should not error")

	assert.True(t, subscription.NotificationType == "OKTA_ANNOUNCEMENT", "There should be a subscription")
}

func TestSubscribeRoleSubscriptionByNotificationType(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	_, err = client.Subscription.SubscribeRoleSubscriptionByNotificationType(ctx, "SUPER_ADMIN", "OKTA_ANNOUNCEMENT")
	require.NoError(t, err, "Subscribing for a role by a notification type should not error")
}

func TestUnsubscribeRoleSubscriptionByNotificationType(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	_, err = client.Subscription.UnsubscribeRoleSubscriptionByNotificationType(ctx, "SUPER_ADMIN", "OKTA_ANNOUNCEMENT")
	require.NoError(t, err, "Unsubscribing for a role by a notification type should not error")
}

func TestSubscribeUserSubscriptionByNotificationType(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	user, _, err := client.User.GetUser(ctx, "me")
	require.NoError(t, err, "Getting the current user should not error")

	_, err = client.Subscription.SubscribeUserSubscriptionByNotificationType(ctx, user.Id, "OKTA_ANNOUNCEMENT")
	require.NoError(t, err, "Subscribing user by a notification type should not error")
}

func TestUnsubscribeUserSubscriptionByNotificationType(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	user, _, err := client.User.GetUser(ctx, "me")
	require.NoError(t, err, "Getting the current user should not error")

	_, err = client.Subscription.UnsubscribeUserSubscriptionByNotificationType(ctx, user.Id, "OKTA_ANNOUNCEMENT")
	require.NoError(t, err, "Unsubscribing user by a notification type should not error")
}
