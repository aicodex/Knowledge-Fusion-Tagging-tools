var electronInstaller = require('electron-winstaller');
var path = require("path");

resultPromise = electronInstaller.createWindowsInstaller({
    appDirectory: path.join('./OutApp/merger-tag-tool-win32-x64'), //打包文件的路径
    outputDirectory: path.join('./OutApp/'), //输出路径
    authors: 'xzp', // 作者名称
    description: 'Merger模块的标注工具。',
    exe: 'merger-tag-tool.exe', //在appDirectory寻找exe的名字
    setupIcon: './icon/icon-128.ico',
    noMsi: true
  });

resultPromise.then(() => console.log("success!"), (e) => console.log(`something wrong: ${e.message}`));