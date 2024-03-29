复杂性
-------

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
 
 ========================================================================================
 
 Modules Should Be Deep
 
 In modular design, a software system is decomposed into a collection of modules that are relatively independent.
  Modules can take many forms, such as classes, subsystems, or services. In an ideal world, each module would be
   completely independent of the others: a developer could work in any of the modules without knowing anything about 
   any of the other modules. In this world, the complexity of a system would be the complexity of its worst module.
   
Unfortunately, this ideal is not achievable. Modules must work together by calling each others’s functions or methods.
 As a result, modules must know something about each other. There will be dependencies between the modules:
 
 In order to manage dependencies, we think of each module in two parts: an interface and an implementation. 
 The interface consists of everything that a developer working in a different module must know in order to use the given module.
  Typically, the interface describes what the module does but not how it does it.
  
  The implementation consists of the code that carries out the promises
  made by the interface. A developer working in a particular module must understand the interface and implementation of
   that module, plus the interfaces of any other modules invoked by the given module. A developer should not need to
    understand the implementations of modules other than the one he or she is working in.
    
   >The best modules are deep: they have a lot of functionality hidden behind a simple interface. A deep module is a good
     abstraction because only a small fraction of its internal complexity is visible to its users.

Deep modules such as Unix I/O (5个函数签名 )and garbage collectors （0个方法签名） provide powerful abstractions because they are easy to use,
 yet they hide significant implementation complexity.


接口复杂度 width
实现复杂度 height
二者形成一个长方形

浅模块|类 可能是不好的信号
...

通用 vs 特定
--------------------------------

模块的接口尽量考虑通用  实现可以特定于某个特殊的当下需求

随着时间的推移 通用性接口 柔性更强
  认知负担比较低 简单 复杂度比较低
  
什么样的接口是通用性|普遍性接口？
> What is the simplest interface that will cover all my current needs? If you reduce the number of methods in an API
 without reducing its overall capabilities, then you are probably creating more general-purpose methods.

pass through
-----------------------
- Pass-through methods

...

装饰器是否有其他考虑实现

- pass-through variable.  

多层传递 但只有特定层才需要

全局对象（X）

> but global variables almost always create other problems. For example, global variables make it impossible to create 
two independent instances of the same system in the same process, since accesses to the global variables will conflict. 
It may seem unlikely that you would need multiple instances in production, but they are often useful in testing.

上下文对象 （ok）
> A context stores all of the application’s global state (anything that would otherwise be a pass-through variable or
 global variable). Most applications have multiple variables in their global state, representing things such as 
 configuration options, shared subsystems, and performance counters. There is one context object per instance of the 
 system. The context allows multiple instances of the system to coexist in a single process, each with its own context.

 >Unfortunately, the context will probably be needed in many places, so it can potentially become a pass-through variable.
  To reduce the number of methods that must be aware of it, a reference to the context can be saved in most of the
   system’s major objects.
   ...
   When a new object is created, the creating method retrieves the context reference from its object and passes it to
    the constructor for the new object. With this approach, the context is available everywhere, but it only appears as 
    an explicit argument in constructors.
    The context object unifies the handling of all system-global information and eliminates the need for pass-through 
    variables. If a new variable needs to be added, it can be added to the context object; no existing code is affected 
    except for the constructor and destructor for the context. The context makes it easy to identify and manage the global
     state of the system, since it is all stored in one place. The context is also convenient for testing: test code can
      change the global configuration of the application by modifying fields in the context. It would be much more difficult 
      to implement such changes if the system used pass-through variables.
      ...
      不完美也没有更好办法了：
      Contexts are far from an ideal solution. The variables stored in a context have most of the disadvantages of global
       variables; for example, it may not be obvious why a particular variable is present, or where it is used. Without
        discipline, a context can turn into a huge grab-bag of data that creates nonobvious dependencies throughout the 
        system. Contexts may also create thread-safety issues; the best way to avoid problems is for variables in a context 
        to be immutable. Unfortunately, I haven’t found a better solution than contexts.

pulls complexity downward.
-------------------------

复杂度下移  
远离使用者 移到下层类|结构|依赖

Configuration parameters are an example of moving complexity upwards instead of down.
程序能计算的就不要暴露给用户去配置


分好 还是合好
-------------------

One of the most fundamental questions in software design is this: given two pieces of functionality, should they be 
implemented together in the same place, or should their implementations be separated? This question applies at all 
levels in a system, such as functions, methods, classes, and services.

When deciding whether to combine or separate, the goal is to reduce the complexity of the system as a whole and 
improve its modularity.
### 合
- Bring together if information is shared

- Bring together if it will simplify the interface

- Bring together to eliminate duplication

### 分
- Separate general-purpose and special-purpose code

Special-purpose code associated with a general-purpose mechanism should normally go in a different module 
(typically one associated with the particular purpose).

In general, the lower layers of a system tend to be more general-purpose and the upper layers more special-purpose. 

