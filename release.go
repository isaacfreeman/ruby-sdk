package main

import (
	"github.com/flowcommerce/tools/executor"
)

func main() {
	executor := executor.Create("ruby-sdk")

	executor = executor.Add("rm -f ./flowcommerce-*.gem")
	executor = executor.Add("apidoc update")
	executor = executor.Add("git commit -m 'Update to latest API' lib/flow_commerce/flow*rb")
	executor = executor.Add("git fetch --tags origin")
	executor = executor.Add("dev tag")
	executor = executor.Add("gem build flowcommerce.gemspec")
	executor = executor.Add("gem push ./flowcommerce-*.gem")
	executor = executor.Add("rm -f ./flowcommerce-*.gem")

	executor.Run()
}
