### 界面截图

![screen shot](https://github.com/aicodex/Knowledge-Fusion-Tagging-tools/blob/master/screen_shots/screen.png?raw=true)

### 使用方法

下载nodejs

使用如下命令安装cnpm

> npm install -g cnpm --registry=https://registry.npm.taobao.org

使用如下命令安装依赖

> cnpm install

执行complile_project来编译项目（修改package.json可以修改编译目标平台，如linux，ia32等）

执行run_test来直接运行

执行build_setup.bat可以打包win安装包（如果要打包32位的，需要先修改package.json，再修改build.js的文件夹名字）

[./split_tools](https://github.com/aicodex/Knowledge-Fusion-Tagging-tools/tree/master/split-tools)里面是分割工具，可以将数据按实体名尽量均分。

[./merge-tag-tool](https://github.com/aicodex/Knowledge-Fusion-Tagging-tools/tree/master/merge-tag-tool)里面是最终多人标注结果的合并、评分程序，是scala的命令。
