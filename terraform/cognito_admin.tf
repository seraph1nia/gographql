resource "aws_iam_user" "cognito_admin" {
  name = "gographql-project-admin"
}

resource "aws_iam_access_key" "cognito_admin" {
  user = aws_iam_user.cognito_admin.name
}

resource "aws_iam_user_policy_attachment" "attach_cognito_power" {
  user       = aws_iam_user.cognito_admin.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonCognitoPowerUser"
}
