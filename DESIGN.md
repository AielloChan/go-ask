# 设计文档

期望的使用方式

> go-ask
No files added to staging! Did you forget to run git add?
> git init
> git add .
> go-ask
cz-cli@4.2.2, cz-conventional-changelog@3.3.0

? Select the type of change that you're committing: (Use arrow keys)
❯ feat:     A new feature 
  fix:      A bug fix 
  docs:     Documentation only changes 
  style:    Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc) 
  refactor: A code change that neither fixes a bug nor adds a feature 
  perf:     A code change that improves performance 
  test:     Adding missing tests or correcting existing tests 
(Move up and down to reveal more choices)
> 
? What is the scope of this change (e.g. component or file name): (press enter to skip)
> (10) 修复首页图片过大问题
? Provide a longer description of the change: (press enter to skip)
>  - 图片过大，前端裁剪
? Are there any breaking changes? (y/N)



- [ ] 支持回到上一步，并且上一步的数据要填充回来
- [x] 文本类的都支持内置关键词替换，支持 shell 模式
- [x] 配置文件