When you encounter a class that includes both general-purpose and special-purpose features for the same abstraction,
 see if the class can be separated into two classes, one containing the general-purpose features, 
 and the other layered on top of it to provide the special-purpose features.
 
 Splitting and joining methods
 -------------
 
 Each method should do one thing and do it completely.
 
 Define Errors Out Of Existence
 ----------------
 Exception handling is one of the worst sources of complexity in software systems.
 错误|异常 处理 带来了复杂性
 A particular piece of code may encounter exceptions in several different ways:
 • A caller may provide bad arguments or configuration information.
 • An invoked method may not be able to complete a requested operation. For example, an I/O operation may fail, or a required resource may not be available.
 • In a distributed system, network packets may be lost or delayed, servers may not respond in a timely fashion, 
 or peers may communicate in unexpected ways.
 • The code may detect bugs, internal inconsistencies, or situations it is not prepared to handle.
 
 When an exception occurs, the programmer can deal with it in two ways, each of which can be complicated. The
 first approach is to move forward and complete the work in progress in spite of the exception. For example, 
 if a network packet is lost, it can be resent; if data is corrupted, perhaps it can be recovered from a redundant copy.
  The second approach is to abort the operation in progress and report the exception upwards. However, aborting can be 
  complicated because the exception may have occurred at a point where system state is inconsistent (a data structure 
  might have been partially initialized); the exception handling code must restore consistency, such as by unwinding any
   changes made before the exception occurred.
   
The exceptions thrown by a class are part of its interface; **classes with lots of exceptions have complex interfaces,
 and they are shallower than classes with fewer exceptions.**   An exception is a particularly complex element of an
 interface. It can propagate up through several stack levels before being caught, so it affects not just the method’s 
 caller, but potentially also higher-level callers (and their interfaces).
 
 Throwing exceptions is easy; handling them is hard. Thus, the complexity of exceptions comes from the exception 
 handling code. The best way to **reduce the complexity damage caused by exception handling is to reduce the number 
 of places where exceptions have to be handled.**
 
 The best way to eliminate exception handling complexity is to define your APIs so that there are no exceptions to handle:
  **define errors out of existence.**
  
  ### Mask exceptions
  掩盖异常
  Exception masking is an example of pulling complexity downward.
  
  ### Exception aggregation
  The idea behind exception aggregation is to handle many exceptions with a single piece of code;
  
  go 有个库： https://godoc.org/go.uber.org/multierr
  
  
  With exceptions, as with many other areas in software design, you must determine what is important and what is not
   important. Things that are not important should be hidden, and the more of them the better. But when something 
   is important, it must be exposed.
   
   ## 11 Design it Twice
   
   Try to pick approaches that are radically different from each other; you’ll
   learn more that way. Even if you are certain that there is only one reasonable approach, consider a second design 
   anyway, no matter how bad you think it will be. It will be instructive to think about the weaknesses of that design
    and contrast them with the features of other designs.
    
   The design-it-twice principle can be applied at many levels in a system. For a module, you can use this approach 
   first to pick the interface, as described above. Then you can apply it again when you are designing the implementation:
   
   if you want to get really great results, you have to consider a second possibility, or perhaps a third, no matter how 
   smart you are. The design of large software systems falls in this category: no-one is good enough to get it right with
    their first try.
    
   12 Why Write Comments?
   ---------------
   process of writing comments, if done correctly, will actually improve a system’s design.
   
13 Comments Should Describe Things that Aren’t Obvious from the Code 
-----------
### Red Flag: Comment Repeats Code
**use different words in the comment from those in the name of the entity being described.**

### Lower-level comments add precision

• What are the units for this variable?
• Are the boundary conditions inclusive or exclusive?
• If a null value is permitted, what does it imply?
• If a variable refers to a resource that must eventually be freed or closed, who is responsible for freeing or closing it?
• Are there certain properties that are always true for the variable (invariants),such as “this list always contains at least one entry”?

When documenting a variable, think nouns, not verbs. In other words, focus on what the variable represents, not how it is manipulated.

评论比代码更详细 级别更低
评论比代码更高级别 更抽象
同级别的就是对代码的重复 不建议！

### Higher-level comments enhance intuition

The second way in which comments can augment code is by providing intuition. These comments are written at a higher level
 than the code. They omit details and help the reader to understand the overall intent and structure of the code.
  This approach is commonly used for comments inside methods, and for interface comments.


Higher-level comments are more difficult to write than lower-level comments because you must think about the code in a different way.

ASK yourself: What is this code trying to do? What is the simplest thing you can say that explains everything in the code?
 What is the most important thing about this code?
 
 Engineers tend to be very detail-oriented. We love details and are good at managing lots of them; this is essential 
 for being a good engineer. But, great software designers can also step back from the details and think about a system at a higher level.

### Interface documentation
**If you want code that presents good abstractions, you must document those abstractions with comments.**

The first step in documenting abstractions is to separate interface comments from implementation comments. 
Interface comments provide information that someone needs to know in order to use a class or method; 
they define the abstraction. Implementation comments describe how a class or method works internally in order to 
implement the abstraction. It’s important to separate these two kinds of comments, so that users of an interface are 
not exposed to implementation details.

**If interface comments must also describe the implementation, then the class or method is shallow.**

### 13.6 Implementation comments: what and why, not how

Implementation comments are the comments that appear inside methods to help readers understand how they work internally.

**The main goal of implementation comments is to help readers understand what the code is doing (not how it does it).**

### Cross-module design decisions

In a perfect world, every important design decision would be encapsulated within a single class. Unfortunately,
 real systems inevitably end up with design decisions that affect multiple classes. For example, the design of a network
  protocol will affect both the sender and the receiver, and these may be implemented in different places. Cross-module 
  decisions are often complex and subtle, and they account for many bugs, so good documentation for them is crucial.
