from book_management import *
from customer_management import *
from sales_management import *

book_list = []
customer_list = []
sales_list = []

def main_menu():
    while True:
        print("\nWelcome to BookMart!")
        print("1. Book Management")
        print("2. Customer Management")
        print("3. Sales Management")
        print("4. Exit")
        choice = input("Enter your choice: ")

        if choice == '1':
            print("1. Add Book")
            print("2. View Books")
            print("3. Search Book")
            sub_choice = input("Enter your choice: ")
            if sub_choice == '1':
                try:
                    title = input("Title: ")
                    author = input("Author: ")
                    price = float(input("Price: "))
                    quantity = int(input("Quantity: "))
                    print(add_book(book_list, title, author, price, quantity))
                except ValueError:
                    print("Invalid input! Price and quantity must be numbers.")
            elif sub_choice == '2':
                print("\n".join(view_books(book_list)))
            elif sub_choice == '3':
                keyword = input("Enter title or author: ")
                print("\n".join(search_book(book_list, keyword)))

        elif choice == '2':
            print("1. Add Customer")
            print("2. View Customers")
            sub_choice = input("Enter your choice: ")
            if sub_choice == '1':
                name = input("Name: ")
                email = input("Email: ")
                phone = input("Phone: ")
                print(add_customer(customer_list, name, email, phone))
            elif sub_choice == '2':
                print("\n".join(view_customers(customer_list)))

        elif choice == '3':
            print("1. Sell Book")
            print("2. View Sales")
            sub_choice = input("Enter your choice: ")
            if sub_choice == '1':
                name = input("Customer Name: ")
                email = input("Email: ")
                phone = input("Phone: ")
                book_title = input("Book Title: ")
                try:
                    quantity = int(input("Quantity: "))
                    print(sell_book(book_list, customer_list, sales_list, name, email, phone, book_title, quantity))
                except ValueError:
                    print("Invalid input! Quantity must be a number.")
            elif sub_choice == '2':
                print("\n".join(view_sales(sales_list)))

        elif choice == '4':
            print("Thank you for using BookMart!")
            break

        else:
            print("Invalid choice, please try again!")

if __name__ == "__main__":
    main_menu()
