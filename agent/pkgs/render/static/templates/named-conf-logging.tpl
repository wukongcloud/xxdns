logging {
        channel "warning" {
                file "/var/log/dns_warnings" versions 3 size 52428800;
                severity warning;
                print-time yes;
                print-severity yes;
                print-category yes;
        };
        channel "general_dns" {
                file "/var/log/dns_logs" versions 9 size 104857600;
                severity info;
                print-time yes;
                print-severity yes;
                print-category yes;
        };
        channel "default_debug" {
                file "/var/log/named.run" versions 9 size 20971520;
                severity dynamic;
        };
        category "default" {
                "warning";
        };
        category "security" {
                "warning";
        };
        category "database" {
                "warning";
        };
        category "queries" {
                "general_dns";
        };
        category "general" {
                "general_dns";
        };
        category "config" {
                "general_dns";
        };
        category "xfer-in" {
                "general_dns";
        };
        category "xfer-out" {
                "general_dns";
        };
        category "notify" {
                "general_dns";
        };
};