## 单例模式

- 模式定义： 确保某一个类只有一个实例，而且自行实例化并向整个系统提供这个实例。
- 使用场景： 

        1、有频繁实例化然后销毁的情况，也就是频繁的 new 对象，可以考虑单例模式；
        2、创建对象时耗时过多或者耗资源过多，但又经常用到的对象；
        3、频繁访问 IO 资源的对象，例如数据库连接池或访问本地文件；