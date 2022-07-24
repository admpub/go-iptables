package iptables

import "io"

// Run runs an iptables command with the given arguments, ignoring
// any stdout output
func (ipt *IPTables) Run(args ...string) error {
	return ipt.run(args...)
}

// runWithOutput runs an iptables command with the given arguments,
// writing any stdout output to the given writer
func (ipt *IPTables) RunWithOutput(args []string, stdout io.Writer) error {
	return ipt.runWithOutput(args, stdout)
}

// GetIptablesCommand returns the correct command for the given protocol, either "iptables" or "ip6tables".
func GetIptablesCommand(proto Protocol) string {
	return getIptablesCommand(proto)
}
