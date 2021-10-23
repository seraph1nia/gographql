output "aws_cognito_secret" {
  value = aws_iam_access_key.cognito_admin.secret
  sensitive = true
}

output "aws_cognito_id" {
  value = aws_iam_access_key.cognito_admin.id
  sensitive = true
}
