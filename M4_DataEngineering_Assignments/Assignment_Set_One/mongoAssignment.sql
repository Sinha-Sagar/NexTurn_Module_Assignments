------------------------MongoDB Exercises:  Online Shopping Platform-----------------------------------

//	CREATING DATABASE AND COLLECTIONS
		mongosh
		use OnlineShoppingDB
		
		db.createCollection("products")
		db.createCollection("users")
		db.createCollection("orders")
		
//	INSERTING SAMPLE DATA IN COLLECTIONS
		db.users.insertMany([
			  {
				"userId": "U001",
				"name": "Sagar Sinha",
				"email": "Sagar.sinha@nexturn.com",
				"age": 21,
				"address": {
				  "city": "Aar",
				  "state": "Bihar",
				  "zip": "802301"
				},
				"createdAt": "2024-01-01T10:00:00Z"
			  },
			  {
				"userId": "U002",
				"name": "Rohit",
				"email": "Rohit@gmail.com",
				"age": 22,
				"address": {
				  "city": "Patna",
				  "state": "Bihar",
				  "zip": "800026"
				},
				"createdAt": "2024-01-02T11:00:00Z"
			  }
		]);
		
		db.orders.insertMany([
			  {
				"orderId": "ORD001",
				"userId": "U001",
				"orderDate": "2024-12-10T14:32:00Z",
				"items": [
				  { "productId": "P001", "quantity": 4, "price": 100 },
				  { "productId": "P002", "quantity": 3, "price": 50 }
				],
				"totalAmount": 250,
				"status": "Delivered"
			  },
			  {
				"orderId": "ORD002",
				"userId": "U002",
				"orderDate": "2024-12-15T09:20:00Z",
				"items": [
				  { "productId": "P002", "quantity": 25, "price": 50 }
				],
				"totalAmount": 150,
				"status": "Delivered"
			  }
		]);
		
		db.products.insertMany([
			  {
				"productId": "P001",
				"name": "Wireless Mouse",
				"category": "Electronics",
				"price": 50,
				"stock": 200,
				"ratings": [
				  { "userId": "U002", "rating": 4.5 },
				  { "userId": "U003", "rating": 3.0 }
				]
			  },
			  {
				"productId": "P002",
				"name": "Keyboard",
				"category": "Electronics",
				"price": 50,
				"stock": 150,
				"ratings": [
				  { "userId": "U001", "rating": 5.0 },
				  { "userId": "U003", "rating": 4.0 }
				]
			  }
			]);
			
//	QUERIES

1.	Find High-Spending Users 

		db.users.aggregate([
			  {
				"$lookup": {
				  "from": "orders",
				  "localField": "userId",
				  "foreignField": "userId",
				  "as": "orders"
				}
			  },
			  {
				"$unwind": "$orders"
			  },
			  {
				"$group": {
				  "_id": "$userId",
				  "totalSpending": { "$sum": "$orders.totalAmount" }
				}
			  },
			  {
				"$match": {
				  "totalSpending": { "$gt": 500 }
				}
			  }
		]);
		
2.	List Popular Products by Average Rating 

		db.products.aggregate([
			  {
				"$unwind": "$ratings"
			  },
			  {
				"$group": {
				  "_id": "$productId",
				  "averageRating": { "$avg": "$ratings.rating" }
				}
			  },
			  {
				"$match": {
				  "averageRating": { "$gte": 4 }
				}
			  }
		]);
		
3.	Search for Orders in a Specific Time Range 
		
		db.orders.aggregate([
			  {
				"$addFields": {
				  "orderDateISO": { "$dateFromString": { "dateString": "$orderDate" } }
				}
			  },
			  {
				"$match": {
				  "orderDateISO": {
					"$gte": ISODate("2024-12-01T00:00:00Z"),
					"$lte": ISODate("2024-12-31T23:59:59Z")
				  }
				}
			  },
			  {
				"$lookup": {
				  "from": "users",
				  "localField": "userId",
				  "foreignField": "userId",
				  "as": "userDetails"
				}
			  },
			  {
				"$unwind": "$userDetails"
			  },
			  {
				"$project": {
				  "orderId": 1,
				  "orderDate": 1,
				  "userId": 1,
				  "userName": "$userDetails.name"
				}
			  }
		]);
		
4.	Update Stock After Order Completion

		db.orders.find({ "status": "Delivered" }).forEach(order => {
			  order.items.forEach(item => {
				db.products.updateOne(
				  { "productId": item.productId },
				  { "$inc": { "stock": -item.quantity } }
				);
			  });
		});
		
		db.products.find();
		
5.	Find Nearest Warehouse

		db.createCollection("warehouses");

		db.warehouses.insertMany([
			  {
				"warehouseId": "W001",
				"location": { "type": "Point", "coordinates": [-74.006, 40.7128] },
				"products": ["P001", "P002", "P003"]
			  },
			  {
				"warehouseId": "W002",
				"location": { "type": "Point", "coordinates": [-118.2437, 34.0522] },
				"products": ["P002", "P003"]
			  }
		]);
		
		db.warehouses.aggregate([
			  {
				"$geoNear": {
				  "near": { "type": "Point", "coordinates": [-60, 30] },
				  "distanceField": "distance",
				  "maxDistance": 50000,
				  "query": { "products": "P001" },
				  "spherical": true
				}
			  }
		]);






