variable "vm_name" {
  type = string
  default = "anka-packer-base-macos"
}

variable "vcpu_count" {
  type = string
  default = ""
}

variable "installer_app" {
  type = string
  default = "/Applications/Install macOS Big Sur.app/"
}

source "veertu-anka-vm-create" "anka-packer-base-macos" {
  installer_app = "${var.installer_app}"
  vm_name = "${var.vm_name}"
  vcpu_count = "${var.vcpu_count}"
  stop_vm = true
}

build {
  sources = [
    "source.veertu-anka-vm-create.anka-packer-base-macos"
  ]

  provisioner "shell" {
    inline = [
      "echo hello world",
      "echo llamas rock"
    ]
  }
}