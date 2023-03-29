# e-com_service

Product Service: provides information about the product like availability, price, category Order service: provides information about the order like orderValue, dispatchDate, orderStatus, prodQuantity

To Run,

    go run main.go

    or

    docker build -t josh-com .

    docker run -it -d -p 8080:8080 josh-com

#todo

update Readme

Add api endpoints

add dockerfile

add sample example result

    1. Create Product cURL:

            curl --location --request POST 'localhost:8080/product' \
            --header 'Content-Type: application/json' \
            --data-raw '{
                "name": "product Josh test",
                "price": 150.00,
                "category": "premium",
                "availability": true
            }'

        Response:

            {
                "status": 201,
                "message": "Product Created successfully",
                "data": {
                    "id": "641840cfd74c6eecc15a2943"
                }
            }

    2. Get Product by its ID cURL:

            curl --location --request GET 'localhost:8080/product/641840cfd74c6eecc15a2943'

        Response:

            {
                "status": 200,
                "message": "The details of the product for requested Id",
                "data": {
                    "id": "641840cfd74c6eecc15a2943",
                    "name": "product Josh test",
                    "price": 150,
                    "category": "premium",
                    "availability": true
                }
            }

    3. Update Product by its ID cURL:

            curl --location --request PUT 'localhost:8080/product/641840cfd74c6eecc15a2943' \
            --header 'Content-Type: application/json' \
            --data-raw '{

                "availability": false

            }'

        Response:

            {
                "status": 200,
                "message": "The product data is updated successfully for given id.",
                "data": {
                    "id": "641840cfd74c6eecc15a2943",
                    "name": "product Josh test",
                    "price": 150,
                    "category": "premium",
                    "availability": false
                }
            }

    4. List all products cURL:

            curl --location --request GET 'localhost:8080/product'

        Response:

            {
                "status": 200,
                "message": "product List success",
                "data": [
                    {
                        "id": "64156c98d27ec97425ada110",
                        "name": "product test order 4",
                        "price": 15,
                        "category": "regular",
                        "availability": true
                    },
                    {
                        "id": "641840cfd74c6eecc15a2943",
                        "name": "product Josh test",
                        "price": 150,
                        "category": "premium",
                        "availability": false
                    }
                ]
            }

    5. Create order cURL:

            curl --location --request POST 'localhost:8080/order' \
            --header 'Content-Type: application/json' \
            --data-raw '{
                "name": "Test order for josh",
                "product" : [
                    {
                        "product_id": "641840cfd74c6eecc15a2943",
                        "quantity": 4

                    },
                    {
                        "product_id": "64156c7dd27ec97425ada10d",
                        "quantity": 3

                    },
                    {
                        "product_id": "64156c98d27ec97425ada110",
                        "quantity": 5

                    },
                    {
                        "product_id": "64156c88d27ec97425ada10e",
                        "quantity": 3

                    }

                ]
            }'


        Response:

            {
                "status": 201,
                "message": "Order Created successfully",
                "data": {
                    "id": "6418463ae3f4077480cf2f69"
                }
            }


    6. Get Order by its ID:

            curl --location --request GET 'localhost:8080/order/6418463ae3f4077480cf2f69'

        Response:

            {
                "status": 200,
                "message": "The details of the order for requested Id,",
                "data": {
                    "id": "6418463ae3f4077480cf2f69",
                    "name": "Test order for josh",
                    "ordervalue": 882.9,
                    "orderstatus": "placed",
                    "products": [
                        {
                            "id": "641840cfd74c6eecc15a2943",
                            "name": "product Josh test",
                            "price": 150,
                            "quantity": 4
                        },
                        {
                            "id": "64156c7dd27ec97425ada10d",
                            "name": "product test order 1",
                            "price": 52,
                            "quantity": 3
                        },
                        {
                            "id": "64156c98d27ec97425ada110",
                            "name": "product test order 4",
                            "price": 15,
                            "quantity": 5
                        },
                        {
                            "id": "64156c88d27ec97425ada10e",
                            "name": "product test order 2",
                            "price": 50,
                            "quantity": 3
                        }
                    ]
                }
            }


    7. Update order status by its ID:

            curl --location --request PUT 'localhost:8080/order/6418463ae3f4077480cf2f69' \
            --header 'Content-Type: application/json' \
            --data-raw '{
                "orderstatus":"dispatched"
            }'

        Response:

            {
                "status": 200,
                "message": "The Order data is updated successfully for given id.",
                "data": {
                    "id": "6418463ae3f4077480cf2f69",
                    "name": "Test order for josh",
                    "ordervalue": 882.9,
                    "orderstatus": "dispatched",
                    "dispatch": "2023-03-23T11:50:50.67Z",
                    "products": [
                        {
                            "product_id": "641840cfd74c6eecc15a2943",
                            "quantity": 4
                        },
                        {
                            "product_id": "64156c7dd27ec97425ada10d",
                            "quantity": 3
                        },
                        {
                            "product_id": "64156c98d27ec97425ada110",
                            "quantity": 5
                        },
                        {
                            "product_id": "64156c88d27ec97425ada10e",
                            "quantity": 3
                        }
                    ]
                }
            }
