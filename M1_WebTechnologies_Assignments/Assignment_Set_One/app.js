document.addEventListener("DOMContentLoaded", () => {
    const taskInput = document.getElementById("taskInput");
    const addTaskButton = document.getElementById("addTaskButton");
    const taskList = document.getElementById("taskList");
    const pendingCount = document.getElementById("pendingCount");

    // Load tasks from localStorage
    const loadTasks = () => {
        const tasks = JSON.parse(localStorage.getItem("tasks")) || [];
        tasks.forEach(task => renderTask(task));
        updatePendingCount();
    };

    // Save tasks to localStorage
    const saveTasks = () => {
        const tasks = Array.from(taskList.children).map(taskItem => ({
            id: taskItem.dataset.id,
            text: taskItem.querySelector(".task-text").textContent,
            completed: taskItem.classList.contains("completed")
        }));
        localStorage.setItem("tasks", JSON.stringify(tasks));
    };

    // Render a task in the list
    const renderTask = task => {
        const taskItem = document.createElement("li");
        taskItem.className = `task ${task.completed ? "completed" : ""}`;
        taskItem.dataset.id = task.id;
        taskItem.draggable = true;
        taskItem.innerHTML = `
            <span class="task-text">${task.text}</span>
            <div class="task-actions">
                <button class="edit">Edit</button>
                <button class="delete">Delete</button>
                <button class="toggle">${task.completed ? "Undo" : "Complete"}</button>
            </div>
        `;

        // Attach event listeners
        taskItem.querySelector(".edit").addEventListener("click", () => editTask(taskItem));
        taskItem.querySelector(".delete").addEventListener("click", () => deleteTask(taskItem));
        taskItem.querySelector(".toggle").addEventListener("click", () => toggleTask(taskItem));
        taskItem.addEventListener("dragstart", handleDragStart);
        taskItem.addEventListener("dragover", handleDragOver);
        taskItem.addEventListener("drop", handleDrop);
        taskItem.addEventListener("dragend", handleDragEnd);

        taskList.appendChild(taskItem);
    };

    // Add a new task
    const addTask = () => {
        const text = taskInput.value.trim();
        if (text === "") return;
        const task = {
            id: Date.now().toString(),
            text,
            completed: false
        };
        renderTask(task);
        saveTasks();
        taskInput.value = "";
        updatePendingCount();
    };

    // Edit a task
    const editTask = taskItem => {
        const newText = prompt("Edit Task", taskItem.querySelector(".task-text").textContent);
        if (newText) {
            taskItem.querySelector(".task-text").textContent = newText;
            saveTasks();
        }
    };

    // Delete a task
    const deleteTask = taskItem => {
        taskList.removeChild(taskItem);
        saveTasks();
        updatePendingCount();
    };

    // Toggle task completion
    const toggleTask = taskItem => {
        taskItem.classList.toggle("completed");
        const toggleButton = taskItem.querySelector(".toggle");
        toggleButton.textContent = taskItem.classList.contains("completed") ? "Undo" : "Complete";
        saveTasks();
        updatePendingCount();
    };

    // Drag and drop functionality
    let draggedItem = null;

    const handleDragStart = e => {
        draggedItem = e.target;
        e.target.style.opacity = "0.5";
    };

    const handleDragOver = e => {
        e.preventDefault();
    };

    const handleDrop = e => {
        e.preventDefault();
        if (draggedItem && draggedItem !== e.target) {
            const items = Array.from(taskList.children);
            const draggedIndex = items.indexOf(draggedItem);
            const targetIndex = items.indexOf(e.target);

            if (draggedIndex > targetIndex) {
                taskList.insertBefore(draggedItem, e.target);
            } else {
                taskList.insertBefore(draggedItem, e.target.nextSibling);
            }
            saveTasks();
        }
    };

    const handleDragEnd = e => {
        e.target.style.opacity = "1";
        draggedItem = null;
    };

    // Update the pending tasks count
    const updatePendingCount = () => {
        const pendingTasks = Array.from(taskList.children).filter(taskItem => !taskItem.classList.contains("completed"));
        pendingCount.textContent = pendingTasks.length;
    };

    // Event listeners
    addTaskButton.addEventListener("click", addTask);
    taskInput.addEventListener("keypress", e => {
        if (e.key === "Enter") addTask();
    });

    // Initialize tasks
    loadTasks();
});
