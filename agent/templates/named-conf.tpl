// This is the primary configuration file for the BIND DNS server named.
//
// Please read /usr/share/doc/bind9/README.Debian.gz for information on the
// structure of BIND configuration files in Debian, *BEFORE* you customize
// this configuration file.

include "/etc/bind/named.conf.options";
include "/etc/bind/named.conf.logging";
include "/etc/bind/cidrs.conf";
include "/etc/bind/views.conf";