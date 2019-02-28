# -*- coding: utf-8 -*-
import os
import datetime
import threading

 


def func():
    os.system("./hrun.sh parser-uq-log2")
    os.system("./hrun.sh parseLog2MysqlTool")
    timer = threading.Timer(86400, func)
    timer.start()


now_time = datetime.datetime.now()
next_time = now_time + datetime.timedelta(days=+1)
next_year = next_time.date().year
next_month = next_time.date().month
next_day = next_time.date().day

next_time = datetime.datetime.strptime(str(next_year)+"-"+str(next_month)+"-"+str(next_day)+" 01:00:00", "%Y-%m-%d %H:%M:%S")


timer_start_time = (next_time - now_time).total_seconds()
print(timer_start_time)


timer = threading.Timer(timer_start_time, func)
timer.start()
