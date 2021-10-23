resource "aws_iam_user" "cognito_admin" {
  name = "gographql-project-admin"
}

resource "aws_iam_access_key" "cognito_admin" {
  user = aws_iam_user.cognito_admin.name
}

resource "aws_iam_user_policy" "cognito_admin" {
  name = "cognito-admin"
  user = aws_iam_user.cognito_admin.name

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "cognito-identity:*",
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
EOF
}
