var Mock = require("mockjs")


let fields =  process.argv[2];
let params=JSON.parse(fields)

// console.log("==========================================")
// console.log(params)
// console.log("==========================================")


let data = Mock.mock(params)
console.log(JSON.stringify(data))

// let data = Mock.mock({
//     "list|10-40":[
//         {
//             'has|1':'@boolean',
//             "id|+1": 1,
//             "name": "@cname",
//             "gender|1": '@string("男女",1)',
//             "age":'@natural(1,100)',
//             "address": "@county(true)",
//             "birth": "@datetime(yyyy-MM-dd)",
//             "email": '@email',
//             'phone|1': /^1[3|4|5|8][0-9]\d{8}$/
//         }
//     ]
// })
//
// console.log(data)


