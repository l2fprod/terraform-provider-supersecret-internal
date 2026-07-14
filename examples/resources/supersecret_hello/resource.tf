resource "supersecret_hello" "example" {
  name = "world"
}

output "greeting" {
  value = supersecret_hello.example.greeting
}
