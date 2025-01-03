document.addEventListener("DOMContentLoaded", () => {
    const expenseForm = document.getElementById("expenseForm");
    const expenseTableBody = document.querySelector("#expenseTable tbody");
    const totalSpendingList = document.getElementById("totalSpending");
    const expenseChartElement = document.getElementById("expenseChart");

    let expenses = JSON.parse(localStorage.getItem("expenses")) || [];

    const updateLocalStorage = () => {
        localStorage.setItem("expenses", JSON.stringify(expenses));
    };

    const renderExpenses = () => {
        expenseTableBody.innerHTML = "";
        expenses.forEach((expense, index) => {
            const row = document.createElement("tr");
            row.innerHTML = `
                <td>${expense.amount}</td>
                <td>${expense.description}</td>
                <td>${expense.category}</td>
                <td><button onclick="deleteExpense(${index})">Delete</button></td>
            `;
            expenseTableBody.appendChild(row);
        });
        updateLocalStorage();
        updateSummary();
        renderChart();
    };

    const updateSummary = () => {
        const categoryTotals = expenses.reduce((acc, expense) => {
            acc[expense.category] = (acc[expense.category] || 0) + parseFloat(expense.amount);
            return acc;
        }, {});

        totalSpendingList.innerHTML = Object.entries(categoryTotals).map(
            ([category, total]) => `<li>${category}: Rs. ${total.toFixed(2)}</li>`
        ).join("");
    };

    const renderChart = () => {
        const categoryTotals = expenses.reduce((acc, expense) => {
            acc[expense.category] = (acc[expense.category] || 0) + parseFloat(expense.amount);
            return acc;
        }, {});

        new Chart(expenseChartElement, {
            type: 'pie',
            data: {
                labels: Object.keys(categoryTotals),
                datasets: [{
                    data: Object.values(categoryTotals),
                    backgroundColor: ['#ff6384', '#36a2eb', '#ffce56']
                }]
            }
        });
    };

    expenseForm.addEventListener("submit", event => {
        event.preventDefault();
        const amount = document.getElementById("amount").value;
        const description = document.getElementById("description").value;
        const category = document.getElementById("category").value;

        expenses.push({ amount, description, category });
        renderExpenses();
        expenseForm.reset();
    });

    window.deleteExpense = index => {
        expenses.splice(index, 1);
        renderExpenses();
    };

    renderExpenses();
});
