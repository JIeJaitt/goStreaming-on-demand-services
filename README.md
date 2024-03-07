# StmSrv: streaming server

> **Note**  
> 本项目目前仅用于学习

流媒体点播服务

![](image/Overall-architecture.png)

https://jiejaitt.blog.csdn.net/article/details/120277314

# LICENSE
[MIT](tps://github.com/JIeJaitt/stmsrv/blob/5aea553bd697a9906484eae470eac5b10123e9f8/LICENSE)




## 常见问题

首先我们先给大家再说一下什么是schedule。Schedule顾名思义的这个调度器，调度什么呢？调度一些任务调度什么样的任务呢？调度我们通过普通的SAPI，它没有办法马上给他结果的任务。那么这些任务都会分发到schedule里面。然后schedule里面，它会根据它的时间的period或者interval来定时给他触发，或者是延时触发无所谓，反正都是类似于这种异步的任务，我才需要用schedule来做的。

我们为什么需要schedule呢？这很顾名思义，因为我们系统中有异步的任务。就比如说在我们今天先要给大家讲述的这个项目里面，我们就有一个非常大的任务叫做延时删除视频。为什么要延时删除视频呢？因为有时候我们整个视频网站会有审的需求，或者是审查的需求，或者是一些数据恢复的需求。


### schedule服务的架构是怎么样的

<img width="591" alt="schedule服务的架构预览" src="https://github.com/JIeJaitt/goStreaming-on-demand-services/assets/77219045/91b1e6d6-a8f6-4a71-ac43-a06ed1348626">

这个就是我们的架构概览，整个大的框架有一个Timer来启动。然后我们的Timer里面有一个task runner ，task runner分为三部分，第一个是Dispatcher，然后第二个就是它的Excutor。实际上也就是我们之前所说的生产消费者Producer和Consumer。然后他们之间通过一个go的原生channel来通信。这样的话，Dispatcher会把他得到的任务内容通过channel发送给Excutor。Excutor就会去读它这些内容，然后去做一些操作。这个就是我们整个schedule我们需要做的什么，以及我们的架构是什么。








