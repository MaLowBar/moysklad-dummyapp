package main

import (
	"github.com/gin-gonic/gin"
	"go-dummyapp-moysklad/internal/app"
	"go-dummyapp-moysklad/internal/config"
)

const dataCtx = `{
  "meta": {
    "href": "https://online.moysklad.ru/api/remap/1.2/entity/employee/b0a02321-13e3-11e9-912f-f3d4002516e3?expand=cashier.retailStore",
    "metadataHref": "https://online.moysklad.ru/api/remap/1.2/entity/employee/metadata",
    "type": "employee",
    "mediaType": "application/json",
    "uuidHref": "https://online.moysklad.ru/app/#employee/edit?id=b0a02321-13e3-11e9-912f-f3d4002516e3"
  },
  "id": "b0a02321-13e3-11e9-912f-f3d4002516e3",
  "accountId": "b0b309ee-13e3-11e9-9109-f8fc0001f188",
  "owner": {
    "meta": {
      "href": "https://online.moysklad.ru/api/remap/1.2/entity/employee/b0a02321-13e3-11e9-912f-f3d4002516e3",
      "metadataHref": "https://online.moysklad.ru/api/remap/1.2/entity/employee/metadata",
      "type": "employee",
      "mediaType": "application/json",
      "uuidHref": "https://online.moysklad.ru/app/#employee/edit?id=b0a02321-13e3-11e9-912f-f3d4002516e3"
    }
  },
  "shared": true,
  "group": {
    "meta": {
      "href": "https://online.moysklad.ru/api/remap/1.2/entity/group/b0b3c289-13e3-11e9-9109-f8fc0001f189",
      "metadataHref": "https://online.moysklad.ru/api/remap/1.2/entity/group/metadata",
      "type": "group",
      "mediaType": "application/json"
    }
  },
  "updated": "2019-12-10 18:37:25.786",
  "name": "Кожевников",
  "externalCode": "Exh56G1wiRTPHpYBc-nx12",
  "archived": false,
  "created": "2019-01-09 10:53:45.202",
  "uid": "admin@bkozhevnikov",
  "email": "bkozhevnikov@moysklad.ru",
  "lastName": "Кожевников",
  "fullName": "Кожевников",
  "shortFio": "Кожевников",
  "cashiers": [
    {
      "meta": {
        "href": "https://online.moysklad.ru/api/remap/1.2/entity/retailstore/b0b7cd8d-13e3-11e9-912f-f3d400251724/cashiers/b0b7d387-13e3-11e9-912f-f3d400251725",
        "type": "cashier",
        "mediaType": "application/json"
      }
    }
  ],
  "permissions": {
    "currency": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL"
    },
    "uom": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL"
    },
    "productfolder": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "print": "ALL"
    },
    "product": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "print": "ALL"
    },
    "bundle": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "print": "ALL"
    },
    "service": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "print": "ALL"
    },
    "consignment": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "print": "ALL"
    },
    "variant": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "print": "ALL"
    },
    "store": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL"
    },
    "counterparty": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "print": "ALL"
    },
    "organization": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL"
    },
    "employee": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL"
    },
    "contract": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "print": "ALL"
    },
    "project": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL"
    },
    "country": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL"
    },
    "customentity": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL"
    },
    "demand": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "approve": "ALL",
      "print": "ALL"
    },
    "customerorder": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "approve": "ALL",
      "print": "ALL"
    },
    "internalorder": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "approve": "ALL",
      "print": "ALL"
    },
    "invoiceout": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "approve": "ALL",
      "print": "ALL"
    },
    "invoicein": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "approve": "ALL",
      "print": "ALL"
    },
    "paymentin": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "approve": "ALL",
      "print": "ALL"
    },
    "paymentout": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "approve": "ALL",
      "print": "ALL"
    },
    "cashin": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "approve": "ALL",
      "print": "ALL"
    },
    "cashout": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "approve": "ALL",
      "print": "ALL"
    },
    "supply": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "approve": "ALL",
      "print": "ALL"
    },
    "salesreturn": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "approve": "ALL",
      "print": "ALL"
    },
    "purchasereturn": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "approve": "ALL",
      "print": "ALL"
    },
    "retailstore": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL"
    },
    "receipttemplate": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL"
    },
    "retailshift": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "print": "ALL"
    },
    "retaildemand": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "approve": "ALL",
      "print": "ALL"
    },
    "retailsalesreturn": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "approve": "ALL",
      "print": "ALL"
    },
    "retaildrawercashin": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "approve": "ALL",
      "print": "ALL"
    },
    "retaildrawercashout": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "approve": "ALL",
      "print": "ALL"
    },
    "prepayment": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "approve": "ALL",
      "print": "ALL"
    },
    "prepaymentreturn": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "approve": "ALL",
      "print": "ALL"
    },
    "purchaseorder": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "approve": "ALL",
      "print": "ALL"
    },
    "move": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "approve": "ALL",
      "print": "ALL"
    },
    "enter": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "approve": "ALL",
      "print": "ALL"
    },
    "loss": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "approve": "ALL",
      "print": "ALL"
    },
    "facturein": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "approve": "ALL",
      "print": "ALL"
    },
    "factureout": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "approve": "ALL",
      "print": "ALL"
    },
    "commissionreportin": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "approve": "ALL",
      "print": "ALL"
    },
    "commissionreportout": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "approve": "ALL",
      "print": "ALL"
    },
    "pricelist": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "approve": "ALL",
      "print": "ALL"
    },
    "processingplanfolder": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL"
    },
    "processingplan": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL"
    },
    "processing": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "approve": "ALL",
      "print": "ALL"
    },
    "processingorder": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "approve": "ALL",
      "print": "ALL"
    },
    "assortment": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "print": "ALL"
    },
    "inventory": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "print": "ALL"
    },
    "bonustransaction": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "approve": "ALL",
      "print": "ALL"
    },
    "crptorder": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "approve": "ALL",
      "print": "ALL"
    },
    "webhook": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL"
    },
    "task": {
      "view": "ALL",
      "create": "ALL",
      "update": "ALL",
      "delete": "ALL",
      "done": "ALL"
    },
    "dashboard": {
      "view": "ALL"
    },
    "stock": {
      "view": "ALL"
    },
    "customAttributes": {
      "view": "ALL"
    },
    "pnl": {
      "view": "ALL"
    },
    "company_crm": {
      "view": "ALL"
    },
    "tariff_crm": {
      "view": "ALL"
    },
    "audit_dashboard": {
      "view": "ALL"
    },
    "admin": {
      "view": "ALL"
    },
    "dashboardMoney": {
      "view": "ALL"
    }
  }
}`
const dataStores = `{
  "context": {
    "employee": {
      "meta": {
        "href": "https://online.moysklad.ru/api/remap/1.2/context/employee",
        "metadataHref": "https://online.moysklad.ru/api/remap/1.2/entity/employee/metadata",
        "type": "employee",
        "mediaType": "application/json"
      }
    }
  },
  "meta": {
    "href": "https://online.moysklad.ru/api/remap/1.2/entity/store/",
    "metadataHref": "https://online.moysklad.ru/api/remap/1.2/entity/store/metadata",
    "type": "store",
    "mediaType": "application/json",
    "size": 5,
    "limit": 1000,
    "offset": 0
  },
  "rows": [
    {
      "meta": {
        "href": "https://online.moysklad.ru/api/remap/1.2/entity/store/caf46ce5-0569-11e6-9464-e4de00000000",
        "metadataHref": "https://online.moysklad.ru/api/remap/1.2/entity/store/metadata",
        "type": "store",
        "mediaType": "application/json"
      },
      "id": "caf46ce5-0569-11e6-9464-e4de00000000",
      "accountId": "84e60e93-f504-11e5-8a84-bae500000008",
      "owner": {
        "meta": {
          "href": "https://online.moysklad.ru/api/remap/1.2/entity/employee/faba7f37-2e58-11e6-8a84-bae500000028",
          "metadataHref": "https://online.moysklad.ru/api/remap/1.2/entity/employee/metadata",
          "type": "employee",
          "mediaType": "application/json"
        }
      },
      "shared": false,
      "group": {
        "meta": {
          "href": "https://online.moysklad.ru/api/remap/1.2/entity/group/f97aa1fb-2e58-11e6-8a84-bae500000002",
          "metadataHref": "https://online.moysklad.ru/api/remap/1.2/entity/group/metadata",
          "type": "group",
          "mediaType": "application/json"
        }
      },
      "updated": "2016-04-18 16:31:01",
      "name": "002",
      "externalCode": "y7ztWINfjXinPToFMqQid2",
      "archived": false,
      "address": "125009, Россия, г Москва, Москва, ул Тверская, 1, 123, addInfo",
      "addressFull": {
        "postalCode": "125009",
        "country": {
          "meta": {
            "href": "https://online.moysklad.ru/api/remap/1.2/entity/country/9df7c2c3-7782-4c5c-a8ed-1102af611608",
            "metadataHref": "https://online.moysklad.ru/api/remap/1.2/entity/country/metadata",
            "type": "country",
            "mediaType": "application/json"
          }
        },
        "region": {
          "meta": {
            "href": "https://online.moysklad.ru/api/remap/1.2/entity/region/00000000-0000-0000-0000-000000000077",
            "metadataHref": "https://online.moysklad.ru/api/remap/1.2/entity/region/metadata",
            "type": "region",
            "mediaType": "application/json"
          }
        },
        "city": "Москва",
        "street": "ул Тверская",
        "house": "1",
        "apartment": "123",
        "addInfo": "addinfo",
        "comment": "some words about address"
      },
      "pathName": ""
    },
    {
      "meta": {
        "href": "https://online.moysklad.ru/api/remap/1.2/entity/store/850ee995-f504-11e5-8a84-bae500000160",
        "metadataHref": "https://online.moysklad.ru/api/remap/1.2/entity/store/metadata",
        "type": "store",
        "mediaType": "application/json"
      },
      "id": "850ee995-f504-11e5-8a84-bae500000160",
      "accountId": "84e60e93-f504-11e5-8a84-bae500000008",
      "owner": {
        "meta": {
          "href": "https://online.moysklad.ru/api/remap/1.2/entity/employee/faba7f37-2e58-11e6-8a84-bae500000028",
          "metadataHref": "https://online.moysklad.ru/api/remap/1.2/entity/employee/metadata",
          "type": "employee",
          "mediaType": "application/json"
        }
      },
      "shared": false,
      "group": {
        "meta": {
          "href": "https://online.moysklad.ru/api/remap/1.2/entity/group/f97aa1fb-2e58-11e6-8a84-bae500000002",
          "metadataHref": "https://online.moysklad.ru/api/remap/1.2/entity/group/metadata",
          "type": "group",
          "mediaType": "application/json"
        }
      },
      "updated": "2016-03-28 19:45:46",
      "name": "Основной склад",
      "externalCode": "OJ8pY2FgjQ3ncLVvvpqyw1",
      "archived": false,
      "pathName": "",
      "address": "125009, Россия, г Москва, Москва, ул Тверская, 1, 123, addInfo",
      "addressFull": {
        "postalCode": "125009",
        "country": {
          "meta": {
            "href": "https://online.moysklad.ru/api/remap/1.2/entity/country/9df7c2c3-7782-4c5c-a8ed-1102af611608",
            "metadataHref": "https://online.moysklad.ru/api/remap/1.2/entity/country/metadata",
            "type": "country",
            "mediaType": "application/json"
          }
        },
        "region": {
          "meta": {
            "href": "https://online.moysklad.ru/api/remap/1.2/entity/region/00000000-0000-0000-0000-000000000077",
            "metadataHref": "https://online.moysklad.ru/api/remap/1.2/entity/region/metadata",
            "type": "region",
            "mediaType": "application/json"
          }
        },
        "city": "Москва",
        "street": "ул Тверская",
        "house": "1",
        "apartment": "123",
        "addInfo": "addinfo",
        "comment": "some words about address"
      },
      "attributes": [
        {
          "meta": {
            "href": "https://online.moysklad.ru/api/remap/1.2/entity/store/metadata/attributes/0cd74e1e-2e59-11e6-8a84-bae50000008a",
            "type": "attributemetadata",
            "mediaType": "application/json"
          },
          "id": "3a85cfe3-12c5-11e6-9464-e4de00000087",
          "name": "Площадь",
          "type": "long",
          "value": 4400
        }
      ]
    },
    {
      "meta": {
        "href": "https://online.moysklad.ru/api/remap/1.2/entity/store/6ebb9094-056a-11e6-9464-e4de000000b4",
        "metadataHref": "https://online.moysklad.ru/api/remap/1.2/entity/store/metadata",
        "type": "store",
        "mediaType": "application/json"
      },
      "id": "6ebb9094-056a-11e6-9464-e4de000000b4",
      "accountId": "84e60e93-f504-11e5-8a84-bae500000008",
      "owner": {
        "meta": {
          "href": "https://online.moysklad.ru/api/remap/1.2/entity/employee/faba7f37-2e58-11e6-8a84-bae500000028",
          "metadataHref": "https://online.moysklad.ru/api/remap/1.2/entity/employee/metadata",
          "type": "employee",
          "mediaType": "application/json"
        }
      },
      "shared": false,
      "group": {
        "meta": {
          "href": "https://online.moysklad.ru/api/remap/1.2/entity/group/f97aa1fb-2e58-11e6-8a84-bae500000002",
          "metadataHref": "https://online.moysklad.ru/api/remap/1.2/entity/group/metadata",
          "type": "group",
          "mediaType": "application/json"
        }
      },
      "updated": "2016-04-18 16:35:36",
      "name": "Подскладик",
      "code": "ZAATY643",
      "externalCode": "d8Ew2hCDiTuJFb0Ya45tH0",
      "archived": false,
      "address": "125009, Россия, г Москва, Москва, ул Тверская, 1, 123, addInfo",
      "addressFull": {
        "postalCode": "125009",
        "country": {
          "meta": {
            "href": "https://online.moysklad.ru/api/remap/1.2/entity/country/9df7c2c3-7782-4c5c-a8ed-1102af611608",
            "metadataHref": "https://online.moysklad.ru/api/remap/1.2/entity/country/metadata",
            "type": "country",
            "mediaType": "application/json"
          }
        },
        "region": {
          "meta": {
            "href": "https://online.moysklad.ru/api/remap/1.2/entity/region/00000000-0000-0000-0000-000000000077",
            "metadataHref": "https://online.moysklad.ru/api/remap/1.2/entity/region/metadata",
            "type": "region",
            "mediaType": "application/json"
          }
        },
        "city": "Москва",
        "street": "ул Тверская",
        "house": "1",
        "apartment": "123",
        "addInfo": "addinfo",
        "comment": "some words about address"
      },
      "pathName": "002"
    },
    {
      "meta": {
        "href": "https://online.moysklad.ru/api/remap/1.2/entity/store/95dcda62-056b-11e6-9464-e4de000000b7",
        "metadataHref": "https://online.moysklad.ru/api/remap/1.2/entity/store/metadata",
        "type": "store",
        "mediaType": "application/json"
      },
      "id": "95dcda62-056b-11e6-9464-e4de000000b7",
      "accountId": "84e60e93-f504-11e5-8a84-bae500000008",
      "owner": {
        "meta": {
          "href": "https://online.moysklad.ru/api/remap/1.2/entity/employee/faba7f37-2e58-11e6-8a84-bae500000028",
          "metadataHref": "https://online.moysklad.ru/api/remap/1.2/entity/employee/metadata",
          "type": "employee",
          "mediaType": "application/json"
        }
      },
      "shared": false,
      "group": {
        "meta": {
          "href": "https://online.moysklad.ru/api/remap/1.2/entity/group/f97aa1fb-2e58-11e6-8a84-bae500000002",
          "metadataHref": "https://online.moysklad.ru/api/remap/1.2/entity/group/metadata",
          "type": "group",
          "mediaType": "application/json"
        }
      },
      "updated": "2016-04-18 16:43:51",
      "name": "Склад1",
      "description": "Основной склад",
      "code": "113AB79",
      "externalCode": "fQPIOtxjg-FaeZNKcLx6B3",
      "archived": false,
      "address": "125009, Россия, г Москва, Москва, ул Тверская, 1, 123, addInfo",
      "addressFull": {
        "postalCode": "125009",
        "country": {
          "meta": {
            "href": "https://online.moysklad.ru/api/remap/1.2/entity/country/9df7c2c3-7782-4c5c-a8ed-1102af611608",
            "metadataHref": "https://online.moysklad.ru/api/remap/1.2/entity/country/metadata",
            "type": "country",
            "mediaType": "application/json"
          }
        },
        "region": {
          "meta": {
            "href": "https://online.moysklad.ru/api/remap/1.2/entity/region/00000000-0000-0000-0000-000000000077",
            "metadataHref": "https://online.moysklad.ru/api/remap/1.2/entity/region/metadata",
            "type": "region",
            "mediaType": "application/json"
          }
        },
        "city": "Москва",
        "street": "ул Тверская",
        "house": "1",
        "apartment": "123",
        "addInfo": "addinfo",
        "comment": "some words about address"
      },
      "pathName": "Основной склад"
    }
  ]
}`

func main() {
	cfg := config.NewConfig("internal/config/config.yml")
	appServer := app.NewServer(cfg)

	router := gin.Default()

	router.PUT(cfg.AppBaseUrl, appServer.ActivateHandler)
	router.DELETE(cfg.AppBaseUrl, appServer.DeleteHandler)
	router.GET(cfg.AppBaseUrl, appServer.StatusHandler)

	//test context
	router.POST("/ctx", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Writer.Write([]byte(dataCtx))
	})
	//end test context

	//test stores
	router.GET("/store", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Writer.Write([]byte(dataStores))
	})
	//end test stores

	router.LoadHTMLFiles("./iframe.html")
	router.GET("/iframe/:appId", appServer.HtmlHandler)
	router.POST("/:appId/update-settings", appServer.UpdateSettingsHandler)
	router.Run(":8080")
}
