package container

/*
	/proc文件系统: 由内核提供，包含系统运行时信息(系统内存,mount设备信息,硬件配置等).只存在于内存中，不存在于外存。以文件系统形式给访问内核数据操作提供接口
	很多系统工具都是去读这个文件系统的文件内容

	/proc/N pid为N的进程信息
	/proc/N/cmdline 进程启动命令
	/proc/N/cwd 链接到进程当前工作目录
	/proc/N/environ 进程环境变量列表
	/proc/N/exe 链接到进程执行命令文件
	/proc/N/fd 包含进程相关的所有文件描述符
	/proc/N/maps 与进程相关的内存映射信息
	/proc/N/mem 进程持有的内存，不可读
	/proc/N/root 链接到进程的根目录
	/proc/N/stat 进程的状态
	/proc/N/statm 进程使用的内存状态
	/proc/N/status 进程状态信息,比/stat/statm更具有可读性
	/proc/self/ 链接到当前正在运行的进程
*/
