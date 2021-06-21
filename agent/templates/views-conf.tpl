view "azure" {
    match-clients { azure; };
    include "/etc/bind/views/azure-zones.conf";

    #common config
    include "/etc/bind/views/common-zones.conf";
};

view "default" {
    match-clients { default; };
    include "/etc/bind/views/default-zones.conf";

    #common config
    include "/etc/bind/views/common-zones.conf";
};