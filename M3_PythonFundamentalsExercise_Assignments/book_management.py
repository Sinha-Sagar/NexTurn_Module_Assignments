class Book:
    def __init__(self, title, author, price, quantity):
        self.title = title
        self.author = author
        self.price = price
        self.quantity = quantity

    def display(self):
        return f"{self.title} by {self.author}, Price: {self.price}, Quantity: {self.quantity}"

def add_book(book_list, title, author, price, quantity):
    book_list.append(Book(title, author, price, quantity))
    return "Book added successfully!"

def view_books(book_list):
    return [book.display() for book in book_list]

def search_book(book_list, keyword):
    return [book.display() for book in book_list if keyword.lower() in book.title.lower() or keyword.lower() in book.author.lower()]