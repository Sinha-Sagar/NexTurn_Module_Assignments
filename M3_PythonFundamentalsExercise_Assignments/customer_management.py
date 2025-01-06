class Customer:
    def __init__(self, name, email, phone):
        self.name = name
        self.email = email
        self.phone = phone

def add_customer(customer_list, name, email, phone):
    customer_list.append(Customer(name, email, phone))
    return "Customer added successfully!"

def view_customers(customer_list):
    return [f"{customer.name}, Email: {customer.email}, Phone: {customer.phone}" for customer in customer_list]