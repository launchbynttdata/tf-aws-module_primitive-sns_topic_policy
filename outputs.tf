// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

output "id" {
  description = "The ID of the resource (same as the topic ARN)."
  value       = aws_sns_topic_policy.topic_policy.id
}

output "arn" {
  description = "The ARN of the SNS topic."
  value       = aws_sns_topic_policy.topic_policy.arn
}

output "owner" {
  description = "The AWS Account ID of the SNS topic owner."
  value       = aws_sns_topic_policy.topic_policy.owner
}
