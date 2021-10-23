resource "aws_cognito_user_pool" "gographql-project" {
  name = "gographql-project"

  account_recovery_setting {
    recovery_mechanism {
      name     = "verified_email"
      priority = 1
    }
  }
}