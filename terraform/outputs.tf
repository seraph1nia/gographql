output "aws_cognito_account_secret" {
  value = aws_iam_access_key.cognito_admin.secret
  sensitive = true
}

output "aws_cognito_account_id" {
  value = aws_iam_access_key.cognito_admin.id
  sensitive = true
}

output "aws_cognito_client_id" {
  value = aws_cognito_user_pool_client.gographql-project-client.id
}

output "aws_user_pool_id" {
  value = aws_cognito_user_pool.gographql-project.id
}
