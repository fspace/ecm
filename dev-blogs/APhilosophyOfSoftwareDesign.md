复杂性

修改 理解代码的难易程度 影响修改，理解系统的一切因素


代码的重量？（行数 心智负担|认知负荷（需要花多长时间来理解) 变更放大|扩散)

不知道的未知因素
变更代码后 你不知道还有什么需要变更 压根不知道这些东西的存在
未记录在文档（文档化）

对于变更放大 修改后你知道会级联修改哪些部分的 即便再多 也是知道的 就怕你不知道这些地方在哪里

> One of the most important goals of good design is for a system to be obvious. 
This is the opposite of high cognitive load and unknown unknowns.

### 原因
  Complexity is caused by two things: dependencies and obscurity.
  
 依赖  
- 修改a 必须修改b 那么b依赖a
- 网络协议  内容格式共识 发送了不同信息 接收者也必须做出调整 反之亦然
- 方法签名  影响实现者和调用者

任何时候你创建一个类 你就创建了一个API对他的依赖 我们设计软件的一个目标就是减少依赖？
并且使依赖尽量简单和明显

晦涩
- 泛化的命名 需要扫描代码才知道意思 类型（弱类型语言 还看不出来单位）
- 伴生于依赖 依赖存在的不明显   如 新的错误状态添加到系统 那么对其的消息也应该添加到一个table中 但消息表的存在对程序员并不明显
- 不一致性 同名变量用于不同目的

缺乏文档

如果系统有干净明确的设计 那么文档就少  
大量文档是坏信号

Dependencies lead to change amplification and a high cognitive load.
 Obscurity creates unknown unknowns, and also contributes to cognitive load.
 If we can find design techniques that minimize dependencies and obscurity, then we can reduce the complexity of software.
 
 ## 复杂度是递增出现的
 累计效应
 每个人都不管复杂度 那么累加完成的代码 可能很复杂
 零容忍  从每个人做起？
 
 
 ====================================================================================
 
 tactical programming
 尽快完成
 
 复杂度是累加的 你这样搞 别人也这样搞 然后...
 
 长远 vs 当下
 重构会影响当前任务的进度
 
 tactical tornado  功能实现很快的人 被视为编码英雄 但对长远发展的代码库 他就是挖坑者
