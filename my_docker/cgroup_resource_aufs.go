package my_docker

import "testing"

/*
	image layer内容存储在/var/lib/docker/aufs/diff  var/lib/docker/aufs/{diff（read-write layer）,mnt(mount目录),containerid（配置文件和metadata）} //docker pull的时候有四层layer，所以有四个文件
	containerId文件夹下有/tmp/newfile文件，存储了在layer层之上rw层的写入修改信息
	docker history可以看到镜像使用的image layer
系统aufs下也会多一个文件:/sys/fs/aufs/si_xxx/*,其中只有最上面的layer是rw权限
	如果要删除一个文件file1，rw层会生成.wh.file1隐藏所有readonly层的file1文件，从而实现文件删除
 */
 */
*/

func TestResourceAufs(t *testing.T) {

}
