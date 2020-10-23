variable "Name" {
  type = string
  description = "The root command's name."
}

variable "ModuleName" {
  type = string
  description = "The Go module's name."
}

variable "BuildCommand" {
  type = string
  description = "A command for building the binary or binaries."
}

variable "Installing" {
  type = string
  description = "The README's installing section."
}

variable "Usage" {
  type = string
  description = "The README's usage section."
}

variable "Developing" {
  type = string
  description = "The README's developing section."
}
