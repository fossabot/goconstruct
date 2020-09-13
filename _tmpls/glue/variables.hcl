variable "name" {
  type = string
  description = "The project's name."
}

variable "readme_installing" {
  type = string
  description = "The README's installing section."
}

variable "readme_usage" {
  type = string
  description = "The README's usage section."
}

variable "readme_developing" {
  type = string
  description = "The README's developing section."
}

variable "module_name" {
  type = string
  description = "The name of the Go module."
}

variable "build_commands" {
  type = list(string)
  description = "A list of commands for building the binary or binaries."
}
