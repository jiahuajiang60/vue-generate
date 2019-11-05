# vue-generate
vue 表单生成器
## 前端表单页面生成使用说明

#### json格式说明
```javascript
let data={
    key:{
        type: 'input', 
        label: '身份证', 
        rule: ["",],
        data,
        notHave
    }
}
```
#### type取值
- input
- select
- checkbox
- input-data

#### rule取值
- number
- phone
- idNumber(身份证号)
