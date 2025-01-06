from customer_management import *

class Transaction(Customer):
    def __init__(self, name, email, phone, book_title, quantity_sold):
        super().__init__(name, email, phone)
        self.book_title = book_title
        self.quantity_sold = quantity_sold

def sell_book(book_list, customer_list, sales_list, name, email, phone, book_title, quantity):
    for book in book_list:
        if book.title == book_title:
            if book.quantity >= quantity:
                book.quantity -= quantity
                sales_list.append(Transaction(name, email, phone, book_title, quantity))
                return f"Sale successful! Remaining quantity: {book.quantity}"
            else:
                return f"Error: Only {book.quantity} copies available."
    return "Book not found."

def view_sales(sales_list):
    return [f"{sale.name} bought {sale.quantity_sold} copies of {sale.book_title}" for sale in sales_list]
