# GCAS
> GCAS全称为Graphite's Cash AA Service

## 说明
这是一个处于开发中的golang练习项目，用于合租室友、大学宿舍室友等人群进行集体账务管理。

## idea来源
相信很多大学生或者刚出社会和同事合租的应届生朋友们都遇到过这样的问题：宿舍/合租房子的公共物品由一个人购买，然后大家均摊该物品的消费。如果只是偶尔发生大额消费相对计算起来比较简单，但一旦经常发生类似的事情（比如笔者合租有时候会去买菜做饭）并且每次的消费金额并不高，如果每次都要计算消费就比较麻烦，而每隔一段时间再统一计算则容易有遗漏

于是，笔者就想到了是否可以开发一个程序，来对室友之间的共同消费进行管理，再者笔者从业于金融软件行业，正好专业对口

## 语言选择
本身笔者的工作是java开发，但是有学习golang的意向，并且每天下班后并不想再看java了，于是选用了并不熟练的golang，正好可以当作练习

## 接口列表
应付款项	/api/gcas/allcreditpayable
获取应付款明细	待定
获取应收款项	/api/gcas/allcreditreceivable
获取流水记录（包括出账记录和aa转账记录）	/api/gcas/statementquery
获取收到的信息	/api/gcas/mailquery
获取成员信息	/api/gcas/userquery

添加出账记录	/api/gcas/recordinput
添加成员	/api/gcas/userinput
添加付款码	/api/gcas/receiptcodeinput
情况款项记录	/api/gcas/recordtruncate
登入	/api/gcas/login
登出	/api/gcas/logout
