#### 历史订单

##### 1) 请求地址

>{haobase_host}/api/v1/base/order/history

##### 2) 调用方式：HTTP GET

##### 3) 接口描述：

> 历史订单,不包含未成交的订单

##### 4) 请求参数:
> 需要登陆
##### Header:
|字段名称       |类型            |必填            |字段说明         |备注     |
| -------------|:--------------:|:--------------:|:--------------:|:------:|
|Token|string|Y|||

##### GET参数:
|字段名称       |类型            |必填            |字段说明         |备注     |
| -------------|:--------------:|:--------------:|:--------------:|:------:|
|symbol|string|Y|||
|order_id|string|N|||
|start_time|string|N|||
|end_time|string|N||默认 500; 最大 1000|
|limit|int|N|||

注意：
 * 如设置 order_id , 订单量将 >= order_id
 * 如果设置 start_time 和 end_time, order_id 就不需要设置。


##### 5) 请求返回结果:

```
{
    "data": [
        {
            "symbol": "usdjpy",
            "order_id": "B23102016253044451073",
            "order_side": "buy",
            "order_type": "limit",
            "price": "89.000",
            "quantity": "10.00",
            "finished_qty": "10.00",
            "finished_amount": "890.000",
            "status": "done",
            "create_time": 1697790330462963000
        },
        {
            "symbol": "usdjpy",
            "order_id": "B23102016243355986072",
            "order_side": "buy",
            "order_type": "market",
            "price": "0.000",
            "quantity": "100.00",
            "finished_qty": "30.00",
            "finished_amount": "30.000",
            "status": "done",
            "create_time": 1697790273583929000
        }
    ],
    "ok": 1
}
```


##### 6) 请求返回结果参数说明:


  
##### END  
  
