document.addEventListener('DOMContentLoaded', () => {
    const addTaskForm = document.getElementById('add-task-form');
    const tasksTableBody = document.querySelector('#tasks-table tbody');
    const updateModal = document.getElementById('update-modal');
    const closeBtn = document.querySelector('.close');
    const updateTaskForm = document.getElementById('update-task-form');
    const updateTaskIdInput = document.getElementById('update-task-id');
    const updateTaskTitleInput = document.getElementById('update-task-title');
    const updateTaskDoneInput = document.getElementById('update-task-done');

    closeBtn.onclick = () => {
        updateModal.style.display = 'none';
    };

    window.onclick = (event) => {
        if (event.target === updateModal) {
            updateModal.style.display = 'none';
        }
    };

    const fetchTasks = async () => {
        try {
            const response = await fetch('/tasks');
            const tasks = await response.json();
            renderTasks(tasks);
        } catch (error) {
            console.error('Ошибка при получении задач:', error);
        }
    };

    const renderTasks = (tasks) => {
        tasksTableBody.innerHTML = '';
        tasks.forEach(task => {
            const row = document.createElement('tr');
            row.innerHTML = `
                <td>${task.id}</td>
                <td>${task.title}</td>
                <td>${task.done ? 'Выполнено' : 'Не выполнено'}</td>
                <td>
                    <button class="update-btn" data-id="${task.id}">Обновить</button>
                    <button class="delete-btn" data-id="${task.id}">Удалить</button>
                </td>
            `;
            tasksTableBody.appendChild(row);
        });

        document.querySelectorAll('.update-btn').forEach(btn => {
            btn.addEventListener('click', () => {
                const taskId = btn.getAttribute('data-id');
                openUpdateModal(taskId);
            });
        });

        document.querySelectorAll('.delete-btn').forEach(btn => {
            btn.addEventListener('click', () => {
                const taskId = btn.getAttribute('data-id');
                deleteTask(taskId);
            });
        });
    };

    const openUpdateModal = async (taskId) => {
        try {
            const response = await fetch(`/tasks/${taskId}`);
            const task = await response.json();
            updateTaskIdInput.value = task.id;
            updateTaskTitleInput.value = task.title;
            updateTaskDoneInput.checked = task.done;
            updateModal.style.display = 'block';
        } catch (error) {
            console.error('Ошибка при получении задачи:', error);
        }
    };

    addTaskForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        const title = document.getElementById('task-title').value;

        try {
            const response = await fetch('/tasks', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ title, done: false })
            });

            if (response.ok) {
                await fetchTasks();
                addTaskForm.reset();
            } else {
                const error = await response.text();
                alert(`Ошибка: ${error}`);
            }
        } catch (error) {
            console.error('Ошибка при добавлении задачи:', error);
        }
    });

    updateTaskForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        const taskId = updateTaskIdInput.value;
        const title = updateTaskTitleInput.value;
        const done = updateTaskDoneInput.checked;

        try {
            const response = await fetch(`/tasks/${taskId}`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ title, done })
            });

            if (response.ok) {
                await fetchTasks();
                updateModal.style.display = 'none';
            } else {
                const error = await response.text();
                alert(`Ошибка: ${error}`);
            }
        } catch (error) {
            console.error('Ошибка при обновлении задачи:', error);
        }
    });

    const deleteTask = async (taskId) => {
        try {
            const response = await fetch(`/tasks/${taskId}`, {
                method: 'DELETE'
            });

            if (response.ok) {
                await fetchTasks();
            } else {
                const error = await response.text();
                alert(`Ошибка: ${error}`);
            }
        } catch (error) {
            console.error('Ошибка при удалении задачи:', error);
        }
    };

    fetchTasks();
});