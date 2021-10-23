resource "aws_cognito_user_pool" "gographql-project" {
  name = "gographql-project"

  account_recovery_setting {
    recovery_mechanism {
      name     = "verified_email"
      priority = 1
    }
  }
}

resource "aws_cognito_user_pool_client" "gographql-project-client" {
  name = "gographql-project-client"

  user_pool_id = aws_cognito_user_pool.gographql-project.id
}