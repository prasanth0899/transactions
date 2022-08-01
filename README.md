# transactions

```
{
    "request": {
        "filters": [
            {
                "filter": "date",
                "title": "fromDate",
                "body": "2022-07-28T13:27:18Z"
            }, {
                "type": "amount", 
                "title": "direction",
                "body": "out/in/all"
            }, {
                "type": "method",
                "title": "ach", 
                "body": "mercuryPayment(ACH)/.../all"
            }, {
                "type": "sentFrom",
                "title": "bank",
                "body": "4000839f-ce91-4bc9-bf69-1bc0c116dffe"
            }, 
            {
                "type": "status",
                "title": "status",
                "body": "posted/pending/failed/canceled/all"
            }, {
                "type": "attachments",
                "title": "attachments",
                "body": "true/false/all"
            }
        ],
        "paging":  {
            "previous":  "2022-07-28T13:27:18Z",
            "size": 100
        }
    }

}



{
    "request_id": "4000839f-ce91-4bc9-bf69-1bc0c116dffe",
    "transactions" : [
        {
            "response" : {
                "date": "Jul 28",
                "to/from": "..",
                "amount": 12.34893,
                "currency": "dollar",
                "account": "ops/payroll",
                "method": "transfer out",
                "debited_at": "Jul 28 at 6:30PM",
                "credited_at": "Jul 28 at 6:35PM",
                "converted_to_amount": 123.46,
                "converted_at_amount": 99.85,
                "converted_to": "GBP",
                "from": "AP",
                "to": "Aliya",
                "notes": "",
                "attachments": ["...."],
                "description": "",
                "recepient_details" : {
                    "name": "Aliyah",
                    "transfer_details": [
                        {
                            "type": "CheckDetails",
                            "details": [
                                {
                                    "title": "address",
                                    "body": "660 Mission St, Floor 4 San Francisco, CA 94105 United States"
                                }
                            ]
                        }
                    ],
                    "last_payment": "July 8"
                }
            },
            "references": [
                {
                    "recepient_id": "4000839f-ce91-4bc9-bf69-1bc0c116dffe",
                    "merchant_id": "4000839f-ce91-4bc9-bf69-1bc0c116dffe",
                }
            ]
        }, 
        {

        }
    ],
    "references" : {
        "accounts": [
            {
                "name": "AP",
                "id": "4000839f-ce91-4bc9-bf69-1bc0c116dffe"
            }
        ]
    }
}

```
