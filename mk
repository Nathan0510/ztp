/tool netwatch
add host=192.168.1.1 interval=10m name=ZTP type=simple up-script=\
    ":local SerialNumber [/system routerboard get serial-number]; /tool fetch u\
    rl=\"http://192.168.1.1:8080/ztp\" http-method=post http-content-type=appl\
    ication/json http-data=\"{\\\"SN\\\":\\\"\$SerialNumber\\\"}\";"
