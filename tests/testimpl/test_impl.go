package testimpl

import (
	"context"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// parseRegionFromTopicARN extracts the region from an SNS topic ARN.
// Format: arn:aws:sns:REGION:ACCOUNT_ID:TOPIC_NAME
func parseRegionFromTopicARN(arn string) string {
	parts := strings.Split(strings.TrimSpace(arn), ":")
	if len(parts) >= 4 {
		return parts[3]
	}
	return "us-east-1"
}

func getSNSClient(t *testing.T, region string) *sns.Client {
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(region))
	require.NoError(t, err, "Failed to load AWS config")
	return sns.NewFromConfig(cfg)
}

func TestComposableComplete(t *testing.T, ctx types.TestContext) {
	t.Run("VerifyTerraformOutputs", func(t *testing.T) {
		opts := ctx.TerratestTerraformOptions()
		id := terraform.Output(t, opts, "id")
		arn := terraform.Output(t, opts, "arn")
		owner := terraform.Output(t, opts, "owner")

		assert.Equal(t, arn, id, "id should equal arn for aws_sns_topic_policy")
		assert.Regexp(t, `^\d{12}$`, owner, "owner must be a 12-digit AWS account ID")
	})

	t.Run("VerifyPolicyViaSNSAPI", func(t *testing.T) {
		opts := ctx.TerratestTerraformOptions()
		topicArn := strings.TrimSpace(terraform.Output(t, opts, "arn"))
		expectedOwner := terraform.Output(t, opts, "owner")

		client := getSNSClient(t, parseRegionFromTopicARN(topicArn))
		result, err := client.GetTopicAttributes(context.Background(), &sns.GetTopicAttributesInput{
			TopicArn: aws.String(topicArn),
		})
		require.NoError(t, err, "GetTopicAttributes must succeed")

		policy, ok := result.Attributes["Policy"]
		require.True(t, ok, "Policy attribute must be present")
		require.NotEmpty(t, policy, "Policy must not be empty")

		actualOwner, ok := result.Attributes["Owner"]
		require.True(t, ok, "Owner attribute must be present")
		assert.Equal(t, expectedOwner, actualOwner, "Owner must match Terraform output")
	})

	t.Run("VerifyPublishSucceeds", func(t *testing.T) {
		opts := ctx.TerratestTerraformOptions()
		topicArn := strings.TrimSpace(terraform.Output(t, opts, "arn"))

		client := getSNSClient(t, parseRegionFromTopicARN(topicArn))
		msg := "Terratest functional verification message"
		_, err := client.Publish(context.Background(), &sns.PublishInput{
			TopicArn: aws.String(topicArn),
			Message:  aws.String(msg),
		})
		require.NoError(t, err, "Publish must succeed to verify policy allows publish")
	})
}

func TestComposableCompleteReadonly(t *testing.T, ctx types.TestContext) {
	t.Run("VerifyTerraformOutputs", func(t *testing.T) {
		opts := ctx.TerratestTerraformOptions()
		id := terraform.Output(t, opts, "id")
		arn := terraform.Output(t, opts, "arn")
		owner := terraform.Output(t, opts, "owner")

		assert.Equal(t, arn, id, "id should equal arn for aws_sns_topic_policy")
		assert.Regexp(t, `^\d{12}$`, owner, "owner must be a 12-digit AWS account ID")
	})

	t.Run("VerifyPolicyViaSNSAPI", func(t *testing.T) {
		opts := ctx.TerratestTerraformOptions()
		topicArn := strings.TrimSpace(terraform.Output(t, opts, "arn"))
		expectedOwner := terraform.Output(t, opts, "owner")

		client := getSNSClient(t, parseRegionFromTopicARN(topicArn))
		result, err := client.GetTopicAttributes(context.Background(), &sns.GetTopicAttributesInput{
			TopicArn: aws.String(topicArn),
		})
		require.NoError(t, err, "GetTopicAttributes must succeed")

		policy, ok := result.Attributes["Policy"]
		require.True(t, ok, "Policy attribute must be present")
		require.NotEmpty(t, policy, "Policy must not be empty")

		actualOwner, ok := result.Attributes["Owner"]
		require.True(t, ok, "Owner attribute must be present")
		assert.Equal(t, expectedOwner, actualOwner, "Owner must match Terraform output")
	})
}
