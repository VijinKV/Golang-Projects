## Bank Server
This is a Server implementation of Bank transfer money functionality in Golang. 
It takes care of all the transaction related issues and aims to deliver consistent 
and reliable apis that maintains the integrity of the banking system.

### Features:
1. Display Account Details like Amount and Account Number.
2. Transaction of Money from sender to receivers.

### Tech used:
1. Golang
2. Mysql

### Check and cases:
  1. One sender can't send money to itself.
  2. One cannot forge the body and transfer the money from others account to others.
  3. The balance cant go in -ve.
  4. Transaction are following ACID properties so that the api don't leave the state of transaction in invalid state.

### APIS & Port
So this server is exposed to default 8080 port.
#### Routes:
1. **/account/{acountId}** - This api will give you all the details of the account and its balance. The support for related
transactions related to the Account is on WIP.

#### Response
```sh
{
    "acc_number": "BANK0001",
    "balance": 920,
    "transactions": null
}
````

2. **/account/{acountId}/tansfer** - This api is used to transfer the from from ones account to the other. 

#### Payload
```sh
{
	"from": "BANK0001",
	"to": "BANK0002",
	"amount": 10
}
```

#### Response
```sh
{
    "id": "71c606b1-86f4-40cc-899e-661ea5b0ce76",
    "from": {
        "Id": "BANK0001",
        "Balance": 700
    },
    "to": {
        "Id": "BANK0002",
        "Balance": 1260
    },
    "transferred": 10,
    "created_datetime": "2023-03-01T16:06:11.9062815+05:30"
}
```

### Tables in the DB
#### Transaction table
```sh
1. transaction_id VARCHAR(255) PRIMARY KEY,
2. from_acc_no VARCHAR(255),
3. to_acc_no VARCHAR(255),
4. transferred_amount FLOAT,
5. created_datetime TIMESTAMP DEFAULT CURRENT_TIMESTAMP
```

### Todo and Good things to have
1. Config file and ENV variable use.
2. Handler interface to reduce coupling.
3. UI to demonstrate the APIs.