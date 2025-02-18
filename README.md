# 自动发布工具
## 目前仅支持录制发布
作者在录制组，这是强需求

## 特点
小巧：编译优化后不到 8MB 体积

自动化：尽可能减少人力工作

跨平台：同时支持Windows/Linux


## 功能 
- 一键发布录制内容

## 安装
依赖：ffmpeg mediainfo

下载并安装：

ffmpeg 下载地址: https://www.ffmpeg.org/download.html

mediainfo 下载地址：https://mediaarea.net/en/MediaInfo

安装好之后将 ffmpeg的bin目录 和 MediaInfo.exe 所在目录加入环境变量

编译本项目,得到二进制文件: pt-auto.exe

在 pt-auto.exe 同级目录新建 manifest/config/config.yaml文件，填写以下内容:

```yaml
tools:
  maxPicNum: 4 # 暂时没用,默认截图 4 张
  torrentMapVideoDirEnabled: true # 是否启用视频文件独立目录,若启用发布完成后会自动将视频文件复制到 torrentMapVideoDir 指定的目录
  torrentMapVideoDir: "D:/Downloads/做种" #  视频文件独立目录
  rapidapi: # 接入google翻译的api, 申请网址: https://rapidapi.com/, 免费用户一个月可以翻译 1000 次
    host: ''
    key: ''
  mteam: # 以下配置咨询管理员
    sourceId: '5' # HDTV/TV
    teamId: '43' # 测试环境使用时,请删除此id
    teamName: 'TPTV'
    categoryId: '402' # categoryId 影劇/綜藝/HD
    URL: '' 
    apiKey: '' 
    uploadImgKey: '' 
    # URL: ''
    # apiKey: ''
    # uploadImgKey: ''
  proxy:
    enabled: true # 是否启用代理访问api
    addr: 'http://127.0.0.1:10810' # 代理地址
```

## 使用方法

将下载好的文件修改好全称，如下所示：
```text
D:\torrent\work\闹元宵.ts
D:\torrent\work\2025内蒙古电视台春节联欢晚会.ts
D:\torrent\work\哈尔滨2025年亚冬会-闭幕式.ts
D:\torrent\work\2025元宵奇妙游.ts
```
执行命令
```shell
pt-auto -i D:\torrent\work\闹元宵.ts
```
程序会发布一个文件

也可以使整个目录同时发布
```shell
pt-auto -i D:\torrent\work
```
程序会自动扫描目录内的所有文件，并进行做种发布

### 指定特别字段值
格式：常规名##年份##副标题##英文名.ts

例如：哈尔滨2025年亚冬会-闭幕式.ts 文件自动翻译出来会有些奇怪

自动翻译：Harbin 2025 Asian Winter Games

专业翻译：The 9th Asian Winter Games Harbin 2025

指定如下所示
```shell
哈尔滨2025年亚冬会-闭幕式##2025##哈尔滨2025年亚冬会-闭幕式##The 9th Asian Winter Games Harbin 2025.ts
```

### 缺省值
假如说只想指定某一个值，可以使用"-"减号占位，比如：
```shell
哈尔滨2025年亚冬会-闭幕式##-##-##The 9th Asian Winter Games Harbin 2025.ts
```
除了中文名不能省略,其余都可以省略，若省略值后面不再有值，则可以忽略当前值减号占位都不需要，这样命名就很灵活，比如：

当使用占位符时，被占的位置默认值如下：

年份：取当前年份

副标题：与文件名一致

英文名：调用google翻译接口，将中文文名翻译后得到

```shell
哈尔滨2025年亚冬会-闭幕式.ts
哈尔滨2025年亚冬会-闭幕式##2025.ts
哈尔滨2025年亚冬会-闭幕式##2025##冬亚会闭幕式.ts
哈尔滨2025年亚冬会-闭幕式##2025##-##The 9th Asian Winter Games Harbin 2025.ts
哈尔滨2025年亚冬会-闭幕式##-##-##The 9th Asian Winter Games Harbin 2025.ts
```

以上命名，就是可以准确识别

这样发布后的固定格式：

标题：The 9Th Asian Winter Games Harbin 2025 Closing Ceremony 2025 1080i H264-TPTV

副标题：哈尔滨2025年亚冬会-闭幕式

MediaInfo: ...很多内容

"The 9Th Asian Winter Games Harbin": 来自指定名称或翻译

2025: 来自指定或自动取当前年

1080i: 来自mediainfo自动识别

H264: 来自mediainfo自动识别

TPTV: 来自配置文件 tools.mteam.teamName

副标题：来自文件名或者指定

MediaInfo: 来自mediainfo.exe识别,通过程序解析后上传到各字段

简介：使用ffmpeg自动截图并上传

# 不完善的地方
1.作者没有官方源码不知道mediainfo内容解析方法与解析度/视频编码/音频编码之间的对应关系，所以这一块内容作者遇到的都能解析对，没有遇到的就留空了，需要手动补齐

2.作者没有发布过其他内容，暂不支持其他内容