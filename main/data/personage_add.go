package data

const JSON = `
{
"number": {
  "type": "input",
    "label": "身份证"
},
"sex": {
  "type": "select",
    "label": "性别",
    "data": [
    {
      "_id": "1",
      "name": "男"
    },
    {
      "_id": "2",
      "name": "女"
    },
    {
      "_id": "3",
      "name": "未知性别"
    },
    {
      "_id": "4",
      "name": "未说明的性别"
    }
  ]
},
"selfPhone": {
  "type": "input",
    "label": "本人电话"
},
"workAddress": {
  "type": "input",
    "label": "工作地址"
},
"contacts": {
  "type": "input",
    "label": "联系人"
},
"contactsPhone": {
  "type": "input",
    "label": "联系电话"
},
"residentType": {
  "type": "select",
    "label": "常驻类型",
    "data": [
    {
      "_id": "1",
      "name": "农村户口"
    },
    {
      "_id": "2",
      "name": "城镇户口"
    }
  ]
},
"famousRace": {
  "type": "select",
    "label": "民族",
    "data": [
    {
      "_id": "1",
      "name": "汉族"
    },
    {
      "_id": "2",
      "name": "少数民族"
    }
  ]
},
"bloodType": {
  "type": "select",
    "label": "血型",
    "data": [
    {
      "_id": "1",
      "name": "A血型"
    },
    {
      "_id": "2",
      "name": "B血型"
    },
    {
      "_id": "3",
      "name": "O血型"
    },
    {
      "_id": "4",
      "name": "不详"
    }
  ]
},
"bloodTypeRh": {
  "type": "select",
    "label": "RH阴性",
    "data": [
    {
      "_id": "1",
      "name": "RH阴性"
    },
    {
      "_id": "2",
      "name": "RH阳性"
    },
    {
      "_id": "3",
      "name": "不详"
    }
  ]
},
"education": {
  "type": "select",
    "label": "教育",
    "data": [
    {
      "_id": "1",
      "name": "研究生"
    },
    {
      "_id": "2",
      "name": "大学本科"
    },
    {
      "_id": "3",
      "name": "大学专科和专科学校"
    },
    {
      "_id": "4",
      "name": "中等专业学校"
    },
    {
      "_id": "5",
      "name": "技工学校"
    },
    {
      "_id": "6",
      "name": "高中"
    },
    {
      "_id": "7",
      "name": "初中"
    },
    {
      "_id": "8",
      "name": "小学"
    },
    {
      "_id": "9",
      "name": "文盲或半文盲"
    },
    {
      "_id": "10",
      "name": "不详"
    }
  ]
},
"work": {
  "type": "select",
    "label": "职业",
    "data": [
    {
      "_id": "1",
      "name": "国家机关,党群组织,企业,事业单位负责人,1专业技术人员"
    },
    {
      "_id": "2",
      "name": "办事人员和有关人员"
    },
    {
      "_id": "3",
      "name": "商业，服务人员"
    },
    {
      "_id": "4",
      "name": "农，林，牧，鱼，水利业生产人员"
    },
    {
      "_id": "5",
      "name": "生产，运输设备操作人员及有关人员"
    },
    {
      "_id": "6",
      "name": "军人"
    },
    {
      "_id": "7",
      "name": "不便分类的其它从业人员"
    },
    {
      "_id": "8",
      "name": "无职业"
    }
  ]
},
"maritalStatus": {
  "type": "select",
    "label": "婚姻状况",
    "data": [
    {
      "_id": "1",
      "name": "未婚"
    },
    {
      "_id": "2",
      "name": "已婚"
    },
    {
      "_id": "3",
      "name": "丧偶"
    },
    {
      "_id": "4",
      "name": "离婚"
    },
    {
      "_id": "5",
      "name": "未说明的婚姻状况"
    }
  ]
},
"paymentOfMedicalExpenses": {
  "type": "checkbox",
    "label": "医疗费用支付方式",
    "notHave": true,
    "data": [
    {
      "_id": "1",
      "name": "城镇职工基本医疗保险"
    },
    {
      "_id": "2",
      "name": "城镇居民基本医疗保险"
    },
    {
      "_id": "3",
      "name": "新型农村合作医疗"
    },
    {
      "_id": "4",
      "name": "贫困救助"
    },
    {
      "_id": "5",
      "name": "商业医疗保险"
    },
    {
      "_id": "6",
      "name": "全公费"
    },
    {
      "_id": "7",
      "name": "全自费"
    },
    {
      "_id": "8",
      "name": "其它"
    }
  ]
},
"historyOfDrugAllergy": {
  "type": "checkbox",
    "label": "药物过敏史",
    "notHave": true,
    "data": [
    {
      "_id": "1",
      "name": "1无"
    },
    {
      "_id": "2",
      "name": "青霉素"
    },
    {
      "_id": "3",
      "name": "磺胺"
    },
    {
      "_id": "4",
      "name": "链霉素"
    },
    {
      "_id": "5",
      "name": "其它"
    }
  ]
},
"exposureHistory": {
  "type": "checkbox",
    "label": "暴露史",
    "data": [
    {
      "_id": "1",
      "name": "无"
    },
    {
      "_id": "2",
      "name": "化学品"
    },
    {
      "_id": "3",
      "name": "毒物"
    },
    {
      "_id": "4",
      "name": "射线"
    }
  ]
},
"disease": {
  "type": "checkbox",
    "label": "既往史疾病",
    "data": [
    {
      "_id": "1",
      "name": "无"
    },
    {
      "_id": "2",
      "name": "高血压"
    },
    {
      "_id": "3",
      "name": "糖尿病"
    },
    {
      "_id": "4",
      "name": "冠心病"
    },
    {
      "_id": "5",
      "name": "慢性阻塞性肺疾病"
    },
    {
      "_id": "6",
      "name": "恶性肿瘤"
    },
    {
      "_id": "7",
      "name": "脑卒中"
    },
    {
      "_id": "8",
      "name": "严重精神障碍"
    },
    {
      "_id": "9",
      "name": "结核病"
    },
    {
      "_id": "10",
      "name": "肝炎"
    },
    {
      "_id": "11",
      "name": "其它法定传染病"
    },
    {
      "_id": "12",
      "name": "职业病"
    },
    {
      "_id": "13",
      "name": "其它"
    }
  ]
},
"surgery": {
  "type": "input-data",
    "label": "手术",
    "notHave": true
},
"trauma": {
  "type": "input-data",
    "label": "请输入外伤名称",
    "notHave": true
},
"geneticDisease": {
  "type": "input-data",
    "label": "遗传病",
    "notHave": true
},
"disabilitySituation": {
  "type": "checkbox",
    "label": "残疾情况",
    "data": [
    {
      "_id": "1",
      "name": "无"
    },
    {
      "_id": "2",
      "name": "视力残疾"
    },
    {
      "_id": "3",
      "name": "听力残疾"
    },
    {
      "_id": "4",
      "name": "言语残疾"
    },
    {
      "_id": "5",
      "name": "肢体残疾"
    },
    {
      "_id": "6",
      "name": "智力残疾"
    },
    {
      "_id": "7",
      "name": "精神残疾"
    },
    {
      "_id": "8",
      "name": "其它残疾"
    }
  ]
},
"kitchenExhaust": {
  "type": "select",
    "label": "厨房排风",
    "data": [
    {
      "_id": "1",
      "name": "无"
    },
    {
      "_id": "2",
      "name": "油烟机"
    },
    {
      "_id": "3",
      "name": "换气扇"
    },
    {
      "_id": "4",
      "name": "烟窗"
    }
  ]
},
"fuelType": {
  "type": "select",
    "label": "燃料类型",
    "data": [
    {
      "_id": "6",
      "name": "其它"
    },
    {
      "_id": "1",
      "name": "液化气"
    },
    {
      "_id": "2",
      "name": "煤"
    },
    {
      "_id": "3",
      "name": "天然气"
    },
    {
      "_id": "4",
      "name": "沼气"
    },
    {
      "_id": "5",
      "name": "柴火"
    }
  ]
},
"drinkingWater": {
  "tyoe": "select",
    "label": "饮水",
    "data": [
    {
      "_id": "1",
      "name": "自来水"
    },
    {
      "_id": "2",
      "name": "经净化过滤的水"
    },
    {
      "_id": "3",
      "name": "井水"
    },
    {
      "_id": "4",
      "name": "河湖水"
    },
    {
      "_id": "5",
      "name": "塘水"
    },
    {
      "_id": "6",
      "name": "其它"
    }
  ]
},
"toilet": {
  "tyoe": "select",
    "label": "厕所",
    "data": [
    {
      "_id": "1",
      "name": "卫生厕所"
    },
    {
      "_id": "2",
      "name": "一格或两格粪池水"
    },
    {
      "_id": "3",
      "name": "马桶"
    },
    {
      "_id": "4",
      "name": "露天粪坑"
    },
    {
      "_id": "5",
      "name": "简易棚厕"
    }
  ]
},
"poultryCorral": {
  "type": "select",
    "label": "禽畜栏",
    "data": [
    {
      "_id": "1",
      "name": "无"
    },
    {
      "_id": "2",
      "name": "单设"
    },
    {
      "_id": "3",
      "name": "室内"
    },
    {
      "_id": "4",
      "name": "室外"
    }
  ]
}
}
`
