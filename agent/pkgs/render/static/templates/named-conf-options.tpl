options {
        #bindkeys-file "/etc/named.iscdlv.key";
        directory "/var/cache/bind";
        #dump-file "/var/named/data/cache_dump.db";
        listen-on port 53 {
                "any";
        };
        #managed-keys-directory "/var/named/dynamic";
        #memstatistics-file "/var/named/data/named_mem_stats.txt";
        #statistics-file "/var/named/data/named_stats.txt";
        recursion no;
        allow-query {
                "any";
        };
        #also-notify {
        #        172.16.42.59 port 53;
        #};
        forward first;
        forwarders{
                202.106.0.20;
                114.114.114.114;
        };
};